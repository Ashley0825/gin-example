package model

type UserInDB struct {
	ID       string
	Username string
	Password string
	VipId    string
	Points   uint
	Coins    uint
}

type VipInDB struct {
	ID       string
	Name     string
	Discount float64
}
