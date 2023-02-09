CREATE TABLE `vs_video`
(
    `video_id`     int(10) unsigned NOT NULL AUTO_INCREMENT,
    `pretty_name`  varchar(255)     NOT NULL DEFAULT '' COMMENT '用来显示的name',
    `meta_name`    varchar(255)     NOT NULL DEFAULT '' COMMENT '分国别，用电影原来的名称，确保唯一',
    `add_datetime` timestamp        NOT NULL DEFAULT current_timestamp(),
    `upt_datetime` timestamp        NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`video_id`) USING BTREE,
    KEY `name` (`pretty_name`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1405
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC