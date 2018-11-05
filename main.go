package main

import (
	"fmt"
	"testing"
	"time"
)

func main() {
	result := testing.Benchmark(func(b *testing.B) { run() })
	fmt.Println(result.T)
}

// 並列処理してないので、終了まで６秒かかる
func run() {
	fmt.Println("Start!")
	process("A")
	process("B")
	process("C")
	fmt.Println("Finish!")
}

// 2秒かかるプロセス
func process(name string) {
	time.Sleep(2 * time.Second)
	fmt.Println(name)
}
