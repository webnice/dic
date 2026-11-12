package dic

import (
	"fmt"
)

// ExampleHeader Пример создания нового справочника заголовков с дополнительны пользовательскими заголовками.
func ExampleHeader() {
	const tplQuestion = "Является ли заголовок %q эквивалентным заголовку %q? Ответ: %t\n"
	// CustomHeader Структура нового справочника с дополнительными пользовательскими заголовками.
	type CustomHeader struct {
		Headers // Встраивание всех уже существующих заголовков.

		// MyCustomHeader1 Заголовок my-custom/header-1.
		MyCustomHeader1 IHeader

		// HeaderSecondCustom Заголовок header-second/custom.
		HeaderSecondCustom IHeader
	}
	var (
		custom *CustomHeader
		mch    IHeader
		hcc    IHeader
		result bool
	)

	// Реализация дополнительных пользовательских заголовков справочника.
	custom = &CustomHeader{
		Headers:            Header(),                          // Подключение всех существующих заголовков.
		MyCustomHeader1:    NewHeader("My-Custom/Header-1"),   // Заголовок my-custom/header-1.
		HeaderSecondCustom: NewHeader("Header-Second/Custom"), // Заголовок header-second/custom.
	}
	// Использование нового справочника.
	mch = custom.MyCustomHeader1
	result = custom.AcceptPatch.IsEqual(mch)
	fmt.Printf(tplQuestion, mch, custom.AcceptPatch, result)
	hcc = custom.HeaderSecondCustom
	result = hcc.IsEqual(mch)
	fmt.Printf(tplQuestion, hcc, mch, result)

	//Output: Является ли заголовок "My-Custom/Header-1" эквивалентным заголовку "Accept-Patch"? Ответ: false
	//Является ли заголовок "Header-Second/Custom" эквивалентным заголовку "My-Custom/Header-1"? Ответ: false
}
