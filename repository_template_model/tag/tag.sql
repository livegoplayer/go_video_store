CREATE TABLE `vs_tag`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT,
    `tag_id`       int(10) unsigned NOT NULL DEFAULT 0 COMMENT '0-100 视频大类 \r\n100-200 电影分类 \r\n200-300 电影国家\r\n300-400 电影语言\r\n1930-3000 电影年份 特殊，预留，有一个电影插入一个\r\n\r\n3000000000-4294967295 电影演员列表\r\n1000000000-200000000 电影导演列表\r\n2000000000-300000000 电影编剧列表\r\n',
    `tag_name`     varchar(255)     NOT NULL DEFAULT '' COMMENT '名称唯一',
    `tag_url`      varchar(255)     NOT NULL DEFAULT '' COMMENT '有些tag是有外链的',
    `add_datetime` timestamp        NOT NULL DEFAULT current_timestamp(),
    `upt_datetime` timestamp        NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    KEY `tag_id` (`tag_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 22647
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC