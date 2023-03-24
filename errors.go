package sgdb

import (
	"fmt"
	"strings"
)

// ResponseError represents an error returned by the SteamGridDB api.
type ResponseError struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

func (r ResponseError) Error() string {
	return fmt.Sprintf("steamgriddb: %s", strings.Join(r.Errors, ", "))
}
