package dic

import (
	"reflect"
	"testing"
)

func TestTError_Singleton(t *testing.T) {
	var (
		f  map[string]bool
		rv reflect.Type
		n  int
		k  string
		ok bool
	)

	f = map[string]bool{
		"Unknown": false,
	}
	rv = reflect.TypeOf(singletonErrors)
	for n = 0; n < rv.NumField(); n++ {
		if _, ok = f[rv.Field(n).Name]; ok {
			f[rv.Field(n).Name] = ok
		}
	}
	ok = true
	for k = range f {
		if !f[k] {
			ok = false
		}
	}
	if !ok {
		t.Fatal("Не найден один или несколько обязательных ошибок.")
		return
	}
}

func TestTError_Error(t *testing.T) {
	var (
		f  map[string]bool
		rv reflect.Type
		n  int
		k  string
		ok bool
	)

	f = map[string]bool{
		"Unknown": false,
	}
	rv = reflect.TypeOf(Error())
	for n = 0; n < rv.NumField(); n++ {
		if _, ok = f[rv.Field(n).Name]; ok {
			f[rv.Field(n).Name] = ok
		}
	}
	ok = true
	for k = range f {
		if !f[k] {
			ok = false
		}
	}
	if !ok {
		t.Fatal("Не найден один или несколько обязательных ошибок.")
		return
	}
}
