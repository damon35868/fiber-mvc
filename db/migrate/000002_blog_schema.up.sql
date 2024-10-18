CREATE TABLE `blogs` (
    `id` int PRIMARY KEY,
    `title` varchar(255) COMMENT '博客名称',
    `desc` text COMMENT '描述',
    `user_id` int COMMENT '用户ID',
    `deleted_at` datetime COMMENT '删除时间',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime COMMENT '更新时间'
);
ALTER TABLE `blogs` COMMENT = '用户博客表';