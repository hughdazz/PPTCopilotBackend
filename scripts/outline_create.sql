-- type Outline struct {
-- 	Id      int
-- 	Outline string
-- }

CREATE TABLE IF NOT EXISTS `outline` (
  `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `outline` text NOT NULL DEFAULT ''
) ENGINE = InnoDB;