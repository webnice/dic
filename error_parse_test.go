package dic

import (
	"errors"
	"strings"
	"testing"
)

func TestParseError_Nil(t *testing.T) {
	const tplIsEqual = "IsEqual(), получено %t, ожидалось %t"
	var (
		ierr IError
		tmpT bool
		errs *TestCustomError
	)

	errs = makeTestCustomErrors()
	ierr = errs.ParseError(ierr, errs)
	if tmpT = Error().Unknown.IsEqual(ierr); !tmpT {
		t.Fatalf(tplIsEqual, tmpT, true)
	}
}

func TestNewError(t *testing.T) {
	const (
		errTwoArgs = "ошибка с двумя аргументами, первый аргумент %q, второй аргумент: %s"
		tplIsEqual = "IsEqual(), получено %t, ожидалось %t"
	)
	var (
		errs *TestCustomError
		ierr IError
		tmpT bool
	)

	errs = makeTestCustomErrors()
	// Создание ошибки с поиском по шаблону ошибки и совпадением по шаблону и аргументам.
	ierr = NewError(errTwoArgs, "арг1", "arg2")
	// Должна быть эквивалентной с аналогичной ошибкой из справочника ошибок.
	if tmpT = errs.TwoArgs.IsEqual(ierr); !tmpT {
		t.Fatalf(tplIsEqual, tmpT, true)
	}
}

func TestParseError_CustomDictionary(t *testing.T) {
	const (
		unknown       = "совершенно неизвестная ошибка"
		tplParseFound = "ParseError(), вернулся объект: %v, ожидался объект эквивалентный: %v"
		tplParseNil   = "ParseError(), вернулся nil, ожидался не nil объект"
		tplParseMsg   = "ParseError(), получено: %q, ожидалось сообщение содержащее: %q"
	)
	var (
		dct  *TestCustomDictionary
		ierr IError
	)

	dct = makeTestCustomDictionary()
	// Поиск ошибки, якорь которой есть в справочнике.
	ierr = dct.ParseError(errors.New(dct.Fatality.Anchor().Error()), dct)
	if ierr == nil || !dct.Fatality.IsEqual(ierr) {
		t.Fatalf(tplParseFound, ierr, dct.Fatality)
	}
	// Поиск ошибки, отсутствующей в справочнике. Парсер обязан безопасно пропустить
	// не инициализированные объекты ошибок, объекты ошибок без якоря и поля не являющиеся
	// ошибками, затем создать новую ошибку на базе переданной.
	ierr = dct.ParseError(errors.New(unknown), dct)
	if ierr == nil {
		t.Fatal(tplParseNil)
	}
	if !strings.Contains(ierr.String(), unknown) {
		t.Fatalf(tplParseMsg, ierr.String(), unknown)
	}
}
