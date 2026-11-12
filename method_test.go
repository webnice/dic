package dic

import (
	"reflect"
	"testing"
)

func TestMethod(t *testing.T) {
	const (
		tplCount    = "справочник HTTP методов содержит %d элементов, ожидалось больше %d"
		tplField    = "справочник HTTP методов, поле %q, тип поля %q, ожидалось: %q"
		typeIMethod = "dic.IMethod"
	)
	var (
		m  Methods
		n  int
		rt reflect.Type
		sf reflect.StructField
	)

	m = Method()
	rt = reflect.TypeOf(m)
	if rt.NumField() <= 0 {
		t.Errorf(tplCount, rt.NumField(), 0)
	}
	for n = 0; n < rt.NumField(); n++ {
		sf = rt.Field(n)
		if sf.Type.String() != typeIMethod {
			t.Errorf(tplField, sf.Name, sf.Type.String(), typeIMethod)
		}
	}
}

func TestTMethod_IsEqual(t *testing.T) {
	const (
		tplNil     = "функция ParseMethod(), вернулся nil, ожидался не nil объект"
		tplParse   = "функция NewMethod(), вернулся объект: %v, ожидался объект: %v"
		tplIsEqual = "функция IsEqualFull(), вернулось значение: %t, ожидалось: %t"
	)
	var (
		post IMethod
		ok   bool
	)

	post = ParseMethod("post")
	if post == nil {
		t.Errorf(tplNil)
	}
	if reflect.ValueOf(post).Pointer() != reflect.ValueOf(Method().Post).Pointer() {
		t.Errorf(tplParse, post, Method().Post)
	}
	if ok = Method().Post.IsEqualFull(post); !ok {
		t.Errorf(tplIsEqual, ok, true)
	}
}

func TestTMethod_IsEqualNil(t *testing.T) {
	const (
		tplParseMethod = "функция ParseMethod(%q), вернулся объект: \"%v\", ожидался \"%v\""
		tplIsEqualObj  = "функция IsEqual(), вернулся объект: \"%v\", ожидался объект: \"%v\""
	)
	var (
		tests []struct {
			text   string
			method IMethod
			eq     bool
		}
		method IMethod
		n      int
	)

	// Таблица тестовых данных.
	tests = []struct {
		text   string
		method IMethod
		eq     bool
	}{
		{"", nil, true},
		{"", Method().Get, false},
		{"vBZ6IC4vUQrPU8Tv2ceh", Method().Get, false},
	}
	for n = range tests {
		method = ParseMethod(tests[n].text)
		if tests[n].method == nil && tests[n].eq && method != tests[n].method {
			t.Errorf(tplParseMethod, tests[n].text, method, nil)
			continue
		}
		if tests[n].method == nil {
			continue
		}
		if tests[n].method.IsEqual(method) != tests[n].eq {
			t.Errorf(tplIsEqualObj, method, tests[n].method)
		}
	}
}

func TestTMethod_IsEqualFull(t *testing.T) {
	const (
		tplParseMethod = "функция ParseMethod(%q), вернулся объект: \"%v\", ожидался \"%v\""
		tplIsEqualObj  = "функция IsEqualFull(), вернулся объект: \"%v\", ожидался объект: \"%v\""
	)
	var (
		tests []struct {
			text   string
			method IMethod
			eq     bool
		}
		method IMethod
		n      int
	)

	// Таблица тестовых данных.
	tests = []struct {
		text   string
		method IMethod
		eq     bool
	}{
		{"get", Method().Get, true},
		{"post", Method().Post, true},
		{"patch", Method().Post, false},
		{"", nil, true},
		{"", Method().Get, false},
		{"vBZ6IC4vUQrPU8Tv2ceh", Method().Get, false},
	}
	for n = range tests {
		method = ParseMethod(tests[n].text)
		if tests[n].method == nil && tests[n].eq && method != tests[n].method {
			t.Errorf(tplParseMethod, tests[n].text, method, nil)
			continue
		}
		if tests[n].method == nil {
			continue
		}
		if tests[n].method.IsEqualFull(method) != tests[n].eq {
			t.Errorf(tplIsEqualObj, method, tests[n].method)
		}
	}
}

// TestTMethod_IsEqualFullBits Проверка сравнения методов с одинаковым именем и разными битовыми масками.
func TestTMethod_IsEqualFullBits(t *testing.T) {
	const tplIsEqualFull = "функция IsEqualFull(), вернулось: %t, ожидалось: %t"
	// Метод с тем же именем, но с другой битовой маской не эквивалентен.
	if Method().Get.IsEqualFull(&tMethod{name: "get", bits: 1 << 20}) {
		t.Errorf(tplIsEqualFull, true, false)
	}
	// Метод с тем же именем и той же битовой маской эквивалентен.
	if !Method().Get.IsEqualFull(&tMethod{name: "GET", bits: 1}) {
		t.Errorf(tplIsEqualFull, false, true)
	}
}

// TestTMethod_IsEqualForeign Проверка сравнения метода с посторонней реализацией интерфейса IMethod.
func TestTMethod_IsEqualForeign(t *testing.T) {
	const tplIsEqual = "функция IsEqual(), вернулось: %t, ожидалось: %t"
	var foreign IMethod = foreignMethod{}

	// Посторонняя реализация интерфейса IMethod не приводит к панике и не эквивалентна методу справочника.
	if Method().Get.IsEqual(foreign) {
		t.Errorf(tplIsEqual, true, false)
	}
	if Method().Get.IsEqualFull(foreign) {
		t.Errorf(tplIsEqual, true, false)
	}
}

// foreignMethod Посторонняя реализация интерфейса IMethod, не являющаяся объектом справочника.
type foreignMethod struct{}

// IsEqual Реализация интерфейса IMethod.
func (foreignMethod) IsEqual(IMethod) bool { return false }

// IsEqualFull Реализация интерфейса IMethod.
func (foreignMethod) IsEqualFull(IMethod) bool { return false }

// String Реализация интерфейса IMethod.
func (foreignMethod) String() string { return "FOREIGN" }
