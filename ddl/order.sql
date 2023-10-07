DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `user_id` int NOT NULL,
                          `total_price` float NOT NULL,
                          `status` int NOT NULL DEFAULT '1',
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`),
                          KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=143 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


