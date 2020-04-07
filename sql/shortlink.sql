SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for hi_short_url
-- ----------------------------
DROP TABLE IF EXISTS `hi_short_url`;
CREATE TABLE `hi_short_url` (
  `id` varchar(32) NOT NULL,
  `short_url` varchar(8) NOT NULL COMMENT '短网址',
  `lang_url` varchar(512) NOT NULL COMMENT '长网址',
  `duration` varchar(12) NOT NULL COMMENT '有效时间',
  `token` varchar(32) DEFAULT NULL COMMENT 'Token',
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for short_url
-- ----------------------------
DROP TABLE IF EXISTS `short_url`;
CREATE TABLE `short_url` (
  `id` varchar(32) NOT NULL,
  `short_url` varchar(8) NOT NULL COMMENT '短网址',
  `lang_url` varchar(512) NOT NULL COMMENT '长网址',
  `duration` varchar(12) NOT NULL COMMENT '有效时间',
  `token` varchar(32) DEFAULT NULL COMMENT 'Token',
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
