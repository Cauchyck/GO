package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

type TraceCode string
type UserId string

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")

	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}

	userKey := UserId("USER_ID")
	userId, ok := ctx.Value(userKey).(int64)
	if !ok {
		fmt.Println("invalid user id")
	}
	
	log.Printf("%s worker func ...", traceCode)
	log.Printf("User id: %d", userId)

LOOP:
	for {
		fmt.Printf("worker, trace code: %s \n", traceCode)
		time.Sleep(time.Millisecond * 10)
		select {
		case <-ctx.Done():
			break LOOP
		default:
		}
	}
	fmt.Println("work done")
	wg.Done()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)

	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "1234")
	ctx = context.WithValue(ctx, UserId("USER_ID"), int64(1))

	wg.Add(1)
	go worker(ctx)

	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Printf("over")
}
