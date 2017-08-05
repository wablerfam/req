package main

import (
	"encoding/json"
	"fmt"
	"flag"
	"path"

	"req/client"
	"req/options"
)

type Result struct {
	Code int `json:"code"`
	Time float64 `json:"time"`
	Body string `json:"body"`
	File string `json:"file"`
}

func main() {
	var (
		flagBody = flag.Bool("b", false, "stream the response body")
		flagOutputFile = flag.String("o", "", "output body to a file")
		flagOutputFileEasy = flag.Bool("O", false, "Output the body to file with no file name")
	)
	
	flag.Parse()

	if *flagOutputFile != "" && *flagOutputFileEasy != false {
		panic("Error: -o and -O can not be used at the same time")
	}

	url := flag.Args()[0]

	code, body, time := client.Request(url)

	result := Result{
		Code: code,
		Time: time.Seconds(),
		Body: "no print",
		File: "no output",
	}

	if *flagBody != false {
		result.Body = body
	}

	if *flagOutputFile != "" {
		options.OutputFile(*flagOutputFile, body)
		result.File = *flagOutputFile
	} else if *flagOutputFileEasy != false {
		_ ,filename := path.Split(url)
		options.OutputFile(filename, body)
		result.File = filename
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonResult))
}
