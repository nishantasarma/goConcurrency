package main

import "fmt"

// start-> stage1 -> stage2 -> end

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}

		close(out)

	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {

		for n := range in {

			out <- n * n
		}
		close(out)
	}()

	return out

}

func main() {
	nums := []int{2, 3, 4, 7}
	//stage1
	dataChannel := sliceToChannel(nums)
	// stage2
	finalChannel := sq(dataChannel)
	// stage3
	for n := range finalChannel {
		fmt.Println(n)
	}
}
