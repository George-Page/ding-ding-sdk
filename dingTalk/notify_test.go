package dingTalk

import (
	"github.com/micro/go-log"
	"testing"
	"github.com/unlonely-river/test_data"
)

func TestNotifyInstance_WorkMessageTextType(t *testing.T) {
	work, err := NewNotifyInstance(test_data.TestAppKey, test_data.TestAppSecret, test_data.TestAppAesKey, test_data.TestAppToken, test_data.TestAppSuiteKey)
	if err != nil {
		log.Fatalf("new notify instance err:%v", err)
	}

	work.SetAgentId(test_data.TestAgentId)
	work.SetUserIdList(test_data.TestUserId)

	msg := TextMessage{
		MsgType: "text",
		Text: TextNotify{
			Content: "这是第一个文本消息测试，代号：message-1000",
		},
	}
	work.SetMsg(msg)
	task, err := work.WorkMessage(test_data.TestWorkNotifyUrl, work.GetAccessToken().AccessToken)
	log.Logf("create notify task:%#v", task)
	log.Logf("err:%v", err)
}

func TestNotifyInstance_WorkMessageOAType(t *testing.T) {
	work, err := NewNotifyInstance(test_data.TestAppKey, test_data.TestAppSecret, test_data.TestAppAesKey, test_data.TestAppToken, test_data.TestAppSuiteKey)
	if err != nil {
		log.Fatalf("new bpms instance err:%v", err)
	}

	work.SetAgentId(test_data.TestAgentId)
	work.SetUserIdList(test_data.TestUserId)

	msg := OAMessage{
		MsgType: "oa",
		OA: oa{
			MessageUrl: "http://www.linewin.cc",
			Head: OaHead{
				BgColor: "FFBBBBBB",
				Text: "测试101",
			},
		},
	}
	form := make([]oaBodyForm, 0)
	form = append(form, oaBodyForm{Key: "姓名：", Value: "王五"})
	form = append(form, oaBodyForm{Key: "年龄：", Value: "19"})
	form = append(form, oaBodyForm{Key: "身高：", Value: "180cm"})
	form = append(form, oaBodyForm{Key: "体重：", Value: "130KG"})
	form = append(form, oaBodyForm{Key: "学历：", Value: "高中"})
	form = append(form, oaBodyForm{Key: "爱好：", Value: "打篮球、听歌"})
	msg.OA.Body = OaBody{Title:"车险快到期了", Form: form}
	work.SetMsg(msg)
	task, err := work.WorkMessage(test_data.TestWorkNotifyUrl, work.GetAccessToken().AccessToken)
	log.Logf("create notify task:%#v", task)
	log.Logf("err:%v", err)
}