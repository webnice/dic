package dic

import (
	"testing"
)

func TestTFile_MimeByFilename(t *testing.T) {
	const (
		tplYes = "функция MimeByFilename(%q), вернулся объект: %v, ожидался объект: %v"
		tplNot = "функция MimeByFilename(%q), вернулся объект: %v, ожидался объект не эквивалентный объекту: %v"
		tplNil = "функция MimeByFilename(%q), вернулся nil, ожидался не nil объект"
	)
	var (
		feo   IFile
		tpl   string
		tests []struct {
			Name  string
			Mime  IMime
			Equal bool
		}
		mime IMime
		n    int
	)

	// Таблица тестовых данных.
	tests = []struct {
		Name  string
		Mime  IMime
		Equal bool
	}{
		{Name: "example.txt", Mime: Mime().TextPlain, Equal: true},
		{Name: "example.html", Mime: Mime().TextHtml, Equal: true},
		{Name: "example.png", Mime: Mime().ImagePng, Equal: true},
		{Name: "example.py", Mime: Mime().TextXPython, Equal: true},
		{Name: "example.md", Mime: Mime().ApplicationXGenesisRom, Equal: false},
		{Name: "example.md", Mime: Mime().TextMarkdown, Equal: true},
		{Name: "example.jsx", Mime: Mime().TextJsx, Equal: true},
		{Name: "example.mk", Mime: Mime().TextXMakefile, Equal: true},
	}
	feo = File()
	for n = range tests {
		if mime = feo.MimeByFilename(tests[n].Name); mime == nil {
			t.Errorf(tplNil, tests[n].Name)
			continue
		}
		switch tests[n].Equal {
		case true:
			tpl = tplYes
		default:
			tpl = tplNot
		}
		if tests[n].Mime.IsEqual(mime) != tests[n].Equal {
			t.Errorf(tpl, tests[n].Name, mime.Mime(), tests[n].Mime.Mime())
		}
	}
}

func TestTFile_MimeByExtension(t *testing.T) {
	const tplNotNil = "функция MimeByExtension(%q), вернулся объект, ожидался nil"
	var mime IMime

	mime = File().MimeByExtension(".My087LjDpZYum")
	if mime != nil {
		t.Errorf(tplNotNil, ".My087LjDpZYum")
	}
}

func TestTFile_MimeByContent(t *testing.T) {
	const tplContent = "функция MimeByContent(), вернулся объект: %v, ожидался объект: %v"
	const tplNotNil = "функция MimeByContent(), вернулся объект, ожидался nil"
	var (
		feo   IFile
		tests []struct {
			content []byte
			mime    IMime
			eq      bool
		}
		mime IMime
		n    int
	)

	// Таблица тестовых данных.
	tests = []struct {
		content []byte
		mime    IMime
		eq      bool
	}{
		{[]byte{}, nil, true},
		{[]byte{}, Mime().ApplicationOctetStream, false},
		{[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8}, Mime().ApplicationOctetStream, true},
		{[]byte(" "), Mime().TextPlain, true},
		{[]byte("GIF89a"), Mime().ImageGif, true},
		{[]byte("PK\x03\x04"), Mime().ApplicationZip, true},
		{[]byte{0x00, 0x61, 0x73, 0x6D}, Mime().ApplicationWasm, true},
		{[]byte{0x00, 0x01, 0x00, 0x00}, Mime().FontTtf, true},
		{[]byte{0x1A, 0x45, 0xDF, 0xA3}, Mime().VideoWebm, true},
		{[]byte("Rar!\x1A\x07\x01\x00"), Mime().ApplicationXRarCompressed, true},
	}
	feo = File()
	for n = range tests {
		mime = feo.MimeByContent(tests[n].content)
		if tests[n].mime == nil && tests[n].eq && mime != nil {
			t.Errorf(tplNotNil)
			continue
		}
		if tests[n].mime == nil {
			continue
		}
		if tests[n].mime.IsEqual(mime) != tests[n].eq {
			t.Errorf(tplContent, mime.Mime(), tests[n].mime.Mime())
		}
	}
}
