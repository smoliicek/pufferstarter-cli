package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func GetAuthToken(clientID, clientSecret, serverIP string) (string, error) {

	grantType := "client_credentials"

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", grantType)

	bodyReader := strings.NewReader(data.Encode())

	req, err := http.NewRequest("POST", "https://"+serverIP+"/oauth2/token", bodyReader)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "pufferstarter-cli/4.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token request failed: %s", resp.Status)
	}

	var tr struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &tr); err != nil {
		return "", err
	}

	if tr.AccessToken == "" {
		return "", fmt.Errorf("access_token not found in response")
	}

	return tr.AccessToken, nil
}
