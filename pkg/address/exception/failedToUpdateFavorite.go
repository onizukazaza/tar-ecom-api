package exception

import "fmt"

type FailedToUpdateFavorite struct {
    ID string
}

func (e *FailedToUpdateFavorite) Error() string {
    return fmt.Sprintf("failed to update favorite status for address with ID '%s'", e.ID)
}
