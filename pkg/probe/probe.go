package probe

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func GetServerStatus(serverIP, authToken, serverID string) (bool, bool, error) {
	apiLink, err := url.JoinPath("https://"+serverIP, "api", "servers", serverID, "status")
	if err != nil {
		return false, false, err
	}

	req, err := http.NewRequest("GET", apiLink, nil)
	if err != nil {
		return false, false, err
	}

	req.Header.Set("User-Agent", "pufferstarter-cli/4.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return false, false, fmt.Errorf("request failed: %s", resp.Status)
	}

	var status struct {
		Installing bool `json:"installing"`
		Running    bool `json:"running"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return false, false, err
	}

	return status.Running, status.Installing, nil
}

func GetAllServers(serverIP, authToken string) (string, error) {
	apiLink, err := url.JoinPath("https://"+serverIP, "api", "servers")
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", apiLink, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "pufferstarter-cli/4.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("request failed: %s", resp.Status)
	}

	var data struct {
		Servers []map[string]interface{} `json:"servers"`
		Paging  interface{}              `json:"paging"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	for _, s := range data.Servers {
		id, ok := s["id"].(string)
		if ok {
			running, installing, _ := GetServerStatus(serverIP, authToken, id)
			s["running"] = running
			s["installing"] = installing
		}
	}

	enrichedBody, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(enrichedBody), nil
}

func GetServerStats(serverIP, authToken, serverID string) (map[string]interface{}, error) {
	apiLink, err := url.JoinPath("https://"+serverIP, "api", "servers", serverID, "stats")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", apiLink, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "pufferstarter-cli/4.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request failed: %s", resp.Status)
	}

	var stats map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, err
	}

	return stats, nil
}

func GetServer(serverIP, authToken, serverID string) (string, error) {
	apiLink, err := url.JoinPath("https://"+serverIP, "api", "servers", serverID)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("GET", apiLink, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "pufferstarter-cli/4.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("request failed: %s", resp.Status)
	}

	var data struct {
		Server map[string]interface{} `json:"server"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	running, installing, _ := GetServerStatus(serverIP, authToken, serverID)
	data.Server["running"] = running
	data.Server["installing"] = installing

	stats, _ := GetServerStats(serverIP, authToken, serverID)
	data.Server["stats"] = stats

	enrichedBody, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	return string(enrichedBody), nil
}
