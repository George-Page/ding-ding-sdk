package dingTalk

import (
	"github.com/micro/go-log"
	"testing"
	"github.com/unlonely-river/test_data"
)

func TestDepartmentInstance_GetDepartment(t *testing.T) {
	c, err := NewDepartmentInstance(test_data.TestAppKey, test_data.TestAppSecret, test_data.TestAppAesKey, test_data.TestAppToken, test_data.TestAppSuiteKey)
	if err != nil {
		log.Fatalf("new bpms instance err:%v", err)
	}
	m := make(map[string]interface{})
	m["access_token"] = c.GetAccessToken().AccessToken
	m["id"] = 23863770
	dept, err := c.GetDepartment(test_data.TestDepartmentGetUrl, m)
	log.Logf("get department instance:%#v", dept)
	log.Logf("err:%v", err)
}

func TestDepartmentInstance_GetDepartmentList(t *testing.T) {

	dept, err := NewDepartmentInstance(test_data.TestAppKey, test_data.TestAppSecret, test_data.TestAppAesKey, test_data.TestAppToken, test_data.TestAppSuiteKey)
	if err != nil {
		log.Fatalf("new department instance err:%v", err)
	}
	m := make(map[string]interface{})
	m["access_token"] = dept.GetAccessToken().AccessToken
	depts, err := dept.ListDepartment(test_data.TestDepartmentListUrl, m)
	log.Logf("department list:%#v", depts)
	log.Logf("err:%v", err)
}