package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)

	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()
	select {
	case <-time.After(time.Second):
		fmt.Println("overSleep")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}