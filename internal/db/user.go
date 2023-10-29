package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/tcarreira/roaw/pkg/types"
)

func ListAllUsers(db *sqlx.DB) (types.Users, error) {
	users := types.Users{}
	err := db.Select(&users, "SELECT * FROM user")
	return users, err
}

func UserCreate(db *sqlx.DB, u *types.User) error {
	if u == nil {
		return fmt.Errorf("cannot create a nil User")
	}
	_, err := db.NamedExec("INSERT INTO user ("+
		"id, name, email, provider, provider_id, "+
		"access_token, refresh_token, avatar_url, "+
		"created_at, updated_at "+
		") VALUES ("+
		"666, :name, :email, :provider, :provider_id, "+
		":access_token, :refresh_token, :avatar_url, "+
		":created_at, :updated_at "+
		")",
		u)
	return err
}
