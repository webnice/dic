package dic

import (
	"reflect"
	"testing"
)

func TestParseMethod(t *testing.T) {
	var post = ParseMethod("post")
	if post == nil {
		t.Errorf("функция ParseMethod(), вернулся nil, ожидался не nil объект")
	}
	if reflect.ValueOf(post).Pointer() != reflect.ValueOf(Method().Post).Pointer() {
		t.Errorf("функция NewMethod(), вернулся объект: %v, ожидался объект: %v", post, Method().Post)
	}
}

func TestNewMethod(t *testing.T) {
	var (
		get IMethod
		fsh IMethod
	)

	// Проверка, должен вернуться объект существующего стандартного метода.
	get = NewMethod("get")
	if get == nil {
		t.Errorf("функция NewMethod(), вернулся nil, ожидался не nil объект")
	}
	if reflect.ValueOf(get).Pointer() != reflect.ValueOf(Method().Get).Pointer() {
		t.Errorf("функция NewMethod(), вернулся объект: %v, ожидался объект: %v", get, Method().Get)
	}
	// Новый, не существующий метод.
	fsh = NewMethod("flush")
	if fsh == nil {
		t.Errorf("функция NewMethod(), вернулся nil, ожидался не nil объект")
	}
	if fsh.(*tMethod).bits == 0 || fsh.(*tMethod).name != "FLUSH" {
		t.Errorf("функция NewMethod(), вернулся объект: %v, ожидался объект: %v", fsh, "FLUSH")
	}
	if fsh.String() != "FLUSH" {
		t.Errorf("функция String(), вернулось значение: %q, ожидалось: %q", fsh.String(), "FLUSH")
	}
}
