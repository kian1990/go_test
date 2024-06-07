DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int,
  `name` varchar(100),
  `gender` varchar(100),
  `sfz` varchar(100)
) ENGINE = InnoDB;

INSERT INTO user VALUES (1,'lily','女','421281199308040016');
INSERT INTO user VALUES (2,'lucy','女','421281199607120018');
INSERT INTO user VALUES (3,'kian','男','421281199008020014');
INSERT INTO user VALUES (4,'alex','男','421281199910020014');
INSERT INTO user VALUES (5,'rose','女','421281199302040018');

SET FOREIGN_KEY_CHECKS = 1;
