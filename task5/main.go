package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// worker запускает 2 горутины: первая пишет случайные целочисленные данные в канал и закрывает канал,
// если получает сигнал от контекста, вторая горутина читает данные из канала, пока он открыт.
func worker(ctx context.Context, ch chan int) {
	go func() {
		for {
			v := rand.Int()

			select{
			case ch <- v:
				fmt.Printf("Send value: %d\n", v)
				time.Sleep(100 * time.Millisecond)
			case <-ctx.Done():
				close(ch)
				return
			}
		}
	}()

	go func() {
		for v := range ch {
			fmt.Printf("Get value: %d\n", v)
		}
	}()
}

func main() {
	// sec задает количество секунд через флаг "-s", по истечению которых программа завершается, по умолчанию 2 секунды
	sec := flag.Int("s", 2, "number of seconds for worker to run")
	flag.Parse()

	ch := make(chan int)

	// создаем конекст с таймаутом и передаем его в worker
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*sec) * time.Second)
	defer cancel()

	worker(ctx, ch)

	<-ctx.Done()
}