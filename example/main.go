package main

import (
	"fmt"
	"hpc_rs/tools"
	"math/rand"
	"time"
)

func main() {
	const N = 10

	a := make([][]float64, N)
	b := make([][]float64, N)

	for i := 0; i < N; i++ {
		a[i] = make([]float64, N)
		b[i] = make([]float64, N)
	}

	// init a & b
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a[i][j] = float64(rand.Float64())
			b[i][j] = float64(rand.Float64())
		}
	}

	bT := time.Now()

	c := tools.MatmulSync(a, b, 10)

	eT := time.Since(bT)

	fmt.Println("WaitGroup method Calculate time: ", eT)
	for i := 0; i < N; i++ {
		fmt.Println(c[i])
	}

	// channel method

	nbT := time.Now() //reset

	d := tools.MatmulCh(a, b, N, 20)

	neT := time.Since(nbT)
	fmt.Println("channel method Calculate time: ", neT)

	for i := 0; i < N; i++ {
		fmt.Println(d[i])
	}
}
