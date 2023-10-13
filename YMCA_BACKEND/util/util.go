package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	"YMCA_BACKEND/model"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// Generate a salt and hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// verifyPassword verifies if the provided password matches the stored hash.
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func CreateDID() (string, error) {
	cmd := exec.Command("curl", "-L", "-X", "POST", "http://localhost:20000/api/createdid", "-F", "did_config={\"type\":0,\"dir\":\"\",\"config\":\"\",\"master_did\":\"\",\"secret\":\"My DID Secret\",\"priv_pwd\":\"mypassword\",\"quorum_pwd\":\"mypassword\",\"img_file\":\"image.png\",\"did_img_file\":\"\",\"pub_img_file\":\"\",\"priv_img_file\":\"\",\"pub_key_file\":\"\",\"priv_key_file\":\"\",\"quorum_pub_key_file\":\"\",\"quorum_priv_key_file\":\"\"}")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("Error running curl: %v", err)
	}

	// Split the output by newline to find the JSON part
	lines := strings.Split(string(output), "\n")

	var jsonResponse string
	for _, line := range lines {
		if strings.HasPrefix(line, "{") {
			jsonResponse = line
			break
		}
	}

	if jsonResponse == "" {
		return "", fmt.Errorf("No JSON response found in curl output")
	}

	// Parse the JSON response
	var response struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
		Result  struct {
			DID string `json:"did"`
		} `json:"result"`
	}

	err = json.NewDecoder(bytes.NewReader([]byte(jsonResponse))).Decode(&response)
	if err != nil {
		return "", fmt.Errorf("Error parsing JSON response: %v", err)
	}

	if !response.Status {
		return "", fmt.Errorf("DID creation failed: %s", response.Message)
	}

	if response.Result.DID == "" {
		return "", fmt.Errorf("DID creation succeeded, but the DID value is empty")
	}

	return response.Result.DID, nil
}

func GenerateSecretKeys() (model.SecretKeys, error) {
	url := "http://localhost:3000/generate-secret-key"

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return model.SecretKeys{}, err
	}

	req.Header.Set("accept", "*/*")

	resp, err := client.Do(req)
	if err != nil {
		return model.SecretKeys{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return model.SecretKeys{}, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	// Read the response body and unmarshal it into the SecretKeys struct.
	var secretKeys model.SecretKeys
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.SecretKeys{}, err
	}

	err = json.Unmarshal(body, &secretKeys)
	if err != nil {
		return model.SecretKeys{}, err
	}

	return secretKeys, nil
}

func CombinePrivateData(data model.PrivateDataEncrypt) (string, error) {
	// Convert the struct to a JSON representation.
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// Convert the JSON data to a string.
	dataString := string(jsonData)

	return dataString, nil
}

func EncryptData(publicKey, plaintext string) (model.EncryptionResponse, error) {
	requestURL := "http://localhost:3000/encrypt"

	// Create an EncryptionRequest object
	requestData := model.EncryptionRequest{
		PublicKey: publicKey,
		Plaintext: plaintext,
	}

	// Convert the request data to a JSON payload
	payload, err := json.Marshal(requestData)
	if err != nil {
		return model.EncryptionResponse{}, err
	}

	// Send the HTTP POST request
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return model.EncryptionResponse{}, err
	}
	defer resp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return model.EncryptionResponse{}, err
	}

	var response model.EncryptionResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return model.EncryptionResponse{}, err
	}

	return response, nil
}
