package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Authentication struct {
	Email    string
	Password string
}

func Request[T any](endpoint string, method string, body map[string]string, authentication Authentication) (T, error) {
	client := &http.Client{}
	var payload io.Reader
	var result T
	if method != http.MethodGet {
		bodyJson, error := json.Marshal(body)
		if error != nil {
			return result, fmt.Errorf("Error coverting body to JSON %s", error)
		}
		payload = bytes.NewBuffer(bodyJson)
	}
	request, error := http.NewRequest(method, endpoint, payload)
	request.Header.Add("Content-Type", "application/json")

	request.SetBasicAuth(authentication.Email, authentication.Password)

	if error != nil {
		return result, fmt.Errorf("Error creating request %s", error)
	}
	response, error := client.Do(request)
	if error != nil {
		fmt.Println(error)
		return result, fmt.Errorf("Error sending request %s", error)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return result, fmt.Errorf("Response returned with error code status %s", response.Status)
	}
	responseBody, error := io.ReadAll(response.Body)

	if error != nil {
		return result, fmt.Errorf("Error reading response %s", error)
	}

	error = json.Unmarshal(responseBody, &result)

	if error != nil {
		return result, fmt.Errorf("Error converting response %s", error)
	}

	return result, nil
}
