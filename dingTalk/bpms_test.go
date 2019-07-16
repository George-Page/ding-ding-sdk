package dingTalk

import (
	"github.com/micro/go-log"
	"testing"
	"github.com/unlonely-river/test_data"
)

func TestBpmsInstance_SetAccessToken(t *testing.T) {
	bpms, err := NewBpmsInstance(test_data.TestAppKey, test_data.TestAppSecret, test_data.TestAppAesKey, test_data.TestAppToken, test_data.TestAppSuiteKey)
	if err != nil {
		log.Fatalf("new bpms instance err:%v", err)
	}
	log.Logf("access token:%#v", bpms.GetAccessToken())
}

func TestBpmsInstance_CreateTaskInstance(t *testing.T) {
	bpms, err := NewBpmsInstance(test_data.TestAppKey, test_data.TestAppSecret, test_data.TestAppAesKey, test_data.TestAppToken, test_data.TestAppSuiteKey)
	if err != nil {
		log.Fatalf("new bpms instance err:%v", err)
	}

	bpms.SetProcessCode(test_data.TestAppProcessCode)
	bpms.SetFormat("json")
	bpms.SetDeptId(60749089)
	bpms.SetOriginatorUserId("06634603251154177")

	bpms.SetApprovers("06634603251154177")
	// 有抄送人时，必传position
	bpms.SetCClist("06634603251154177")
	bpms.SetCCPosition(CC_Position_Start)

	fv := make([]FormValues, 0)
	item1 := FormValues{
		Name:     "审批模块",
		Value:    "喝水去",
		ExtValue: "go to wc",
	}
	item2 := FormValues{
		Name:     "审批内容",
		Value:    "起床了",
		ExtValue: "invalid approval",
	}
	fv = append(fv, item1, item2)
	bpms.SetFormComponent(fv)

	task, err := bpms.CreateTaskInstance(test_data.TestCreateTaskUrl, bpms.GetAccessToken().AccessToken)
	log.Logf("create task:%#v", task)
	log.Logf("err:%v", err)
}

func TestBpmsInstance_GetTaskInstance(t *testing.T) {
	bpms, err := NewBpmsInstance(test_data.TestAppKey, test_data.TestAppSecret, test_data.TestAppAesKey, test_data.TestAppToken, test_data.TestAppSuiteKey)
	if err != nil {
		log.Fatalf("new bpms instance err:%v", err)
	}

	bpms.SetFormat("json")
	bpms.SetProcessInstanceId(test_data.TestProcessInstanceId)

	task, err := bpms.GetTaskInstance(test_data.TestQueryProcessInstanceUrl, bpms.GetAccessToken().AccessToken)
	log.Logf("get task instance:%#v", task)
	log.Logf("err:%v", err)
}
