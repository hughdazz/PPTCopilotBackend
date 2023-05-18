-- type Favorite struct {
-- 	Id      int       `orm:"auto;pk"`
-- 	User    *User     `orm:"rel(fk)"` // 设置一对多的反向关系
-- 	Project *Project  `orm:"rel(fk)"` // 设置一对多的反向关系
-- 	Created time.Time `orm:"auto_now_add;type(datetime)"`
-- }



CREATE TABLE IF NOT EXISTS `favorite` (
  `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `user_id` integer NOT NULL,
  `project_id` integer NOT NULL,
  `created` datetime NOT NULL
) ENGINE = InnoDB;