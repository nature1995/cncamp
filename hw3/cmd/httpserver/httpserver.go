package httpserver

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"hw1/internal/logger"
)

func NewHttpServer(addr string) *http.Server {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/version", getVersion)
	r.HandleFunc("/healthz", health)
	r.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	logger.Logger.Infof("[HTTP] http server listen: %s", addr)

	return &http.Server{
		Addr:    ":" + addr,
		Handler: r,
	}
}

// health check
func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, fmt.Sprintf("healthz return code: %d", http.StatusOK))
	logger.Logger.Infof("healthz return code: %d", http.StatusOK)
	if err != nil {
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	for s, strings := range r.Header {
		for _, v := range strings {
			w.Header().Add(s, v)
		}
	}
	getVersion(w, r)
	fmt.Printf("Access client ip is: %s, return code is: %d\n", getRemoteIP(r), http.StatusOK)
	logger.Logger.Infof("Access client ip is: %s, return code is: %d\n", getRemoteIP(r), http.StatusOK)
}

// getVersion 获取版本
func getVersion(w http.ResponseWriter, r *http.Request) {
	version := os.Getenv("VERSION")
	if version != "" {
		w.Header().Set("VERSION", version)
		_, _ = io.WriteString(w, fmt.Sprintf("VERSION is: %s", version))
	}
}

func getRemoteIP(req *http.Request) string {
	rAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		rAddr = ip
	} else if ip = req.Header.Get("X-Real-IP"); ip != "" {
		rAddr = ip
	} else {
		rAddr, _, _ = net.SplitHostPort(rAddr)
	}

	if rAddr == "::1" {
		rAddr = "127.0.0.1"
	}

	return rAddr
}
