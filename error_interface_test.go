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
	// TwoArgs ошибка с двумя аргументами, первый аргумент %q, второй аргумент: %w.
	TwoArgs IError
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
		NotFound:        NewError(errNotFound, "название объекта"),
		Fatality:        NewError(errFatality),
		AnotherNilError: nil,
		ExitCodeNotNull: NewError(errExitCodeNotNull).CodeU8().Set(99),
		TwoArgs:         NewError(errTwoArgs, "аргумент 1", "вложенная ошибка"),
	}
	ret.Errors.Unknown = Error().Unknown

	return
}
