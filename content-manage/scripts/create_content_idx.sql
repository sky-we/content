CREATE TABLE IF NOT EXISTS t_idx_content_details (
                                                     id BIGINT NOT NULL AUTO_INCREMENT,
                                                     content_id VARCHAR(255) DEFAULT '' COMMENT '内容ID',
    title VARCHAR(255) DEFAULT '' COMMENT '内容标题',
    author VARCHAR(255) DEFAULT '' COMMENT '作者',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '内容更新时间',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '内容创建时间',
    PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;