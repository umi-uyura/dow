package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/goodsign/monday"
)

func main() {
	date := time.Now()

	if len(os.Args) > 1 {
		datestring := os.Args[1]
		datestring = strings.ReplaceAll(datestring, "/", "-")
		datestring = strings.ReplaceAll(datestring, ".", "-")
		
		d, err := time.Parse("2006-1-2", datestring)
		if err == nil {
			date = d
		} else {
			fmt.Println("Invalid date format.")
			fmt.Println("Valid date formats: 'yyyy-mm-dd', 'yyyy/mm/dd', 'yyyy.mm.dd'.")
			os.Exit(1)
		}
	}

	var locale monday.Locale = monday.LocaleEnUS
	lang := os.Getenv("LANG")

	if len(lang) > 0 {
		lang1 := strings.Split(lang, ".")
		for _, l := range monday.ListLocales() {
			if lang1[0] == string(l) {
				locale = l
				break
			}
		}
	}

	fmt.Println(monday.Format(date, "Mon", locale))
}
