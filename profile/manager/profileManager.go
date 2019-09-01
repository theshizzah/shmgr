package manager

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "profupdates.db"

// ProfileManager manages the profile metadata
type profileManager struct {
	db *sql.DB
}

// New is a ProfileManager constructor
func New(profilePath string) *profileManager {
	dbPath := filepath.Join(profilePath, dbName)
	dsn := fmt.Sprintf("file:%s?cache=private&mode=rwc", dbPath)

	pm := new(profileManager)
	_, statErr := os.Stat(dbPath)
	db, openErr := sql.Open("sqlite3", dsn)
	if openErr != nil {
		fmt.Printf("error connecting to db at %s: %s", dbPath, openErr)
	}
	if statErr != nil {
		_, execErr := db.Exec(initSQL())
		if execErr != nil {
			fmt.Printf("error init'ing db: %s", execErr)
		}
	}

	pm.db = db
	return pm
}

func initSQL() string {
	return `
CREATE TABLE IF NOT EXISTS updates (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	cmd TEXT NOT NULL,
	ts TEXT NOT NULL
);`
}
