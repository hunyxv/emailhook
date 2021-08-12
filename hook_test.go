package emailhook

import (
	"errors"
	"testing"

	"github.com/hunyxv/utils/emailnotify"
	"github.com/sirupsen/logrus"
)



func Init() {
	logrus.SetReportCaller(true) // 日志文件中将记录简单的堆栈（打印日志的地方、函数）

	emailHook := NewMailHook("testapp", emailnotify.WithAccountUserName("hunyxv@gmail"),
		emailnotify.WithAccountPasswd("xxxxxxxxxx"), emailnotify.WithSMTPHost("smtp.gmail.com"),
		emailnotify.WithSMTPPort(465))
	emailHook.AddNotifyMem("hunyxv@gmail")
	logrus.AddHook(emailHook)
}

func TestLogger(t *testing.T) {
	Init()

	// errors 使用github.com/pkg/errors， 发送的邮件中将携带详细堆栈信息（日志文件中不会有）
	logrus.WithError(errors.New("a error")).Error("test err")
}
