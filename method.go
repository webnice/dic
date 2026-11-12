package dic

import "strings"

// Method Справочник HTTP методов запросов.
func Method() Methods { return singletonMethod }

// IsEqual Истина, если HTTP методы эквивалентны.
func (mdo *tMethod) IsEqual(m IMethod) (ret bool) {
	var tmo *tMethod

	// Проверка типа объекта, сравнение возможно только с объектом справочника.
	if tmo, ret = m.(*tMethod); !ret {
		return
	}
	ret = strings.EqualFold(mdo.name, tmo.name)

	return
}

// IsEqualFull Истина, если HTTP методы эквивалентны, сравнивается и метод и битовая маска.
func (mdo *tMethod) IsEqualFull(m IMethod) (ret bool) {
	var tmo *tMethod

	// Проверка типа объекта, сравнение возможно только с объектом справочника.
	if tmo, ret = m.(*tMethod); !ret {
		return
	}
	ret = mdo.IsEqual(m) && mdo.bits == tmo.bits

	return
}

// String Интерфейс Stringify.
func (mdo *tMethod) String() string { return mdo.name }
