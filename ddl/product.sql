DROP TABLE IF EXISTS `products`;
CREATE TABLE `products` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `name` varchar(200) NOT NULL,
                         `code` varchar(20) NOT NULL,
                         `color` varchar(20) NOT NULL,
                         `size` int NOT NULL,
                         `desc` text,
                         `price` float NOT NULL,
                         `inventory_id` int NOT NULL,
                         `category_id` int NOT NULL,
                         `discount_id` int DEFAULT NULL,
                         `images` json DEFAULT NULL,
                         `liked_count` int DEFAULT '0',
                         `status` int NOT NULL DEFAULT '1',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         KEY `inventory_id` (`inventory_id`) USING BTREE,
                         KEY `category_id` (`category_id`) USING BTREE,
                         KEY `discount_id` (`discount_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=131 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

