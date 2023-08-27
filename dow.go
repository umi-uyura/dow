package dow

import (
	"strings"
	"time"

	"github.com/goodsign/monday"
)

func GetName(date time.Time) string {
	return GetLocaleName(date, "")
}

func GetLocaleName(date time.Time, lang string) string {
	var locale = getLocale(lang)
	return format(date, locale, false)
}

func GetNameWithDate(date time.Time) string {
	var locale = getLocale("")
	return format(date, locale, true)
}

func GetLocaleNameWithDate(date time.Time, lang string) string {
	var locale = getLocale(lang)
	return format(date, locale, true)
}

func format(date time.Time,
	locale monday.Locale,
	withDate bool) string {

	if !withDate {
		return monday.Format(date, "Mon", locale)
	} else {
		return monday.Format(date, "2006-1-2(Mon)", locale)
	}
}

func getLocale(lang string) monday.Locale {
	var locale monday.Locale = monday.LocaleEnUS

	if len(lang) > 0 {
		lang1 := strings.Split(lang, ".")
		for _, l := range monday.ListLocales() {
			if lang1[0] == string(l) {
				locale = l
				break
			}
		}
	}

	return locale
}
