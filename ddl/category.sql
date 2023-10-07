DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `name` varchar(100) NOT NULL,
                              `desc` text,
                              `icon` json DEFAULT NULL,
                              `status` int NOT NULL DEFAULT '1',
                              `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=302 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
