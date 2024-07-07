package tools

import (
	"sync"
)

type Number interface {
	float32 | float64 | int | int16 | int32 | int64
}

func MatmulSync[T Number](a [][]T, b [][]T, N int) [][]T {
	var wg sync.WaitGroup

	c := make([][]T, N)

	for i := 0; i < N; i++ {
		c[i] = make([]T, N)
	}

	for i := 0; i < N; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			for j := 0; j < N; j++ {
				for k := 0; k < N; k++ {
					c[i][j] += a[i][k] * b[k][j]
				}
			}
		}(i)
	}

	wg.Wait()

	return c
}

func MatmulCh[T Number](a [][]T, b [][]T, N int, max_ch int) [][]T {
	c := make([][]T, N)

	for i := 0; i < N; i++ {
		c[i] = make([]T, N)
	}

	ch := make(chan struct{}, max_ch) //利用channel控制并发数 为20

	for i := 0; i < N; i++ {
		go func(i int) {
			ch <- struct{}{} // 一个"协程"入队
			for j := 0; j < N; j++ {
				for k := 0; k < N; k++ {
					c[i][j] += a[i][k] * b[k][j]
				}
			}
			<-ch //一个"协程"出队
		}(i)
	}

	return c
}
