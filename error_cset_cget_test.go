package dic

import (
	"errors"
	"math"
	"strconv"
	"testing"
)

func TestTError_CSet_CGet(t *testing.T) {
	const (
		errNoIError    = "объект стандартной ошибки интерфейса error не содержит IError"
		tplCodeValue   = "не верное значение кода, получено %d, ожидалось %d"
		tplCodeValue64 = "не верное значение кода, получено %d, ожидалось %s"
		tplIsEqualCode = "IsEqualCode(), получено %t, ожидалось %t"
	)
	var (
		errs *TestCustomError
		errI IError
		err1 error
		err2 error
		tmpT bool
	)

	errs = makeTestCustomErrors()
	// Значение определённое в тестах по умолчанию.
	err1 = errs.ExitCodeNotNull.Bind()
	if !errors.As(err1, &errI) {
		t.Fatal(errNoIError)
	}
	if errI.CodeI().Get() != 99 {
		t.Fatalf(tplCodeValue, errI.CodeI().Get(), 99)
	}
	if errI.CodeI8().Get() != 99 {
		t.Fatalf(tplCodeValue, errI.CodeI8().Get(), 99)
	}
	if errI.CodeI16().Get() != 99 {
		t.Fatalf(tplCodeValue, errI.CodeI16().Get(), 99)
	}
	if errI.CodeI32().Get() != 99 {
		t.Fatalf(tplCodeValue, errI.CodeI32().Get(), 99)
	}
	if errI.CodeI64().Get() != 99 {
		t.Fatalf(tplCodeValue, errI.CodeI64().Get(), 99)
	}
	if errI.CodeU().Get() != 99 {
		t.Fatalf(tplCodeValue, errI.CodeU().Get(), 99)
	}
	if errI.CodeU8().Get() != 99 {
		t.Fatalf(tplCodeValue, errI.CodeU8().Get(), 99)
	}
	if errI.CodeU16().Get() != 99 {
		t.Fatalf(tplCodeValue, errI.CodeU16().Get(), 99)
	}
	if errI.CodeU32().Get() != 99 {
		t.Fatalf(tplCodeValue, errI.CodeU32().Get(), 99)
	}
	if errI.CodeU64().Get() != 99 {
		t.Fatalf(tplCodeValue, errI.CodeU64().Get(), 99)
	}

	// uint8 Изменение значения в справочнике и создание новой ошибки.
	err1 = errs.ExitCodeNotNull.CodeU8().Set(math.MaxUint8).Bind()
	errors.As(err1, &errI)
	if errI.CodeU8().Get() != math.MaxUint8 {
		t.Fatalf(tplCodeValue, errI.CodeU8().Get(), math.MaxUint8)
	}
	// int Изменение значения в справочнике и создание новой ошибки.
	err1 = errs.ExitCodeNotNull.CodeI().Set(math.MaxInt).Bind()
	errors.As(err1, &errI)
	if errI.CodeI().Get() != math.MaxInt {
		t.Fatalf(tplCodeValue, errI.CodeI().Get(), math.MaxInt)
	}
	// int64 Изменение значения в справочнике и создание новой ошибки.
	err1 = errs.ExitCodeNotNull.CodeI64().Set(math.MaxInt64).Bind()
	errors.As(err1, &errI)
	if errI.CodeI64().Get() != math.MaxInt64 {
		t.Fatalf(tplCodeValue, errI.CodeI64().Get(), math.MaxInt64)
	}
	// uint64 Изменение значения в справочнике и создание новой ошибки.
	err1 = errs.ExitCodeNotNull.CodeU64().Set(math.MaxUint64).Bind()
	errors.As(err1, &errI)
	if errI.CodeU64().Get() != math.MaxUint64 {
		t.Fatalf(tplCodeValue64, errI.CodeU64().Get(), strconv.FormatUint(math.MaxUint64, 10))
	}
	// Сравнение ошибок по коду ошибки.
	err1 = errs.Fatality.CodeI64().Set(math.MaxInt64 - 1).Bind()
	err2 = errs.ExitCodeNotNull.CodeI64().Set(math.MaxInt64 - 1).Bind()
	errI = errs.Unbind(err1)
	if tmpT = errs.Unbind(err2).IsEqualCode(errI); !tmpT {
		t.Fatalf(tplIsEqualCode, tmpT, true)
	}
	// Сравнение по коду ошибки с nil объектом ошибки.
	if tmpT = errs.Fatality.IsEqualCode(nil); tmpT {
		t.Fatalf(tplIsEqualCode, tmpT, false)
	}
}
