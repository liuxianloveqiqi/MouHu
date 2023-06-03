

CREATE TABLE `follow` (
                          `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                          `follower_id` INT UNSIGNED NOT NULL,
                          `followee_id` INT UNSIGNED NOT NULL,
                          `create_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          `update_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          `delete_time`  COMMENT '删除时间',
                          PRIMARY KEY (`id`),
                          INDEX `follower_id_index` (`follower_id`),
                          INDEX `followee_id_index` (`followee_id`)
) ENGINE=InnoDB;