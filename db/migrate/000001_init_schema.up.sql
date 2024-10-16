CREATE TABLE `users` (
  `id` int PRIMARY KEY,
  `nickname` varchar(255) COMMENT '用户名',
  `password` varchar(255) COMMENT '密码',
  `gender` tinyint COMMENT '用户性别',
  `age` int COMMENT '年龄',
  `avatar` varchar(255) COMMENT '头像',
  `created_at` date COMMENT '创建时间',
  `updated_at` date COMMENT '更新时间'
);

ALTER TABLE `users` COMMENT = '用户信息表';
