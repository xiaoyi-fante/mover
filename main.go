package main

import (
	"math/rand"
	"time"

	"github.com/xiaoyi-fante/mover/cmd"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd.Execute()
}
