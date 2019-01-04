# The DataBase source file, if you want use if, please import to your MySQL.
# My MySQL version is 5.5.53, it's also can be use the high version.
# I plan change MySQL to TIDB on future, now first use MySQL.
# DATABASE: MYSQL 5.5.53
# DBHOST: 127.0.0.1
# DBPORT: 3306
# DBNAME: wechat-management
# -------------------------------------------------------------------------------------------------------------------------------------------------------------------------
# -------------------------------------------------------------------------------------------------------------------------------------------------------------------------
# The table of administrator
DROP TABLE IF EXISTS `wt_user`;
CREATE TABLE `wt_user`(
`id` INT AUTO_INCREMENT PRIMARY KEY,
`username` VARCHAR(32) NOT NULL UNIQUE COMMENT "user login name",
`cellphone` VARCHAR(11) NOT NULL DEFAULT "" COMMENT "user's mobilephone No.",
`realname` VARCHAR(32) NOT NULL DEFAULT "" COMMENT "user's realname",
`gender` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT "user's gender, the value of '0' indicate this user's gender is unknow, the value of '1' indicate this user is male, the value of '2' indicate this user is female",
`roleid` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "user's role id, basic on RBAC",
`createtime` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "this user's register time",
`logintime` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "the user last login time",
`loginip` VARCHAR(15) NOT NULL DEFAULT "" COMMENT "the user last login client IP",
`avatar` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "the user's head portrait url",
`signature` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "the user's signature",
`isblacklist` TINYINT UNSIGNED NOT NULL DEFAULT 0 COMMENT "this user is in blacklis or not. value of '0' indicate not in blacklist, another value indicate in blacklist"
)ENGINE = INNODB DEFAULT CHARSET = UTF8;

# The table of administrator role
DROP TABLE IF EXISTS `wt_user_role`;
CREATE TABLE `wt_user_role`(
`id` INT AUTO_INCREMENT PRIMARY KEY,
`rolename` VARCHAR(32) NOT NULL UNIQUE COMMENT "role name",
`authids` VARCHAR(255) NOT NULL DEFAULT '' COMMENT "authorities of this role",
`isshow` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT "is or not show it, value of '0' indicate not show it and value of '1' indicate show it",
`createtime` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "role create time",
`modifytime` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "this record changing time"
)ENGINE = INNODB DEFAULT CHARSET = UTF8;

# The table of role's authorities
DROP TABLE IF EXISTS `wt_user_authority`;
CREATE TABLE `wt_user_authority`(
`id` INT AUTO_INCREMENT PRIMARY KEY,
`authname` VARCHAR(32) NOT NULL UNIQUE COMMENT "authority name",
`isvalid` TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT "this authority is or not in valid, value of '0' indicate not in valid and value of '1' indicate in valid",
`createtime` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "this authority create time",
`modifytime` INT UNSIGNED NOT NULL DEFAULT 0 COMMENT "this record changing time"
)ENGINE = INNODB DEFAULT CHARSET = UTF8;