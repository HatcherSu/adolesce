package code

import (
	"encoding/xml"
	"strconv"
)

type OpenConf struct {
	AppId       string `json:"app_id"`
	VerifyToken string `json:"verify_token"`
	SecretKey   string `json:"secret_key"`
}

type EncryptMsg struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string   `xml:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature"`
	TimeStamp    string   `xml:"TimeStamp"`
	Nonce        string   `xml:"Nonce"`
}

// 加密
func Encrypt(openConf *OpenConf, xmlData string, timeStamp string, nonce int) (string, error) {
	e, err := NewWechatCryptor(openConf.AppId, openConf.VerifyToken, openConf.SecretKey)
	if err != nil {
		return "", err
	}
	str, err := e.EncryptMsg(xmlData, timeStamp, strconv.Itoa(nonce))
	if err != nil {
		return "", err
	}

	return str, nil
}

// 解密 只提供方法，参数必须传入与上方加密方法相同的openConf才可以对加密文件进行解密
func Decrypt(openConf *OpenConf, xmlData string) (string, error) {
	e, err := NewWechatCryptor(openConf.AppId, openConf.VerifyToken, openConf.SecretKey)
	if err != nil {
		return "", err
	}
	encryptMsg := EncryptMsg{}
	err = xml.Unmarshal([]byte(xmlData), &encryptMsg)
	if err != nil {
		return "", err
	}
	s, err := e.DecryptMsg(encryptMsg.MsgSignature, encryptMsg.TimeStamp, encryptMsg.Nonce, xmlData)
	if err != nil {
		return "", err
	}

	return s, nil
}
