package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var NUMBER_OF_SITES = 5
var OK_STATUS_CODE = 200

func main() {

	sites := []string{"http://google.com", "http://facebook.com", "http://amazon.com", "http://stackoverflow.com", "http://linkedin.com"}

	// Without go routine or WaitGroup
	var ini time.Time
	fmt.Println("Without go routineor WaitGroup:")
	ini = time.Now()
	for _, v := range sites {
		healthCheck(v)
	}
	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")

	// Using WaitGroup
	fmt.Println("Using go WaitGroup:")
	ini = time.Now()
	wg := &sync.WaitGroup{}
	wg.Add(NUMBER_OF_SITES)
	for _, v := range sites {
		go healthCheckUsingWaitGroup(wg, v)
	}
	wg.Wait()
	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")

	// Using go routine
	fmt.Println("Using go routine1:")
	ini = time.Now()
	for _, v := range sites {
		result := healthCheckUsingGoRoutine(v)
		for {
			msg := <-result
			if msg == "" {
				break
			}
			fmt.Println(msg)
		}
	}
	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")

	fmt.Println("Using go routine2:")
	ini = time.Now()
	result := make(chan string)
	for _, v := range sites {
		go healthCheckUsingGoRoutine2(v, result)
	}
	counter := 0
	for msg := range result {
		fmt.Println(msg)
		counter++
		if counter == NUMBER_OF_SITES {
			break
		}
	}
	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")

	fmt.Println("Merging go routine:")
	t1 := healthCheckList(sites[:3])
	t2 := healthCheckList(sites[2:5])
	r := healthCheckMergingGoRoutines(t1, t2)
	for {
		msg := <-r
		if msg == "" {
			break
		}

		fmt.Println(msg)
	}
	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")
}

func healthCheck(l string) {
	resp, err := http.Get(l)
	if err != nil {
		fmt.Printf("Error accessing site %s \n", l)
	}
	if resp.StatusCode == OK_STATUS_CODE {
		fmt.Printf("%s OK\n", l)
	} else {
		fmt.Printf("Error for site %s\n", l)
	}
}

func healthCheckUsingGoRoutine(l string) <-chan string {
	result := make(chan string)
	go func() {
		resp, err := http.Get(l)
		if err != nil {
			result <- fmt.Sprintf("Error accessing site %s \n", l)
		}
		if resp.StatusCode == OK_STATUS_CODE {
			result <- fmt.Sprintf("%s OK\n", l)
		} else {
			result <- fmt.Sprintf("Error for site %s\n", l)
		}
		close(result)
	}()
	return result
}

func healthCheckUsingGoRoutine2(l string, c chan<- string) {
	resp, err := http.Get(l)
	if err != nil {
		c <- fmt.Sprintf("Error accessing site %s \n", l)
	}
	if resp.StatusCode == OK_STATUS_CODE {
		c <- fmt.Sprintf("%s OK\n", l)
	} else {
		c <- fmt.Sprintf("Error for site %s\n", l)
	}
}

func healthCheckList(sites []string) <-chan string {
	result := make(chan string)
	go func() {
		for _, v := range sites {
			resp, err := http.Get(v)
			if err != nil {
				result <- fmt.Sprintf("Error accessing site %s \n", v)
			}
			if resp.StatusCode == OK_STATUS_CODE {
				result <- fmt.Sprintf("%s OK\n", v)
			} else {
				result <- fmt.Sprintf("Error for site %s\n", v)
			}
		}
		close(result)
	}()
	return result
}

func healthCheckMergingGoRoutines(task1, task2 <-chan string) <-chan string {
	result := make(chan string)
	go func() {
		for {
			select {
			case msg := <-task1:
				result <- msg
			case msg := <-task2:
				result <- msg
			}
		}
	}()
	return result
}

func healthCheckUsingWaitGroup(wg *sync.WaitGroup, l string) {
	defer wg.Done()
	resp, err := http.Get(l)
	if err != nil {
		fmt.Printf("Error accessing site %s \n", l)
	}
	if resp.StatusCode == OK_STATUS_CODE {
		fmt.Printf("%s OK\n", l)
	} else {
		fmt.Printf("Error for site %s\n", l)
	}

}
