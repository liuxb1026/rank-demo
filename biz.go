package main

import "time"

type RankDemoBiz struct {
	dao RankDemoDao
}

var MaxTimeStamp = int64(5000000000000)

func (rank RankDemoBiz) UpdateUserIntegral(userId, integral int64) bool {
	timeStamp := time.Now().UnixMilli()
	timeDifference := MaxTimeStamp - timeStamp
	userIntegral := rank.dao.GetUserIntegral(userId)
	if userIntegral >= 10000 {
		return false
	}
	if userIntegral+integral > 10000 {
		integral = 10000
	} else {
		integral = userIntegral + integral
	}
	integral = integral*10000000000000 + timeDifference
	return rank.dao.UpdateUserIntegral(userId, integral)
}

func (rank RankDemoBiz) UserMonthlySortByIntegral() ([]UserRankInfo, error) {
	//todo 不需一次性获取全部数据，根据前端分页来获取指定偏移量数据，再对score的值进行时间戳处理，得到具体积分。
	//todo 具体排名可依据offset计算，或遍历查询redis，或前端自行计算。
	return rank.dao.UserMonthlySortByIntegral()
}

func (rank RankDemoBiz) UserRankQuery(userId int64) ([]UserRankInfo, error) {
	userRank := rank.dao.GetUserRank(userId)
	//todo 根据当前userRank再去获取其前后十名用户
	resp := []UserRankInfo{}
	resp, err := rank.dao.GetUponUser(userRank)
	return resp, err
}
