package dic

import (
	"reflect"
	"testing"
)

func TestParseMethod(t *testing.T) {
	const (
		tplNil   = "функция ParseMethod(), вернулся nil, ожидался не nil объект"
		tplParse = "функция NewMethod(), вернулся объект: %v, ожидался объект: %v"
	)
	var post IMethod

	post = ParseMethod("post")
	if post == nil {
		t.Errorf(tplNil)
	}
	if reflect.ValueOf(post).Pointer() != reflect.ValueOf(Method().Post).Pointer() {
		t.Errorf(tplParse, post, Method().Post)
	}
}

func TestNewMethod(t *testing.T) {
	const (
		tplNil       = "функция NewMethod(), вернулся nil, ожидался не nil объект"
		tplNewMethod = "функция NewMethod(), вернулся объект: %v, ожидался объект: %v"
		tplString    = "функция String(), вернулось значение: %q, ожидалось: %q"
	)
	var (
		get IMethod
		fsh IMethod
	)

	// Проверка, должен вернуться объект существующего стандартного метода.
	get = NewMethod("get")
	if get == nil {
		t.Errorf(tplNil)
	}
	if reflect.ValueOf(get).Pointer() != reflect.ValueOf(Method().Get).Pointer() {
		t.Errorf(tplNewMethod, get, Method().Get)
	}
	// Новый, не существующий метод.
	fsh = NewMethod("flush")
	if fsh == nil {
		t.Errorf(tplNil)
	}
	if fsh.(*tMethod).bits == 0 || fsh.(*tMethod).name != "FLUSH" {
		t.Errorf(tplNewMethod, fsh, "FLUSH")
	}
	if fsh.String() != "FLUSH" {
		t.Errorf(tplString, fsh.String(), "FLUSH")
	}
}
