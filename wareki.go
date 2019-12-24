package wareki

import (
	"fmt"
	"time"
)

const(
	Version = "0.1"
)

const (
	JISX0301Short = "M01.02.03"
	JISX0301Mid   = "明01.02.03"
	JISX0301Long  = "明治01.02.03"
	JISX0301LongKanji = "明治01年02月03日"
)

const (
	stdJapaneseEraShort = iota + 1
	stdJapaneseEraMid
	stdJapaneseEraLong
)

const (
	Meiji = iota
	Taisho
	Showa
	Heisei
	Reiwa
)

func parseLayout(layout string) (prefix string, std int) {
	if len(layout) < 1 {
		return
	}

	r := []rune(layout)

	if r[0] == 'M' {
		return JISX0301Short, stdJapaneseEraShort
	}

	if r[0] == '明' {
		if r[1] == '治' {
			if r[4] == '年'{
				return JISX0301LongKanji, stdJapaneseEraLong
			}
			return JISX0301Long, stdJapaneseEraLong
		}

		return JISX0301Mid, stdJapaneseEraMid
	}

	return "", 0 // error
}

func parseDate(value string) (int, int, int) {
	if len(value) < 8 {
		// FIXME: Error handling
	}

	runes := []rune(value)
	y := int(runes[0]-'0')*10 + int(runes[1]-'0')
	m := int(runes[3]-'0')*10 + int(runes[4]-'0')
	d := int(runes[6]-'0')*10 + int(runes[7]-'0')
	return y, m, d
}

func convertMeiji(y, m, d int) (time.Time, error) {
	var year, month, day int

	// FIXME: 旧暦対応(1868-1872)
	year = 1867 + y // 1873 + y - 6
	month = m
	day = d

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
}

func convertWareki(era, y, m, d int) (time.Time, error) {
	year := warekiTable[era].startDate.Year() + y - 1
	month := m
	day := d

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local), nil
}

func Parse(layout, value string) (time.Time, error) {
	tmpPrefix, std := parseLayout(layout)
	if std == 0 {
		return time.Time{}, &time.ParseError{Layout: layout, Value: value, LayoutElem: tmpPrefix, ValueElem: value}
	}

	var prefix string
	switch std {
	case stdJapaneseEraShort:
		prefix = string(value[0])
	case stdJapaneseEraMid:
		prefix = string([]rune(value)[0])
	case stdJapaneseEraLong:
		prefix = string([]rune(value)[0:2])
	}

	era := parseEra(prefix)
	prefixLen := len([]rune(prefix))
	y, m, d := parseDate(string([]rune(value)[prefixLen:]))

	switch era {
	case Meiji:
		return convertMeiji(y, m, d)
	default:
		return convertWareki(era, y, m, d)
	}
}

func getSeparator(layout string) (string, string, string) {
	switch layout {
	case JISX0301Short, JISX0301Mid, JISX0301Long:
		return ".", ".", ""
	case JISX0301LongKanji:
		return "年", "月", "日"
	default:
		return "", "", ""
	}
}

func format(era int, dt time.Time, layout string) string {
	var buf string

	tmpLayout, std := parseLayout(layout)
	w := warekiTable[era]

	switch std {
	case stdJapaneseEraShort:
		buf += w.Short
	case stdJapaneseEraMid:
		buf += w.Mid
	case stdJapaneseEraLong:
		buf += w.Long
	}

	ys, ms, ds := getSeparator(tmpLayout)

	y := dt.Year() - w.startDate.Year() + 1
	buf += fmt.Sprintf("%02d", y)
	buf += ys
	buf += fmt.Sprintf("%02d", int(dt.Month()))
	buf += ms
	buf += fmt.Sprintf("%02d", dt.Day())
	buf += ds
	return buf
}

func Format(dt time.Time, layout string) string {
	era := lookUpWareki(dt)
	return format(era, dt, layout)
}
