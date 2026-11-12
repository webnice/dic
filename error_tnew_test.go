package dic

import (
	"reflect"
	"testing"
)

func TestTError_TNew(t *testing.T) {
	const (
		errT       = "Когда не можешь ничего поделать, просто прими это (c) %q."
		errA       = "Доктор Хаус"
		tplAnchor  = "Anchor(), получено %q, ожидалось %q"
		tplTArg    = "TArg(), получено %v, ожидалось %v"
		tplIsEqual = "IsEqual(), получено %t, ожидалось %t"
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
		t.Fatalf(tplIsEqual, tmpT, true)
	}
	// Изменение ошибки в справочнике ошибок.
	_ = errs.Fatality.TNew(errT, []string{errA}...)
	// Сравнение шаблона и аргументов.
	if tmpS = errs.Fatality.Anchor().Error(); tmpS == originT {
		t.Fatalf(tplAnchor, tmpS, errT)
	}
	if tmpA = errs.Fatality.TArg(); reflect.DeepEqual(tmpA, originA) {
		t.Fatalf(tplTArg, tmpA, []string{errA})
	}
	// Сравнение эквивалентности ошибки.
	err2 = errs.Fatality.Bind()
	errI = errs.ParseError(err2)
	if tmpT = errI.IsEqual(ParseError(err1)); tmpT {
		t.Fatalf(tplIsEqual, tmpT, false)
	}
}
