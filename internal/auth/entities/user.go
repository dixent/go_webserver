package entities

type User struct {
	Id       int64  `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

func (u *User) GetJwtClaims() map[string]any {
	return map[string]any{
		"id": u.Id,
		"email": u.Email,
	}
}
