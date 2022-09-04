package main

type RankDemoService struct {
	biz RankDemoBiz
}

func (r RankDemoService) UpdateUserIntegral(userId, integral int64) bool {
	return r.biz.UpdateUserIntegral(userId, integral)
}

func (r RankDemoService) UserMonthlySortByIntegral() ([]UserRankInfo, error) {
	return r.biz.UserMonthlySortByIntegral()
}

func (r RankDemoService) UserRankQuery(userId int64) ([]UserRankInfo, error) {
	return r.biz.UserRankQuery(userId)
}
