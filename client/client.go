package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// BaseURL is the main franz api url
const (
	apiURL       string = "http://localhost:4005"
	jobsEndpoint string = apiURL + "/jobs"
)

// payload sent to franz api
type payload struct {
	Sha1          string `json:"sha1"`
	Branch        string `json:"branch"`
	CommitMessage string `json:"commit_message"`
	Author        string `json:"author"`
	RawInput      string `json:"raw_input"`

	// TODO: add resources
	// TODO: add runtime
	// TODO: add ci
}

// SendReport method sends the authenticated payload to the franz api
func SendReport(input []byte) error {

	p := payload{
		Sha1:          "sha1",
		Branch:        "branch",
		CommitMessage: "commit message",
		Author:        "author",
		RawInput:      base64.StdEncoding.EncodeToString(input),
	}

	client := http.Client{}

	j, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, jobsEndpoint, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")

	// TODO: add franz api key

	if err != nil {
		// TODO: handle
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("benchmark report sucessfully sent to franz")
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Print(string(b))
	} else {
		fmt.Println("failed")
	}

	return nil
}
