CREATE TABLE `vs_video_detail`
(
    `id`                int(10) unsigned      NOT NULL AUTO_INCREMENT,
    `video_id`          int(10) unsigned      NOT NULL DEFAULT 0,
    `season`            int(2) unsigned       NOT NULL DEFAULT 1 COMMENT '季 * 100 如果要得出正确的季/100 就行了 \r\n有编号的是特别篇',
    `video_description` text                  NOT NULL COMMENT '视频简介',
    `video_poster_url`  varchar(255)          NOT NULL DEFAULT '' COMMENT '视频的海报',
    `official_website`  varchar(255)          NOT NULL DEFAULT '' COMMENT '视频官网地址',
    `description_url`   varchar(255)          NOT NULL DEFAULT '' COMMENT '豆瓣介绍页面或者其他',
    `per_update_time`   varchar(50)           NOT NULL DEFAULT '' COMMENT '每次更新时间 ddys: 周几(深夜/凌晨) ',
    `released_date_str` text                  NOT NULL COMMENT '上映时间 str tag库中的release_date是去除地区的，这里做个缓存',
    `year`              int(5) unsigned       NOT NULL DEFAULT 0 COMMENT '年份，tag库存储1+个年份，第一个是该年份，其他年份取自release_date',
    `set_num`           int(5) unsigned       NOT NULL DEFAULT 0 COMMENT '本季集数，仅第一次爬取时候的集数，仅豆瓣有的时候显示为最大， 为0的话表示没有该信息，前端可以显示为当前集数',
    `score`             float(10, 1) unsigned NOT NULL DEFAULT 0.0 COMMENT '豆瓣评分，仅爬取的时候的分数',
    `time_str`          varchar(50)           NOT NULL DEFAULT '' COMMENT '时间str',
    `poster_cache`      varchar(255)          NOT NULL DEFAULT '' COMMENT '豆瓣缓存封面图片',
    `poster_list_url`   varchar(255)          NOT NULL DEFAULT '' COMMENT '豆瓣电影剧照页面',
    `add_datetime`      timestamp             NOT NULL DEFAULT current_timestamp(),
    `upt_datetime`      timestamp             NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    KEY `fetch` (`video_id`, `season`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 1662
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC