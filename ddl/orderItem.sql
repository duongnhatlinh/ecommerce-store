DROP TABLE IF EXISTS `order_items`;
CREATE TABLE `order_items` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `order_id` int NOT NULL,
                                 `product_id` int NOT NULL,
                                 `price` float NOT NULL,
                                 `quantity` int NOT NULL,
                                 `status` int NOT NULL DEFAULT '1',
                                 `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`),
                                 KEY `order_id` (`order_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=129 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
