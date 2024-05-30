package flutterwave

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func marshallToJSON(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("%v: %v", errors.New("error marshalling data to json"), err)
	}

	return jsonData, nil
}

func (f *Flutterwave) buildPostRequest(jsonData []byte, url string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error building post request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+f.secretKey)

	return req, nil
}

func (f *Flutterwave) sendRequest(reqData interface{}, url string) (*http.Response, error) {
	postData, err := marshallToJSON(reqData)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	req, err := f.buildPostRequest(postData, url)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request %v", err)
	}

	return resp, nil
}

func (f *Flutterwave) readJSONResponse(response *http.Response, target interface{}) error {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}
	err = json.Unmarshal(body, target)
	if err != nil {
		return fmt.Errorf("error unmarshalling response body: %v", err)
	}
	return nil
}
