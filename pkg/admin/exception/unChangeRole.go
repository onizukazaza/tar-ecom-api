package exception

import "fmt"

type UnChangeRole struct {
	UserID string
	Role   string
}

func (e *UnChangeRole) Error() string {
	return fmt.Sprintf("failed to update role for user %s to %s", e.UserID, e.Role)
}
