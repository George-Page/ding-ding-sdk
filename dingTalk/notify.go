package dingTalk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type NotifyInstance struct {
	Client         Client
	format         string                 // http Content-type
	createInstance map[string]interface{} // 发送消息参数
}

func NewNotifyInstance(appKey, appSecret, encodingAesKey, token, suiteKey string) (*NotifyInstance, error) {
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
	return &NotifyInstance{
		Client:         *c,
		format: Format,
		createInstance:  make(map[string]interface{}, 0),
	}, nil
}

// 必传 设置应用agentId 必传
func (n *NotifyInstance) SetAgentId(agentId uint32) {
	n.createInstance["agent_id"] = agentId
}

// 必传 设置接收者的用户userid列表，最大列表长度：100
func (n *NotifyInstance) SetUserIdList(userIdList string) {
	n.createInstance["userid_list"] = userIdList
}

// 可选 接收者的部门id列表，最大列表长度：20,  接收者是部门id下(包括子部门下)的所有用户
func (n *NotifyInstance) SetDeptIdList(deptIdList string) {
	n.createInstance["dept_id_list"] = deptIdList
}

// 可选 是否发送给企业全部用户(ISV不能设置true)
func (n *BpmsInstance) SetToAllUser(toAllUser string) {
	n.createInstance["to_all_user"] = toAllUser
}

func (n *NotifyInstance) SetMsg(msg interface{}) {
	n.createInstance["msg"] = msg
}

func (n *NotifyInstance) GetAccessToken() AccessTokenRsp {
	params := make(map[string]interface{}, 0)
	params["appkey"] = n.Client.Config.AppKey
	params["appsecret"] = n.Client.Config.AppSecret
	reqUrl := BuildHttpGetParams(n.Client.Config.AccessTokenUrl, params)
	rsp, err := n.Client.Conn.doRequest(http.MethodGet, reqUrl, nil, nil)
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

func (n *NotifyInstance) WorkMessage(gateWay, accessToken string) (notify *MessageNotify, err error) {

	p, err := json.Marshal(n.createInstance)
	n.createInstance = make(map[string]interface{}, 0)
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string, 0)
	headers[HTTPHeaderContentType] = ContentTypeJson
	params := make(map[string]interface{}, 0)
	params["access_token"] = accessToken
	reqUrl := BuildHttpGetParams(gateWay, params)

	reader := bytes.NewReader(p)
	rsp, err := n.Client.Conn.doRequest(http.MethodPost, reqUrl, headers, reader)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(r, &notify)
	if err != nil {
		return nil, err
	}
	return notify, nil
}