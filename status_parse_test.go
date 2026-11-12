package dic

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestParseStatusString(t *testing.T) {
	const (
		tplNil         = "функция ParseStatusString(), вернулся nil, ожидался не nil объект"
		tplParseStatus = "функция ParseStatusString(), вернулся объект: %v, ожидался объект: %v"
		tplCode        = "функция ParseStatusString(), код статуса: %d, ожидался: %d"
	)
	var (
		tmp    string
		status IStatus
	)

	tmp = strings.ToLower(http.StatusText(http.StatusBadRequest))
	status = ParseStatusString(tmp)
	if status == nil {
		t.Fatalf(tplNil)
	}
	if reflect.ValueOf(status).Pointer() != reflect.ValueOf(Status().BadRequest).Pointer() {
		t.Fatalf(tplParseStatus, status, Status().BadRequest)
	}
	if status.Code() != http.StatusBadRequest {
		t.Fatalf(tplCode, status.Code(), http.StatusBadRequest)
	}
}

func TestParseStatusCodeNil(t *testing.T) {
	const tplNotNil = "функция ParseStatusCode(), вернулся объект: %v, ожидался nil"
	var status IStatus

	if status = ParseStatusCode(-1); status != nil {
		t.Fatalf(tplNotNil, status)
	}
}

func TestParseStatusCode(t *testing.T) {
	const (
		tplNil         = "функция ParseStatusCode(), вернулся nil, ожидался не nil объект"
		tplParseStatus = "функция ParseStatusCode(), вернулся объект: %v, ожидался объект: %v"
		tplCode        = "функция ParseStatusCode(), код статуса: %d, ожидался: %d"
	)
	var status IStatus

	status = ParseStatusCode(400)
	if status == nil {
		t.Fatalf(tplNil)
	}
	if reflect.ValueOf(status).Pointer() != reflect.ValueOf(Status().BadRequest).Pointer() {
		t.Fatalf(tplParseStatus, status, Status().BadRequest)
	}
	if status.Code() != http.StatusBadRequest {
		t.Fatalf(tplCode, status.Code(), http.StatusBadRequest)
	}
}

func TestNewStatus(t *testing.T) {
	const (
		newStatusText, newStatusCode = "xm9N3uCr0e6Rvs58eH3S", 999
		tplNil                       = "функция ParseStatusString(), вернулся nil, ожидался не nil объект"
		tplParseStatus               = "функция ParseStatusString(), вернулся объект: %v, ожидался объект: %v"
		tplCode                      = "функция ParseStatusString(), код статуса: %d, ожидался: %d"
	)
	var status IStatus

	status = NewStatus("", http.StatusBadRequest)
	if status == nil {
		t.Fatalf(tplNil)
	}
	if reflect.ValueOf(status).Pointer() != reflect.ValueOf(Status().BadRequest).Pointer() {
		t.Fatalf(tplParseStatus, status, Status().BadRequest)
	}
	if status.Code() != http.StatusBadRequest {
		t.Fatalf(tplCode, status.Code(), http.StatusBadRequest)
	}
	// Совершенно новый неизвестный HTTP статус.
	status = NewStatus(newStatusText, newStatusCode)
	if status == nil {
		t.Fatalf(tplNil)
	}
}
