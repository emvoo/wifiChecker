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
				t:    time.Date(2019, 4, 1, 1, 1, 1, 1, time.UTC),
				from: time.Date(2019, 4, 1, 1, 1, 1, 1, time.UTC),
				to:   time.Date(2019, 4, 1, 1, 1, 1, 1, time.UTC),
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

// func Test_isConnected(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		want bool
// 	}{
// 		{},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := isConnected(); got != tt.want {
// 				t.Errorf("isConnected() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_commandRunner(t *testing.T) {
// 	type args struct {
// 		command string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := commandRunner(tt.args.command); got != tt.want {
// 				t.Errorf("commandRunner() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
//
// func Test_toCommand(t *testing.T) {
// 	type args struct {
// 		input string
// 	}
// 	tests := []struct {
// 		name  string
// 		args  args
// 		want  string
// 		want1 []string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, got1 := toCommand(tt.args.input)
// 			if got != tt.want {
// 				t.Errorf("toCommand() got = %v, want %v", got, tt.want)
// 			}
// 			if !reflect.DeepEqual(got1, tt.want1) {
// 				t.Errorf("toCommand() got1 = %v, want %v", got1, tt.want1)
// 			}
// 		})
// 	}
// }
