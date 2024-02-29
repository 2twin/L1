package main

import (
	"context"
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	// первый вариант - через конекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()

	<-ctx.Done()

	// второй вариант
	// <-time.After(d)
}

func main() {
	d := 2 * time.Second
	now := time.Now()

	sleep(d)

	fmt.Println(time.Since(now))
}