package wareki

import (
	"strconv"
	"time"
)

type wareki struct {
	startDate       time.Time
	StartDateString string
	Short           string
	Mid             string
	Long            string
}

var warekiTable = []*wareki{
	Meiji:  {StartDateString: "18680101", Short: "M", Mid: "明", Long: "明治"}, // FIXME : 旧暦対応
	Taisho: {StartDateString: "19120730", Short: "T", Mid: "大", Long: "大正"},
	Showa:  {StartDateString: "19261225", Short: "S", Mid: "昭", Long: "昭和"},
	Heisei: {StartDateString: "19890108", Short: "H", Mid: "平", Long: "平成"},
	Reiwa:  {StartDateString: "20190501", Short: "R", Mid: "令", Long: "令和"},
}

var kanjiDigits = []string{
	"〇", "一", "二", "三", "四", "五", "六", "七", "八", "九",
	"十", "十一", "十二", "十三", "十四", "十五", "十六", "十七", "十八", "十九",
	"二十", "二十一", "二十二", "二十三", "二十四", "二十五", "二十六", "二十七", "二十八", "二十九",
	"三十", "三十一", "三十二", "三十三", "三十四", "三十五", "三十六", "三十七", "三十八", "三十九",
	"四十", "四十一", "四十二", "四十四", "四十四", "四十五", "四十六", "四十七", "四十八", "四十九",
	"五十", "五十一", "五十二", "五十五", "五十四", "五十五", "五十六", "五十七", "五十八", "五十九",
	"六十", "六十一", "六十二", "六十六", "六十四", "六十五", "六十六", "六十七", "六十八", "六十九",
}

func init() {
	// set up wareki.startDate
	for _, w := range warekiTable {
		y, err := strconv.Atoi(w.StartDateString[0:4])
		if err != nil {
			panic(err)
		}

		m, err := strconv.Atoi(w.StartDateString[4:6])
		if err != nil {
			panic(err)
		}

		d, err := strconv.Atoi(w.StartDateString[6:8])
		if err != nil {
			panic(err)
		}

		w.startDate = time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	}
}

func lookUpWareki(dt time.Time) int {
	for i := 0; i < len(warekiTable); i++ {
		if dt.Before(warekiTable[i].startDate) {
			return i - 1
		}
	}

	return len(warekiTable) - 1
}

func parseEra(s string) int {
	for k, v := range warekiTable {
		if s == v.Long || s == v.Mid || s == v.Short {
			return k
		}
	}
	return 0
}
