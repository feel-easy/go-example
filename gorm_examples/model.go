package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// gorm.Model 定义
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// 将字段 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` 注入到 `User` 模型中
type User struct {
	gorm.Model
	Name string
}

// 声明 gorm.Model 模型
// type User struct {
// 	ID   int
// 	Name string
// }
