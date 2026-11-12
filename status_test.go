package dic

import (
	"net/http"
	"reflect"
	"testing"
)

func TestStatus(t *testing.T) {
	const (
		tplCount    = "справочник статусов HTTP ответов содержит %d элементов, ожидалось больше %d"
		tplField    = "справочник статусов HTTP ответов, поле %q, тип поля %q, ожидалось: %q"
		typeIStatus = "dic.IStatus"
	)
	var (
		s  Statuses
		n  int
		rt reflect.Type
		sf reflect.StructField
	)

	s = Status()
	rt = reflect.TypeOf(s)
	if rt.NumField() <= 0 {
		t.Errorf(tplCount, rt.NumField(), 0)
	}
	for n = 0; n < rt.NumField(); n++ {
		sf = rt.Field(n)
		if sf.Type.String() != typeIStatus {
			t.Errorf(tplField, sf.Name, sf.Type.String(), typeIStatus)
		}
	}
}

func TestTStatus_IsEqual(t *testing.T) {
	const (
		tplParseStatus = "функция ParseStatusString(), вернулся объект: \"%v\", ожидался \"%v\""
		tplIsEqualObj  = "функция IsEqual(), вернулся объект: \"%v\", ожидался объект: \"%v\""
	)
	var (
		tests []struct {
			text   string
			status IStatus
			eq     bool
		}
		status IStatus
		n      int
	)

	// Таблица тестовых данных.
	tests = []struct {
		text   string
		status IStatus
		eq     bool
	}{
		{"", Status().Unknown, true},
		{"HjzT89pyKWUfSvLJFsMZ", nil, true},
		{"HjzT89pyKWUfSvLJFsMZ", Status().Unknown, false},
	}
	for n = range tests {
		status = ParseStatusString(tests[n].text)
		if tests[n].status == nil && tests[n].eq && status != tests[n].status {
			t.Errorf(tplParseStatus, status, nil)
			continue
		}
		if tests[n].status == nil {
			continue
		}
		if tests[n].status.IsEqual(status) != tests[n].eq {
			t.Errorf(tplIsEqualObj, status, tests[n].status)
		}
	}
}

func TestTStatus_IsEqualCode(t *testing.T) {
	const (
		newStatusText, newStatusCode = "xm9N3uCr0e6Rvs58eH3S", http.StatusOK
		tplIsEqualCode               = "функция IsEqualCode(), вернулся: %t, ожидался: %t"
		tplIsEqual                   = "функция IsEqual(), вернулся: %t, ожидался: %t"
	)
	var (
		t1 IStatus
		t2 IStatus
		ok bool
	)

	t1 = NewStatus("", http.StatusAccepted)
	if ok = Status().Accepted.IsEqualCode(http.StatusAccepted); !ok {
		t.Errorf(tplIsEqualCode, ok, true)
	}
	if ok = Status().Accepted.IsEqual(t1); !ok {
		t.Errorf(tplIsEqual, ok, true)
	}
	t2 = NewStatus(newStatusText, newStatusCode)
	if ok = Status().Accepted.IsEqual(t2); ok {
		t.Errorf(tplIsEqual, ok, false)
	}
}

func TestTStatus_Bytes(t *testing.T) {
	const errBytesString = "ошибка в функции Bytes() или String()"
	var t1 IStatus

	t1 = NewStatus("", http.StatusOK)
	if string(t1.Bytes()) != t1.String() {
		t.Errorf(errBytesString)
	}
}
