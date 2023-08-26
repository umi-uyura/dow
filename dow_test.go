package dow

import (
	"testing"
	"time"
)

func TestGetName(t *testing.T) {
	tests := []struct {
		testdate string
		expected string
	}{
		{"2023-1-1", "Sun"},
		{"2023-1-2", "Mon"},
		{"2023-1-3", "Tue"},
		{"2023-1-4", "Wed"},
		{"2023-1-5", "Thu"},
		{"2023-1-6", "Fri"},
		{"2023-1-7", "Sat"},
		{"2023-01-08", "Sun"},
		{"2023-01-09", "Mon"},
		{"2023-01-10", "Tue"},
		{"2023-01-11", "Wed"},
		{"2023-01-12", "Thu"},
		{"2023-01-13", "Fri"},
		{"2023-01-14", "Sat"},
	}

	for _, tt := range tests {
		d, err := time.Parse("2006-1-2", tt.testdate)
		if err != nil {
			t.Errorf("Failed parse date: %s", tt.testdate)
			continue
		}

		t.Run(tt.testdate, func(t *testing.T) {
			actual := GetName(d)
			if actual != tt.expected {
				t.Errorf("Invalid day of week name: %s -> %s, expected: %s",
					tt.testdate, actual, tt.expected)
			}
		})
	}
}

func TestGetLocaleNameEnUS(t *testing.T) {
	tests := []struct {
		testdate string
		expected string
	}{
		{"2023-1-1", "Sun"},
		{"2023-1-2", "Mon"},
		{"2023-1-3", "Tue"},
		{"2023-1-4", "Wed"},
		{"2023-1-5", "Thu"},
		{"2023-1-6", "Fri"},
		{"2023-1-7", "Sat"},
		{"2023-01-08", "Sun"},
		{"2023-01-09", "Mon"},
		{"2023-01-10", "Tue"},
		{"2023-01-11", "Wed"},
		{"2023-01-12", "Thu"},
		{"2023-01-13", "Fri"},
		{"2023-01-14", "Sat"},
	}

	for _, tt := range tests {
		d, err := time.Parse("2006-1-2", tt.testdate)
		if err != nil {
			t.Errorf("Failed parse date: %s", tt.testdate)
			continue
		}

		t.Run(tt.testdate, func(t *testing.T) {
			actual := GetLocaleName(d, "en_US")
			if actual != tt.expected {
				t.Errorf("Invalid day of week name (en_US): %s -> %s, expected: %s",
					tt.testdate, actual, tt.expected)
			}
		})
	}
}

func TestGetLocaleNameJaJP(t *testing.T) {
	tests := []struct {
		testdate string
		expected string
	}{
		{"2023-1-1", "日"},
		{"2023-1-2", "月"},
		{"2023-1-3", "火"},
		{"2023-1-4", "水"},
		{"2023-1-5", "木"},
		{"2023-1-6", "金"},
		{"2023-1-7", "土"},
		{"2023-01-08", "日"},
		{"2023-01-09", "月"},
		{"2023-01-10", "火"},
		{"2023-01-11", "水"},
		{"2023-01-12", "木"},
		{"2023-01-13", "金"},
		{"2023-01-14", "土"},
	}

	for _, tt := range tests {
		d, err := time.Parse("2006-1-2", tt.testdate)
		if err != nil {
			t.Errorf("Failed parse date: %s", tt.testdate)
			continue
		}

		t.Run(tt.testdate, func(t *testing.T) {
			actual := GetLocaleName(d, "ja_JP")
			if actual != tt.expected {
				t.Errorf("Invalid day of week name (ja_JP): %s -> %s, expected: %s",
					tt.testdate, actual, tt.expected)
			}
		})
	}
}

