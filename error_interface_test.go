package dic

// TestCustomError Структура объекта справочника тестовых ошибок.
type TestCustomError struct {
	// Встраивание предопределённых ошибок.
	Errors
	// NotFound объект %q не найден.
	NotFound IError
	// Fatality ошибка без аргументов.
	Fatality IError
	// AnotherNilError Объект ошибки не инициализирован и оставлен nil.
	AnotherNilError IError
	// ExitCodeNotNull ошибка с указанием кода.
	ExitCodeNotNull IError
	// TwoArgs ошибка с двумя аргументами, первый аргумент %q, второй аргумент: %s.
	TwoArgs IError
}

// TestCustomDictionary Структура справочника ошибок с нестандартными полями, используется для тестирования парсера ошибок.
type TestCustomDictionary struct {
	// Встраивание предопределённых ошибок.
	Errors
	// Fatality обычная ошибка справочника.
	Fatality IError
	// NilError объект ошибки не инициализирован и оставлен nil.
	NilError IError
	// Anchorless объект ошибки без якоря.
	Anchorless IError
	// NotAnError строковое поле, не являющееся ошибкой.
	NotAnError string
}

// Создание объекта тестового справочника ошибок с нестандартными полями.
func makeTestCustomDictionary() (ret *TestCustomDictionary) {
	ret = &TestCustomDictionary{
		Errors:     Error(),
		Fatality:   NewError("ошибка без аргументов"),
		NilError:   nil,
		Anchorless: &tError{},
		NotAnError: "не ошибка",
	}

	return
}

// Создание объекта тестового справочника ошибок.
func makeTestCustomErrors() (ret *TestCustomError) {
	const (
		errNotFound        = "объект %q не найден"
		errFatality        = "ошибка без аргументов"
		errExitCodeNotNull = "ошибка с указанием кода"
		errTwoArgs         = "ошибка с двумя аргументами, первый аргумент %q, второй аргумент: %s"
	)

	ret = &TestCustomError{
		Errors:          Error(),
		NotFound:        NewError(errNotFound, "название объекта"),
		Fatality:        NewError(errFatality),
		AnotherNilError: nil,
		ExitCodeNotNull: NewError(errExitCodeNotNull).CodeU8().Set(99),
		TwoArgs:         NewError(errTwoArgs, "аргумент 1", "вложенная ошибка"),
	}

	return
}
