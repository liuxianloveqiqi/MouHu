CREATE TABLE `like` (
                        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                        `user_id` INT UNSIGNED NOT NULL,
                        `target_type` VARCHAR(20) NOT NULL,
                        `target_id` INT UNSIGNED NOT NULL,
                        `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                        `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        `delete_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
                        PRIMARY KEY (`id`),
                        INDEX `user_id_index` (`user_id`),
                        INDEX `target_index` (`target_type`, `target_id`)
) ENGINE=InnoDB;

CREATE TABLE `follow` (
                          `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                          `follower_id` INT UNSIGNED NOT NULL,
                          `followee_id` INT UNSIGNED NOT NULL,
                          `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          `delete_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
                          PRIMARY KEY (`id`),
                          INDEX `follower_id_index` (`follower_id`),
                          INDEX `followee_id_index` (`followee_id`)
) ENGINE=InnoDB;