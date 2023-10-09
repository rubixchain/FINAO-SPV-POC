package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

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

func main() {
	did, err := CreateDID()
	if err != nil {
		fmt.Printf("Error creating DID: %v\n", err)
		return
	}

	fmt.Printf("Created DID: %s\n", did)
}
