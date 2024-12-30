package dic

import (
	"errors"
	"math"
	"strconv"
	"testing"
)

func TestTError_CSet_CGet(t *testing.T) {
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
		t.Fatal("объект стандартной ошибки интерфейса error не содержит IError")
		return
	}
	if errI.CodeI().Get() != 99 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeI().Get(), 99)
		return
	}
	if errI.CodeI8().Get() != 99 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeI8().Get(), 99)
		return
	}
	if errI.CodeI16().Get() != 99 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeI16().Get(), 99)
		return
	}
	if errI.CodeI32().Get() != 99 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeI32().Get(), 99)
		return
	}
	if errI.CodeI64().Get() != 99 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeI64().Get(), 99)
		return
	}
	if errI.CodeU().Get() != 99 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeU().Get(), 99)
		return
	}
	if errI.CodeU8().Get() != 99 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeU8().Get(), 99)
		return
	}
	if errI.CodeU16().Get() != 99 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeU16().Get(), 99)
		return
	}
	if errI.CodeU32().Get() != 99 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeU32().Get(), 99)
		return
	}
	if errI.CodeU64().Get() != 99 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeU64().Get(), 99)
		return
	}

	// uint8 Изменение значения в справочнике и создание новой ошибки.
	err1 = errs.ExitCodeNotNull.CodeU8().Set(math.MaxUint8).Bind()
	errors.As(err1, &errI)
	if errI.CodeU8().Get() != math.MaxUint8 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeU8().Get(), math.MaxUint8)
		return
	}
	// int Изменение значения в справочнике и создание новой ошибки.
	err1 = errs.ExitCodeNotNull.CodeI().Set(math.MaxInt).Bind()
	errors.As(err1, &errI)
	if errI.CodeI().Get() != math.MaxInt {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeI().Get(), math.MaxInt)
		return
	}
	// int64 Изменение значения в справочнике и создание новой ошибки.
	err1 = errs.ExitCodeNotNull.CodeI64().Set(math.MaxInt64).Bind()
	errors.As(err1, &errI)
	if errI.CodeI64().Get() != math.MaxInt64 {
		t.Fatalf("не верное значение кода, получено %d, ошидалось %d", errI.CodeI64().Get(), math.MaxInt64)
		return
	}
	// uint64 Изменение значения в справочнике и создание новой ошибки.
	err1 = errs.ExitCodeNotNull.CodeU64().Set(math.MaxUint64).Bind()
	errors.As(err1, &errI)
	if errI.CodeU64().Get() != math.MaxUint64 {
		t.Fatalf(
			"не верное значение кода, получено %d, ошидалось %s",
			errI.CodeU64().Get(),
			strconv.FormatUint(math.MaxUint64, 10),
		)
		return
	}
	// Сравнение ошибок по коду ошибки.
	err1 = errs.Fatality.CodeI64().Set(math.MaxInt64 - 1).Bind()
	err2 = errs.ExitCodeNotNull.CodeI64().Set(math.MaxInt64 - 1).Bind()
	errI = errs.Unbind(err1)
	if tmpT = errs.Unbind(err2).IsEqualCode(errI); !tmpT {
		t.Fatalf("IsEqualCode(), получено %t, ошидалось %t", tmpT, true)
		return
	}
}
