CREATE TABLE authors
(
    `id`         BIGINT       NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '自增ID',
    `name`       varchar(100) NOT NULL COMMENT '作者名字',
    `nick_name`  varchar(100) NOT NULL COMMENT '作者昵称',
    `created_at` datetime     NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime     NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间'
);


CREATE TABLE author_articles
(
    `id`              BIGINT   NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '自增ID',
    `author_id`       BIGINT   NOT NULL COMMENT '作者id',
    `article_content` text     NOT NULL,
    `created_at`      datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at`      datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间'
);

