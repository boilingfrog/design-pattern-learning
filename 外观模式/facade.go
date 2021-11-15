package 外观模式

import "fmt"

type User struct {
}

func (u *User) GetUser(userId int) {
	fmt.Println("获取用户的信息")
}

type GoldCoin struct {
}

func (g *GoldCoin) GetUserGoldCoin(userId int) {
	fmt.Println("获取用户金币的信息")
}

type Order struct {
}

func (o *Order) GetUserOrder(userId int) {
	fmt.Println("获取用户订单信息")
}

func GetUserInfo(userId int) {
	user := User{}
	user.GetUser(userId)

	goldCoin := GoldCoin{}
	goldCoin.GetUserGoldCoin(userId)

	order := Order{}
	order.GetUserOrder(userId)
}
