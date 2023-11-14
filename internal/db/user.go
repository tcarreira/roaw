package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tcarreira/roaw/pkg/types"
)

func ListAllUsers(db *sqlx.DB) (types.Users, error) {
	users := types.Users{}
	err := db.Select(&users, "SELECT * FROM roaw_user")
	return users, err
}

func UserCreate(db *sqlx.DB, u *types.User) error {
	if u == nil {
		return fmt.Errorf("cannot create a nil User")
	}
	_, err := db.Exec("INSERT INTO roaw_user ("+
		"id, name, email, provider, provider_id, "+
		"access_token, refresh_token, avatar_url, "+
		"created_at, updated_at "+
		") VALUES ("+
		"?, ?, ?, ?, ?, "+
		"?, ?, ?, "+
		"?, ?, "+
		")",
		u.ID, u.Name, u.Email, u.Provider, u.ProviderID,
		u.AccessToken, u.RefreshToken, u.AvatarURL,
		time.Now(), u.UpdatedAt,
	)
	return err
}
