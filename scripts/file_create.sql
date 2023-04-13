-- type File struct {
-- 	Id   int
-- 	Name string `orm:"size(100)"`
-- 	// Description string    `orm:"size(100)"`
-- 	Project *Project  `orm:"rel(fk)"` // 设置一对多的反向关系
-- 	Created time.Time `orm:"auto_now_add;type(datetime)"`
-- 	Updated time.Time `orm:"auto_now;type(datetime)"`
-- }
CREATE TABLE IF NOT EXISTS `file` (
  `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `name` varchar(100) NOT NULL DEFAULT '',
  `project_id` integer NOT NULL,
  `created` datetime NOT NULL,
  `updated` datetime NOT NULL
) ENGINE = InnoDB;
