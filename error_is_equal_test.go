package dic

import (
	"errors"
	"strings"
	"testing"
)

func TestTError_IsEqual(t *testing.T) {
	const errTpl = "ошибка без аргументов"
	var (
		errs *TestCustomError
		ier1 IError
		ier2 IError
		tmpT bool
		err  error
		err1 error
		err2 error
	)

	errs = makeTestCustomErrors()
	// Тест не корректного вызова.
	if tmpT = errs.Fatality.IsEqual(ier1); tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, false)
		return
	}
	// Ошибка совпадающая по якорю.
	err = errs.Fatality.Bind()
	if tmpT = errs.Fatality.IsEqual(ParseError(err)); !tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, true)
		return
	}
	// Ошибка не совпадающая по якорю, но полностью совпадающая по контенту.
	err = errors.New(errTpl)
	// Проверка отсутствия совпадения по якорю.
	if tmpT = errors.Is(errs.Fatality.Anchor(), err); tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, false)
		return
	}
	// Проверка полного по контенту.
	if tmpT = errs.Fatality.IsEqual(ParseError(err)); !tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, true)
		return
	}
	// Сравнение шаблонизированной ошибки TwoArgs.
	err1 = errs.Fatality.Bind()
	err = errs.TwoArgs.Bind("вложена ошибка Fatality", err1)
	ier1 = errs.ParseError(err)
	err2 = errs.ExitCodeNotNull.Bind()
	err = errs.TwoArgs.Bind("вложена ошибка ExitCodeNotNull", err2)
	ier2 = errs.ParseError(err)
	// Ошибки должны совпадать.
	if tmpT = ier1.IsEqual(ier2); !tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, true)
		return
	}
	// Но текст ошибок разный.
	if tmpT = strings.EqualFold(ier1.String(), ier2.String()); tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, false)
		return
	}
	// Ошибки так же не должны совпадать с вложенными ошибками.
	if tmpT = ier1.IsEqual(ParseError(err1)); tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, false)
		return
	}
	if tmpT = ier2.IsEqual(ParseError(err2)); tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, false)
		return
	}
	// Ошибка созданная на основе якоря ошибки.
	err = errs.Fatality.Anchor()
	ier1 = errs.ParseError(err, errs)
	if tmpT = ier1.IsEqual(errs.ParseError(err)); !tmpT {
		t.Fatalf("IsEqual(), получено %t, ошидалось %t", tmpT, true)
		return
	}
}
