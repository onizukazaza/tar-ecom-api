package exception

import "fmt"

type UserEditing struct {
	UserID string
}

func (e *UserEditing) Error() string {
	return fmt.Sprintf("editing user with ID '%s' failed", e.UserID)
}
