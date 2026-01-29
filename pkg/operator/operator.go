package operator

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ChangeServerStatus(serverIP, authToken, serverID, status string) (string, error) {

	baseURL := fmt.Sprintf("https://%s/api/servers/%s/", serverIP, serverID)

	var apiLink string
	var bodyReader io.Reader = nil

	switch status {
	case "on":
		apiLink = baseURL + "start"
	case "off":
		apiLink = baseURL + "stop"
	case "restart":
		apiLink = baseURL + "restart"
	case "kill":
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
	fmt.Printf("Status: %d\n", resp.StatusCode)

	os.Exit(255)

	return string(body), nil
}
