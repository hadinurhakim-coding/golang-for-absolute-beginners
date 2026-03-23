package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("halo io")
	data, err := io.ReadAll(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	var buf bytes.Buffer
	fmt.Fprint(&buf, "tulis ke buffer")
	fmt.Println(buf.String())
}
