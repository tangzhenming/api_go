package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint // 用于在数据库中存储帖子和用户之间的关联关系。它是一个外键，指向用户表中的主键（通常是 ID 字段）。当你在查询帖子时，Gorm 会使用这个外键来自动加载关联的用户对象。
	User    User // 用于在程序中存储关联的用户对象。它不会被存储到数据库中，只是用来在程序运行时临时存储数据。
}
