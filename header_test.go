package dic

import (
	"reflect"
	"testing"
)

func TestHeader(t *testing.T) {
	const (
		tplCount    = "справочник заголовков содержит %d элементов, ожидалось больше %d"
		tplField    = "справочник заголовков, поле %q, тип поля %q, ожидалось: %q"
		typeIHeader = "dic.IHeader"
	)
	var (
		h  Headers
		n  int
		rt reflect.Type
		sf reflect.StructField
	)

	h = Header()
	rt = reflect.TypeOf(h)
	if rt.NumField() <= 0 {
		t.Errorf(tplCount, rt.NumField(), 0)
	}
	for n = 0; n < rt.NumField(); n++ {
		sf = rt.Field(n)
		if sf.Type.String() != typeIHeader {
			t.Errorf(tplField, sf.Name, sf.Type.String(), typeIHeader)
		}
	}
}

func TestTHeader_IsEqual(t *testing.T) {
	const (
		newHeader1, newHeader2 = "X-Completely-Bind-Header", "X-Bind-Header"
		tplNil                 = "функция NewHeader(), вернулся nil, ожидался не nil объект"
		tplIsEqual             = "функция IsEqual(), вернулся: %t, ожидался: %t"
	)
	var (
		h1, h2 IHeader
		ok     bool
	)

	if h1, h2 = NewHeader(newHeader1), NewHeader(newHeader2); h1 == nil || h2 == nil {
		t.Errorf(tplNil)
	}
	if ok = h1.IsEqual(h2); ok {
		t.Errorf(tplIsEqual, ok, false)
	}
	h2 = NewHeader(newHeader1)
	if ok = h1.IsEqual(h2); !ok {
		t.Errorf(tplIsEqual, ok, true)
	}
}

func TestTHeader_IsEqualNil(t *testing.T) {
	const tplParseHeader = "функция ParseHeader(%q), вернулся объект: \"%v\", ожидался \"%v\""
	const tplIsEqualObj = "функция IsEqual(), вернулся объект: \"%v\", ожидался объект: \"%v\""
	var (
		tests []struct {
			text   string
			header IHeader
			eq     bool
		}
		header IHeader
		n      int
	)

	// Таблица тестовых данных.
	tests = []struct {
		text   string
		header IHeader
		eq     bool
	}{
		{"", nil, true},
		{"", Header().ContentType, false},
	}
	for n = range tests {
		header = ParseHeader(tests[n].text)
		if tests[n].header == nil && tests[n].eq && header != tests[n].header {
			t.Errorf(tplParseHeader, tests[n].text, header, nil)
			continue
		}
		if tests[n].header == nil {
			continue
		}
		if tests[n].header.IsEqual(header) != tests[n].eq {
			t.Errorf(tplIsEqualObj, header, tests[n].header)
		}
	}
}
