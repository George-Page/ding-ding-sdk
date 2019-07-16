package dingTalk

import (
	"github.com/micro/go-log"
	"testing"
	"github.com/unlonely-river/test_data"
)


func TestUserInstance_GetUser(t *testing.T) {
	c, err := NewUserInstance(test_data.TestAppKey, test_data.TestAppSecret, test_data.TestAppAesKey, test_data.TestAppToken, test_data.TestAppSuiteKey)
	if err != nil {
		log.Fatalf("new bpms instance err:%v", err)
	}

	m := make(map[string]interface{}, 0)
	m["access_token"] = c.GetAccessToken().AccessToken
	m["userid"] =  "42374520993566"
	user, err := c.GetUser(test_data.TestUserGetUrl, m)
	log.Logf("get user instance:%#v", user)
	log.Logf("err:%v", err)
}


func TestUserInstance_ListUser(t *testing.T) {
	c, err := NewUserInstance(test_data.TestAppKey, test_data.TestAppSecret, test_data.TestAppAesKey, test_data.TestAppToken, test_data.TestAppSuiteKey)
	if err != nil {
		log.Fatalf("new user instance err:%v", err)
	}

	m := make(map[string]interface{}, 0)
	m["access_token"] = c.GetAccessToken().AccessToken
	m["department_id"] =  1
	m["offset"] = 0
	m["size"] = 100
	list, err := c.ListUser(test_data.TestUserListUrl, m)
	log.Logf("list user:%#v", list)
	log.Logf("err:%v", err)
}
