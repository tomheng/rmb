package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	filePath := args[0]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	readLen := 3
	bs := make([]byte, readLen)
	for i := 0; ; i++ {
		n, err := file.Read(bs)
		if err != nil {
			log.Fatalln(err)
		}
		if n < 1 {
			break
		}
		fmt.Printf("%08d\t", i*readLen)
		for _, b := range bs {
			fmt.Printf(" %08b", b)
		}
		fmt.Println("")
		if n != readLen {
			break
		}
	}
}
