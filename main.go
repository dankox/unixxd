package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

// cli flags
var (
	uni = flag.Bool("unicode", false, "display unicode code points in hex")
)

func main() {
	flag.Parse()
	var file string
	if flag.NArg() >= 1 {
		file = flag.Arg(0)
	}
	if file == "" {
		fmt.Println("no file specified")
		os.Exit(1)
	}
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if *uni {
		displayUnicodeTable(b)
	} else {
		displayHexTable(b)
	}
}

func displayHexTable(b []byte) {
	i := 0
	hexPart := strings.Builder{}
	strPart := strings.Builder{}
	for _, ch := range b {
		if i%4 == 0 && i > 0 && i%16 != 0 {
			hexPart.WriteString(" ")
		}
		if i%16 == 0 && i > 0 {
			fmt.Printf("%08x:  %s  %s\n", i-16, hexPart.String(), strPart.String())
			hexPart.Reset()
			strPart.Reset()
		}
		if unicode.IsGraphic(rune(ch)) {
			strPart.WriteRune(rune(ch))
		} else {
			strPart.WriteRune('.')
		}
		hexPart.WriteString(fmt.Sprintf("%02x ", ch))
		i++
	}
	if hexPart.Len() > 0 {
		hexlen := 3*16 + 3
		fmt.Printf("%08x:  %s%s  %s\n", i-16, hexPart.String(), strings.Repeat(" ", hexlen-len(hexPart.String())), strPart.String())
	}
}

func displayUnicodeTable(b []byte) {
	data := string(b)
	i := 0
	hexPart := strings.Builder{}
	strPart := strings.Builder{}
	for _, ch := range data {
		if i%4 == 0 && i > 0 && i%16 != 0 {
			hexPart.WriteString(" ")
		}
		if i%16 == 0 && i > 0 {
			fmt.Printf("%08x:  %s  %s\n", i-16, hexPart.String(), strPart.String())
			hexPart.Reset()
			strPart.Reset()
		}
		if unicode.IsGraphic(ch) {
			strPart.WriteRune(ch)
		} else {
			strPart.WriteRune('.')
		}
		ucp := fmt.Sprintf("%U", ch)
		if len(ucp) > 6 {
			ucp = ucp[2:]
		} else {
			ucp = fmt.Sprintf("0%s", ucp[2:])
		}
		hexPart.WriteString(ucp)
		hexPart.WriteString(" ")
		i++
	}
	if hexPart.Len() > 0 {
		hexlen := 6*16 + 3 // default for unicode
		fmt.Printf("%08x:  %s%s  %s\n", i-16, hexPart.String(), strings.Repeat(" ", hexlen-len(hexPart.String())), strPart.String())
	}
	// fmt.Printf("\n\n%d characters\n", i-1)
	// fmt.Println("bytes length:", len(b))
	// fmt.Println("number of runes in bytes:", utf8.RuneCount(b))
	// fmt.Println("number of runes in string:", utf8.RuneCountInString(data))
}

// 😁
