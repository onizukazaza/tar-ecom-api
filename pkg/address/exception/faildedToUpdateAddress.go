package exception

import "fmt"

type FailedToUpdateAddress struct {
    AddressID string
}

func (e *FailedToUpdateAddress) Error() string {
    return fmt.Sprintf("Failed to update address with ID '%s'", e.AddressID)
}
