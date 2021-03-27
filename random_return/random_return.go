package main

import (
	"math/rand"
	"time"
)

var hex = []rune("0123456798abcdefgh")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func random_int(lower, upper int) (int){
	if lower > upper{
		lower, upper = upper, lower
	}
	return rand.Intn((upper - lower + 1) + lower)
}

func random_hex(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = hex[rand.Intn(len(hex))]
	}
	return string(b)
}

func main() {

}


