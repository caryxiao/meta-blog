-- 创建 users 表
CREATE TABLE users (
                       id BIGINT PRIMARY KEY AUTO_INCREMENT,
                       created_at DATETIME(3) DEFAULT NULL,
                       updated_at DATETIME(3) DEFAULT NULL,
                       deleted_at DATETIME(3) DEFAULT NULL,
                       username VARCHAR(50) NOT NULL UNIQUE,
                       password VARCHAR(255) NOT NULL,
                       email VARCHAR(100) NOT NULL UNIQUE
);

-- 创建 posts 表（user_id 逻辑关联 users.id）
CREATE TABLE posts (
                       id BIGINT PRIMARY KEY AUTO_INCREMENT,
                       created_at DATETIME(3) DEFAULT NULL,
                       updated_at DATETIME(3) DEFAULT NULL,
                       deleted_at DATETIME(3) DEFAULT NULL,
                       title VARCHAR(200) NOT NULL,
                       content TEXT NOT NULL,
                       user_id BIGINT NOT NULL
);

-- 创建 comments 表（user_id 逻辑关联 users.id，post_id 逻辑关联 posts.id）
CREATE TABLE comments (
                          id BIGINT PRIMARY KEY AUTO_INCREMENT,
                          created_at DATETIME(3) DEFAULT NULL,
                          updated_at DATETIME(3) DEFAULT NULL,
                          deleted_at DATETIME(3) DEFAULT NULL,
                          content TEXT NOT NULL,
                          user_id BIGINT NOT NULL,
                          post_id BIGINT NOT NULL
);

CREATE TABLE logs (
                      id BIGINT PRIMARY KEY AUTO_INCREMENT,
                      created_at DATETIME(3) DEFAULT NULL,        -- 操作时间
                      user_id BIGINT,                         -- 操作人（可选）
                      action VARCHAR(100) NOT NULL,          -- 操作动作（如 login, create_post）
                      target_type VARCHAR(100),              -- 目标对象类型（如 post, comment）
                      target_id BIGINT,                      -- 目标对象 ID
                      description TEXT,                      -- 描述信息（如请求内容）
                      ip_address VARCHAR(45),                -- 操作 IP（支持 IPv6）
                      user_agent VARCHAR(255)               -- 浏览器/客户端信息
);