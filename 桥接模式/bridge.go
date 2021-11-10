package 桥接模式

// 手机软件
type HandsetSoft interface {
	Run() string
}

// 手机游戏
type HandsetGame struct {
}

func (hg *HandsetGame) Run() string {
	return "运行手机游戏"
}

// 手机通讯录
type HandsetAddressList struct {
}

func (hg *HandsetAddressList) Run() string {
	return "运行手机通讯录"
}

// 手机品牌
type HandsetBrand interface {
	SetHandsetSoft(HandsetSoft)
}

// 华为手机
type Huawei struct {
	HandsetSoft
}

func NewHuawei() *Huawei {
	return &Huawei{}
}

func (hw *Huawei) SetHandsetSoft(soft HandsetSoft) {
	hw.HandsetSoft = soft
}

func (hw *Huawei) Run() string {
	return "Huawei手机-" + hw.HandsetSoft.Run()
}

// 苹果手机
type Apple struct {
	HandsetSoft
}

func NewApple() *Apple {
	return &Apple{}
}

func (ap *Apple) SetHandsetSoft(soft HandsetSoft) {
	ap.HandsetSoft = soft
}

func (ap *Apple) Run() string {
	return "Apple手机-" + ap.HandsetSoft.Run()
}
