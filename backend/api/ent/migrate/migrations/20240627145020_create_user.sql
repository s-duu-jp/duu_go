-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `sid` varchar(255) NOT NULL, `uid` varchar(255) NOT NULL, `name` varchar(255) NOT NULL, `email` varchar(255) NOT NULL, `password` varchar(255) NULL, `role_type` varchar(255) NOT NULL, `status_type` varchar(255) NOT NULL, `oauth_type` varchar(255) NOT NULL, `sub` varchar(255) NULL, PRIMARY KEY (`id`), UNIQUE INDEX `uid` (`uid`), UNIQUE INDEX `email` (`email`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;