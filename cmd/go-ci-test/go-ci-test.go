package main

import (
	"log"

	"github.com/Zouizoui78/simple-ddns/internal/add"
)

func main() {
	arg1 := 1
	arg2 := 2
	res := add.Add(arg1, arg2)
	log.Printf("result of %v + %v is %v", arg1, arg2, res)
}
