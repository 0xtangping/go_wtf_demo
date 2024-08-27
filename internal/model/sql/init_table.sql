CREATE DATABASE wtf_demo CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;

USE wtf_demo;

CREATE TABLE article (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    author VARCHAR(32) NOT NULL,
    title VARCHAR(256) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- goctl model mysql datasource -url="root:QWEqwe1234@tcp(107.174.218.167:3306)/wtf_demo" -table="articles" -dir internal/model
-- 
