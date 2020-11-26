USE userdb;

ALTER TABLE course_registration
    DROP FOREIGN KEY course_registration_ibfk_3,
    ADD CONSTRAINT course_registration_ibfk_4 FOREIGN KEY (course_ingress_solicitation_id)
        REFERENCES course_ingress_solicitation (id)
        ON DELETE CASCADE;
