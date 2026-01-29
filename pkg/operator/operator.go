package operator

import (
	"fmt"
	"io"
	"net/http"
)

func ChangeServerStatus(serverIP, authToken, serverID, status string) (string, error) {

	baseURL := fmt.Sprintf("https://%s/api/servers/%s/", serverIP, serverID)

	var apiLink string
	var bodyReader io.Reader = nil

	switch status {
	case "on":
		fmt.Println("Turning server on")
		apiLink = baseURL + "start"
	case "off":
		fmt.Println("Turning server off")
		apiLink = baseURL + "stop"
	case "restart":
		fmt.Println("Restarting server")
		apiLink = baseURL + "restart"
	case "kill":
		fmt.Println("Killing server")
		apiLink = baseURL + "kill"
	default:
		return "", fmt.Errorf("invalid status: %s, use on, off, restart or kill", status)
	}

	req, err := http.NewRequest("POST", apiLink, bodyReader)
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
