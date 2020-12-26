package tgql

import (
	"github.com/zhanghup/go-tools/database/txorm"
	"reflect"
	"sync"
	"time"

	"github.com/zhanghup/go-tools"
)

type SliceLoader struct {
	sql          string
	param        map[string]interface{}
	db           txorm.IEngine
	keyField     string
	resultField  string
	requestTable reflect.Type
	resultTable  reflect.Type

	cache map[string]interface{}
	batch *sliceLoaderBatch
	sync  *sync.RWMutex
}

type sliceLoaderBatch struct {
	keys    []string
	data    map[string]interface{}
	error   error
	closing bool
	done    chan struct{}
}

func (this *SliceLoader) fetch(keys []string) (map[string]interface{}, error) {
	query := map[string]interface{}{}
	for k, v := range this.param {
		query[k] = v
	}
	query["keys"] = keys
	if this.requestTable.Kind() != reflect.Struct {
		panic("requestTable必须为struct")
	}
	if this.resultTable.Kind() != reflect.Struct {
		panic("resultTable必须为struct")
	}
	vl := reflect.New(reflect.SliceOf(this.requestTable))

	err := this.db.SF(this.sql, query).Find(vl.Interface())
	if err != nil {
		return nil, err
	}
	result := map[string]interface{}{}

	for i := 0; i < vl.Elem().Len(); i++ {
		vv := vl.Elem().Index(i)
		tools.Rft.DeepGet(vv.Interface(), func(t reflect.Type, v reflect.Value, tf reflect.StructField) bool {
			if tf.Name == this.keyField {
				if v.Kind() == reflect.Ptr && v.Pointer() != 0 {
					t = t.Elem()
					v = v.Elem()
				}

				if v.Kind() == reflect.String {
					if ls, ok := result[v.String()]; ok {
						var lss reflect.Value
						if len(this.resultField) > 0 {
							vvv := vv.FieldByName(this.resultField)
							if !vv.CanInterface() {
								return false
							}
							lss = reflect.Append(reflect.ValueOf(ls), vvv)
						} else {
							lss = reflect.Append(reflect.ValueOf(ls), vv)
						}
						result[v.String()] = lss.Interface()
					} else {
						ls := reflect.New(reflect.SliceOf(this.resultTable)).Elem()
						if len(this.resultField) > 0 {
							tools.Rft.DeepGet(vv.Interface(), func(t reflect.Type, v reflect.Value, tf reflect.StructField) bool {
								if tf.Name == this.resultField {
									ls = reflect.Append(ls, v)
									return false
								}
								return true
							})
						} else {
							ls = reflect.Append(ls, vv)
						}
						result[v.String()] = ls.Interface()
					}
				}
				return false
			}
			return true
		})
	}
	return result, nil
}

func (l *SliceLoader) Load(key string, result interface{}) error {
	i, err := l.LoadThunk(key)()
	if err != nil {
		return err
	}
	if i == nil {
		return nil
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(i))
	return nil
}

func (l *SliceLoader) LoadThunk(key string) func() (interface{}, error) {
	l.sync.Lock()
	if it, ok := l.cache[key]; ok {
		l.sync.Unlock()
		return func() (interface{}, error) {
			return it, nil
		}
	}
	if l.batch == nil {
		l.batch = &sliceLoaderBatch{done: make(chan struct{})}
	} else if l.batch.closing {
		l.batch.keys = nil
		l.batch.data = nil
		l.batch.error = nil
		l.batch.closing = false
		l.batch.done = make(chan struct{})
	}
	batch := l.batch
	batch.keyIndex(l, key)
	l.sync.Unlock()

	return func() (interface{}, error) {
		<-batch.done

		if batch.error == nil {
			l.sync.Lock()
			l.unsafeSet(key, batch.data[key])
			l.sync.Unlock()
		}

		return batch.data[key], batch.error
	}
}

func (l *SliceLoader) unsafeSet(key string, value interface{}) {
	if l.cache == nil {
		l.cache = map[string]interface{}{}
	}
	l.cache[key] = value
}

func (b *sliceLoaderBatch) keyIndex(l *SliceLoader, key string) {
	for _, existingKey := range b.keys {
		if key == existingKey {
			return
		}
	}

	pos := len(b.keys)
	b.keys = append(b.keys, key)
	if pos == 0 {
		go b.startTimer(l)
	}

	return
}

func (b *sliceLoaderBatch) startTimer(l *SliceLoader) {
	time.Sleep(time.Millisecond * 5)
	l.sync.Lock()

	if b.closing {
		l.sync.Unlock()
		return
	}

	l.sync.Unlock()
	b.end(l)
}

func (b *sliceLoaderBatch) end(l *SliceLoader) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
	b.closing = true

}
