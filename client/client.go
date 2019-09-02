package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kinche/franz-agent/ci"
)

// BaseURL is the main franz api url
const (
	apiURL       string = "http://franz-api-staging.kinche.co"
	jobsEndpoint string = apiURL + "/jobs"
)

// payload sent to franz api
type payload struct {
	Sha1          string `json:"sha1"`
	Branch        string `json:"branch"`
	CommitMessage string `json:"commit_message"`
	Author        string `json:"author"`
	RawInput      string `json:"raw_input"`
	CI            string `json:"ci"`

	// TODO: add resources
	// TODO: add runtime
}

// SendReport method sends the authenticated payload to the franz api
func SendReport(input []byte, cinfo ci.Info) error {

	p := payload{
		Sha1:          cinfo.Sha1,
		Branch:        cinfo.Branch,
		CI:            cinfo.CI,
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
