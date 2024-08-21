package main

import (
	"fmt"
	"gin-ranking/router"
)

func main() {
	r := router.Router()

	// defer recover panic
	defer fmt.Println("111")
	r.Run(":8080")
}
