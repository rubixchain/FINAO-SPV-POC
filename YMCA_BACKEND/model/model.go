package model

import "time"

type User struct {
	UserID      int       `db:"user_Id primarykey" json:"user_Id"`
	Name        string    `db:"name" json:"name"`
	Email       string    `db:"email" json:"email"`
	Password    string    `db:"password" json:"password"`
	DateOfBirth time.Time `db:"date_of_birth" json:"date_of_birth"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Address     string    `db:"address" json:"address"`
	DID         string    `db:"did" json:"did"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// PublicData represents public data associated with a user.
type PublicData struct {
	PubDataID   int       `json:"pub_data_id" db:"pub_data_id primarykey"`
	FocusArea   string    `json:"focus_area" db:"focus_area"`
	Communities string    `json:"communities" db:"communities"`
	UserID      int       `json:"user_id" db:"user_id foreignkey"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// PrivateData represents private data associated with a user.
type PrivateData struct {
	PvtDataID  int       `json:"pvt_data_id" db:"pvt_data_id primarykey"`
	Capsule    string    `json:"capsule" db:"capsule"`
	CipherText string    `json:"cipher_text" db:"cipher_text"`
	UserID     int       `json:"user_id" db:"user_id foreignkey"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" db:"updated_at"`
}

// AccessSheet represents access information for private data.
type AccessSheet struct {
	AccessID      int       `json:"access_id" db:"access_id primarykey"`
	PvtDataID     int       `json:"pvt_data_id" db:"pvt_data_id foreignkey"`
	DecryptUserID int       `json:"decrypt_user_id" db:"decrypt_user_id"`
	OwnerUserID   int       `josn:"owner_user_id" db:"owner_user_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type SignUpRequest struct {
	Name        string `db:"name" json:"name"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	DateOfBirth string `db:"date_of_birth" json:"date_of_birth"`
	PhoneNumber string `db:"phone_number" json:"phone_number"`
}

type SignUpResponse struct {
	UserID  int    `json:"user_id"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type LogInRequest struct {
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}

type LogInResponse struct {
	UserID  int    `json:"user_id"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type PublicDataResponse struct {
	FocusArea   string `json:"focus_area" db:"focus_area"`
	Communities string `json:"communities" db:"communities"`
	UserID      int    `json:"user_id" db:"user_id"`
}

type PrivateDataResponse struct {
	Capsule    string `json:"capsule" db:"capsule"`
	CipherText string `json:"cipher_text" db:"cipher_text"`
	UserID     int    `json:"user_id" db:"user_id"`
}

type PublicDataInputReq struct {
	FocusArea   string `json:"focus_area" db:"focus_area"`
	Communities string `json:"communities" db:"communities"`
	UserID      int    `json:"user_id" db:"user_id foreignkey"`
}

type PrivateDataInputReq struct {
	FocusArea     string `json:"focus_area" db:"focus_area"`
	Communities   string `json:"communities" db:"communities"`
	UserID        int    `json:"user_id" db:"user_id foreignkey"`
	DecryptUserID int    `json:"decrypt_user_id" db:"decrypt_user_id"`
}

type PrivateDataEncrypt struct {
	FocusArea   string `json:"focus_area" db:"focus_area"`
	Communities string `json:"communities" db:"communities"`
}

type AddPrivateDataResponse struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	PvtDataID int    `json:"pvt_data_id" `
	AccessID  int    `json:"access_id"`
}

type AddPublicDataResponse struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	PubDataID int    `json:"pub_data_id"`
}

type BasicResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	UserID  int    `json:user_id"`
	DID     string `json:did"`
}

type PvtDataResponse struct {
	Status      bool          `json:"status"`
	Message     string        `json:"message"`
	PrivateData []PrivateData `json:"privateData"`
}

type SecretKeys struct {
	SecretKey string `json:"secretKey"`
	PublicKey string `json:"publicKey"`
}

type SecretKeyData struct {
	KeyID     int    `json:"key_id" db:"key_id primarykey"`
	SecretKey string `json:"secretKey" db:"secret_key"`
	PublicKey string `json:"publicKey" db:"public_key"`
	UserID    int    `json:"user_id" db:"user_id foreignkey"`
}

type EncryptionRequest struct {
	PublicKey string `json:"public_key"`
	Plaintext string `json:"plaintext"`
}

type EncryptionResponse struct {
	Capsule    string `json:"capsule"`
	Ciphertext string `json:"ciphertext"`
}

type DecryptDataRequest struct {
	Capsule    string `json:"capsule"`
	Ciphertext string `json:"ciphertext"`
	UserID     int    `json:"user_id"`
}

type DecryptDataResponse struct {
	FocusArea   string `json:"focus_area"`
	Communities string `json:"communities"`
}

type DecryptServerInput struct {
	Capsule    string `json:"capsule"`
	Ciphertext string `json:"ciphertext"`
	SecretKey  string `json:"secretKey"`
}

type DecryptServerResponse struct {
	PlainText string `json:"plaintext"`
}