func TestGetNameWithDate(t *testing.T) {
	tests := []struct {
		testdate string
		expected string
	}{
		{"2023-1-1", "2023-1-1(Sun)"},
		{"2023-1-2", "2023-1-2(Mon)"},
		{"2023-1-3", "2023-1-3(Tue)"},
		{"2023-1-4", "2023-1-4(Wed)"},
		{"2023-1-5", "2023-1-5(Thu)"},
		{"2023-1-6", "2023-1-6(Fri)"},
		{"2023-1-7", "2023-1-7(Sat)"},
		{"2023-01-08", "2023-1-8(Sun)"},
		{"2023-01-09", "2023-1-9(Mon)"},
		{"2023-01-10", "2023-1-10(Tue)"},
		{"2023-01-11", "2023-1-11(Wed)"},
		{"2023-01-12", "2023-1-12(Thu)"},
		{"2023-01-13", "2023-1-13(Fri)"},
		{"2023-01-14", "2023-1-14(Sat)"},
	}

	for _, tt := range tests {
		d, err := time.Parse("2006-1-2", tt.testdate)
		if err != nil {
			t.Errorf("Failed parse date: %s", tt.testdate)
			continue
		}

		t.Run(tt.testdate, func(t *testing.T) {
			actual := GetNameWithDate(d)
			if actual != tt.expected {
				t.Errorf("Invalid day of week name: %s -> %s, expected: %s",
					tt.testdate, actual, tt.expected)
			}
		})
	}
}

func TestGetLocaleNameWithDateEnUS(t *testing.T) {
	tests := []struct {
		testdate string
		expected string
	}{
		{"2023-1-1", "2023-1-1(Sun)"},
		{"2023-1-2", "2023-1-2(Mon)"},
		{"2023-1-3", "2023-1-3(Tue)"},
		{"2023-1-4", "2023-1-4(Wed)"},
		{"2023-1-5", "2023-1-5(Thu)"},
		{"2023-1-6", "2023-1-6(Fri)"},
		{"2023-1-7", "2023-1-7(Sat)"},
		{"2023-01-08", "2023-1-8(Sun)"},
		{"2023-01-09", "2023-1-9(Mon)"},
		{"2023-01-10", "2023-1-10(Tue)"},
		{"2023-01-11", "2023-1-11(Wed)"},
		{"2023-01-12", "2023-1-12(Thu)"},
		{"2023-01-13", "2023-1-13(Fri)"},
		{"2023-01-14", "2023-1-14(Sat)"},
	}

	for _, tt := range tests {
		d, err := time.Parse("2006-1-2", tt.testdate)
		if err != nil {
			t.Errorf("Failed parse date: %s", tt.testdate)
			continue
		}

		t.Run(tt.testdate, func(t *testing.T) {
			actual := GetLocaleNameWithDate(d, "en_US")
			if actual != tt.expected {
				t.Errorf("Invalid day of week name (en_US): %s -> %s, expected: %s",
					tt.testdate, actual, tt.expected)
			}
		})
	}
}

func TestGetLocaleNameWithDateJaJP(t *testing.T) {
	tests := []struct {
		testdate string
		expected string
	}{
		{"2023-1-1", "2023-1-1(日)"},
		{"2023-1-2", "2023-1-2(月)"},
		{"2023-1-3", "2023-1-3(火)"},
		{"2023-1-4", "2023-1-4(水)"},
		{"2023-1-5", "2023-1-5(木)"},
		{"2023-1-6", "2023-1-6(金)"},
		{"2023-1-7", "2023-1-7(土)"},
		{"2023-01-08", "2023-1-8(日)"},
		{"2023-01-09", "2023-1-9(月)"},
		{"2023-01-10", "2023-1-10(火)"},
		{"2023-01-11", "2023-1-11(水)"},
		{"2023-01-12", "2023-1-12(木)"},
		{"2023-01-13", "2023-1-13(金)"},
		{"2023-01-14", "2023-1-14(土)"},
	}

	for _, tt := range tests {
		d, err := time.Parse("2006-1-2", tt.testdate)
		if err != nil {
			t.Errorf("Failed parse date: %s", tt.testdate)
			continue
		}

		t.Run(tt.testdate, func(t *testing.T) {
			actual := GetLocaleNameWithDate(d, "ja_JP")
			if actual != tt.expected {
				t.Errorf("Invalid day of week name (ja_JP): %s -> %s, expected: %s",
					tt.testdate, actual, tt.expected)
			}
		})
	}
}
