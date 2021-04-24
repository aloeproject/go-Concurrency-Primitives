package pipeline

func create(done chan struct{}, value ...int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for _, v := range value {
			select {
			case <-done:
				return
			case result <- v:
			}
		}
	}()
	return result
}

func add(done chan struct{}, value <-chan int, addVal int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for v := range value {
			select {
			case <-done:
				return
			case result <- v + addVal:
			}
		}
	}()
	return result
}

func multiply(done chan struct{}, value <-chan int, mulVal int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)
		for v := range value {
			select {
			case <-done:
				return
			case result <- v * mulVal:
			}
		}
	}()
	return result
}

