USE userdb;

CREATE TABLE category (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE course (
    id INT AUTO_INCREMENT PRIMARY KEY,
    teacher_id INT,
    name VARCHAR(100),
    description VARCHAR(500),
    category_id INT,
    max_students INT,
    classes INT,
    start_day date,
    type VARCHAR(10),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (teacher_id) REFERENCES teacher(id),
    FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE class_date (
    id INT AUTO_INCREMENT PRIMARY KEY,
    course_id INT,
    day INT,
    hour TIME,
    FOREIGN KEY (course_id) REFERENCES course(id)
);

CREATE TABLE rescheduled_class (
    id INT AUTO_INCREMENT PRIMARY KEY,
    course_id INT,
    class_date_id INT,
    datetime DATETIME,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (course_id) REFERENCES course(id),
    FOREIGN KEY (class_date_id) REFERENCES class_date(id)
);