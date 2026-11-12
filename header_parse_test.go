package dic

import (
	"reflect"
	"testing"
)

func TestParseHeader(t *testing.T) {
	const tplNil = "функция ParseHeader(), вернулся nil, ожидался не nil объект"
	const tplParseHeader = "функция ParseHeader(), вернулся объект: %v, ожидался объект: %v"
	var cth IHeader

	cth = ParseHeader("content-Type")
	if cth == nil {
		t.Errorf(tplNil)
	}
	if reflect.ValueOf(cth).Pointer() != reflect.ValueOf(Header().ContentType).Pointer() {
		t.Errorf(tplParseHeader, cth, Header().ContentType)
	}
}

func TestNewHeader(t *testing.T) {
	const (
		newHeader    = "X-Completely-Bind-Header"
		tplNil       = "функция NewHeader(), вернулся nil, ожидался не nil объект"
		tplNewHeader = "функция NewHeader(), вернулся объект: %v, ожидался объект: %v"
		tplString    = "функция String(), вернулось значение: %q, ожидалось: %q"
	)
	var (
		cth IHeader
		xnw IHeader
	)

	// Проверка, должен вернуться объект существующего стандартного заголовка.
	cth = NewHeader("content-Type")
	if cth == nil {
		t.Errorf(tplNil)
	}
	if reflect.ValueOf(cth).Pointer() != reflect.ValueOf(Header().ContentType).Pointer() {
		t.Errorf(tplNewHeader, cth, Header().ContentType)
	}
	// Новый, не существующий заголовок.
	xnw = NewHeader(newHeader)
	if xnw == nil {
		t.Errorf(tplNil)
	}
	if xnw.(*tHeader).header != newHeader {
		t.Errorf(tplNewHeader, xnw, newHeader)
	}
	if xnw.String() != newHeader {
		t.Errorf(tplString, xnw.String(), newHeader)
	}
}
