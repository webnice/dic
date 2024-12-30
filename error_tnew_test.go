package dic

import (
	"reflect"
	"testing"
)

func TestTError_TNew(t *testing.T) {
	const (
		errT = "Когда не можешь ничего поделать, просто прими это (c) %q."
		errA = "Доктор Хаус"
	)
	var (
		errs    *TestCustomError
		errI    IError
		err1    error
		err2    error
		originT string
		originA []string
		tmpT    bool
		tmpS    string
		tmpA    []string
	)

	errs = makeTestCustomErrors()
	// Получение оригинального шаблона ошибки и аргументов ошибки.
	originT = errs.Fatality.Anchor().Error()
	originA = errs.Fatality.TArg()
	// Создание объекта ошибки.
	err1 = errs.Fatality.Bind()
	// Проверка эквивалентности ошибки.
	if tmpT = errs.Fatality.IsEqual(ParseError(err1)); !tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, true)
		return
	}
	// Изменение ошибки в справочнике ошибок.
	_ = errs.Fatality.TNew(errT, []string{errA}...)
	// Сравнение шаблона и аргументов.
	if tmpS = errs.Fatality.Anchor().Error(); tmpS == originT {
		t.Fatalf("Anchor(), получено %q, ошидалось %q", tmpS, errT)
		return
	}
	if tmpA = errs.Fatality.TArg(); reflect.DeepEqual(tmpA, originA) {
		t.Fatalf("TArg(), получено %v, ошидалось %v", tmpA, []string{errA})
		return
	}
	// Сравнение эквивалентности ошибки.
	err2 = errs.Fatality.Bind()
	errI = errs.ParseError(err2)
	if tmpT = errI.IsEqual(ParseError(err1)); tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, false)
		return
	}
}
