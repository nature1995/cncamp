package httpserver

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"hw4/internal/logger"
)

func init() {
	_ = os.Setenv("VERSION", "1.0.0-test")
	// 初始化日志
	_, err := logger.NewLogger(logger.Config{
		Path:  "./logs",
		Level: "DEBUG",
	})
	if err != nil {
		fmt.Printf("Failed to create logger: %s", err)
		os.Exit(1)
	}

	// 初始化服务
	server := NewHttpServer("8080")
	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}

func checkHttpStatusCode(t *testing.T, got, expect int) {
	if got != expect {
		t.Errorf("status code is not match, expect %d, got %d", expect, got)
	}
}

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
		{
			name: "测试",
			args: args{
				req: &http.Request{
					RemoteAddr: "",
				},
			},
			want: "",
		},
		{
			name: "测试X-Forwarded-For",
			args: args{
				req: &http.Request{
					RemoteAddr: "",
					Header: map[string][]string{
						"X-Forwarded-For": {"127.0.0.1"},
					},
				},
			},
			want: "127.0.0.1",
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

func Test_getVersion(t *testing.T) {
	r, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	srv := http.HandlerFunc(getVersion)
	srv.ServeHTTP(w, r)

	checkHttpStatusCode(t, w.Code, http.StatusOK)
	if w.Header().Get("Version") == "" {
		t.Error("Invalid version info")
	}
}

func Test_health(t *testing.T) {
	r, err := http.NewRequest("GET", "/healthz", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	srv := http.HandlerFunc(health)
	srv.ServeHTTP(w, r)

	checkHttpStatusCode(t, w.Code, http.StatusOK)
}

func Test_handler(t *testing.T) {
	r, err := http.NewRequest("GET", "/", nil)
	r.Header.Add("key", "value")
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	srv := http.HandlerFunc(handler)
	srv.ServeHTTP(w, r)

	checkHttpStatusCode(t, w.Code, http.StatusOK)
	if w.Header().Get("Version") == "" {
		t.Error("Invalid version info")
	}
}

func Test_ready(t *testing.T) {
	r, err := http.NewRequest("GET", "/ready", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	srv := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	srv.ServeHTTP(w, r)

	checkHttpStatusCode(t, w.Code, http.StatusOK)
}

func TestNewHttpServer(t *testing.T) {
	_ = os.Setenv("VERSION", "1.0.0-test")
	// 初始化日志
	_, err := logger.NewLogger(logger.Config{
		Path:  "./logs",
		Level: "DEBUG",
	})
	if err != nil {
		fmt.Printf("Failed to create logger: %s", err)
		os.Exit(1)
	}
	server := NewHttpServer("18080")
	go func() {
		if err = server.ListenAndServe(); err != nil {
			t.Error(err)
			return
		}
	}()
	time.Sleep(5 * time.Second)
}
