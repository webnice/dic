package dic

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestParseStatusString(t *testing.T) {
	var (
		tmp    string
		status IStatus
	)

	tmp = strings.ToLower(http.StatusText(http.StatusBadRequest))
	status = ParseStatusString(tmp)
	if status == nil {
		t.Fatalf("функция ParseStatusString(), вернулся nil, ожидался не nil объект")
	}
	if reflect.ValueOf(status).Pointer() != reflect.ValueOf(Status().BadRequest).Pointer() {
		t.Fatalf(
			"функция ParseStatusString(), вернулся объект: %v, ожидался объект: %v",
			status, Status().BadRequest,
		)
	}
	if status.Code() != http.StatusBadRequest {
		t.Fatalf("функция ParseStatusString(), код статуса: %d, ожидался: %d", status.Code(), http.StatusBadRequest)
	}
}

func TestParseStatusCodeNil(t *testing.T) {
	if status := ParseStatusCode(-1); status != nil {
		t.Fatalf("функция ParseStatusCode(), вернулся объект: %v, ожидался не nil объект", status)
	}
}

func TestParseStatusCode(t *testing.T) {
	var status IStatus

	status = ParseStatusCode(400)
	if status == nil {
		t.Fatalf("функция ParseStatusCode(), вернулся nil, ожидался не nil объект")
	}
	if reflect.ValueOf(status).Pointer() != reflect.ValueOf(Status().BadRequest).Pointer() {
		t.Fatalf(
			"функция ParseStatusCode(), вернулся объект: %v, ожидался объект: %v",
			status, Status().BadRequest,
		)
	}
	if status.Code() != http.StatusBadRequest {
		t.Fatalf("функция ParseStatusCode(), код статуса: %d, ожидался: %d", status.Code(), http.StatusBadRequest)
	}
}

func TestNewStatus(t *testing.T) {
	const newStatusText, newStatusCode = "xm9N3uCr0e6Rvs58eH3S", 999
	var (
		status IStatus
	)

	status = NewStatus("", http.StatusBadRequest)
	if status == nil {
		t.Fatalf("функция ParseStatusString(), вернулся nil, ожидался не nil объект")
	}
	if reflect.ValueOf(status).Pointer() != reflect.ValueOf(Status().BadRequest).Pointer() {
		t.Fatalf(
			"функция ParseStatusString(), вернулся объект: %v, ожидался объект: %v",
			status, Status().BadRequest,
		)
	}
	if status.Code() != http.StatusBadRequest {
		t.Fatalf("функция ParseStatusString(), код статуса: %d, ожидался: %d", status.Code(), http.StatusBadRequest)
	}
	// Совершенно новый неизвестный HTTP статус.
	status = NewStatus(newStatusText, newStatusCode)
	if status == nil {
		t.Fatalf("функция ParseStatusString(), вернулся nil, ожидался не nil объект")
	}
}
