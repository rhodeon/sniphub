package main

import (
	"testing"
	"time"
)

func Test_formattedDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2022, 03, 9, 22, 10, 0, 0, time.UTC),
			want: "Mar 09, 2022 at 22:10",
		},
		{
			name: "zero-time",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2022, 03, 9, 22, 10, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "Mar 09, 2022 at 21:10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formattedDate(tt.tm)
			if got != tt.want {
				t.Errorf("\nGot:\t%s\nWant:\t%s", got, tt.want)
			}
		})
	}
}
