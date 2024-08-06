package dic

import (
	"net/http"
	"reflect"
	"testing"
)

func TestStatus(t *testing.T) {
	var (
		s  Statuses
		n  int
		rt reflect.Type
		sf reflect.StructField
	)

	s = Status()
	rt = reflect.TypeOf(s)
	if rt.NumField() <= 0 {
		t.Errorf("справочник статусов HTTP ответов содержит %d элементов, ожидалось больше %d", rt.NumField(), 0)
	}
	for n = 0; n < rt.NumField(); n++ {
		sf = rt.Field(n)
		if sf.Type.String() != "dic.IStatus" {
			t.Errorf(
				"справочник статусов HTTP ответов, поле %q, тип поля %q, ожидалось: %q",
				sf.Name, sf.Type.String(), "dic.IStatus",
			)
		}
	}
}

func TestTStatus_IsEqual(t *testing.T) {
	var (
		tests = []struct {
			text   string
			status IStatus
			eq     bool
		}{
			{"", Status().Unknown, true},
			{"HjzT89pyKWUfSvLJFsMZ", nil, true},
			{"HjzT89pyKWUfSvLJFsMZ", Status().Unknown, false},
		}
		status IStatus
		n      int
	)

	for n = range tests {
		status = ParseStatusString(tests[n].text)
		if tests[n].status == nil && tests[n].eq && status != tests[n].status {
			t.Errorf("функция ParseStatusString(), вернулся объект: \"%v\", ожидался \"%v\"", status, nil)
			continue
		}
		if tests[n].status == nil {
			continue
		}
		if tests[n].status.IsEqual(status) != tests[n].eq {
			t.Errorf("функция IsEqual(), вернулся объект: \"%v\", ожидался объект: \"%v\"", status, tests[n].status)
		}
	}
}

func TestTStatus_IsEqualCode(t *testing.T) {
	const newStatusText, newStatusCode = "xm9N3uCr0e6Rvs58eH3S", http.StatusOK
	var (
		t1 IStatus
		t2 IStatus
		ok bool
	)

	t1 = NewStatus("", http.StatusAccepted)
	if ok = Status().Accepted.IsEqualCode(http.StatusAccepted); !ok {
		t.Errorf("функция IsEqualCode(), вернулся: %t, ожидался: %t", ok, true)
	}
	if ok = Status().Accepted.IsEqual(t1); !ok {
		t.Errorf("функция IsEqual(), вернулся: %t, ожидался: %t", ok, true)
	}
	t2 = NewStatus(newStatusText, newStatusCode)
	if ok = Status().Accepted.IsEqual(t2); ok {
		t.Errorf("функция IsEqual(), вернулся: %t, ожидался: %t", ok, false)
	}
}

func TestTStatus_Bytes(t *testing.T) {
	var (
		t1 IStatus
	)

	t1 = NewStatus("", http.StatusOK)
	if string(t1.Bytes()) != t1.String() {
		t.Errorf("ошибка в функции Bytes() или String()")
	}
}
