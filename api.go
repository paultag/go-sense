package sense

import (
	"fmt"
	"net/http"

	"encoding/json"
)

func (s Sense) Request(method, path string) (*http.Response, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", s.serverBase, path), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", s.userAgent)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.token.AccessToken))
	client := &http.Client{}
	return client.Do(req)
}

func (s Sense) Timeline(when string) (*Timeline, error) {
	resp, err := s.Request("GET", fmt.Sprintf("v2/timeline/%s", when))
	if err != nil {
		return nil, err
	}
	timeline := Timeline{}
	if err := json.NewDecoder(resp.Body).Decode(&timeline); err != nil {
		return nil, err
	}
	return &timeline, nil
}
