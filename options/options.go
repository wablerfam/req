package options

import (
	"os"
)

func Output(filename string, body string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	file.Write(([]byte)(body))

	defer file.Close()

	return
}