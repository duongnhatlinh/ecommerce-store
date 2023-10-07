DROP TABLE IF EXISTS `product_likes`;
CREATE TABLE `product_likes` (
                              `user_id` int NOT NULL,
                              `product_id` int NOT NULL,
                              `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`user_id`,`product_id`),
                              KEY `product_id` (`product_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
