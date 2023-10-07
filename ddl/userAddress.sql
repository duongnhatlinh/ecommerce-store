DROP TABLE IF EXISTS `user_addresses`;
CREATE TABLE `user_addresses` (
                                  `id` int NOT NULL AUTO_INCREMENT,
                                  `user_id` int NOT NULL,
                                  `country` varchar(100) NOT NULL,
                                  `city` varchar(100) NOT NULL,
                                  `address` varchar(255) NOT NULL,
                                  `postal_code` varchar(20) NOT NULL,
                                  `phone` varchar(20) NOT NULL,
                                  `lat` double DEFAULT NULL, -- chua trien khai
                                  `lng` double DEFAULT NULL, -- chua trien khai
                                  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                  PRIMARY KEY (`id`),
                                  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=202 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
