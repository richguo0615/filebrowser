package storage

import (
	"github.com/richguo0615/filebrowser/auth"
	"github.com/richguo0615/filebrowser/settings"
	"github.com/richguo0615/filebrowser/share"
	"github.com/richguo0615/filebrowser/users"
)

// Storage is a storage powered by a Backend whih makes the neccessary
// verifications when fetching and saving data to ensure consistency.
type Storage struct {
	Users    *users.Storage
	Share    *share.Storage
	Auth     *auth.Storage
	Settings *settings.Storage
}
