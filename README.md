# dic

[![GoDoc](https://godoc.org/github.com/webnice/dic?status.png)](http://godoc.org/github.com/webnice/dic)
[![Coverage Status](https://coveralls.io/repos/github/webnice/dic/badge.svg?branch=master%0Av1)](https://coveralls.io/github/webnice/dic?branch=master%0Av1)

Библиотека справочников для Go.

## Описание

`dic` — это библиотека предопределённых справочников: MIME типов, HTTP методов
запросов, заголовков, статусов HTTP ответов, MIME типов расширений имён файлов
и ошибок. Каждый справочник предоставляет типизированный доступ к значениям,
разбор строковых представлений и сравнение элементов без учёта регистра
символов.

Справочники реализованы таким образом, чтобы их можно было дополнять
пользовательскими данными через встраивание. Получившийся пользовательский
справочник обладает всеми свойствами справочника библиотеки и принимается
теми же функциями, которые используют исходный справочник.

### Особенности

- Справочник MIME типов с опциями, кодовыми страницами и сравнением с учётом
  и без учёта атрибутов;
- Справочник HTTP методов с битовыми масками для компактного хранения наборов
  методов;
- Справочник заголовков HTTP запросов и ответов с описанием и ссылками на RFC;
- Справочник статусов HTTP ответов с кодами, названиями и описанием;
- Определение MIME типа контента файла по имени файла, расширению и содержимому;
- Справочник ошибок с шаблонами, аргументами, числовыми кодами и
  эквивалентностью по якорю;
- Нулевые зависимости, покрытие тестами 100%.

## Требования

- Go 1.18 или новее (используются generics).

## Установка

```bash
go get github.com/webnice/dic
```

Подключение:

```go
import (
	"github.com/webnice/dic"
)
```

### Зависимости

    NONE

## Справочники

| Справочник | Функция доступа | Интерфейс элемента |
|------------|-----------------|--------------------|
| MIME типы | `dic.Mime()` | `dic.IMime` |
| HTTP методы | `dic.Method()` | `dic.IMethod` |
| Заголовки | `dic.Header()` | `dic.IHeader` |
| Статусы HTTP ответов | `dic.Status()` | `dic.IStatus` |
| MIME типы файлов | `dic.File()` | `dic.IFile` |
| Ошибки | `dic.Error()` | `dic.IError` |

## Примеры использования

### MIME типы

```go
m := dic.Mime().ApplicationJson
m.Main()    // "application"
m.Sub()     // "json"
m.Charset() // "utf-8"
m.String()  // "application/json; charset=utf-8"

// Разбор строки и создание нового MIME типа.
m = dic.ParseMime("text/plain; charset=utf-8")
m = dic.NewMime("application", "json", map[string]string{"charset": "utf-8"})
```

### HTTP методы

```go
m := dic.Method().Get
m.IsEqual(dic.ParseMethod("get")) // true
m.IsEqualFull(dic.Method().Head)  // false

// Создание нового метода.
m = dic.NewMethod("FLUSH") // HTTP метод с именем "FLUSH"
```

### Заголовки

```go
h := dic.Header().ContentType
h.IsEqual(dic.ParseHeader("content-type")) // true

// Создание нового заголовка.
h = dic.NewHeader("X-Custom-Header")
```

### Статусы HTTP ответов

```go
s := dic.Status().NotFound
s.Code()                            // 404
s.IsEqualCode(404)                  // true
s.IsEqual(dic.ParseStatusCode(404)) // true

// Создание нового статуса.
s = dic.NewStatus("Custom Status", 799)
```

### MIME типы файлов

```go
f := dic.File()
f.MimeByFilename("example.txt") // text/plain
f.MimeByExtension(".png")       // image/png
f.MimeByContent(content)        // по первым 512 байтам контента
```

### Ошибки

```go
const errNotFound = "объект %q не найден"

// myErrors Пользовательский справочник ошибок.
type myErrors struct {
	dic.Errors
	NotFound dic.IError
}

var Errors = &myErrors{
	Errors:   dic.Error(),
	NotFound: dic.NewError(errNotFound, "название объекта"),
}

err := Errors.NotFound.Bind("файл конфигурации")
errors.Is(err, Errors.NotFound)               // true
errors.Is(err, Errors.NotFound.Bind("другой")) // true, эквивалентность по якорю

// Числовые коды ошибок.
Errors.NotFound.CodeU8().Set(10)
Errors.NotFound.CodeU8().Get() // 10
```

Копии ошибок с разными значениями аргументов эквивалентны при сравнении через
`errors.Is`, так как сравнение выполняется по якорю — константе ошибки
с фиксированным адресом.

### Дополнение справочника пользовательскими данными

```go
// CustomHeader Структура нового справочника с дополнительными пользовательскими заголовками.
type CustomHeader struct {
	dic.Headers // Встраивание всех уже существующих заголовков.

	// MyCustomHeader1 Заголовок my-custom/header-1.
	MyCustomHeader1 dic.IHeader

	// HeaderSecondCustom Заголовок header-second/custom.
	HeaderSecondCustom dic.IHeader
}

// Реализация дополнительных пользовательских заголовков справочника.
var custom = &CustomHeader{
	Headers:            dic.Header(),                          // Подключение всех существующих заголовков.
	MyCustomHeader1:    dic.NewHeader("My-Custom/Header-1"),   // Заголовок my-custom/header-1.
	HeaderSecondCustom: dic.NewHeader("Header-Second/Custom"), // Заголовок header-second/custom.
}

// Использование нового справочника.
func main() {
	var mch = custom.MyCustomHeader1
	var result = custom.AcceptPatch.IsEqual(mch)

	fmt.Printf("Является ли заголовок %q эквивалентным заголовку %q? Ответ: %t\n", mch, custom.AcceptPatch, result)
	hcc := custom.HeaderSecondCustom
	result = hcc.IsEqual(mch)
	fmt.Printf("Является ли заголовок %q эквивалентным заголовку %q? Ответ: %t\n", hcc, mch, result)
}
```

Результат:

```text
Является ли заголовок "My-Custom/Header-1" эквивалентным заголовку "Accept-Patch"? Ответ: false
Является ли заголовок "Header-Second/Custom" эквивалентным заголовку "My-Custom/Header-1"? Ответ: false
```

## Генерация кода

Справочники заголовков, MIME типов, статусов HTTP ответов и расширений файлов
генерируются из исходных данных командой:

```bash
go generate ./...
```

Исходные данные находятся в файлах `header.txt`, `mime.txt`, `status.txt`
и в директории `file/`.

## Тестирование

```bash
# Стандартный прогон с покрытием.
go test -cover ./...

# Прогон через Makefile (агрегированный профиль покрытия).
make test
```

## Лицензия

© WEB NICE, распространяется под [лицензией MIT](LICENSE).
