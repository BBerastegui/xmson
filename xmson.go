package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	xj "github.com/basgys/goxml2json"
)

func main() {

	// Flags
	prettifyPtr := flag.Bool("p", false, "Prettify output.")
	flag.Parse()

	xmlData := bufio.NewReader(os.Stdin)
	jsonData, err := xj.Convert(xmlData)

	if err != nil {
		panic(err)
	}

	var result string
	if *prettifyPtr {
		result = string(prettify(jsonData)[:])
	} else {
		result = jsonData.String()
	}

	// Print final result
	fmt.Println(result)
}

func prettify(jsonData *bytes.Buffer) []byte {
	var dat map[string]interface{}
	b, _ := ioutil.ReadAll(jsonData)

	if err := json.Unmarshal(b, &dat); err != nil {
		panic(err)
	}

	result, _ := json.MarshalIndent(dat, "", "    ")

	return result
}
