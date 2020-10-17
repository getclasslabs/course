USE userdb;

CREATE TABLE category (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE course (
    id INT AUTO_INCREMENT PRIMARY KEY,
    teacher_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    category_id INT NOT NULL,
    max_students INT,
    classes INT DEFAULT NULL,
    periods VARCHAR(200),
    price DECIMAL(5,2),
    start_day date,
    type VARCHAR(10) NOT NULL,
    place VARCHAR(200) DEFAULT NULL,
    class_open BOOLEAN DEFAULT FALSE,
    classes_given INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (teacher_id) REFERENCES teacher(id),
    FOREIGN KEY (category_id) REFERENCES category(id)
);