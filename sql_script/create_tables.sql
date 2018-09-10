CREATE TABLE IF NOT EXISTS User (
    userID VARCHAR(64) PRIMARY KEY,
    userType VARCHAR(32) NOT NULL,
    name VARCHAR(64)
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS Accounting (
    ID INT UNSIGNED PRIMARY KEY,
    userID VARCHAR(64) UNIQUE KEY,
    kind VARCHAR(64),
    createdTime TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    accountingTime DATETIME DEFAULT CURRENT_TIMESTAMP,
    price DECIMAL(19, 4),
    purpose VARCHAR(128)
) ENGINE=InnoDB CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;