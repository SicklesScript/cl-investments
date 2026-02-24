package cli

import (
	"github.com/SicklesScript/cl-investments/internal/database"
)

// Struct to contain queries for access in cli.go
type State struct {
	DBQueries   *database.Queries
	CurrentUser string
}
