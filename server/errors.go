package server

var (
	errJobNotFound         = responseError{"this could return info on a position in your company; contact me!"}
	errStudiesNotFound     = responseError{"haven't yet looked into this"}
	errRefNotFound         = responseError{"still cannot put this person as reference"}
	errAchievementNotFound = responseError{"achievement not yet performed"}
	errSkillNotFound       = responseError{"skill not found, yet"}
)
