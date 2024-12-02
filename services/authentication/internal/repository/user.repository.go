package repository

import (
	"database/sql"
	"dolott_authentication/internal/models"
	"dolott_authentication/internal/types"
	"fmt"
	"time"
)

func (c *authenticationRepository) UserExistsByEmail(email string) (bool, *types.Error) {
	query := `SELECT 1 FROM "users" WHERE email = $1`
	var result int32
	err := c.db.QueryRow(query, email).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, types.NewInternalError("internal issue, error code #1001")
	}
	return true, nil
}

func (c *authenticationRepository) UserExistsByPhoneNumber(phone_number string) (bool, *types.Error) {
	query := `SELECT 1 FROM "users" WHERE phone_number = $1`
	var result int32
	err := c.db.QueryRow(query, phone_number).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, types.NewInternalError("internal issue, error code #1002")
	}
	return true, nil
}

func (c *authenticationRepository) UserExistsByAccountUsername(username string) (bool, *types.Error) {
	query := `SELECT 1 FROM "users" WHERE account_username = $1`
	var result int32
	err := c.db.QueryRow(query, username).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, types.NewInternalError("internal issue, error code #1042")
	}
	return true, nil
}

func (c *authenticationRepository) UserExistsByUserId(user_id int32) (bool, *types.Error) {
	query := `SELECT 1 FROM "users" WHERE user_id = $1`
	var result int32
	err := c.db.QueryRow(query, user_id).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, types.NewInternalError("internal issue, error code #1003")
	}
	return true, nil
}

func (c *authenticationRepository) GetUserByUserId(user_id int32) (*models.User, *types.Error) {
	query := `SELECT user_id, phone_number, email, account_username, user_role, user_status, provider, is_sso, created_at FROM "users" WHERE user_id = $1`
	var user models.User
	err := c.db.QueryRow(query, user_id).Scan(&user.UserId, &user.PhoneNumber, &user.Email, &user.AccountUsername, &user.UserRole, &user.UserStatus, &user.Provider, &user.IsSSO, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("user not found, error code #1004")
		}
		return nil, types.NewInternalError("internal issue, error code #1005")

	}
	return &user, nil
}

func (c *authenticationRepository) GetRoleByUserId(user_id int32) (string, *types.Error) {
	query := `SELECT user_role FROM "users" WHERE user_id = $1`
	var userRole string
	err := c.db.QueryRow(query, user_id).Scan(&userRole)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", types.NewNotFoundError("user not found, error code #1006")
		}
		return "", types.NewInternalError("internal issue, error code #1007")
	}

	return userRole, nil
}

func (c *authenticationRepository) AddUser(data *models.UserDTO) (*models.User, *types.Error) {
	fmt.Println(data.Email, data.AccountUsername)
	query := `INSERT INTO "users" (email, account_username, user_role, user_status, provider, is_sso, phone_number, created_at) VALUES ($1,$2,$3,$4,$5,$6,'',NOW()) RETURNING user_id`
	row := c.db.QueryRow(query, data.Email, data.AccountUsername, data.UserRole, data.UserStatus, data.Provider, data.IsSSO)
	if row.Err() != nil {
		fmt.Println(row.Err())
		return nil, types.NewInternalError("internal issue, error code #1008")
	}
	var result int32
	if err := row.Scan(&result); err != nil {
		return nil, types.NewInternalError("internal issue, error code #1009")
	}
	return &models.User{
		UserId:      result,
		PhoneNumber: "",
		Email:       data.Email,
		UserRole:    data.UserRole,
		UserStatus:  data.UserStatus,
		CreatedAt:   time.Now(),
	}, nil
}

func (c *authenticationRepository) GetUserByEmail(email string) (*models.User, *types.Error) {
	query := `SELECT user_id, phone_number, email, account_username, user_role, user_status, provider, is_sso, created_at FROM "users" WHERE email = $1`
	var user models.User
	err := c.db.QueryRow(query, email).Scan(&user.UserId, &user.PhoneNumber, &user.Email, &user.AccountUsername, &user.UserRole, &user.UserStatus, &user.Provider, &user.IsSSO, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("user not found, error code #1010")
		}
		fmt.Println(err)
		return nil, types.NewInternalError("internal issue, error code #1011")

	}
	return &user, nil
}

func (c *authenticationRepository) GetUserByAccountUsername(username string) (*models.User, *types.Error) {
	query := `SELECT user_id, phone_number, email, account_username, user_role, user_status, provider, is_sso, created_at FROM "users" WHERE account_username = $1`
	var user models.User
	err := c.db.QueryRow(query, username).Scan(&user.UserId, &user.PhoneNumber, &user.Email, &user.AccountUsername, &user.UserRole, &user.UserStatus, &user.Provider, &user.IsSSO, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, types.NewNotFoundError("user not found, error code #1043")
		}
		fmt.Println(err)
		return nil, types.NewInternalError("internal issue, error code #1044")

	}
	return &user, nil
}
func (c *authenticationRepository) GetUsers(data *models.Pagination) ([]models.User, *types.Error) {
	query := `SELECT user_id, phone_number, email, account_username, user_role, user_status, provider, is_sso, created_at FROM "users" OFFSET $1 LIMIT $2`

	rows, err := c.db.Query(query, data.Offset, data.Limit)
	if err != nil {
		return nil, types.NewInternalError("internal issue, error code #1017")
	}

	defer rows.Close()
	users := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(
			&user.UserId,
			&user.PhoneNumber,
			&user.Email,
			&user.AccountUsername,
			&user.UserRole,
			&user.UserStatus,
			&user.Provider,
			&user.IsSSO,
			&user.CreatedAt,
		); err != nil {
			return nil, types.NewInternalError("internal issue, error code #1018")
		}
		users = append(users, user)
	}
	return users, nil
}

func (c *authenticationRepository) GetUserCount() (int32, *types.Error) {
	query := `SELECT count(*) FROM "users"`
	var totalCount int32
	err := c.db.QueryRow(query).Scan(&totalCount)
	if err != nil {
		return 0, types.NewInternalError("internal issue, error code #1019")
	}
	return totalCount, nil
}

func (c *authenticationRepository) ChangeUserStatus(userId int32, newStatus string) *types.Error {
	query := `UPDATE "users" SET user_status = $1 WHERE user_id = $2`

	_, err := c.db.Exec(query, newStatus, userId)
	if err != nil {
		return types.NewInternalError("internal issue, error code #1020")
	}
	return nil
}
