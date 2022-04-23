package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

// HW1
// 编写一个 HTTP服务器。
// 1.接收客户端request，并将request中带的 header写入response header
// 2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
// 3.Server端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
// 4.当访问localhost/healthz 时，应返回 200
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	mux.HandleFunc("/healthz", health)
	log.Println("[HTTP] http server listen: ", 80)

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Println("[HTTP] http server listen: ", 80, "error: ", err)
		os.Exit(1)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, fmt.Sprintf("healthz return code: %d", http.StatusOK))
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
	getVersion(w)
	fmt.Printf("Access client ip is: %s, return code is: %d\n", getRemoteIP(r), http.StatusOK)
}

func getVersion(w http.ResponseWriter) {
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
