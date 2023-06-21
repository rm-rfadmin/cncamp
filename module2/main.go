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

	err := http.ListenAndServe(":80", nil)

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
	fmt.Println("entering root handler")
	user := r.URL.Query().Get("user")
	env := getEnvVersion()
	if user != "" {
		io.WriteString(w, fmt.Sprintf("hello [%s]\n", user))
	} else {
		io.WriteString(w, "hello [stranger]\n")
	}

	// 1. header
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}

	// 2. get env
	if env != "" {
		io.WriteString(w, fmt.Sprintf("version: %s\n", env))
	} else {
		io.WriteString(w, "version: v1\n")
	}

	// 3. get ip
	ip := r.RemoteAddr
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ", ")
		ip = ips[len(ips)-1]
	}
	fmt.Fprintf(w, "客户端 IP: %s\n", ip)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "http 状态码: %d\n", http.StatusOK)
}
