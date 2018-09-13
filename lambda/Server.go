// The main entry point for services of the lambda
package lambda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
)

// Result is the json response of the code execution with any error
type Result struct {
	Result string `json:"codeValue"`
	Error  string `json:"error"`
}

func LambdaEndpointHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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

		result := Result{
			Result: stdOut.String(),
			Error:  stdErr.String(),
		}

		response, err := json.Marshal(result)
		if err != nil {
			fmt.Println("Something went wrong with marshaling")
			return
		}

		w.Write(response)
		return
	}
}
