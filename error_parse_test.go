package dic

import (
	"errors"
	"strings"
	"testing"
)

func TestParseError_Nil(t *testing.T) {
	var (
		ierr IError
		tmpT bool
		errs *TestCustomError
	)

	errs = makeTestCustomErrors()
	ierr = errs.ParseError(ierr, errs)
	if tmpT = Error().Unknown.IsEqual(ierr); !tmpT {
		t.Fatalf("IsEqual(), получено %t, ожидалось %t", tmpT, true)
		return
	}
}

func TestNewError(t *testing.T) {
	const errTwoArgs = "ошибка с двумя аргументами, первый аргумент %q, второй аргумент: %s"
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
		t.Fatalf("IsEqual(), получено %t, ожидалось %t", tmpT, true)
		return
	}
}

func TestParseError_CustomDictionary(t *testing.T) {
	const unknown = "совершенно неизвестная ошибка"
	var (
		dct  *TestCustomDictionary
		ierr IError
	)

	dct = makeTestCustomDictionary()
	// Поиск ошибки, якорь которой есть в справочнике.
	ierr = dct.ParseError(errors.New(dct.Fatality.Anchor().Error()), dct)
	if ierr == nil || !dct.Fatality.IsEqual(ierr) {
		t.Fatalf("ParseError(), вернулся объект: %v, ожидался объект эквивалентный: %v", ierr, dct.Fatality)
	}
	// Поиск ошибки, отсутствующей в справочнике. Парсер обязан безопасно пропустить
	// не инициализированные объекты ошибок, объекты ошибок без якоря и поля не являющиеся
	// ошибками, затем создать новую ошибку на базе переданной.
	ierr = dct.ParseError(errors.New(unknown), dct)
	if ierr == nil {
		t.Fatal("ParseError(), вернулся nil, ожидался не nil объект")
	}
	if !strings.Contains(ierr.String(), unknown) {
		t.Fatalf("ParseError(), получено: %q, ожидалось сообщение содержащее: %q", ierr.String(), unknown)
	}
}
