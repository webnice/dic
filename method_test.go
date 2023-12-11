package dic

import (
	"reflect"
	"testing"
)

func TestMethod(t *testing.T) {
	var (
		m  Methods
		n  int
		rt reflect.Type
		sf reflect.StructField
	)

	m = Method()
	rt = reflect.TypeOf(m)
	if rt.NumField() <= 0 {
		t.Errorf("справочник HTTP методов содержит %d элементов, ожидалось больше %d", rt.NumField(), 0)
	}
	for n = 0; n < rt.NumField(); n++ {
		sf = rt.Field(n)
		if sf.Type.String() != "dic.IMethod" {
			t.Errorf(
				"справочник HTTP методов, поле %q, тип поля %q, ожидалось: %q",
				sf.Name, sf.Type.String(), "dic.IMethod",
			)
		}
	}
}

func TestTMethod_IsEqual(t *testing.T) {
	var (
		post IMethod
		ok   bool
	)

	post = ParseMethod("post")
	if post == nil {
		t.Errorf("функция ParseMethod(), вернулся nil, ожидался не nil объект")
	}
	if reflect.ValueOf(post).Pointer() != reflect.ValueOf(Method().Post).Pointer() {
		t.Errorf("функция NewMethod(), вернулся объект: %v, ожидался объект: %v", post, Method().Post)
	}
	if ok = Method().Post.IsEqualFull(post); !ok {
		t.Errorf("функция IsEqualFull(), вернулось значение: %t, ожидалось: %t", ok, true)
	}
}

func TestTMethod_IsEqualNil(t *testing.T) {
	var (
		tests = []struct {
			text   string
			method IMethod
			eq     bool
		}{
			{"", nil, true},
			{"", Method().Get, false},
			{"vBZ6IC4vUQrPU8Tv2ceh", Method().Get, false},
		}
		method IMethod
		n      int
	)

	for n = range tests {
		method = ParseMethod(tests[n].text)
		if tests[n].method == nil && tests[n].eq && method != tests[n].method {
			t.Errorf("функция ParseMethod(%q), вернулся объект: \"%v\", ожидался \"%v\"", tests[n].text, method, nil)
			continue
		}
		if tests[n].method == nil {
			continue
		}
		if tests[n].method.IsEqual(method) != tests[n].eq {
			t.Errorf("функция IsEqual(), вернулся объект: \"%v\", ожидался объект: \"%v\"", method, tests[n].method)
		}
	}
}

func TestTMethod_IsEqualFull(t *testing.T) {
	var (
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
		method IMethod
		n      int
	)

	for n = range tests {
		method = ParseMethod(tests[n].text)
		if tests[n].method == nil && tests[n].eq && method != tests[n].method {
			t.Errorf("функция ParseMethod(%q), вернулся объект: \"%v\", ожидался \"%v\"", tests[n].text, method, nil)
			continue
		}
		if tests[n].method == nil {
			continue
		}
		if tests[n].method.IsEqualFull(method) != tests[n].eq {
			t.Errorf("функция IsEqualFull(), вернулся объект: \"%v\", ожидался объект: \"%v\"", method, tests[n].method)
		}
	}
}
