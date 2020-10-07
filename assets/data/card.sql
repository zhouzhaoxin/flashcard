CREATE TABLE `cards` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `front` text NOT NULL,
  `back` text NOT NULL,
  `known` int(10) DEFAULT '1' COMMENT '1:未知，2:已知',
  `tp` int(11) DEFAULT '-1',
  `uid` int(11) DEFAULT '-1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `card_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tp` int(11) DEFAULT '-1',
  `name` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT '',
  `password` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;