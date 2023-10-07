DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `email` varchar(50) NOT NULL,
                         `password` varchar(60) NOT NULL,
                         `salt` varchar(50) DEFAULT NULL,
                         `last_name` varchar(50) NOT NULL,
                         `first_name` varchar(50) NOT NULL,
                         `phone` varchar(20) DEFAULT NULL,
                         `avatar` json DEFAULT NULL,
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=111 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
