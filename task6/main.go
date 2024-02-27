package main

import (
	"context"
	"fmt"
	"time"
)

// stopWithCtx завершает свою работу при получении сигнала от контекста.
func stopWithCtx(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println("stopWithCtx is still running")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// stopWithChan завершает свою работу при записи данных в канал
func stopWithChan(done chan struct{}) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("stopWithChan is still running")
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	// создаем контекст с отменой и вызываем CancelFunc спустя 400 милисекунд
	ctx, cancel := context.WithCancel(context.Background())
	go stopWithCtx(ctx)
	time.Sleep(400 * time.Millisecond)
	cancel()

	// создаем канал и записываем в него значение спустя 300 милисекунд
	done := make(chan struct{})
	defer close(done)

	go stopWithChan(done)
	time.Sleep(300 * time.Millisecond)
	done <- struct{}{}
}
