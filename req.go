package main

import (
	"encoding/json"
	"fmt"
	"flag"
	"path"

	"req/client"
	"req/options"
)

type Out struct {
	Code int `json:"code"`
	Time float64 `json:"time"`
	Body string `json:"body"`
	Output string `json:"output"`
}

func main() {
	var flagBody = flag.Bool("b", false, "stream the response body")
	var flagOutput = flag.String("o", "", "output body to a file")
	var flagOutputEasy = flag.Bool("O", false, "Output the body to file with no file name")

	flag.Parse()

	if *flagOutput != "" && *flagOutputEasy != false {
		panic("Error: -o and -O can not be used at the same time")
	}

	url := flag.Args()[0]

	code, body, time := client.Exec(url)

	out := Out{
		Code: code,
		Time: time.Seconds(),
		Body: "no print",
		Output: "no output",
	}

	if *flagBody != false {
		out.Body = body
	}

	if *flagOutput != "" {
		options.Output(*flagOutput, body)
		out.Output = *flagOutput
	} else if *flagOutputEasy != false {
		_ ,filename := path.Split(url)
		options.Output(filename, body)
		out.Output = filename
	}

	jsonOut, err := json.Marshal(out)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(jsonOut))
}
