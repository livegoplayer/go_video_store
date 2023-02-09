CREATE TABLE `vs_video_resources`
(
    `id`                 int(10) unsigned NOT NULL AUTO_INCREMENT,
    `video_id`           int(10) unsigned NOT NULL DEFAULT 0,
    `season`             int(2) unsigned  NOT NULL DEFAULT 1 COMMENT '季 * 100 如果要得出正确的季/100 就行了 \r\n有编号的是特别篇',
    `episode`            int(2) unsigned  NOT NULL DEFAULT 1 COMMENT '集，分属季的集',
    `video_refer_domain` varchar(255)     NOT NULL DEFAULT '' COMMENT '视频原来的网站地址',
    `video_refer`        int(5) unsigned  NOT NULL DEFAULT 0 COMMENT '1： ddys.art ',
    `video_url`          varchar(255)     NOT NULL DEFAULT '' COMMENT '视频原来的播放地址',
    `name`               varchar(255)     NOT NULL DEFAULT '' COMMENT '源网站对该集或者视频的播放item描述',
    `last_update_time`   int(10) unsigned NOT NULL COMMENT '上次更新时间，一般获取某个video_id的最后一集的更新时间就是该视频的更新时间，其他的没用，这个是用来判断是否爬取最新信息的依据之一，其他还有剧集的集数等等， 网站的最后更新时间解析出来都是一样的，所以直接用大小判断即可',
    `data`               text             NOT NULL COMMENT '数据，一般存储的是有益于获取视频播放源的数据    \r\nddys.art:  有 id           ',
    `add_datetime`       timestamp        NOT NULL DEFAULT current_timestamp(),
    `upt_datetime`       timestamp        NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
    PRIMARY KEY (`id`) USING BTREE,
    KEY `video_index` (`video_id`, `season`, `episode`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 11877
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC