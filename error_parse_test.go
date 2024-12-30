package dic

import "testing"

func TestParseError_Nil(t *testing.T) {
	var (
		ierr IError
		tmpT bool
		errs *TestCustomError
	)

	errs = makeTestCustomErrors()
	ierr = errs.ParseError(ierr, errs)
	if tmpT = Error().Unknown.IsEqual(ierr); !tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, true)
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
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, true)
		return
	}
}
