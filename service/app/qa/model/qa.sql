-- 问题表（Question）
CREATE TABLE `question` (
                            `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                            `user_id` INT UNSIGNED NOT NULL,
                            `title` VARCHAR(255) NOT NULL DEFAULT '',
                            `content` TEXT NOT NULL,
                            `create_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                            `update_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                            `delete_time` datetime        COMMENT '删除时间',
                            PRIMARY KEY (`id`),
                            INDEX `user_id_index` (`user_id`)
) ENGINE=InnoDB;


-- 问答表（Q&A）
CREATE TABLE `answer` (
                          `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
                          `user_id` INT UNSIGNED NOT NULL,
                          `question_id` INT UNSIGNED NOT NULL,
                          `parent_id` INT UNSIGNED DEFAULT 0 ,
                          `content` TEXT NOT NULL,
                          `is_answer` TINYINT(1) NOT NULL DEFAULT 0,
                          `create_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                          `update_time`  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          `delete_time` datetime       COMMENT '删除时间',
                          PRIMARY KEY (`id`),
                          INDEX `user_id_index` (`user_id`),
                          INDEX `question_id_index` (`question_id`),
                          INDEX `parent_id_index` (`parent_id`)
) ENGINE=InnoDB;