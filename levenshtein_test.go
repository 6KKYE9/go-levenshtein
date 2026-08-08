package main

import "testing"

func TestDistance(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"kitten", "sitting", 3},
		{"flaw", "lawn", 2},
		{"same", "same", 0},
		{"", "abc", 3},
		{"中文", "中文", 0},
		{"中文", "中文字", 1},
	}
	for _, c := range cases {
		if got := Distance(c.a, c.b); got != c.want {
			t.Errorf("Distance(%q,%q)=%d 想要 %d", c.a, c.b, got, c.want)
		}
	}
}
