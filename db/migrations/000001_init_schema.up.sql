-- 创建 users 表
CREATE TABLE users (
                       id BIGINT PRIMARY KEY AUTO_INCREMENT,
                       username VARCHAR(50) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       created_at DATETIME(3) NOT NULL,
                       updated_at DATETIME(3) NOT NULL
);

-- 创建 posts 表（user_id 逻辑关联 users.id）
CREATE TABLE posts (
                       id BIGINT PRIMARY KEY AUTO_INCREMENT,
                       title VARCHAR(200) NOT NULL,
                       content TEXT NOT NULL,
                       user_id BIGINT NOT NULL,
                       created_at DATETIME(3) NOT NULL,
                       updated_at DATETIME(3) NOT NULL
);

-- 创建 comments 表（user_id 逻辑关联 users.id，post_id 逻辑关联 posts.id）
CREATE TABLE comments (
                          id BIGINT PRIMARY KEY AUTO_INCREMENT,
                          content TEXT NOT NULL,
                          user_id BIGINT NOT NULL,
                          post_id BIGINT NOT NULL,
                          created_at DATETIME(3) NOT NULL
);
