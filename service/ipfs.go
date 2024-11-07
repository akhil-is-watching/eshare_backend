package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akhil-is-watching/encryptedFileSharing/types"
)

type PinataIPFS struct {
	apiKey    string
	apiSecret string
	endpoint  string
	gateway   string
}

func NewPinataIPFS(apiKey, apiSecret string) *PinataIPFS {
	return &PinataIPFS{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		endpoint:  "https://api.pinata.cloud/pinning/pinJSONToIPFS",
		gateway:   "https://rose-international-cicada-831.mypinata.cloud/ipfs/",
	}
}

func (p *PinataIPFS) Publish(document types.DocumentPackage) (string, error) {
	// Create a request body with pinata metadata
	requestBody := map[string]interface{}{
		"pinataContent": document,
		"pinataMetadata": map[string]interface{}{
			"name": document.Metadata.FileName,
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal document: %w", err)
	}

	req, err := http.NewRequest("POST", p.endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("pinata_api_key", p.apiKey)
	req.Header.Set("pinata_secret_api_key", p.apiSecret)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("pinata returned non-200 status: %d", resp.StatusCode)
	}

	var pinataResp struct {
		IpfsHash string `json:"IpfsHash"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&pinataResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return pinataResp.IpfsHash, nil
}

func (p *PinataIPFS) Retrieve(cid string) (*types.DocumentPackage, error) {
	url := p.gateway + cid

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("gateway returned non-200 status: %d", resp.StatusCode)
	}

	var document types.DocumentPackage
	if err := json.NewDecoder(resp.Body).Decode(&document); err != nil {
		return nil, fmt.Errorf("failed to decode document: %w", err)
	}

	return &document, nil
}
