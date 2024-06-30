-- Create "organizations" table
CREATE TABLE `organizations` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Modify "users" table
ALTER TABLE `users` ADD COLUMN `organization_users` bigint NULL, ADD CONSTRAINT `users_organizations_users` FOREIGN KEY (`organization_users`) REFERENCES `organizations` (`id`) ON DELETE SET NULL;
-- Create "photos" table
CREATE TABLE `photos` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `url` varchar(255) NOT NULL, `user_photos` bigint NULL, PRIMARY KEY (`id`), CONSTRAINT `photos_users_photos` FOREIGN KEY (`user_photos`) REFERENCES `users` (`id`) ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
