CREATE DATABASE `cms_content` DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
use cms_content;
create table content_details_0
(
    id              bigint auto_increment comment '自增ID'    primary key,
    content_id      varchar(255)                              null comment '内容ID',
    title           varchar(255) default ''                   null comment '内容标题',
    description     text                                      null comment '内容描述',
    author          varchar(255) default ''                   null comment '作者',
    video_url       varchar(255) default ''                   null comment '视频播放URL',
    thumbnail       varchar(255) default ''                   null comment '封面图URL',
    category        varchar(255) default ''                   null comment '内容分类',
    duration        bigint       default 0                    null comment '内容时长',
    resolution      varchar(255) default ''                   null comment '分辨率 如720p、1080p',
    file_size       bigint       default 0                    null comment '文件大小',
    format          varchar(255) default ''                   null comment '文件格式 如MP4、AVI',
    quality         int          default 0                    null comment '视频质量 1-高清 2-标清',
    approval_status int          default 1                    null comment '审核状态 1-审核中 2-审核通过 3-审核不通过',
    updated_at      timestamp(6) default CURRENT_TIMESTAMP(6) not null on update CURRENT_TIMESTAMP(6) comment '内容更新时间',
    created_at      timestamp(6) default CURRENT_TIMESTAMP(6) not null comment '内容创建时间'
)
    comment '内容详情';

create table content_details_1
(
    id              bigint auto_increment comment '自增ID'    primary key,
    content_id      varchar(255)                              null comment '内容ID',
    title           varchar(255) default ''                   null comment '内容标题',
    description     text                                      null comment '内容描述',
    author          varchar(255) default ''                   null comment '作者',
    video_url       varchar(255) default ''                   null comment '视频播放URL',
    thumbnail       varchar(255) default ''                   null comment '封面图URL',
    category        varchar(255) default ''                   null comment '内容分类',
    duration        bigint       default 0                    null comment '内容时长',
    resolution      varchar(255) default ''                   null comment '分辨率 如720p、1080p',
    file_size       bigint       default 0                    null comment '文件大小',
    format          varchar(255) default ''                   null comment '文件格式 如MP4、AVI',
    quality         int          default 0                    null comment '视频质量 1-高清 2-标清',
    approval_status int          default 1                    null comment '审核状态 1-审核中 2-审核通过 3-审核不通过',
    updated_at      timestamp(6) default CURRENT_TIMESTAMP(6) not null on update CURRENT_TIMESTAMP(6) comment '内容更新时间',
    created_at      timestamp(6) default CURRENT_TIMESTAMP(6) not null comment '内容创建时间'
)
    comment '内容详情';

create table content_details_2
(
    id              bigint auto_increment comment '自增ID'    primary key,
    content_id      varchar(255)                              null comment '内容ID',
    title           varchar(255) default ''                   null comment '内容标题',
    description     text                                      null comment '内容描述',
    author          varchar(255) default ''                   null comment '作者',
    video_url       varchar(255) default ''                   null comment '视频播放URL',
    thumbnail       varchar(255) default ''                   null comment '封面图URL',
    category        varchar(255) default ''                   null comment '内容分类',
    duration        bigint       default 0                    null comment '内容时长',
    resolution      varchar(255) default ''                   null comment '分辨率 如720p、1080p',
    file_size       bigint       default 0                    null comment '文件大小',
    format          varchar(255) default ''                   null comment '文件格式 如MP4、AVI',
    quality         int          default 0                    null comment '视频质量 1-高清 2-标清',
    approval_status int          default 1                    null comment '审核状态 1-审核中 2-审核通过 3-审核不通过',
    updated_at      timestamp(6) default CURRENT_TIMESTAMP(6) not null on update CURRENT_TIMESTAMP(6) comment '内容更新时间',
    created_at      timestamp(6) default CURRENT_TIMESTAMP(6) not null comment '内容创建时间'
)
    comment '内容详情';

create table content_details_3
(
    id              bigint auto_increment comment '自增ID'    primary key,
    content_id      varchar(255)                              null comment '内容ID',
    title           varchar(255) default ''                   null comment '内容标题',
    description     text                                      null comment '内容描述',
    author          varchar(255) default ''                   null comment '作者',
    video_url       varchar(255) default ''                   null comment '视频播放URL',
    thumbnail       varchar(255) default ''                   null comment '封面图URL',
    category        varchar(255) default ''                   null comment '内容分类',
    duration        bigint       default 0                    null comment '内容时长',
    resolution      varchar(255) default ''                   null comment '分辨率 如720p、1080p',
    file_size       bigint       default 0                    null comment '文件大小',
    format          varchar(255) default ''                   null comment '文件格式 如MP4、AVI',
    quality         int          default 0                    null comment '视频质量 1-高清 2-标清',
    approval_status int          default 1                    null comment '审核状态 1-审核中 2-审核通过 3-审核不通过',
    updated_at      timestamp(6) default CURRENT_TIMESTAMP(6) not null on update CURRENT_TIMESTAMP(6) comment '内容更新时间',
    created_at      timestamp(6) default CURRENT_TIMESTAMP(6) not null comment '内容创建时间'
)
    comment '内容详情';


CREATE TABLE IF NOT EXISTS idx_content_details (
                                                   id BIGINT NOT NULL AUTO_INCREMENT,
                                                   content_id VARCHAR(255) DEFAULT '' COMMENT '内容ID',
                                                   title VARCHAR(255) DEFAULT '' COMMENT '内容标题',
                                                   author VARCHAR(255) DEFAULT '' COMMENT '作者',
                                                   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '内容更新时间',
                                                   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '内容创建时间',
                                                   PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;