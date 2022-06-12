package httpserver

import (
	"net/http"
	"testing"
)

func Test_getRemoteIP(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "测试localhost",
			args: args{
				req: &http.Request{
					RemoteAddr: "127.0.0.1:8080",
				},
			},
			want: "127.0.0.1",
		},
		{
			name: "测试no localhost",
			args: args{
				req: &http.Request{
					RemoteAddr: "192.168.1.2:8080",
				},
			},
			want: "192.168.1.2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRemoteIP(tt.args.req); got != tt.want {
				t.Errorf("getRemoteIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
