package dic

import (
	"strings"
	"testing"
)

func TestTError_Bind(t *testing.T) {
	const (
		arg1 = "значение первого аргумента"
		arg2 = "значение второго аргумента"
	)
	var (
		errs *TestCustomError
		err  error
		str  string
	)

	errs = makeTestCustomErrors()
	// Полный набор аргументов.
	err = errs.TwoArgs.Bind(arg1, arg2)
	str = ParseError(err).String()
	if !strings.Contains(str, arg1) || !strings.Contains(str, arg2) {
		t.Fatalf("Bind(), получено: %q, ожидалось что оба аргумента будут подставлены в шаблон ошибки", str)
	}
	// Частичный набор аргументов, второй аргумент подставляется из описания аргументов шаблона.
	err = errs.TwoArgs.Bind(arg1)
	str = ParseError(err).String()
	if !strings.Contains(str, arg1) || !strings.Contains(str, "вложенная ошибка") {
		t.Fatalf("Bind(), получено: %q, ожидалось что аргументы будут подставлены в шаблон ошибки", str)
	}
	// Отсутствие аргументов, оба аргумента подставляются из описания аргументов шаблона.
	err = errs.TwoArgs.Bind()
	str = ParseError(err).String()
	if !strings.Contains(str, "аргумент 1") || !strings.Contains(str, "вложенная ошибка") {
		t.Fatalf("Bind(), получено: %q, ожидалось что аргументы шаблона будут подставлены в шаблон ошибки", str)
	}
}
