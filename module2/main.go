package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "net/http/pprof"
)

func main() {
	// flag.Set("v", "4")
	// glog.V(2).Info("Starting http server...")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)

	err := http.ListenAndServe(":8080", nil)

	// mux.HandleFunc("/healthz", healthz)

	if err != nil {
		log.Fatal(err)
	}

}

// response 200
func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
}

func getEnvVersion() string {
	version := os.Getenv("VERSION")
	return version
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	// 需要先设置响应头，再写入响应体
	// 1. set header
	w.Header().Set("Content-Length", "1024")
	for key, values := range r.Header {
		for _, value := range values {
			// fmt.Println(key, value)
			w.Header().Add(key, value)
		}
	}

	// 2. get env
	env := getEnvVersion()
	if env != "" {
		w.Header().Add("version", env+"\n")
	} else {
		w.Header().Add("version", "v1\n")
	}
	w.WriteHeader(http.StatusOK)

	// 3. get ip
	ip := r.RemoteAddr
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ", ")
		ip = ips[len(ips)-1]
	}
	fmt.Fprintf(w, "客户端 IP: %s\n", ip)

	fmt.Fprintf(w, "http 状态码: %d\n", http.StatusOK)

	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")

	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}

}
