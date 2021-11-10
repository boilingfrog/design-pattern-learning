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

// M品牌手机手机
type HandsetBrandM struct {
	HandsetSoft
}

func NewHandsetBrandM() *HandsetBrandM {
	return &HandsetBrandM{}
}

func (hw *HandsetBrandM) SetHandsetSoft(soft HandsetSoft) {
	hw.HandsetSoft = soft
}

func (hw *HandsetBrandM) Run() string {
	return "M品牌的手机-" + hw.HandsetSoft.Run()
}

// N品牌的手机
type HandsetBrandN struct {
	HandsetSoft
}

func NewHandsetBrandN() *HandsetBrandN {
	return &HandsetBrandN{}
}

func (ap *HandsetBrandN) SetHandsetSoft(soft HandsetSoft) {
	ap.HandsetSoft = soft
}

func (ap *HandsetBrandN) Run() string {
	return "N品牌的手机-" + ap.HandsetSoft.Run()
}
