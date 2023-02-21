DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `username`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Username',
    `password`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Password',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User account create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User account update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User account delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_username` (`username`) COMMENT 'Username index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='User account table';

DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `author_id`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'AuthorID',
    `publish_at`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'Video publish time',
    `play_url`   varchar(128) NOT NULL DEFAULT '' COMMENT 'PlayURL',
    `cover_url`   varchar(128) NOT NULL DEFAULT '' COMMENT 'CoverURL',
    `title`   varchar(128) NOT NULL DEFAULT '' COMMENT 'Title',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Video info create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Video info update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Video info delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_author_id` (`author_id`) COMMENT 'author_id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Video info table';

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `video_id`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'VideoID',
    `user_id`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'UserID',
    `comment_text`   varchar(128) NOT NULL DEFAULT '' COMMENT 'CommentText',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Comment create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Comment update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Comment delete time',
    PRIMARY KEY (`id`),
    KEY          `idx_user_id` (`user_id`) COMMENT 'user_id index',
    KEY          `idx_video_id` (`video_id`) COMMENT 'video_id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Comment table';

DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `video_id`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'VideoID',
    `user_id`  bigint unsigned NOT NULL DEFAULT 0 COMMENT 'UserID',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Favorite create time',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Favorite update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Favorite delete time',
    PRIMARY KEY (`id`),
    UNIQUE KEY   `idx_user_video` (`user_id`,`video_id`) COMMENT 'user_video unique index',
    KEY          `idx_video_id` (`video_id`) COMMENT 'video_id index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Favorite table';