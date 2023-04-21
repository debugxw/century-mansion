package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go Cbw.startServer(&wg)
	time.Sleep(time.Second)
	conn := Cbw.startClient(&wg)
	t1 := time.Now()
	wg.Wait()
	elapsed := time.Since(t1)
	_ = conn.Close()
	fmt.Println("耗时: ", elapsed)
}
