package dic

import (
	"reflect"
	"testing"
)

func TestMime(t *testing.T) {
	var (
		m  Mimes
		n  int
		rt reflect.Type
		sf reflect.StructField
	)

	m = Mime()
	rt = reflect.TypeOf(m)
	if rt.NumField() <= 0 {
		t.Errorf("справочник MIME типов содержит %d элементов, ожидалось больше %d", rt.NumField(), 0)
		return
	}
	for n = 0; n < rt.NumField(); n++ {
		sf = rt.Field(n)
		if sf.Type.String() != "dic.IMime" {
			t.Errorf(
				"справочник MIME, поле %q, тип поля %q, ожидалось: %q",
				sf.Name, sf.Type.String(), "dic.IMime",
			)
			return
		}
	}
}

func TestTMime_Charset(t *testing.T) {
	const vCharset = "3CQz3vWx3sdGe8GkR9"
	var mime IMime

	if mime = NewMime("", "", map[string]string{}); mime.Charset() != "" {
		t.Errorf("не верный MIME тип: %q, ожидалось: %q", mime, "")
	}
	mime = NewMime("1111", "2222", map[string]string{mimeCharset: vCharset})
	if mime.Charset() != vCharset {
		t.Errorf("функция Charset(), результат: %q, ожидалось: %q", mime.Charset(), vCharset)
	}
}

func TestTMime_Main(t *testing.T) {
	const vMain = "OoS817cY1xzz4xQOZu"
	var mime IMime

	if mime = NewMime("", "", map[string]string{}); mime.Charset() != "" {
		t.Errorf("не верный MIME тип: %q, ожидалось: %q", mime, "")
	}
	mime = NewMime(vMain, "", map[string]string{})
	if mime.Main() != vMain {
		t.Errorf("функция Main(), результат: %q, ожидалось: %q", mime.Main(), vMain)
	}
}

func TestTMime_Sub(t *testing.T) {
	const vSub = "APq3UCGaQTJh27nX5zjH"
	var mime IMime

	if mime = NewMime("", "", map[string]string{}); mime.Charset() != "" {
		t.Errorf("не верный MIME тип: %q, ожидалось: %q", mime, "")
	}
	mime = NewMime("application", vSub, map[string]string{})
	if mime.Sub() != vSub {
		t.Errorf("функция Sub(), результат: %q, ожидалось: %q", mime.Sub(), vSub)
	}
}

func TestTMime_IsEqual(t *testing.T) {
	var (
		n        int
		testCase = []struct {
			m1 IMime
			m2 IMime
			eq bool
		}{
			{m1: NewMime("", "", map[string]string{}), m2: NewMime("", "", map[string]string{}), eq: true},
			{m1: NewMime("audio", "basic", map[string]string{}), m2: NewMime("audio", "", map[string]string{}), eq: false},
			{m1: NewMime("audio", "basic", map[string]string{}), m2: NewMime("audio", "basic", map[string]string{}), eq: true},
			{m1: NewMime("audio", "basic", map[string]string{}), m2: NewMime("audio", "basic", map[string]string{"charset": "utf-8"}), eq: true},
			{m1: NewMime("", "", map[string]string{}), m2: nil, eq: false},
		}
		ok bool
	)

	for n = range testCase {
		if ok = testCase[n].m1.IsEqual(testCase[n].m2); ok != testCase[n].eq {
			t.Errorf(
				"функция IsEqual(), сравнение %q = \"%v\", результат: %t, ожидалось: %t",
				testCase[n].m1, testCase[n].m2, ok, testCase[n].eq,
			)
		}
	}
}

func TestTMime_IsEqualFull(t *testing.T) {
	var (
		n        int
		testCase = []struct {
			m1 IMime
			m2 IMime
			eq bool
		}{
			{m1: NewMime("", "", map[string]string{}), m2: NewMime("", "", map[string]string{}), eq: true},
			{m1: NewMime("", "", map[string]string{}), m2: nil, eq: false},
			{m1: NewMime("audio", "basic", map[string]string{}), m2: NewMime("audio", "basic", map[string]string{"charset": "utf-8"}), eq: false},
		}
		ok bool
	)

	for n = range testCase {
		if ok = testCase[n].m1.IsEqualFull(testCase[n].m2); ok != testCase[n].eq {
			t.Errorf(
				"функция IsEqualFull(), сравнение %q = \"%v\", результат: %t, ожидалось: %t",
				testCase[n].m1, testCase[n].m2, ok, testCase[n].eq,
			)
		}
	}
}
