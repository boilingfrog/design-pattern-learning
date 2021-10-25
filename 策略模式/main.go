package main

import (
	"crypto/rsa"
	"fmt"
	"net/http"
)

func main() {
	wx := WeixinPay{}
	fmt.Println(wx.Name(), wx.Order())

	al := AliPay{}
	fmt.Println(al.Name(), al.Order())
}

// Payment interface
type Payment interface {
	Name() string
	Order() string
}

type WeixinPay struct {
	conf      map[string]string
	notifyURL string
	client    *http.Client
}

func (w *WeixinPay) Name() string {

	return "我是：微信支付；"
}

func (w *WeixinPay) Order() string {

	return "订单号：wx-001；"
}

type AliPay struct {
	privateKey   *rsa.PrivateKey
	aliPublicKey *rsa.PublicKey
	notifyURL    string
	appID        string
	partner      string
	sellerID     string
}

func (a *AliPay) Name() string {

	return "我是：阿里支付；"
}

func (a *AliPay) Order() string {

	return "订单号：al-0010；"
}
