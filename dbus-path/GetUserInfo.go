package main

import (
	"encoding/json"
	"fmt"
	"github.com/godbus/dbus/v5"
	"log"
)

func main() {
	conn, err := dbus.SessionBus()
	if err != nil {
		return
	}

	caller := conn.Object("com.deepin.deepinid", "/com/deepin/deepinid")
	var dsess map[string]interface{}
	var deepinidSess []byte
	err = caller.Call("com.deepin.deepinid.Get", 0).Store(&deepinidSess)
	if err != nil {
		log.Printf("get deepinid userinfo error %v", err)
		return
	}
	log.Printf("get deepinid userinfo data %v", string(deepinidSess))
	err = json.Unmarshal(deepinidSess, &dsess)
	if err != nil {
		log.Printf("decode deepinid userinfo error %v", err)
		return
	}
	var newNum float64
	_, _ = fmt.Sscanf(fmt.Sprint(dsess["UserID"]), "%e", &newNum)
	var sess = map[string]dbus.Variant{
		"uid":           dbus.MakeVariant(int64(newNum)),
		"session_id":    dbus.MakeVariant(dsess["Region"]),
		"region":        dbus.MakeVariant(dsess["SessionID"]),
		"nickname":      dbus.MakeVariant("nickname"),
		"profile_image": dbus.MakeVariant(dsess["ProfileImage"]),
		"username":      dbus.MakeVariant(dsess["Username"]),
	}
	if dsess["Mobile"] == nil {
		sess["mobile"] = dbus.MakeVariant("")
	}
	if dsess["Email"] == nil {
		sess["email"] = dbus.MakeVariant("")
	}
	if dsess["WechatNickname"] == nil {
		sess["wechat_nickname"] = dbus.MakeVariant("")
	}
	err = caller.Call("com.deepin.deepinid.Set", 0, sess).Store()
	if err != nil {
		log.Printf("update deepinid userinfo error %v", err)
		return
	}
}
