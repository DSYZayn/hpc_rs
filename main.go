package main

import (
	"fmt"
	"math/rand"

	// "sync"
	"time"
)

func main() {
	const N = 10
	// var wg sync.WaitGroup

	a := make([][]float64, N)
	b := make([][]float64, N)
	// c := make([][]float64, N)
	d := make([][]float64, N)

	for i := 0; i < N; i++ {
		a[i] = make([]float64, N)
		b[i] = make([]float64, N)
		// c[i] = make([]float64, N)
		d[i] = make([]float64, N)
	}

	// init a & b
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a[i][j] = float64(rand.Float64())
			b[i][j] = float64(rand.Float64())
		}
	}

	// bT := time.Now()

	// for i := 0; i < N; i++ {
	// 	wg.Add(1)

	// 	go func(i int) {
	// 		defer wg.Done()

	// 		for j := 0; j < N; j++ {
	// 			for k := 0; k < N; k++ {
	// 				c[i][j] += a[i][k] * b[k][j]
	// 			}
	// 		}
	// 	}(i)
	// }

	// wg.Wait()

	// eT := time.Since(bT)

	// fmt.Println("WaitGroup method Calculate time: ", eT)

	// channel method

	nbT := time.Now() //reset

	ch := make(chan struct{}, 20) //利用channel控制并发数 为20

	for i := 0; i < N; i++ {
		go func(i int) {
			ch <- struct{}{} // 一个"协程"入队
			for j := 0; j < N; j++ {
				for k := 0; k < N; k++ {
					d[i][j] += a[i][k] * b[k][j]
				}
			}
			<-ch //一个"协程"出队
		}(i)
	}

	neT := time.Since(nbT)
	fmt.Println("channel method Calculate time: ", neT)

	for i := 0; i < N; i++ {
		fmt.Println(d[i])
	}
}
