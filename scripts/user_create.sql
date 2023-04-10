/*
 创建user表的sql
 ```go
 // 用户信息
 type User struct {
 Id       int
 Username string `orm:"size(100)"`
 Password string `orm:"size(100)"`
 Email    string `orm:"size(100)"`
 }
 ```
 */
CREATE TABLE IF NOT EXISTS `user` (
  `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `username` varchar(100) NOT NULL DEFAULT '',
  `password` varchar(100) NOT NULL DEFAULT '',
  `email` varchar(100) NOT NULL DEFAULT ''
) ENGINE = InnoDB;