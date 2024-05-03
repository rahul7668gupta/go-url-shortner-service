package utils

import (
	"testing"
)

func TestGetShortCodeFromId(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				num: 1,
			},
			want: "1",
		},
		{
			name: "2D",
			args: args{
				num: 137,
			},
			want: "2D",
		},
		{
			name: "1000",
			args: args{
				num: 238328,
			},
			want: "1000",
		},
		{
			name: "10000",
			args: args{
				num: 14776336,
			},
			want: "10000",
		},
		{
			name: "yKSe",
			args: args{
				num: 14378336,
			},
			want: "yKSe",
		},
		{
			name: "16iki",
			args: args{
				num: 16378336,
			},
			want: "16iki",
		},
		{
			name: "100QXA",
			args: args{
				num: 916234832,
			},
			want: "100QXA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetShortCodeFromId(tt.args.num); got != tt.want {
				t.Errorf("GetShortCodeFromId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetUrlDomain(t *testing.T) {
	type args struct {
		inputUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "https://amazon.in",
			args:    args{inputUrl: "https://amazon.in"},
			want:    "amazon.in",
			wantErr: false,
		},
		{
			name:    "https://www.youtube.in",
			args:    args{inputUrl: "https://www.youtube.in"},
			want:    "www.youtube.in",
			wantErr: false,
		},
		{
			name:    "https://www.udemy.com",
			args:    args{inputUrl: "https://www.udemy.com"},
			want:    "www.udemy.com",
			wantErr: false,
		},
		{
			name:    "https://www.infracloudtech.in",
			args:    args{inputUrl: "https://www.infracloudtech.in"},
			want:    "www.infracloudtech.in",
			wantErr: false,
		},
		{
			name:    "https://www.amazon.in",
			args:    args{inputUrl: "https://www.amazon.in"},
			want:    "www.amazon.in",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUrlDomain(tt.args.inputUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUrlDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUrlDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
