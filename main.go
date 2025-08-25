package main

import (
	"fmt"
	"strings"
	"sync"
)


func isPalindrome(s string) bool {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}


func say(lang, text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%s: %s\n", lang, text)
}

func factorial(n uint64) uint64 {
	var res uint64 = 1
	for i := uint64(2); i <= n; i++ {
		res *= i
	}
	return res
}

func main() {
	fmt.Println(isPalindrome("level")) // true
	fmt.Println(isPalindrome("hello")) // false

	hellos := map[string]string{
		"English": "Hello",
		"Spanish": "Hola",
		"French":  "Bonjour",
		"Hindi":   "Namaste",
	}

	var wg sync.WaitGroup
	wg.Add(len(hellos))
	for lang, text := range hellos {
		go say(lang, text, &wg)
	}
	wg.Wait()

	fmt.Println(factorial(5)) 
	fmt.Println(factorial(10))

	text := "Go is fun and Go is fast"
	counts := map[string]int{}
	for _, w := range strings.Fields(text) {
		counts[strings.ToLower(w)]++
	}
	fmt.Println(counts) 
}
