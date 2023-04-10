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
CREATE TABLE `user` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `Username` varchar(100) DEFAULT NULL,
  `Password` varchar(100) DEFAULT NULL,
  `Email` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;