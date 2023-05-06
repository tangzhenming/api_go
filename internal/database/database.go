package database

import "database/sql"

// 初始化一个 database/sql 的引用；这个变量在 database 包中数据全局变量，只要引用了 database 这个 package 的模块都可以访问到
var DB *sql.DB
