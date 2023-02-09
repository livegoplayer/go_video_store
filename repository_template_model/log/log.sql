CREATE TABLE `vs_log`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT,
    `cat`          int(10) unsigned NOT NULL DEFAULT 0 COMMENT '0: 普通日志\r\n1：需要管理员介入的日志',
    `level`        int(5) unsigned  NOT NULL DEFAULT 0 COMMENT '存入level所代表的数字',
    `message`      varchar(500)     NOT NULL COMMENT '日志记录',
    `data`         text             NOT NULL COMMENT '数据存根，使用json',
    `add_datetime` datetime(3)      NOT NULL DEFAULT current_timestamp(3) COMMENT '操作时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `level` (`level`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 3203
  DEFAULT CHARSET = utf8mb4
  ROW_FORMAT = DYNAMIC