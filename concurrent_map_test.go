package gotool

import (
	"math/rand"
	"testing"
)

func TestBasicOPs(t *testing.T) {
	for i := 0; i < 10; i++ {
		testV := rand.Intn(1000)
		m := CreateConcurrentMap(99)
		v, ok := m.Get(StrKey("Hello"))
		if v != nil || ok != false {
			t.Error("init/get failed")
		}
		m.Set(StrKey("Hello"), testV)
		v, ok = m.Get(StrKey("Hello"))
		if v.(int) != testV || ok != true {
			t.Error("set/get failed.")
		}
		m.Del(StrKey("Hello"))
		v, ok = m.Get(StrKey("Hello"))
		if v != nil || ok != false {
			t.Error("del failed")
		}
	}
}
