package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

// writer пишет случайные целочисленные значения в канал с интервалом в 200 милисекунд.
// Когда приходит сигнал от контекста, worker завершает свою работу.
func writer(ctx context.Context, in chan<- int) {
	for {
		select {
		case in <- rand.Int():
			time.Sleep(200 * time.Millisecond)
		case <-ctx.Done():
			close(in)
			return
		}
	}
}

// worker читает значение из канала и выводит его в stdout.
// Когда канал закрывается, worker завершает свою работу.
func worker(out <-chan int) {
	for val := range out {
		fmt.Println("Get value: ", val)
	}
}

func main() {
	// numOfWorkers задает количество воркеров через флаг "-n", по умолчанию 2 воркера
	numOfWorkers := flag.Int("n", 2, "number of worker goroutines")
	flag.Parse()

	fmt.Printf("Started program with number of workers: %d\n", *numOfWorkers)

	ch := make(chan int)

	// Контекст, который будет завершен сигналом Ctrl+C
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	go writer(ctx, ch)

	for i := 0; i < *numOfWorkers; i++ {
		go worker(ch)
	}

	// Дожидаемся сигнала о завершении программы
	<-ctx.Done()
}
