CREATE TABLE `connects` (
    `id` int PRIMARY KEY,
    `address` varchar(255) COMMENT '用户地址',
    `user_id` int COMMENT '用户ID',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime COMMENT '更新时间'
);
ALTER TABLE `connects` COMMENT = '用户信息表';