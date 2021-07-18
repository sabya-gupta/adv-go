package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println("Hi!!!!!")
	// structs.ProduceSomeTeamMembers()
	// slices.WorkWithSlices()
	// maps.MakeSomeMaps()
	// sets.UseMySet()
	// slices.AnotherSliceFunc()
	// mainFunc()
	// mainFuncWithCh()
	// sayHelloMultipleTimes()
	scClMain()
	// scClMainWithDefault()
}

func mainFunc() {
	go waitFor2Secs("Sub Function")
	time.Sleep(3 * time.Second)
	fmt.Println("Main Function")
}

func waitFor2Secs(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}

func mainFuncWithCh() {
	ch := make(chan string)
	go waitAndSay(ch)
	str := "to send"
	ch <- str
	str2 := <-ch
	log(str2)

}

func waitAndSay(c chan string) {
	b := <-c
	fmt.Println("Received " + b)
	time.Sleep(2 * time.Second)
	c <- "OK " + b

}

func log(s string) {
	fmt.Println(s)
}

func sayHelloMultipleTimes() {
	ch := make(chan string)

	go func(c chan string) {
		for i := 0; i < 5; i++ {
			c <- "hello"
		}
		close(c)
	}(ch)

	for s := range ch {
		log(s)
	}

	v, ok := <-ch

	fmt.Println(v, ok)

}

//SELECT
var sc = map[string]int{
	"id1": 6,
	"id2": 7,
	"id3": 5,
}

func findSecCl(id string, ch chan int) {
	fmt.Println("Recd ", id)
	time.Sleep(time.Duration(rand.Intn(7)) * time.Second)
	ch <- sc[id]
}

func scClMain() {
	rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int)
	ch2 := make(chan int)

	id1 := "id1"
	id2 := "id2"

	go findSecCl(id1, ch1)
	go findSecCl(id2, ch2)

	select {
	case sc := <-ch1:
		fmt.Println("ch 1 SecC = ", sc)

	case sc := <-ch2:
		fmt.Println("ch 2 SecC = ", sc)
	case <-time.After(4 * time.Second):
		fmt.Println("time out")
	}

}

func scClMainWithDefault() {
	rand.Seed(time.Now().UnixNano())

	ch1 := make(chan int)
	ch2 := make(chan int)

	id1 := "id1"
	id2 := "id2"

	go findSecCl(id1, ch1)
	go findSecCl(id2, ch2)

	select {
	case sc := <-ch1:
		fmt.Println("ch 1 SecC = ", sc)

	case sc := <-ch2:
		fmt.Println("ch 2 SecC = ", sc)
	default:
		fmt.Println("In Default")
	}

}
