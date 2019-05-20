package main

import (
	"testing"
	"time"
)

func Test_isAllowed(t *testing.T) {
	type args struct {
		t    time.Time
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test monday 6am",
			args: args{
				t:    time.Date(2019, 4, 1, 6, 0, 0, 0, time.UTC),
				from: time.Date(2019, 4, 1, 8, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 4, 1, 19, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "test monday 9am",
			args: args{
				t:    time.Date(2019, 4, 1, 9, 0, 0, 0, time.UTC),
				from: time.Date(2019, 4, 1, 8, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 4, 1, 19, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "test friday 23:59pm",
			args: args{
				t:    time.Date(2019, 4, 5, 23, 59, 0, 0, time.UTC),
				from: time.Date(2019, 4, 5, 8, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 4, 5, 19, 0, 0, 0, time.UTC),
			},
			want: false,
		},
		{
			name: "test saturday 00:01am",
			args: args{
				t:    time.Date(2019, 4, 6, 0, 1, 0, 0, time.UTC),
				from: time.Date(2019, 4, 6, 8, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 4, 6, 19, 0, 0, 0, time.UTC),
			},
			want: true, // weekend already
		},
		{
			name: "test weekend 6am",
			args: args{
				t:    time.Date(2019, 4, 6, 6, 0, 0, 0, time.UTC),
				from: time.Date(2019, 4, 6, 8, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 4, 6, 19, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "test weekend 8pm",
			args: args{
				t:    time.Date(2019, 4, 6, 20, 0, 0, 0, time.UTC),
				from: time.Date(2019, 4, 6, 8, 0, 0, 0, time.UTC),
				to:   time.Date(2019, 4, 6, 19, 0, 0, 0, time.UTC),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAllowed(tt.args.t, tt.args.from, tt.args.to); got != tt.want {
				t.Errorf("isAllowed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isWeekend(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test monday",
			args: args{
				t: time.Date(2019, 4, 1, 1, 1, 1, 1, time.UTC),
			},
			want: false,
		},
		{
			name: "test tuesday",
			args: args{
				t: time.Date(2019, 4, 2, 1, 1, 1, 1, time.UTC),
			},
			want: false,
		},
		{
			name: "test wednesday",
			args: args{
				t: time.Date(2019, 4, 3, 1, 1, 1, 1, time.UTC),
			},
			want: false,
		},
		{
			name: "test thursday",
			args: args{
				t: time.Date(2019, 4, 4, 1, 1, 1, 1, time.UTC),
			},
			want: false,
		},
		{
			name: "test friday",
			args: args{
				t: time.Date(2019, 4, 5, 1, 1, 1, 1, time.UTC),
			},
			want: false,
		},
		{
			name: "test saturday",
			args: args{
				t: time.Date(2019, 4, 6, 1, 1, 1, 1, time.UTC),
			},
			want: true,
		},
		{
			name: "test sunday",
			args: args{
				t: time.Date(2019, 4, 7, 1, 1, 1, 1, time.UTC),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isWeekend(tt.args.t); got != tt.want {
				t.Errorf("isWeekend() = %v, want %v", got, tt.want)
			}
		})
	}
}
