package dic

import (
	"reflect"
	"testing"
)

func TestTError_Singleton(t *testing.T) {
	const errRequired = "Не найден один или несколько обязательных ошибок."
	var (
		f  map[string]bool
		rv reflect.Type
		n  int
		k  string
		ok bool
	)

	// Список обязательных ошибок справочника.
	f = map[string]bool{"Unknown": false}
	// Проверка наличия обязательных ошибок в справочнике.
	rv = reflect.TypeOf(singletonErrors)
	for n = 0; n < rv.NumField(); n++ {
		if _, ok = f[rv.Field(n).Name]; ok {
			f[rv.Field(n).Name] = ok
		}
	}
	// Контроль результата проверки.
	ok = true
	for k = range f {
		if !f[k] {
			ok = false
		}
	}
	if !ok {
		t.Fatal(errRequired)
	}
}

func TestTError_Error(t *testing.T) {
	const errRequired = "Не найден один или несколько обязательных ошибок."
	var (
		f  map[string]bool
		rv reflect.Type
		n  int
		k  string
		ok bool
	)

	// Список обязательных ошибок справочника.
	f = map[string]bool{"Unknown": false}
	// Проверка наличия обязательных ошибок в справочнике.
	rv = reflect.TypeOf(Error())
	for n = 0; n < rv.NumField(); n++ {
		if _, ok = f[rv.Field(n).Name]; ok {
			f[rv.Field(n).Name] = ok
		}
	}
	// Контроль результата проверки.
	ok = true
	for k = range f {
		if !f[k] {
			ok = false
		}
	}
	if !ok {
		t.Fatal(errRequired)
	}
}
