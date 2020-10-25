USE userdb;

CREATE TABLE bank_infos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(250) NOT NULL,
    cpf  VARCHAR(20) NOT NULL,
    bank_name VARCHAR(100) NOT NULL,
    agency VARCHAR(20) NOT NULL,
    type VARCHAR(30)
);

ALTER TABLE teacher ADD bank_infos_id INT;
ALTER TABLE teacher
ADD FOREIGN KEY (bank_infos_id) REFERENCES bank_infos(id);

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
