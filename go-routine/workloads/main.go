package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go task0(&wg)
	go task1(&wg)
	go task1(&wg)
	go task1(&wg)
	go task2(&wg)
	go task3(&wg)
	go task4(&wg)
	go task0(&wg)
	go task0(&wg)
	go task0(&wg)
	wg.Wait()
}

func task0(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get("http://localhost:9090")
	if err != nil {
		log.Fatalf("Could not Fetch data from api")
	}

	fmt.Println("Task 0 done")
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	response, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		log.Fatalf("Could not Fetch data from api")
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("countries.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString(string(data))
	fmt.Println("Task 1 done")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	var count int
	for i := 0; i < 1_000_000_000; i++ {
		count = count + i
	}
	fmt.Println("Task 2 done")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 500)
	fmt.Println("Task 3 done")
}

func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 600)
	fmt.Println("Task 4 done")
}
