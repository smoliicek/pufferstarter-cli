package probe

import (
	"fmt"
	"io"
	"net/http"
)

func GetAllServers(serverIP, authToken string) (string, error) {
	apiLink := fmt.Sprintf("https://%s/api/servers/", serverIP)
	var bodyReader io.Reader = nil

	req, err := http.NewRequest("GET", apiLink, bodyReader)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "pufferstarter-cli/4.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+authToken)

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

	if resp.StatusCode != 200 && resp.StatusCode != 202 && resp.StatusCode != 204 {
		return "", fmt.Errorf("request failed: %s", resp.Status)
	}

	if resp.StatusCode == 400 {
		return "", fmt.Errorf("bad request: %s", body)
	}

	if resp.StatusCode == 401 {
		return "", fmt.Errorf("unauthorized: %s", body)
	}

	return string(body), nil
}

func GetServer(serverIP, authToken, serverID string) (string, error) {
	apiLink := fmt.Sprintf("https://%s/api/servers/%s", serverIP, serverID)

	var bodyReader io.Reader = nil

	req, err := http.NewRequest("GET", apiLink, bodyReader)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "pufferstarter-cli/4.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Bearer "+authToken)

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

	if resp.StatusCode != 200 && resp.StatusCode != 202 && resp.StatusCode != 204 {
		return "", fmt.Errorf("request failed: %s", resp.Status)
	}

	if resp.StatusCode == 400 {
		return "", fmt.Errorf("bad request: %s", body)
	}

	if resp.StatusCode == 401 {
		return "", fmt.Errorf("unauthorized: %s", body)
	}

	return string(body), nil
}
