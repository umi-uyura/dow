package dow_test

import (
	"fmt"
	"time"

	"github.com/umi-uyura/dow"
)

func ExampleGetName() {
	d, _ := time.Parse("2006-1-2", "2023-1-1")
	weekday := dow.GetName(d)
	fmt.Println(weekday)
	// Output: Sun
}

func ExampleGetLocaleName() {
	d, _ := time.Parse("2006-1-2", "2023-1-1")
	weekday := dow.GetLocaleName(d, "ja_JP")
	fmt.Println(weekday)
	// Output: 日
}

func ExampleGetNameWithDate() {
	d, _ := time.Parse("2006-1-2", "2023-1-1")
	weekday := dow.GetNameWithDate(d)
	fmt.Println(weekday)
	// Output: 2023-1-1(Sun)
}

func ExampleGetLocaleNameWithDate() {
	d, _ := time.Parse("2006-1-2", "2023-1-1")
	weekday := dow.GetLocaleNameWithDate(d, "ja_JP")
	fmt.Println(weekday)
	// Output: 2023-1-1(日)
}
