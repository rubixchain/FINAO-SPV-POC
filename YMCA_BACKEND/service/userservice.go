package service

import (
	"YMCA_BACKEND/model"
	"YMCA_BACKEND/util"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// signup Return user data
// @Summary Return user data
// @Description This endpoint is used to when new user signs up
// @Accept json
// @Produce json
// @Param user body model.SignUpRequest true "enter email and phone number"
// @Success 200 {object} model.SignUpResponse
// @Router /signup [post]
func (s *Service) SignUp(w http.ResponseWriter, r *http.Request) {
	var signUpReq model.SignUpRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&signUpReq); err != nil {
		http.Error(w, "Failed to parse JSON request: "+err.Error(), http.StatusBadRequest)
		return
	}

	signUpRes := &model.SignUpResponse{
		Status: false,
	}

	dateParts := strings.Split(signUpReq.DateOfBirth, "-")
	if len(dateParts) != 3 {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	day, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	year, _ := strconv.Atoi(dateParts[2])

	hashedPassword, err := util.HashPassword(signUpReq.Password)
	if err != nil {
		http.Error(w, "error storing details", http.StatusBadRequest)
		return
	}

	newUser := &model.User{
		Name:        signUpReq.Name,
		Email:       signUpReq.Email,
		Password:    hashedPassword,
		DateOfBirth: time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
		PhoneNumber: signUpReq.PhoneNumber,
	}

	userID, err := s.storage.InsertUser(*newUser)
	if err != nil {
		s.log.Println("Error creating new entry for user in DB, ,", err)
		signUpRes.Message = "Error creating user entry" + err.Error()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	signUpRes.Status = true
	signUpRes.UserID = userID
	signUpRes.Message = "User created successfully"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(signUpRes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go func() {
		result, err := util.CreateDID()
		if err != nil {
			// Handle errors from createDID if needed
			s.log.Println("Error creating DID:", err)
		} else {
			s.log.Println("DID created successfully:", result)
			// You can take further actions based on the result here
		}
		//fmt.Println(result)
		err = s.storage.UpdateUserDID(userID, result)
		if err != nil {
			s.log.Println("Failed to update DID for user ", userID)
		}

	}()

}

// login Return user data
// @Summary Return user data
// @Description This endpoint is used to authenticate existing user log in
// @Accept json
// @Produce json
// @Param user body model.LogInRequest true "enter email and password"
// @Success 200 {object} model.LogInResponse
// @Router /login [post]
func (s *Service) LogIn(w http.ResponseWriter, r *http.Request) {
	var logInReq model.LogInRequest

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&logInReq); err != nil {
		http.Error(w, "Failed to parse JSON request: "+err.Error(), http.StatusBadRequest)
		return
	}

	loginRes := &model.LogInResponse{
		Status: false,
	}

	userDetails, err := s.storage.GetUserByEmail(logInReq.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = util.VerifyPassword(userDetails.Password, logInReq.Password)
	if err != nil {
		loginRes.Message = "Password Does not match"
	} else {
		loginRes.Status = true
		loginRes.Message = "User authenticated successfully"
		loginRes.UserID = userDetails.UserID
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(loginRes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// @Summary Return user public data by DID
// @Description Get public data for a user by their DID
// @Accept json
// @Produce json
// @Param did query string true "User's DID"
// @Success 200 {object} DataResponse
// @Router /getAllPublicDataByDID [get]
func (s *Service) GetAllPublicDataByDID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	did := vars["did"]

	// TODO: Query the database to fetch public data for the user with the given DID
	// Replace the following line with your database query logic
	publicData, err := s.storage.GetPublicDataByDID(did)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a PublicDataResponse from the retrieved data
	response := &model.DataResponse{
		FocusArea:   publicData.FocusArea,
		Communities: publicData.Communities,
		UserID:      publicData.UserID,
		DID:         did,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// @Summary Return user private data by DID
// @Description Get private data for a user by their DID
// @Accept json
// @Produce json
// @Param did query string true "User's DID"
// @Success 200 {object} DataResponse
// @Router /getAllPrivateDataByDID [get]
func (s *Service) GetAllPrivateDataByDID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	did := vars["did"]

	// TODO: Query the database to fetch public data for the user with the given DID
	// Replace the following line with your database query logic
	publicData, err := s.storage.GetPrivateDataByDID(did)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a PublicDataResponse from the retrieved data
	response := &model.DataResponse{
		FocusArea:   publicData.FocusArea,
		Communities: publicData.Communities,
		UserID:      publicData.UserID,
		DID:         did,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// @Summary Return user private data that has been given access to a  DID
// @Description Get rivate data that has been given access to a  DID
// @Accept json
// @Produce json
// @Param did query string true "User's DID"
// @Success 200 {object} DataResponse
// @Router /getAllAccessDatabyDID [get]
func (s *Service) GetAllAccessDataByDID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	did := vars["did"]

	// Query the database to fetch private data for the user with the given DID
	privateDataList, err := s.storage.GetAllAccessDataByDID(did)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a response slice for private data
	var responseList []model.DataResponse

	// Convert the retrieved private data into the desired response format
	for _, privateData := range privateDataList {
		response := model.DataResponse{
			FocusArea:   privateData.FocusArea,
			Communities: privateData.Communities,
			UserID:      privateData.UserID,
			DID:         did,
		}
		responseList = append(responseList, response)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseList)
}
