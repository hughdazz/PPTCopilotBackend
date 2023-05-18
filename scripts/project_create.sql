/*
 type Project struct {
 Id          int
 Name        string    `orm:"size(100)"`
 Description string    `orm:"size(100)"`
 Creator     *User     `orm:"rel(fk)"` // 设置一对多的反向关系
 Created     time.Time `orm:"auto_now_add;type(datetime)"`
 Updated     time.Time `orm:"auto_now;type(datetime)"`
 }
 */
CREATE TABLE IF NOT EXISTS `project` (
  `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `name` varchar(100) NOT NULL DEFAULT '',
  `description` varchar(100) NOT NULL DEFAULT '',
  `creator_id` integer NOT NULL,
  `star_count` integer NOT NULL DEFAULT 0,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL
) ENGINE = InnoDB;