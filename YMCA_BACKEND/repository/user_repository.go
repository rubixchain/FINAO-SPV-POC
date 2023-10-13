package repository

import (
	"YMCA_BACKEND/model"
	"database/sql"
	"fmt"
	"time"
)

// InsertUser inserts a new user into the database and returns the ID of the created user.
func (repository *Repository) InsertUser(user model.User) (int, error) {
	stmt := `INSERT INTO user (name,email, password, date_of_birth, phone_number, address, did, created_at, updated_at)
            VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	now := time.Now().UTC()
	result, err := repository.db.Exec(stmt, user.Name, user.Email, user.Password, user.DateOfBirth, user.PhoneNumber, user.Address, user.DID, now, now)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(userID), nil
}

// GetUserByEmail retrieves a user from the database based on their email address.
func (repository *Repository) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	// Define the SQL query to retrieve the user based on email
	query := "SELECT * FROM user WHERE email = ?"

	// Execute the query and scan the result into the user struct
	err := repository.db.QueryRow(query, email).Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.DateOfBirth,
		&user.PhoneNumber,
		&user.Address,
		&user.DID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	// Check for errors
	if err != nil {
		if err == sql.ErrNoRows {
			// Handle the case where no user with the given email was found
			return user, fmt.Errorf("user not found")
		}
		// Handle other database query errors
		return user, err
	}

	// User with the provided email was found and is stored in the 'user' variable
	return user, nil
}

func (repository *Repository) UpdateUserDID(userID int, did string) error {
	// Define the SQL query to update the DID for the user with the given userID
	query := "UPDATE user SET did = ? WHERE user_Id = ?"

	// Execute the SQL query with the provided arguments
	_, err := repository.db.Exec(query, did, userID)
	if err != nil {
		return fmt.Errorf("failed to update user's DID: %v", err)
	}

	return nil
}

// GetPublicDataByDID fetches the public data of a user by their DID, sorted by created time in descending order.
func (repo *Repository) GetPublicDataByID(userID int) (*model.PublicData, error) {
	// Define the SQL query to fetch public data by DID and sort by created time in descending order
	query := `
		SELECT pub_data_id, focus_area, communities, user_id, created_at, updated_at
		FROM publicdata
		WHERE user_id = ?
		ORDER BY created_at DESC;`

	// Execute the query and retrieve the public data
	var publicData model.PublicData
	err := repo.db.QueryRow(query, userID).Scan(
		&publicData.PubDataID,
		&publicData.FocusArea,
		&publicData.Communities,
		&publicData.UserID,
		&publicData.CreatedAt,
		&publicData.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("public data not found for user ID: %s", userID)
		}
		return nil, err
	}

	return &publicData, nil
}

// GetPrivateataByDID fetches the public data of a user by their DID, sorted by created time in descending order.
func (repo *Repository) GetPrivateDataByID(userID int) (*model.PrivateData, error) {
	// Define the SQL query to fetch public data by DID and sort by created time in descending order
	query := `
		SELECT pvt_data_id, focus_area, communities, user_id, created_at, updated_at
		FROM privatedata
		WHERE user_id = ?
		ORDER BY created_at DESC;`

	// Execute the query and retrieve the public data
	var privateData model.PrivateData
	err := repo.db.QueryRow(query, userID).Scan(
		&privateData.PvtDataID,
		&privateData.Capsule,
		&privateData.CipherText,
		&privateData.UserID,
		&privateData.CreatedAt,
		&privateData.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("public data not found for userID : %s", userID)
		}
		return nil, err
	}

	return &privateData, nil
}

// GetAllAccessDataByDID retrieves private data by DID.
func (repo *Repository) GetAllAccessDataByID(userID int) ([]model.PrivateData, error) {
	// Define the SQL query to fetch private data by DID
	query := `
		SELECT pd.pvt_data_id, pd.capsule, pd.cipher_text, pd.user_id, pd.created_at, pd.updated_at
		FROM privatedata pd
		INNER JOIN accesssheet as acs ON pd.pvt_data_id = acs.pvt_data_id
		INNER JOIN users u ON acs.decrypt_user_id = u.user_id
		WHERE u.user_id = ?;`

	// Execute the query and retrieve the private data
	rows, err := repo.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var privateDataList []model.PrivateData

	// Iterate through the rows and populate the private data list
	for rows.Next() {
		var privateData model.PrivateData
		err := rows.Scan(
			&privateData.PvtDataID,
			&privateData.Capsule,
			&privateData.CipherText,
			&privateData.UserID,
			&privateData.CreatedAt,
			&privateData.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		privateDataList = append(privateDataList, privateData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return privateDataList, nil
}

// AddPublicData inserts a new public data entry into the PublicData table.
func (r *Repository) AddPublicData(data *model.PublicData) (int, error) {
	// Define the SQL query to insert a new public data entry.
	query := `
        INSERT INTO publicdata (focus_area, communities, user_id, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?)
    `

	// Execute the SQL query to insert the new public data entry.
	result, err := r.db.Exec(
		query,
		data.FocusArea,
		data.Communities,
		data.UserID,
		time.Now(), // CreatedAt
		time.Now(), // UpdatedAt
	)

	pubDataID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(pubDataID), nil
}

// AddPrivateData inserts a new private data entry into the PrivateData table.
func (r *Repository) AddPrivateData(data *model.PrivateData) (int, error) {
	// Define the SQL query to insert a new private data entry.
	query := `
        INSERT INTO privatedata (capsule, cipher_text, user_id, created_at, updated_at)
        VALUES (?, ?, ?, ?, ?)
    `

	// Execute the SQL query to insert the new private data entry.
	result, err := r.db.Exec(
		query,
		data.Capsule,
		data.CipherText,
		data.UserID,
		time.Now(), // CreatedAt
		time.Now(), // UpdatedAt
	)
	if err != nil {
		return 0, err
	}

	pvtDataID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	r.log.Println("pvtDataID ", pvtDataID)
	return int(pvtDataID), nil
}

func (r *Repository) AddAccess(accessSheet *model.AccessSheet) (int, error) {
	// Define the SQL query to insert an access sheet entry
	query := `
		INSERT INTO accesssheet (pvt_data_id, decrypt_user_id)
		VALUES (?, ?);
	`

	// Execute the query to insert the access sheet entry
	result, err := r.db.Exec(query, accessSheet.PvtDataID, accessSheet.DecryptUserID)
	if err != nil {
		return 0, err
	}

	// Retrieve the last inserted ID (access ID)
	accessID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(accessID), nil
}

func (r *Repository) GetUserIDByDID(did string) (int, error) {
	// Define the SQL query to retrieve the user ID by DID
	r.log.Println("did ", did)
	query := "SELECT user_id FROM user WHERE did = ?;"

	var userID int
	err := r.db.QueryRow(query, did).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *Repository) GetDIDByUserID(userID int) (string, error) {
	// Define the SQL query to retrieve the DID by user ID
	query := "SELECT did FROM user WHERE user_id = ?;"

	var did string
	err := r.db.QueryRow(query, userID).Scan(&did)
	if err != nil {
		return "", err
	}

	return did, nil
}

func (r *Repository) GetPvtDataByUserID(userID int) ([]model.PrivateData, error) {
	// Define the SQL query to fetch private data by user ID
	query := `
        SELECT pd.pvt_data_id, pd.capsule, pd.cipher_text
        FROM privatedata pd
        INNER JOIN accesssheet as ON pd.pvt_data_id = as.pvt_data_id
        WHERE as.decrypt_user_id = ?;`

	// Execute the query and retrieve the private data
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var privateDataList []model.PrivateData

	// Iterate through the rows and populate the private data list
	for rows.Next() {
		var privateData model.PrivateData
		err := rows.Scan(
			&privateData.PvtDataID,
			&privateData.Capsule,
			&privateData.CipherText,
		)
		if err != nil {
			return nil, err
		}
		privateDataList = append(privateDataList, privateData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return privateDataList, nil
}

func (r *Repository) AddSecretKeys(secretKeyData model.SecretKeyData) error {

	query := "INSERT INTO secretkeydata (secret_key, public_key, user_id) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, secretKeyData.SecretKey, secretKeyData.PublicKey, secretKeyData.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetKeyDetails(userID int) (model.SecretKeyData, error) {
	var keyDetails model.SecretKeyData
	query := "SELECT * FROM secretkeydata WHERE user_id = ?"

	err := r.db.QueryRow(query, userID).Scan(
		&keyDetails.KeyID,
		&keyDetails.SecretKey,
		&keyDetails.PublicKey,
		&keyDetails.UserID,
	)

	if err != nil {
		return model.SecretKeyData{}, err
	}

	return keyDetails, nil
}
