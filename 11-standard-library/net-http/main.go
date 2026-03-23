package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	h := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ok")
	}
	srv := httptest.NewServer(http.HandlerFunc(h))
	defer srv.Close()

	resp, err := http.Get(srv.URL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("status:", resp.Status, "body:", string(body))
}
