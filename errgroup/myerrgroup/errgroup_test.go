package myerrgroup

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestErrGroup_Go(t *testing.T) {
	var g ErrGroup
	g.Go(func(ctx context.Context) error {
		name := "task 1"
		fmt.Println("run ", name)
		time.Sleep(1 * time.Second)
		return nil
	})

	g.Go(func(ctx context.Context) error {
		name := "task 2"
		fmt.Println("run ", name)
		time.Sleep(2 * time.Second)
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

func TestErrGroup_WithContext(t *testing.T) {
	ctx := context.Background()
	g := WithContext(ctx)
	g.Go(func(ctx context.Context) error {
		name := "task 1"
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel ", name)
				return nil
			default:
				fmt.Println("run ", name)
				time.Sleep(1 * time.Second)
				return fmt.Errorf("error %s", name)
			}
		}
	})

	g.Go(func(ctx context.Context) error {
		name := "task 2"
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel ", name)
				return nil
			default:
				fmt.Println("run ", name)
				time.Sleep(3 * time.Second)
				//				return nil
			}
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

func TestErrGroup_Panic(t *testing.T) {
	ctx := context.Background()
	g := WithContext(ctx)
	g.Go(func(ctx context.Context) error {
		name := "task 1"
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel ", name)
				return nil
			default:
				fmt.Println("run ", name)
				time.Sleep(1 * time.Second)
				panic("panic")
				//return fmt.Errorf("error %s", name)
			}
		}
	})

	g.Go(func(ctx context.Context) error {
		name := "task 2"
		for {
			select {
			case <-ctx.Done():
				fmt.Println("cancel ", name)
				return nil
			default:
				fmt.Println("run ", name)
				time.Sleep(3 * time.Second)
				//				return nil
			}
		}
	})

	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}
