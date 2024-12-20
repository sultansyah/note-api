CREATE TABLE notes(
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    note TEXT NOT NULL,
    status varchar(255) NOT NULL,
    priority varchar(255) NOT NULL,
    category varchar(255) NOT NULL,
    tags varchar(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
) ENGINE = InnoDB;