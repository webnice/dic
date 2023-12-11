package dic

import (
	"reflect"
	"testing"
)

func TestParseHeader(t *testing.T) {
	var cth = ParseHeader("content-Type")

	if cth == nil {
		t.Errorf("функция ParseHeader(), вернулся nil, ожидался не nil объект")
	}
	if reflect.ValueOf(cth).Pointer() != reflect.ValueOf(Header().ContentType).Pointer() {
		t.Errorf("функция ParseHeader(), вернулся объект: %v, ожидался объект: %v", cth, Header().ContentType)
	}
}

func TestNewHeader(t *testing.T) {
	const newHeader = "X-Completely-New-Header"
	var (
		cth IHeader
		xnw IHeader
	)

	// Проверка, должен вернуться объект существующего стандартного метода.
	cth = NewHeader("content-Type")
	if cth == nil {
		t.Errorf("функция NewHeader(), вернулся nil, ожидался не nil объект")
	}
	if reflect.ValueOf(cth).Pointer() != reflect.ValueOf(Header().ContentType).Pointer() {
		t.Errorf("функция NewHeader(), вернулся объект: %v, ожидался объект: %v", cth, Header().ContentType)
	}
	// Новый, не существующий заголовок.
	xnw = NewHeader(newHeader)
	if xnw == nil {
		t.Errorf("функция NewHeader(), вернулся nil, ожидался не nil объект")
	}
	if xnw.(*tHeader).header != newHeader {
		t.Errorf("функция NewHeader(), вернулся объект: %v, ожидался объект: %v", xnw, newHeader)
	}
	if xnw.String() != newHeader {
		t.Errorf("функция String(), вернулось значение: %q, ожидалось: %q", xnw.String(), newHeader)
	}
}
