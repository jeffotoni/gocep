package env

import (
	"testing"
	"time"
)

func TestGetInt(t *testing.T) {
	type args struct {
		name         string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test_getint_", args{name: "MONGO_PORT", defaultValue: 10}, 10},
		{"test_getint_", args{name: "MONGO_PORT", defaultValue: 5}, 5},
		{"test_getint_", args{name: "MONGO_PORT", defaultValue: 2}, 2},
		{"test_getint_", args{name: "MONGO_PORT", defaultValue: -100}, -100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt(tt.args.name, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDuration(t *testing.T) {
	type args struct {
		name         string
		defaultValue time.Duration
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{"test_getduration_", args{name: "MONGO_TIME",
			defaultValue: time.Duration(time.Second) * 2}, time.Duration(time.Second) * 2},
		{"test_getduration_", args{name: "MONGO_TIME",
			defaultValue: time.Duration(time.Second) * 60}, time.Duration(time.Second) * 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDuration(tt.args.name, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetString(t *testing.T) {
	type args struct {
		name         string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test_getstring_", args{name: "MONGO_USER",
			defaultValue: "mongo-user"}, "mongo-user"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.name, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	type args struct {
		name         string
		defaultValue bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test_getstring_", args{name: "MONGO_DB_EXIST",
			defaultValue: true}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBool(tt.args.name, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
