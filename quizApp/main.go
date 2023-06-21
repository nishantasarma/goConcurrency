package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"
)

func startTimer(timeout chan bool) {
	fmt.Println("Starting timer")
	time.Sleep(time.Second * 3)
	fmt.Println("3 Seconds over")

	close(timeout)
	fmt.Println("3 Seconds over")

}

func main() {

	bytes, _ := ioutil.ReadFile("./quiz")
	questions := strings.Split(fmt.Sprintf(string(bytes)), "\n")
	fmt.Println(questions)
	fmt.Println(reflect.TypeOf(questions))
	var ans string
	timeout := make(chan bool)
	go startTimer(timeout)
	for i := 0; i < len(questions); i += 2 {
		select {
		case <-timeout:
			fmt.Println("Timed out")
			return
		default:
			fmt.Println(questions[i])

			fmt.Scan(&ans)

			if questions[i+1] != ans {
				fmt.Println("Wrong Answer")
				break
			}

		}

	}

}
