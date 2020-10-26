USE userdb;

CREATE TABLE course_ingress_solicitation (
    id INT AUTO_INCREMENT PRIMARY KEY,
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    text VARCHAR(500),
    image VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (course_id) REFERENCES course(id)
);

CREATE TABLE course_registration (
    id INT AUTO_INCREMENT PRIMARY KEY,
    student_id INT NOT NULL,
    course_id INT NOT NULL,
    course_ingress_solicitation_id INT NOT NULL,
    valid BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id),
    FOREIGN KEY (course_id) REFERENCES course(id),
    FOREIGN KEY (course_ingress_solicitation_id) REFERENCES course_ingress_solicitation(id)
);
