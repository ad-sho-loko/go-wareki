package wareki

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func _d(y, m, d int) time.Time {
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
}

func TestNextStdChunk(t *testing.T) {
	_, std := parseLayout(JISX0301Short)
	assert.Equal(t, stdJapaneseEraShort, std)

	_, std = parseLayout(JISX0301Mid)
	assert.Equal(t, stdJapaneseEraMid, std)

	_, std = parseLayout(JISX0301Long)
	assert.Equal(t, stdJapaneseEraLong, std)
}

func TestParseDate(t *testing.T) {
	y, m, d := parseDate("01.02.03")
	assert.Equal(t, 1, y)
	assert.Equal(t, 2, m)
	assert.Equal(t, 3, d)

	y, m, d = parseDate("10.11.12")
	assert.Equal(t, 10, y)
	assert.Equal(t, 11, m)
	assert.Equal(t, 12, d)
}

func TestParseEra(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"M", Meiji}, {"明", Meiji}, {"明治", Meiji},
		{"T", Taisho}, {"大", Taisho}, {"大正", Taisho},
		{"S", Showa}, {"昭", Showa}, {"昭和", Showa},
		{"H", Heisei}, {"平", Heisei}, {"平成", Heisei},
		{"R", Reiwa}, {"令", Reiwa}, {"令和", Reiwa},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.want, parseEra(tt.input))
	}
}

func TestMeiji(t *testing.T) {
	dt, err := convertMeiji(6, 1, 1)
	if err != nil {
		log.Fatal(err)
	}

	y, m, d := dt.Year(), int(dt.Month()), dt.Day()
	assert.Equal(t, 1873, y)
	assert.Equal(t, 1, m)
	assert.Equal(t, 1, d)
}

func TestParseShort(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"M06.01.01", _d(1873, 1, 1)},
		{"M45.07.29", _d(1912, 7, 29)},
		{"T01.07.30", _d(1912, 7, 30)},
		{"T15.12.24", _d(1926, 12, 24)},
		{"S01.12.25", _d(1926, 12, 25)},
		{"S64.01.07", _d(1989, 1, 7)},
		{"H01.01.08", _d(1989, 1, 8)},
		{"H31.04.30", _d(2019, 4, 30)},
		{"R01.05.01", _d(2019, 5, 1)},
		{"R01.12.24", _d(2019, 12, 24)},
	}

	for _, tt := range tests {
		got, err := Parse(JISX0301Short, tt.input)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, tt.want.Year(), got.Year())
		assert.Equal(t, tt.want.Month(), got.Month())
		assert.Equal(t, tt.want.Day(), got.Day())
	}
}

func TestParseMid(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"明06.01.01", _d(1873, 1, 1)},
		{"大01.07.30", _d(1912, 7, 30)},
		{"昭01.12.25", _d(1926, 12, 25)},
		{"平01.01.08", _d(1989, 1, 8)},
		{"令01.05.01", _d(2019, 5, 1)},
	}

	for _, tt := range tests {
		got, err := Parse(JISX0301Mid, tt.input)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, tt.want.Year(), got.Year())
		assert.Equal(t, tt.want.Month(), got.Month())
		assert.Equal(t, tt.want.Day(), got.Day())
	}
}

func TestParseLong(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"明治06.01.01", _d(1873, 1, 1)},
		{"大正01.07.30", _d(1912, 7, 30)},
		{"昭和01.12.25", _d(1926, 12, 25)},
		{"平成01.01.08", _d(1989, 1, 8)},
		{"令和01.05.01", _d(2019, 5, 1)},
	}

	for _, tt := range tests {
		got, err := Parse(JISX0301Long, tt.input)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, tt.want.Year(), got.Year())
		assert.Equal(t, tt.want.Month(), got.Month())
		assert.Equal(t, tt.want.Day(), got.Day())
	}
}

func TestParseLongKanji(t *testing.T) {
	tests := []struct {
		input string
		want  time.Time
	}{
		{"明治06年01月01日", _d(1873, 1, 1)},
		{"大正01年07月30日", _d(1912, 7, 30)},
		{"昭和01年12月25日", _d(1926, 12, 25)},
		{"平成01年01月08日", _d(1989, 1, 8)},
		{"令和01年05月01日", _d(2019, 5, 1)},
	}

	for _, tt := range tests {
		got, err := Parse(JISX0301LongKanji, tt.input)
		if err != nil {
			log.Fatal(err)
		}

		assert.Equal(t, tt.want.Year(), got.Year())
		assert.Equal(t, tt.want.Month(), got.Month())
		assert.Equal(t, tt.want.Day(), got.Day())
	}
}


func TestFormatShort(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{_d(1873, 1, 1), "M06.01.01"},
		{_d(1912, 7, 29), "M45.07.29"},
		{_d(1912, 7, 30), "T01.07.30"},
		{_d(1926, 12, 24), "T15.12.24"},
		{_d(1926, 12, 25), "S01.12.25"},
		{_d(1989, 1, 7), "S64.01.07"},
		{_d(1989, 1, 8), "H01.01.08"},
		{_d(2019, 4, 30), "H31.04.30"},
		{_d(2019, 5, 1), "R01.05.01"},
		{_d(2019, 12, 24), "R01.12.24"},
	}

	for _, tt := range tests {
		got := Format(tt.input, JISX0301Short)
		assert.Equal(t, tt.want, got)
	}
}

func TestFormatMid(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{_d(1873, 1, 1), "明06.01.01"},
		{_d(1912, 7, 30), "大01.07.30"},
		{_d(1926, 12, 25), "昭01.12.25"},
		{_d(1989, 1, 8), "平01.01.08"},
		{_d(2019, 5, 1), "令01.05.01"},
	}

	for _, tt := range tests {
		got := Format(tt.input, JISX0301Mid)
		assert.Equal(t, tt.want, got)
	}
}

func TestFormatLong(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{_d(1873, 1, 1), "明治06.01.01"},
		{_d(1912, 7, 30), "大正01.07.30"},
		{_d(1926, 12, 25), "昭和01.12.25"},
		{_d(1989, 1, 8), "平成01.01.08"},
		{_d(2019, 5, 1), "令和01.05.01"},
	}

	for _, tt := range tests {
		got := Format(tt.input, JISX0301Long)
		assert.Equal(t, tt.want, got)
	}
}

func TestFormatLongKanji(t *testing.T) {
	tests := []struct {
		input time.Time
		want  string
	}{
		{_d(1873, 1, 1), "明治06年01月01日"},
		{_d(1912, 7, 30), "大正01年07月30日"},
		{_d(1926, 12, 25), "昭和01年12月25日"},
		{_d(1989, 1, 8), "平成01年01月08日"},
		{_d(2019, 5, 1), "令和01年05月01日"},
	}

	for _, tt := range tests {
		got := Format(tt.input, JISX0301LongKanji)
		assert.Equal(t, tt.want, got)
	}
}
