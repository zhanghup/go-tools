package loader

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"reflect"
	"time"
)

var _cache = tools.NewCache[any]()

func Load[T any](id string, fetch func(keys []string) (map[string]T, error)) IObject[T] {
	snc := tools.Mutex("51e761c0-d4ff-478d-923a-14fb5b2bd0af,f3fe7357-2908-4758-8652-1778bb764b27")

	ty := reflect.TypeOf(new(T))
	key := fmt.Sprintf("%s,%s,%s,%s,%s", ty.PkgPath(), ty.Name(), ty.String(), ty.Kind().String(), id)

	snc.Lock()
	defer snc.Unlock()
	obj, ok := _cache.Get(key)
	if ok {
		return obj.(IObject[T])
	}

	oo := NewObjectLoader[T](fetch)
	_cache.Set(key, oo, time.Now().Unix()+86400)
	return oo
}
