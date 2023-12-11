package dic

import (
	"strings"
	"testing"
)

func TestParseMime(t *testing.T) {
	var (
		n        int
		testCase = []struct {
			src string
			mim IMime
		}{
			{src: "", mim: &tMime{main: "", subt: "",
				opts: map[string]string{}}},
			{src: "auto-replied", mim: &tMime{main: "auto-replied", subt: "",
				opts: map[string]string{}}},
			{src: "message/delivery-method", mim: &tMime{main: "message", subt: "delivery-method",
				opts: map[string]string{}}},
			{src: "message/imdn+xml;charset=UTF-8", mim: &tMime{main: "message", subt: "imdn+xml",
				opts: map[string]string{"charset": "utf-8"}}},
			{src: "application/json;encoding=base64;charset=UTF-8", mim: &tMime{main: "application", subt: "json",
				opts: map[string]string{"charset": "utf-8", "encoding": "base64"}}},
		}
		obj IMime
	)

	for n = range testCase {
		obj = ParseMime(testCase[n].src)
		if !testCase[n].mim.IsEqualFull(obj) {
			t.Errorf(
				"разбор %q выполнен не верно, получено: %q, ожидалось: %q",
				testCase[n].src, obj, testCase[n].mim,
			)
		}
	}
}

func TestNewMime(t *testing.T) {
	var (
		n, j     int
		opt      map[string]string
		obj      IMime
		testCase = []struct {
			m    string
			s    string
			o    string
			mime IMime
		}{
			{m: "", s: "", o: "",
				mime: &tMime{main: "", subt: "", opts: map[string]string{}}},
			{m: "video", s: "", o: "",
				mime: &tMime{main: "video", subt: "", opts: map[string]string{}}},
			{m: "application", s: "json", o: "charset=utf-8",
				mime: &tMime{main: "application", subt: "json", opts: map[string]string{"charset": "utf-8"}}},
			{m: "application", s: "json", o: "charset=utf-8;=base64",
				mime: &tMime{main: "application", subt: "json", opts: map[string]string{"charset": "utf-8"}}},
			{m: "application", s: "json", o: "charset=utf-8;encoding=",
				mime: &tMime{main: "application", subt: "json", opts: map[string]string{"charset": "utf-8"}}},
		}
	)

	for n = range testCase {
		opt = make(map[string]string)
		tmp := strings.Split(testCase[n].o, ";")
		for j = range tmp {
			if kv := strings.Split(tmp[j], "="); len(kv) == 2 {
				opt[kv[0]] = kv[1]
			}
		}
		if obj = NewMime(testCase[n].m, testCase[n].s, opt); !testCase[n].mime.IsEqualFull(obj) {
			t.Errorf("MIME тип создан с ошибкой, результат: %q, ожидалось: %q", obj, testCase[n].mime)
			return
		}
	}
}
