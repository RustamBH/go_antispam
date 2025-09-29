// http_task.go
package main

import (
	"fmt"
	"strings"
)

// MaskURL маскирует часть URL после протокола
func MaskURL(url string) string {
	if !strings.HasPrefix(url, "http://") {
		return url
	}	

	out := []rune(url)
	for i := 7; i < len(out); i++ {
		out[i] = '*'
	}
	return string(out)
}

func main() {
	url := "http://hehefouls.netHAHAHA"
	masked := MaskURL(url)
	fmt.Println(masked)
}


