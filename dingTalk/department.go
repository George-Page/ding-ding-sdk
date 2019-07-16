package dingTalk

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DepartmentInstance struct {
	Client        Client
}

func NewDepartmentInstance(appKey, appSecret, encodingAesKey, token, suiteKey string) (*DepartmentInstance, error) {
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
	return &DepartmentInstance{
		Client:        *c,
	}, nil
}

func (d *DepartmentInstance) GetAccessToken() AccessTokenRsp {
	params := make(map[string]interface{}, 0)
	params["appkey"] = d.Client.Config.AppKey
	params["appsecret"] = d.Client.Config.AppSecret
	reqUrl := BuildHttpGetParams(d.Client.Config.AccessTokenUrl, params)
	rsp, err := d.Client.Conn.doRequest(http.MethodGet, reqUrl, nil, nil)
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

func (d *DepartmentInstance) ListDepartment(gateWay string, parameters map[string]interface{}) (dept *DepartmentList, err error) {

	reqUrl := BuildHttpGetParams(gateWay, parameters)
	rsp, err := d.Client.Conn.doRequest(http.MethodGet, reqUrl, nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(r, &dept)
	if err != nil {
		return nil, err
	}
	return dept, nil
}

func (d *DepartmentInstance) GetDepartment(gateWay string, parameters map[string]interface{}) (dept *Department, err error) {

	reqUrl := BuildHttpGetParams(gateWay, parameters)
	rsp, err := d.Client.Conn.doRequest(http.MethodGet, reqUrl, nil, nil)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(r, &dept)
	if err != nil {
		return nil, err
	}
	return dept, nil
}
