package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type ContractInputRequest struct {
	Port              string `json:"port"`
	SmartContractHash string `json:"smart_contract_hash"` //port should also be added here, so that the api can understand which node.
}

type RubixResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type SmartContractInput struct {
	Did            string `json:"did"`
	BinaryCodePath string `json:"binaryCodePath"`
	RawCodePath    string `json:"rawCodePath"`
	SchemaFilePath string `json:"schemaFilePath"`
	Port           string `json:"port"`
}

type DeploySmartContractInput struct {
	Comment            string `json:"comment"`
	DeployerAddress    string `json:"deployerAddress"`
	QuorumType         int    `json:"quorumType"`
	RbtAmount          int    `json:"rbtAmount"`
	SmartContractToken string `json:"smartContractToken"`
	Port               string `json:"port"`
}

type ExecuteSmartContractInput struct {
	Comment            string `json:"comment"`
	ExecutorAddress    string `json:"executorAddress"`
	QuorumType         int    `json:"quorumType"`
	SmartContractData  string `json:"smartContractData"`
	SmartContractToken string `json:"smartContractToken"`
	Port               string `json:"port"`
}

type GetSmartContractDataInput struct {
	Port  string `json:"port"`
	Token string `json:"token"`
}

func GenerateSmartContract(w http.ResponseWriter, r *http.Request) {
	var inputReq SmartContractInput
	err1 := json.NewDecoder(r.Body).Decode(&inputReq) //decode request into struct
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Error reading response body: %s\n", err1)
		return
	}

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	// Add the form fields
	_ = writer.WriteField("did", inputReq.Did)

	// Add the binaryCodePath field
	file, _ := os.Open(inputReq.BinaryCodePath)
	defer file.Close()
	binaryPart, _ := writer.CreateFormFile("binaryCodePath", inputReq.BinaryCodePath)
	_, _ = io.Copy(binaryPart, file)

	// Add the rawCodePath field
	rawFile, _ := os.Open(inputReq.RawCodePath)
	defer rawFile.Close()
	rawPart, _ := writer.CreateFormFile("rawCodePath", inputReq.RawCodePath)
	_, _ = io.Copy(rawPart, rawFile)

	// Add the schemaFilePath field
	schemaFile, _ := os.Open(inputReq.SchemaFilePath)
	defer schemaFile.Close()
	schemaPart, _ := writer.CreateFormFile("schemaFilePath", inputReq.SchemaFilePath)
	_, _ = io.Copy(schemaPart, schemaFile)

	// Close the writer
	writer.Close()

	// Create the HTTP request
	url := fmt.Sprintf("http://localhost:%s/api/generate-smart-contract", inputReq.Port)
	req, _ := http.NewRequest("POST", url, &requestBody)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	_, _ = io.Copy(binaryPart, file)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	var response RubixResponse
	data2, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}
	err2 := json.Unmarshal([]byte(data2), &response)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}

	// Process the data as needed
	fmt.Println("Response Body in Generate Smart Contract :", string(data2))

	// Process the response as needed
	json.NewEncoder(w).Encode(response)
}

func DeploySmartContract(w http.ResponseWriter, r *http.Request) {
	var inputReq DeploySmartContractInput
	err1 := json.NewDecoder(r.Body).Decode(&inputReq) //decode request into struct
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Error reading response body: %s\n", err1)
		return
	}
	data := map[string]interface{}{
		"comment":            inputReq.Comment,
		"deployerAddr":       inputReq.DeployerAddress,
		"quorumType":         inputReq.QuorumType,
		"rbtAmount":          inputReq.RbtAmount,
		"smartContractToken": inputReq.SmartContractToken,
	}
	fmt.Println("inputReq.Port =", inputReq.Port)
	bodyJSON, err := json.Marshal(data)
	fmt.Println(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
	}
	url := fmt.Sprintf("http://localhost:%s/api/deploy-smart-contract", inputReq.Port)
	fmt.Println("inputReq.Port =", inputReq.Port)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
	}
	fmt.Println("Response Status:", resp.Status)
	data2, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
	}
	// Process the data as needed
	fmt.Println("Response Body in deploy smart contract:", string(data2))
	var response map[string]interface{}
	err3 := json.Unmarshal(data2, &response)
	if err3 != nil {
		fmt.Println("Error unmarshaling response:", err3)
	}

	result := response["result"].(map[string]interface{})
	id := result["id"].(string)

	defer resp.Body.Close()
	SignatureResponse := SignatureResponse(id, inputReq.Port)
	fmt.Println("Signature Response:", SignatureResponse)
	var finalResponse RubixResponse
	err2 := json.Unmarshal([]byte(SignatureResponse), &finalResponse)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	json.NewEncoder(w).Encode(finalResponse)
}

func ExecuteSmartContract(w http.ResponseWriter, r *http.Request) {
	var inputReq ExecuteSmartContractInput
	err1 := json.NewDecoder(r.Body).Decode(&inputReq) //decode request into struct
	if err1 != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Printf("Error reading response body: %s\n", err1)
		return
	}
	data := map[string]interface{}{
		"comment":            inputReq.Comment,
		"executorAddr":       inputReq.ExecutorAddress,
		"quorumType":         inputReq.QuorumType,
		"smartContractData":  inputReq.SmartContractData,
		"smartContractToken": inputReq.SmartContractToken,
	}
	bodyJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	url := fmt.Sprintf("http://localhost:%s/api/execute-smart-contract", inputReq.Port)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}
	fmt.Println("Response Status:", resp.Status)
	data2, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		return
	}
	// Process the data as needed
	fmt.Println("Response Body in execute smart contract :", string(data2))
	var response map[string]interface{}
	err3 := json.Unmarshal(data2, &response)
	if err3 != nil {
		fmt.Println("Error unmarshaling response:", err3)
	}

	result := response["result"].(map[string]interface{})
	id := result["id"].(string)
	SignatureResponse := SignatureResponse(id, inputReq.Port)
	fmt.Println("Signature Response:", SignatureResponse)
	var finalResponse RubixResponse
	err2 := json.Unmarshal([]byte(SignatureResponse), &finalResponse)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	json.NewEncoder(w).Encode(finalResponse)

	defer resp.Body.Close()

}

func SignatureResponse(requestId string, port string) string {
	data := map[string]interface{}{
		"id":       requestId,
		"mode":     0,
		"password": "mypassword",
	}

	bodyJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		//	return
	}
	url := fmt.Sprintf("http://localhost:%s/api/signature-response", port)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJSON))
	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		//return
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		//return
	}
	fmt.Println("Response Status:", resp.Status)
	data2, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %s\n", err)
		//return
	}
	// Process the data as needed
	fmt.Println("Response Body in signature response :", string(data2))
	//json encode string
	defer resp.Body.Close()
	return string(data2)
}
