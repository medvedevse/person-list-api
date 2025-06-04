package entity

type Filters struct {
	Gender      string `form:"gender"`
	Age         int    `form:"age"`
	Nationality string `form:"nationality"`
}
