package flakygo

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func flaky(n int) bool {
	return rand.Intn(n) == 0
}
