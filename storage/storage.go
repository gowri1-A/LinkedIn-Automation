package storage

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

type Data struct {
	SentConnections []string          `json:"sent_connections"`
	SentMessages    map[string]string `json:"sent_messages"`
}

var FileName = "storage.json"

// Load data from storage.json
func Load() (*Data, error) {
	data := &Data{
		SentConnections: []string{},
		SentMessages:    map[string]string{},
	}

	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		return data, nil // File does not exist, return empty
	}

	file, err := os.ReadFile(FileName)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(file, data); err != nil {
		return nil, err
	}

	return data, nil
}

// Save data to storage.json
func Save(data *Data) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(FileName, file, 0644); err != nil {
		return err
	}

	return nil
}

// Add a sent connection
func AddConnection(profile string) error {
	data, _ := Load()
	for _, p := range data.SentConnections {
		if p == profile {
			return nil // already sent
		}
	}
	data.SentConnections = append(data.SentConnections, profile)
	return Save(data)
}

// Add a sent message
func AddMessage(profile string, message string) error {
	data, _ := Load()
	data.SentMessages[profile] = message
	return Save(data)
}

// Check if connection already sent
func IsConnectionSent(profile string) bool {
	data, _ := Load()
	for _, p := range data.SentConnections {
		if p == profile {
			return true
		}
	}
	return false
}

func SaveCookies(page *rod.Page, path string) error {
	// Get all cookies
	cookies, err := proto.NetworkGetAllCookies{}.Call(page)
	if err != nil {
		return err
	}

	// Save as JSON
	data, err := json.MarshalIndent(cookies.Cookies, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadCookies(page *rod.Page, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var cookies []proto.NetworkCookie
	if err := json.Unmarshal(data, &cookies); err != nil {
		return err
	}

	// Convert to NetworkCookieParam
	var params []*proto.NetworkCookieParam
	for _, c := range cookies {
		params = append(params, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure:   c.Secure,
			SameSite: c.SameSite,
		})
	}

	// Set cookies
	return proto.NetworkSetCookies{Cookies: params}.Call(page)
}
