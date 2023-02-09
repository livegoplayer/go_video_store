CREATE TABLE `vs_video_tag`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT,
    `video_id`     int(10) unsigned NOT NULL DEFAULT 0,
    `tag_id`       int(10) unsigned NOT NULL DEFAULT 0,
    `add_datetime` timestamp        NOT NULL DEFAULT current_timestamp(),
    `upt_datetime` timestamp        NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uni` (`video_id`, `tag_id`) USING BTREE,
    KEY `get videos` (`tag_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 43525
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC