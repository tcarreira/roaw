package types

import (
	"time"
)

type Users []User
type User struct {
	ID           string    `json:"id" form:"id" query:"id" db:"id"`
	Name         string    `json:"name" form:"name" query:"name" db:"name"`
	Email        string    `json:"email" form:"email" query:"email" db:"email"`
	Provider     string    `json:"provider" form:"provider" query:"provider" db:"provider"`
	ProviderID   string    `json:"provider_id" form:"provider_id" query:"provider_id" db:"provider_id"`
	AccessToken  string    `json:"-" db:"access_token"`
	RefreshToken string    `json:"-" db:"refresh_token"`
	AvatarURL    string    `json:"avatar_url" form:"avatar_url" query:"avatar_url" db:"avatar_url"`
	CreatedAt    time.Time `json:"created_at" form:"created_at" query:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at" query:"updated_at" db:"updated_at"`
}
