Create TABLE  student
(
    id              BIGINT UNSIGNED PRIMARY KEY,
    name            varchar(50)    NOT NULL,
    age             int(3)         NOT NULL,
    gender          int(2)         NOT NULL,
    class_hour      int(10)        not null,
    left_class_hour int(10)        not null,
    sign_amount     decimal(10, 2) not null,
    open_id         varchar(50),
    extra           text,
    INDEX idx_name (name) USING HASH
) COMMENT 'Student' CHARACTER SET =utf8mb4;

ALTER TABLE student ADD COLUMN create_at datetime DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE student ADD COLUMN modify_at datetime DEFAULT CURRENT_TIMESTAMP;
ALTER TABLE student ADD COLUMN mobile VARCHAR(15);
ALTER TABLE student ADD INDEX idx_mobile(mobile) USING HASH;
ALTER TABLE student ADD COLUMN status int(2) DEFAULT 1 COMMENT '学生状态 1:未体验 10:在学；20:毕业 90:退学';