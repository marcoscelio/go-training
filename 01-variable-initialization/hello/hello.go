package hello

import (
	"fmt"
	"time"
)

// func Greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World! %s", time.Now())
// }

func Hello() {
	fmt.Printf("Hello World! %s", time.Now())
}