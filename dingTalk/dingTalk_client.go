package dingTalk

import (
	"crypto/aes"
	"encoding/base64"
	"crypto/cipher"
	"encoding/binary"
	"sort"
	"fmt"
	"time"
)

type (
	// dingTalk client
	Client struct {
		Config *Config
		Conn   *Conn
	}
)

type EncryptMsg struct {
	Signature string `json:"msg_signature"`
	Encrypt   string `json:"encrypt"`
	TimeStamp string `json:"timeStamp"`
	Nonce     string `json:"nonce"`
}

func NewClient(appKey, appSecret, encodingAesKey, token, suiteKey string) (*Client, error) {

	config := getDefaultDingTalkConfig()
	config.AppKey = appKey
	config.AppSecret = appSecret
	config.AppAesKey = encodingAesKey
	config.Token = token
	config.SuiteKey = suiteKey

	var err error
	// dingTalk client
	c := &Client{
		Config: config,
	}
	c.Config.Key, err = base64.StdEncoding.DecodeString(encodingAesKey + "=")
	if err != nil {
		return c, err
	}
	c.Config.Block, err = aes.NewCipher(c.Config.Key)
	if err != nil {
		return c, err
	}

	return c, nil
}

// 签名
func (c *Client) EncryptMsg(plain, timeStamp, nonce string) (uint64, *EncryptMsg) {
	randomStr := GetRandomString(DefaultRandomNum)
	plainSize := make([]byte, 4)
	binary.BigEndian.PutUint32(plainSize, uint32(len(plain)))
	plain = fmt.Sprintf("%s%s%s%s", randomStr, plainSize, plain, c.Config.AppKey)
	text := Encode([]byte(plain), c.Config.Block.BlockSize())
	if len(text)%aes.BlockSize != 0 {
		return EncryptAESError, nil
	}
	bm := cipher.NewCBCEncrypter(c.Config.Block, c.Config.Key[:c.Config.Block.BlockSize()])
	cipherText := make([]byte, len(text))
	bm.CryptBlocks(cipherText, text)
	enStr := base64.StdEncoding.EncodeToString(cipherText)
	if len(timeStamp) == 0 {
		timeStamp = fmt.Sprintf("%d", time.Now().Unix())
	}
	signature := c.getSha1(c.Config.Token, timeStamp, nonce, string(enStr))
	encrypt := &EncryptMsg{
		Signature: signature,
		Encrypt:   enStr,
		TimeStamp: timeStamp,
		Nonce:     nonce,
	}
	return OK, encrypt
}

// 密文解密
// 参数详情: signature 签名字符串, timeStamp 时间戳, nonce 随机字符串, encrypt 密文
func (c *Client) DecryptMsg(signature, timeStamp, nonce, encrypt string) (uint64, string) {
	// 验证签名
	if c.getSha1(c.Config.Token, timeStamp, nonce, encrypt) != signature {
		return ValidateSignatureError, ""
	}
	decode, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return ComputeSignatureError, ""
	}
	if len(decode) < aes.BlockSize {
		return DecryptAESError, ""
	}
	bm := cipher.NewCBCDecrypter(c.Config.Block, c.Config.Key[:c.Config.Block.BlockSize()])
	pt := make([]byte, len(decode))
	bm.CryptBlocks(pt, decode)
	pt = Decode(pt)
	size := binary.BigEndian.Uint32(pt[16 : 16+4])
	pt = pt[16+4:]
	corpid := pt[size:]
	if string(corpid) != c.Config.AppKey {
		return ValidateSuiteKeyError, ""
	}

	return OK, string(pt[:size])
}

// 生成sha1
func (c *Client) getSha1(token, timeStamp, nonce, encryptMsg string) string {
	// 先将参数值进行排序
	params := make([]string, 0)
	params = append(params, token)
	params = append(params, encryptMsg)
	params = append(params, timeStamp)
	params = append(params, nonce)
	sort.Strings(params)
	return Sha1Sign(params[0] + params[1] + params[2] + params[3])
}



