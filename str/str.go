package str

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

func Uid() string {
	machineId := func() []byte {
		var sum [3]byte
		id := sum[:]
		hostname, err1 := os.Hostname()
		if err1 != nil {
			n := uint32(time.Now().UnixNano())
			sum[0] = byte(n >> 0)
			sum[1] = byte(n >> 8)
			sum[2] = byte(n >> 16)
			return id
		}
		hw := md5.New()
		hw.Write([]byte(hostname))
		copy(id, hw.Sum(nil))
		return id
	}()
	processId := os.Getpid()
	objectIdCounter := uint32(time.Now().UnixNano())

	var b [12]byte
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineId[0]
	b[5] = machineId[1]
	b[6] = machineId[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	b[7] = byte(processId >> 8)
	b[8] = byte(processId)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectIdCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return hex.EncodeToString(b[:])
}

// 以json格式输出struct对象
func JSONString(obj interface{}, format ...bool) string {
	var datas []byte
	if len(format) > 0 && format[0] {

		r, err := json.MarshalIndent(obj, "\t", "\t")
		if err != nil {
			datas = []byte("数据格式化异常")
		} else {
			datas = r
		}
	} else {
		r, err := json.Marshal(obj)
		if err != nil {
			datas = []byte("数据格式化异常")
		} else {
			datas = r
		}
	}
	return string(datas)
}

func Println(obj interface{}, format ...bool) {

	_, _ = os.Stdout.Write([]byte(JSONString(obj, format...)))
}

func Contains(src []string, tag string) bool {
	for _, s := range src {
		if s == tag {
			return true
		}
	}
	return false
}

func Rand(l int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	res := make([]byte, l)
	for i := 0; i < l; i++ {
		b := r.Intn(26) + 65
		res[i] = byte(b)
	}
	return string(res)
}
