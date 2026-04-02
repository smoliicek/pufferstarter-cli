package operator

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func ChangeServerStatus(serverIP, authToken, serverID, status string) (string, error) {

	var action string
	switch status {
	case "on":
		fmt.Println("Turning server on")
		action = "start"
	case "off":
		fmt.Println("Turning server off")
		action = "stop"
	case "restart":
		fmt.Println("Restarting server")
		action = "restart"
	case "kill":
		fmt.Println("Killing server")
		action = "kill"
	default:
		return "", fmt.Errorf("invalid status: %s, use on, off, restart or kill", status)
	}

	apiLink, err := url.JoinPath("https://"+serverIP, "api", "servers", serverID, action)
	if err != nil {
		return "", err
	}

	var bodyReader io.Reader = nil

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
