package database

import (
    "database/sql"
	"errors"
    _ "github.com/mattn/go-sqlite3"
    "noahhefner/notes/models"
)

var (
    // ErrUserNotFound is returned when a user is not found in the database
    ErrUserNotFound = errors.New("user not found")
)

var db *sql.DB

// Init initializes the SQLite database
func Init(dbPath string) error {
    var err error
    db, err = sql.Open("sqlite3", dbPath)
    if err != nil {
        return err
    }

    // Create users table if not exists
    _, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        username TEXT PRIMARY KEY,
        password TEXT
    )`)
    if err != nil {
        return err
    }

    return nil
}

// Close closes the database connection
func Close() error {
    if db != nil {
        return db.Close()
    }
    return nil
}

// UserExists checks if a user with the given username exists in the database
func UserExists(username string) bool {
    var exists bool
    err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = ?)", username).Scan(&exists)
    if err != nil {
        return false
    }
    return exists
}

// InsertUser inserts a new user into the database
func InsertUser(user models.User) error {
    _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
    return err
}

// GetUserByUsername retrieves a user from the database by username
func GetUserByUsername(username string) (models.User, error) {
    var user models.User
    err := db.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.Username, &user.Password)
    if err != nil {
        if err == sql.ErrNoRows {
            return models.User{}, ErrUserNotFound
        }
        return models.User{}, err
    }
    return user, nil
}