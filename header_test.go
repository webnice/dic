package dic

import (
	"reflect"
	"testing"
)

func TestHeader(t *testing.T) {
	var (
		h  Headers
		n  int
		rt reflect.Type
		sf reflect.StructField
	)

	h = Header()
	rt = reflect.TypeOf(h)
	if rt.NumField() <= 0 {
		t.Errorf("справочник заголовков содержит %d элементов, ожидалось больше %d", rt.NumField(), 0)
	}
	for n = 0; n < rt.NumField(); n++ {
		sf = rt.Field(n)
		if sf.Type.String() != "dic.IHeader" {
			t.Errorf(
				"справочник заголовков, поле %q, тип поля %q, ожидалось: %q",
				sf.Name, sf.Type.String(), "dic.IHeader",
			)
		}
	}
}

func TestTHeader_IsEqual(t *testing.T) {
	const newHeader1, newHeader2 = "X-Completely-Bind-Header", "X-Bind-Header"
	var (
		h1, h2 IHeader
		ok     bool
	)

	if h1, h2 = NewHeader(newHeader1), NewHeader(newHeader2); h1 == nil || h2 == nil {
		t.Errorf("функция NewHeader(), вернулся nil, ожидался не nil объект")
	}
	if ok = h1.IsEqual(h2); ok {
		t.Errorf("функция IsEqual(), вернулся: %t, ожидался: %t", ok, false)
	}
	h2 = NewHeader(newHeader1)
	if ok = h1.IsEqual(h2); !ok {
		t.Errorf("функция IsEqual(), вернулся: %t, ожидался: %t", ok, true)
	}
}

func TestTHeader_IsEqualNil(t *testing.T) {
	var (
		tests = []struct {
			text   string
			header IHeader
			eq     bool
		}{
			{"", nil, true},
			{"", Header().ContentType, false},
		}
		header IHeader
		n      int
	)

	for n = range tests {
		header = ParseHeader(tests[n].text)
		if tests[n].header == nil && tests[n].eq && header != tests[n].header {
			t.Errorf("функция ParseHeader(%q), вернулся объект: \"%v\", ожидался \"%v\"", tests[n].text, header, nil)
			continue
		}
		if tests[n].header == nil {
			continue
		}
		if tests[n].header.IsEqual(header) != tests[n].eq {
			t.Errorf("функция IsEqual(), вернулся объект: \"%v\", ожидался объект: \"%v\"", header, tests[n].header)
		}
	}
}
