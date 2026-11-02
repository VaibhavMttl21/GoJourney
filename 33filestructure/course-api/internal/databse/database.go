package database

import (
    "fmt"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)
// what dsn info give
//host     → where PostgreSQL is
// user     → who is connecting
// password → password
// dbname   → which database
// port     → PostgreSQL port
// sslmode  → SSL configuration
func Connect() (*gorm.DB, error) {

    dsn := "host=localhost user=postgres password=postgres dbname=course_db port=5432 sslmode=disable"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %w", err)
    }

    return db, nil
}