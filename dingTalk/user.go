package dingTalk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type UserInstance struct {
	Client       Client
}

func NewUserInstance(appKey, appSecret, encodingAesKey, token, suiteKey string) (*UserInstance, error) {
	c, err := NewClient(appKey, appSecret, encodingAesKey, token, suiteKey)
	if err != nil {
		return nil, err
	}
	// http connect
	conn := &Conn{config: c.Config}
	c.Conn = conn
	err = conn.Init(c.Config)
	if err != nil {
		return nil, err
	}
	return &UserInstance{
		Client:       *c,
	}, nil
}

func (u *UserInstance) GetAccessToken() AccessTokenRsp {
	params := make(map[string]interface{}, 0)
	params["appkey"] = u.Client.Config.AppKey
	params["appsecret"] = u.Client.Config.AppSecret
	reqUrl := BuildHttpGetParams(u.Client.Config.AccessTokenUrl, params)
	rsp, err := u.Client.Conn.doRequest(http.MethodGet, reqUrl, nil, nil)
	if err != nil {
		return AccessTokenRsp{}
	}
	r, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return AccessTokenRsp{}
	}
	at := AccessTokenRsp{}
	err = json.Unmarshal(r, &at)
	if err != nil {
		return AccessTokenRsp{}
	}
	return at
}

func (u *UserInstance) GetUser(gateWay string, parameters map[string]interface{}) (user *User, err error) {

	reqUrl := BuildHttpGetParams(gateWay, parameters)
	rsp, err := u.Client.Conn.doRequest(http.MethodGet, reqUrl, nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(r, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserInstance) ListUser(gateWay string, parameters map[string]interface{}) (list *UserList, err error) {
	reqUrl := BuildHttpGetParams(gateWay, parameters)

	rsp, err := u.Client.Conn.doRequest(http.MethodGet, reqUrl, nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(r, &list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
