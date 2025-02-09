package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func InitDB(connStr string) (*sql.DB, error) {
    db, err := sql.Open("mysql", connStr)
    if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

    // Create tasks table if it doesn't exist
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS tasks (
            id INT AUTO_INCREMENT PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            description TEXT,
            status VARCHAR(50) DEFAULT 'pending',
            due_date DATETIME,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
            cost DECIMAL(10,2),
            currency VARCHAR(3)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
    `)

    return db, err
}