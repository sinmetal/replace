package main

import (
	"bufio"
	"bytes"
	"flag"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("file", "", "")
	o := flag.String("old", "", "")
	n := flag.String("new", "", "")

	flag.Parse()

	oldText := *o
	if strings.HasPrefix(*o, "$") {
		oldText = getOSEnv(*o)
	}
	newText := *n
	if strings.HasPrefix(*n, "$") {
		newText = getOSEnv(*n)
	}

	buf, err := replaceFile(*fileName, oldText, newText)
	if err != nil {
		panic(err)
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

func replaceFile(fileName string, old string, new string) (bytes.Buffer, error) {
	buf := bytes.Buffer{}
	f, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		return buf, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := strings.Replace(scanner.Text(), old, new, -1)
		buf.WriteString(l)
		buf.WriteString("\n")
	}
	if err := scanner.Err(); err != nil {
		return buf, err
	}
	return buf, nil
}

func getOSEnv(key string) string {
	return os.Getenv(strings.Replace(key, "$", "", -1))
}
