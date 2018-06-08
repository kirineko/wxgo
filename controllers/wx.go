package controllers

import (
	"crypto/sha1"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"sort"
	"strings"
)

type WxConnectController struct {
	beego.Controller
}

const Token = "hellowx2018"

func (c *WxConnectController) Get() {

	timestamp, nonce, signatureIn := c.GetString("timestamp"), c.GetString("nonce"), c.GetString("signature")
	signatureGen := makeSignature(timestamp, nonce)

	if signatureGen != signatureIn {
		fmt.Printf("signatureGen != signatureIn signatureGen=%s,signatureIn=%s\n", signatureGen, signatureIn)
		c.Ctx.WriteString("")

	} else {
		//如果请求来自于微信，则原样返回echostr参数内容 以上完成后，接入验证就会生效，开发者配置提交就会成功。
		echostr := c.GetString("echostr")
		c.Ctx.WriteString(echostr)
	}
}

func makeSignature(timestamp, nonce string) string {

	//1. 将 plat_token、timestamp、nonce三个参数进行字典序排序
	sl := []string{Token, timestamp, nonce}
	sort.Strings(sl)
	//2. 将三个参数字符串拼接成一个字符串进行sha1加密
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))

	return fmt.Sprintf("%x", s.Sum(nil))
}
