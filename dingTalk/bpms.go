package dingTalk

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type BpmsInstance struct {
	Client         Client
	format         string                 // http Content-type
	createInstance map[string]interface{} // 普通审批流表单
	queryInstance  map[string]interface{}  // 查询审批实例
}

func NewBpmsInstance(appKey, appSecret, encodingAesKey, token, suiteKey string) (*BpmsInstance, error) {
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
	return &BpmsInstance{
		Client:         *c,
		createInstance: make(map[string]interface{}, 0),
		queryInstance:  make(map[string]interface{}, 0),
	}, nil
}

func (b *BpmsInstance) SetFormat(format string) {
	b.format = format
}

// // 审批流的唯一码
func (b *BpmsInstance) SetProcessCode(code string) {
	b.createInstance["process_code"] = code
}

// // 发起人
func (b *BpmsInstance) SetDeptId(deptId uint32) {
	b.createInstance["dept_id"] = deptId
}

//
func (b *BpmsInstance) SetOriginatorUserId(uid string) {
	b.createInstance["originator_user_id"] = uid
}

// 审批人
func (b *BpmsInstance) SetApprovers(approvers string) {
	b.createInstance["approvers"] = approvers
}

// 审批人列表，支持会签/或签，优先级高于approvers变量
func (b *BpmsInstance) SetApproversV2(vo ProcessInstanceApproverVo) {
	b.createInstance["approvers_v2"] = vo
}

// starting -------------------
// 抄送人userid列表，最大列表长度：20。多个抄送人用逗号分隔。该参数需要与cc_position参数一起传，抄送人才能生效。
func (b *BpmsInstance) SetCClist(cclist string) {
	b.createInstance["cc_list"] = cclist
}

// 抄送时间，分为（START, FINISH, START_FINISH）
func (b *BpmsInstance) SetCCPosition(action string) {
	b.createInstance["cc_position"] = action
}
// ending ---------------------

func (b *BpmsInstance) SetFormComponent(values []FormValues) {
	b.createInstance["form_component_values"] = values
}

func (b *BpmsInstance) SetProcessInstanceId(id string) {
	b.queryInstance["process_instance_id"] = id
}

func (b *BpmsInstance) GetAccessToken() AccessTokenRsp {
	params := make(map[string]interface{}, 0)
	params["appkey"] = b.Client.Config.AppKey
	params["appsecret"] = b.Client.Config.AppSecret
	reqUrl := BuildHttpGetParams(b.Client.Config.AccessTokenUrl, params)
	rsp, err := b.Client.Conn.doRequest(http.MethodGet, reqUrl, nil, nil)
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

func (b *BpmsInstance) CreateTaskInstance(gateWay, accessToken string) (task *BpmsInstanceTaskRsp, err error) {

	p, err := json.Marshal(b.createInstance)
	b.createInstance = make(map[string]interface{}, 0)
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string, 0)
	if b.format == Format {
		headers[HTTPHeaderContentType] = ContentTypeJson
	}
	params := make(map[string]interface{}, 0)
	params["access_token"] = accessToken
	reqUrl := BuildHttpGetParams(gateWay, params)

	reader := bytes.NewReader(p)
	rsp, err := b.Client.Conn.doRequest(http.MethodPost, reqUrl, headers, reader)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(r, &task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (b *BpmsInstance) GetTaskInstance(gateWay, accessToken string) (taskDetail *BpmsInstanceTaskDetailRsp, err error) {

	p, err := json.Marshal(b.queryInstance)
	if err != nil {
		return nil, err
	}
	headers := make(map[string]string, 0)
	if b.format == Format {
		headers[HTTPHeaderContentType] = ContentTypeJson
	}
	params := make(map[string]interface{}, 0)
	params["access_token"] = accessToken
	reqUrl := BuildHttpGetParams(gateWay, params)

	reader := bytes.NewReader(p)
	rsp, err := b.Client.Conn.doRequest(http.MethodPost, reqUrl, headers, reader)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(r, &taskDetail)
	if err != nil {
		return nil, err
	}
	return taskDetail, nil
}
