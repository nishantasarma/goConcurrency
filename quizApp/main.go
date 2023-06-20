package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"time"
)

func startTimer() {
	time.Sleep(time.Second * 3)
	close(timeout)
}

func main() {

	bytes, _ := ioutil.ReadFile("./quiz")
	questions := strings.Split(fmt.Sprintf(string(bytes)), "\n")
	fmt.Println(questions)
	fmt.Println(reflect.TypeOf(questions))
	var ans string
	timeout := make(chan bool)
	mytimer := make(chan bool)

	for i := 0; i < len(questions); i += 2 {
		// fmt.Println(questions[i])

		fmt.Scan(&ans)

		if questions[i+1] != ans {
			fmt.Println("Wrong Answer")
			break

		}
	}

}
