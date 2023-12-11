//go:build ignore
// +build ignore

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

const (
	sourceFile     = "status.txt"
	resultFile     = "status_constants.go"
	resultTemplate = `// НЕ РЕДАКТИРОВАТЬ! Изменения будут перезаписаны при следующей кодогенерации.
// Code generated by go generate; DO NOT EDIT.

package dic

// Statuses Структура справочника статусов HTTP ответов.
type Statuses struct {
{{- range .}}
	// {{.Name}} Статус: "{{.Source}}", код: {{.Code}}. {{.Description}}
	{{.Name}}	IStatus
{{- end}}
}

func init() {
	singletonStatus = Statuses{
{{- range .}}
		{{.Name}}: &tStatus{status: "{{.Source}}", code: {{.Code}}},
{{- end}}
	}
}
`
)

type source struct {
	Name        string
	Source      string
	Code        int
	Description string
}

func main() {
	var (
		err    error
		src    []*source
		result *bytes.Buffer
		tpl    *template.Template
		data   []byte
	)

	if src, err = readSource(sourceFile); err != nil {
		log.Fatalf("чтение файла %q прервано ошибкой: %s", sourceFile, err)
		return
	}
	sort.SliceStable(src, func(i, j int) bool { return src[i].Code < src[j].Code })
	result = &bytes.Buffer{}
	tpl = template.Must(template.New(resultFile).Parse(resultTemplate))
	if err = tpl.Execute(result, src); err != nil {
		log.Fatalf("шаблонизатор завершился с ошибкой: %s", err)
		return
	}
	if data, err = format.Source(result.Bytes()); err != nil {
		log.Fatalf("форматирование созданного .go файла прервано ошибкой: %s", err)
		return
	}
	if err = os.WriteFile(resultFile, data, os.FileMode(0644)); err != nil {
		log.Fatalf("создание файла %q прервано ошибкой: %s", resultFile, err)
		return
	}
}

func readSource(f string) (ret []*source, err error) {
	const sep = "|"
	var (
		data    []byte
		scanner *bufio.Scanner
		line    []string
		code    int64
	)

	if data, err = os.ReadFile(f); err != nil {
		return
	}
	scanner = bufio.NewScanner(bytes.NewBuffer(data))
	for scanner.Scan() {
		if line = strings.Split(scanner.Text(), sep); len(line) != 4 {
			continue
		}
		if code, err = strconv.ParseInt(strings.TrimSpace(line[2]), 10, 64); err != nil {
			err = fmt.Errorf("преобразование строки %q в число, прервано ошибкой: %s", line[2], err)
			return
		}
		ret = append(ret, &source{
			Name:        strings.TrimSpace(line[0]),
			Source:      strings.TrimSpace(line[1]),
			Code:        int(code),
			Description: strings.TrimSpace(line[3]),
		})
	}

	return
}
