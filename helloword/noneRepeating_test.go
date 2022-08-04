package main

import "testing"

func TestSubstr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabc", 3},
		{"bbb", 1},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("got %d for input %s; expected %d", actual, tt.s, tt.ans)
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	s := "黑化肥"
	ans := 3

	for i := 0; i < b.N; i++ {
		if actual := lengthOfNonRepeatingSubStr(s); actual != ans {
			b.Errorf("got %d for input %s; expected %d", actual, s, ans)
		}
	}
}
