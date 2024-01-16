package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func xor(input, key string) string {
	var result string
	for i := 0; i < len(input); i += 8 {
		if i+8 > len(input) {
			break
		}
		inp, _ := strconv.ParseUint(input[i:i+8], 2, 8)
		k, _ := strconv.ParseUint(key[i%len(key):i%len(key)+8], 2, 8)
		result += fmt.Sprintf("%08b", byte(inp)^byte(k))
	}
	return result
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		fmt.Println("Usage: xor textfile keyfile")
		os.Exit(1)
	}

	textfile := args[0]
	keyfile := args[1]

	text, err := ioutil.ReadFile(textfile)
	if err != nil {
		fmt.Println("Error reading text file:", err)
		os.Exit(1)
	}

	key, err := ioutil.ReadFile(keyfile)
	if err != nil {
		fmt.Println("Error reading key file:", err)
		os.Exit(1)
	}

	if len(key) < len(text) {
		fmt.Println("The key is shorter than the text!")
		os.Exit(1)
	}

	fmt.Println(xor(string(text), string(key)))
}

