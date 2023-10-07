DROP TABLE IF EXISTS `inventories`;
CREATE TABLE `inventories` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `quantity` int NOT NULL,
                         `status` int NOT NULL DEFAULT '1',
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=157 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
