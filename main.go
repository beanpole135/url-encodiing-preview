package main

import (
	"bufio"
	"fmt"
	"os"
	"net/url"
	"strings"
)

func main() {
	//Prompt for the text to encode
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Text to convert: ")
	txt, _ := reader.ReadString('\n')
	txt = strings.TrimSuffix(txt,"\n")
	txt = strings.TrimSuffix(txt,"\r") //windows
	//Now Print out the results:
	clean, _ := url.PathUnescape(txt)
	fmt.Println("---------------")
	fmt.Println("Clean Text: ", clean)
	fmt.Println("URL Encoded: ", url.PathEscape(txt))
	fmt.Println("---------------")
}