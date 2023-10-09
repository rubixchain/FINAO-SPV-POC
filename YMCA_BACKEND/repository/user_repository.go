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
func (repo *Repository) GetPublicDataByDID(did string) (*model.PublicData, error) {
	// Define the SQL query to fetch public data by DID and sort by created time in descending order
	query := `
		SELECT pub_data_id, focus_area, communities, user_id, created_at, updated_at
		FROM public_data
		WHERE user_id = (
			SELECT user_id
			FROM users
			WHERE did = ?
		)
		ORDER BY created_at DESC;`

	// Execute the query and retrieve the public data
	var publicData model.PublicData
	err := repo.db.QueryRow(query, did).Scan(
		&publicData.PubDataID,
		&publicData.FocusArea,
		&publicData.Communities,
		&publicData.UserID,
		&publicData.CreatedAt,
		&publicData.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("public data not found for DID: %s", did)
		}
		return nil, err
	}

	return &publicData, nil
}

// GetPrivateataByDID fetches the public data of a user by their DID, sorted by created time in descending order.
func (repo *Repository) GetPrivateDataByDID(did string) (*model.PrivateData, error) {
	// Define the SQL query to fetch public data by DID and sort by created time in descending order
	query := `
		SELECT pvt_data_id, focus_area, communities, user_id, created_at, updated_at
		FROM private_data
		WHERE user_id = (
			SELECT user_id
			FROM users
			WHERE did = ?
		)
		ORDER BY created_at DESC;`

	// Execute the query and retrieve the public data
	var privateData model.PrivateData
	err := repo.db.QueryRow(query, did).Scan(
		&privateData.PvtDataID,
		&privateData.Capsule,
		&privateData.CipherText,
		&privateData.UserID,
		&privateData.CreatedAt,
		&privateData.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("public data not found for DID: %s", did)
		}
		return nil, err
	}

	return &privateData, nil
}

// GetAllAccessDataByDID retrieves private data by DID.
func (repo *Repository) GetAllAccessDataByDID(did string) ([]model.PrivateData, error) {
	// Define the SQL query to fetch private data by DID
	query := `
		SELECT pd.pvt_data_id, pd.focus_area, pd.communities, pd.user_id, pd.created_at, pd.updated_at
		FROM private_data pd
		INNER JOIN access_sheet as ON pd.pvt_data_id = as.pvt_data_id
		INNER JOIN users u ON as.decrypt_user_id = u.user_id
		WHERE u.did = ?;`

	// Execute the query and retrieve the private data
	rows, err := repo.db.Query(query, did)
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
