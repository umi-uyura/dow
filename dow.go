// Dow package provide functions to get the name of the day of the week.
//
// The locales supported by this package depend on github.com/goodsign/monday.
// https://pkg.go.dev/github.com/goodsign/monday#pkg-constants
package dow

import (
	"strings"
	"time"

	"github.com/goodsign/monday"
)

// Get the name of the day of the week for a given date.
//
// Return the name of the day of the week in en_US locale.
func GetName(date time.Time) string {
	return GetLocaleName(date, "")
}

// Get the locale name of the day of the week for a given date
//
// If the specified locale is not supported,
// returns the name of the day of the week in the en_US locale.
func GetLocaleName(date time.Time, lang string) string {
	var locale = getLocale(lang)
	return format(date, locale, false)
}

// Get the name of the day of the week for a given date with the date.
//
// Return the name of the day of the week in en_US locale.
func GetNameWithDate(date time.Time) string {
	var locale = getLocale("")
	return format(date, locale, true)
}

// Get the locale name of the day of the week for a given date with the date
//
// If the specified locale is not supported,
// returns the name of the day of the week in the en_US locale.
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
