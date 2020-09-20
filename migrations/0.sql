USE userdb;

CREATE TABLE category (
    id INT(6) AUTO_INCREMENT PRIMARY KEY,
    name varchar(100)
);

CREATE TABLE course (
    id INT(6) AUTO_INCREMENT PRIMARY KEY,
    teacher_id INT(6),
    name varchar(100),
    category_id int(6),
    max_students int(4),
    classes int(4),
    start_day date,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (teacher_id) REFERENCES teacher(id),
    FOREIGN KEY (category_id) REFERENCES category(id)
);

CREATE TABLE class_date (
    id INT(6) AUTO_INCREMENT PRIMARY KEY,
    course_id INT(6),
    day int(1),
    hour TIME,
    FOREIGN KEY (course_id) REFERENCES course(id)
);

CREATE TABLE rescheduled_class (
    id INT(6) AUTO_INCREMENT PRIMARY KEY,
    course_id INT(6),
    class_date_id INT(6),
    datetime DATETIME,
    created_at timestamp default CURRENT_TIMESTAMP,
    FOREIGN KEY (course_id) REFERENCES course(id),
    FOREIGN KEY (class_date_id) REFERENCES class_date(id)
);