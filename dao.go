package main

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"log"
)

type RankDemoRepo struct {
	DB  *gorm.DB
	rdb *redis.Client
	log *log.Logger
}

type RankDemoDao interface {
	GetUserIntegral(userId int64) int64
	UpdateUserIntegral(userId, integral int64) bool //采用延时双删策略，保证数据一致
	UserMonthlySortByIntegral() ([]UserRankInfo, error)
	GetUserRank(userId int64) int64
	GetUponUser(userId int64) ([]UserRankInfo, error)
}

type UserRankInfo struct {
	Name  string
	Score int64
	Rank  int64
}
