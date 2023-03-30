package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(5)
	go First_task1(&wg)
	go Second_task1(&wg)
	go Third_task1(&wg)
	go Fourth_task1(&wg)
	go Fifth_task1(&wg)

	wg.Wait()

}

func First_task1(wg *sync.WaitGroup) {
	defer wg.Done()
	data, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(data)
	fmt.Println("First Task")
}

func Second_task1(wg *sync.WaitGroup) {
	defer wg.Done()
	var count int
	for i := 0; i < 1_000_000_000; i++ {
		count = count + i
	}
	fmt.Println("Second Task")
}

func Third_task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Third Task")
}

func Fourth_task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fourth Task")
}

func Fifth_task1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Fifth Task")
}
