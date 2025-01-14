package repository


type AdminRepository interface {
	UpdateUserRole(userID string, role string) error

}

