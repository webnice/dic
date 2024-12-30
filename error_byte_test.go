package dic

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestTError_Bytes(t *testing.T) {
	const (
		tpl  = "выполнение действия %q прервано ошибкой: %s"
		tplA = "func(W7L1O8BC1BBOL3CL9A89)"
		tplB = "P7V3RRE3F2186U6UMYS8R9NMPWENJ4V28STUJ9XL"
	)
	var (
		ier  IError
		err  error
		buf  *bytes.Buffer
		tmpT bool
	)

	ier = NewError(tpl, "название действия", "описание возникшей ошибки").CodeI64().Set(816)
	err = ier.Bind(tplA, tplB)
	ier = ParseError(err)
	if tmpT = strings.EqualFold(ier.String(), fmt.Sprintf(tpl, tplA, tplB)); !tmpT {
		t.Fatalf("String(), получено %t, ошидалось %t", tmpT, true)
		return
	}
	buf = bytes.NewBufferString(fmt.Sprintf(tpl, tplA, tplB))
	if tmpT = bytes.EqualFold(ier.Bytes(), buf.Bytes()); !tmpT {
		t.Fatalf("Bytes(), получено %t, ошидалось %t", tmpT, true)
		return
	}
}
