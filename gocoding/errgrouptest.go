package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	group, cxt := errgroup.WithContext(context.Background())
	data := make(chan int)
	group.Go(func() error {
		defer close(data)
		for i := 0; i < 5; i++ {
			fmt.Println("sending ", i)

			println(i)
			data <- i
		}
		return fmt.Errorf("hhhhh")
	})
	group.Go(func() error {
		for {
			select {
			case <-cxt.Done():
				return cxt.Err()
			case num := <-data:
				fmt.Println("receiving ", num)
			}
		}
	})
	err := group.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
