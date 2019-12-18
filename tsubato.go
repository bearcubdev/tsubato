package tsubato

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const modifiedSinceFormat string = "Mon, 2 Jan 2006 15:04:05 MST"

var (
	previousRequests = map[string]time.Time{}
	lastRequest      = time.Time{}
)

func hasQueriedRecently(url string) bool {
	elapsed := time.Now().Sub(lastRequest)
	if elapsed.Seconds() >= 1 {
		lastRequest = time.Now()
		return false
	}
	return true
}

func get(queryPath string) ([]byte, error) {
	url := fmt.Sprintf(`https://a.4cdn.org%s`, queryPath)
	if hasQueriedRecently(url) {
		return nil, fmt.Errorf("%s requested too recently", url)
	}
	lastRequestTime, hasRequestedBefore := previousRequests[url]

	var client http.Client
	req, err := http.NewRequest("GET", url, nil)
	if hasRequestedBefore {
		req.Header.Add("If-Modified-Since", lastRequestTime.Format(modifiedSinceFormat))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return nil, err
		}
		return bodyBytes, nil
	}
	return nil, fmt.Errorf("Error: %d", resp.StatusCode)
}
