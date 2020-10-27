package main

import (
	"math"
	"testing"
)

func Test_getEntropy(t *testing.T) {
	tests := []struct {
		name string
		pass string
		want float64
	}{
		{name: "toto", pass: "toto", want: 18.802},
		{name: "password123", pass: "password123", want: 56.869},
		{name: "xkcd", pass: "correcthorsebatterystable", want: 89.308},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math.Round(getEntropy(tt.pass)*1000) / 1000; got != tt.want {
				t.Errorf("getEntropy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCrackDuration(t *testing.T) {
	tests := []struct {
		name string
		e    float64
		g    int64
		want string
	}{
		{name: "seconds", e: 32.903, g: 1000000000, want: "4s"},
		{name: "minutes", e: 37.604, g: 1000000000, want: "2m 17s"},
		{name: "hours", e: 42.304, g: 1000000000, want: "1h 13m 18s"},
		{name: "days", e: 51.699, g: 1000000000, want: "26d 1h 29m 59s"},
		{name: "years", e: 62.039, g: 1000000000, want: "146y 85d 23h 53m 38s"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCrackDuration(tt.e, tt.g); got != tt.want {
				t.Errorf("getCrackDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSHA1Sum(t *testing.T) {
	tests := []struct {
		name string
		p    string
		want string
	}{
		{name: "toto", p: "toto", want: "0b9c2625dc21ef05f6ad4ddf47c5f203837aa32c"},
		{name: "password123", p: "password123", want: "cbfdac6008f9cab4083784cbd1874f76618d2a97"},
		{name: "xkcd", p: "correcthorsebatterystable", want: "ca16ce18e0877508a32cd678de36c44982930e29"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSHA1Sum(tt.p); got != tt.want {
				t.Errorf("getSHA1Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
