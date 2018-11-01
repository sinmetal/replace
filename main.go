package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("file", "", "")
	src := flag.String("src", "", "")
	dst := flag.String("dst", "", "")
	flag.Parse()

	fmt.Printf("param %s:%s\n", *src, *dst)
	fmt.Println("$SECRET_TOKEN =", os.Getenv("$SECRET_TOKEN"))
	fmt.Println("$$SECRET_TOKEN =", os.Getenv("$$SECRET_TOKEN"))

	srcText := *src
	if strings.HasPrefix(*src, "$") {
		srcText = os.Getenv(*src)
	}
	dstText := *dst
	if strings.HasPrefix(*dst, "$") {
		dstText = os.Getenv(*dst)
	}

	buf := bytes.Buffer{}
	f, err := os.OpenFile(*fileName, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("%s:%s:%s\n", scanner.Text(), srcText, dstText)
		l := strings.Replace(scanner.Text(), srcText, dstText, -1)
		buf.WriteString(l)
		buf.WriteString("\n")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	wf, err := os.Create(*fileName)
	if err != nil {
		panic(err)
	}
	_, err = buf.WriteTo(wf)
	if err != nil {
		panic(err)
	}
}