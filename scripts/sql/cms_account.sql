CREATE DATABASE `cms_account` DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
use cms_account;
create table account
(
    id         int auto_increment comment '主键ID'    primary key,
    user_id    varchar(64) default ''                null comment '用户id',
    pass_word  varchar(64) default ''                null comment '密码',
    nick_name  varchar(64) default ''                null comment '昵称',
    created_at timestamp   default CURRENT_TIMESTAMP not null comment '创建时间',
    updated_at timestamp   default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    constraint uk_user_id
        unique (user_id)
)
    comment 'cms账号信息';


