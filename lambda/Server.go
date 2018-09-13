// The main entry point for services of the lambda
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

// CodeResult is the json response of the code execution
type CodeResult struct {
	Result string
	Error  string
}

const (
	defaultPort = ":8080"
)

func main() {
	port := defaultPort
	if args := os.Args; len(args) > 1 {
		port = args[0]
	}

	http.HandleFunc("/lambda", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Error reading body: %s\n", err.Error())
			return
		}

		code := string(body)
		fmt.Printf("code: %s\n", code)

		cmd := exec.Command("/usr/bin/python3", "-c", code)

		stdOut, stdErr := new(bytes.Buffer), new(bytes.Buffer)
		cmd.Stdout = stdOut
		cmd.Stderr = stdErr
		cmd.Run()
		w.Write(stdOut.Bytes())
		return
	})

	http.ListenAndServe(port, nil)

}
