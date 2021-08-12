package emailhook

import (
	"fmt"
	"strings"

	"github.com/hunyxv/utils/emailnotify"
	"github.com/sirupsen/logrus"
)

var _ logrus.Hook = (*MailHook)(nil)

// MailHook logrus 发送邮件的 hook
type MailHook struct {
	AppName string            // 应用名称
	mail    *emailnotify.Mail // 发送邮件
}

// NewMailHook .
func NewMailHook(appName string, opts ...emailnotify.Option) *MailHook {
	mail := emailnotify.NewMail(opts...)
	return &MailHook{
		AppName: appName,
		mail:    mail,
	}
}

// AddNotifyMem 添加通知成员
func (hook *MailHook) AddNotifyMem(to ...string) error {
	for _, email := range to {
		nick := strings.Split(email, "@")[0]
		if err := hook.mail.To.Add(nick, email); err != nil {
			return err
		}
	}
	return nil
}

// FlushedNotifyMem 刷新通知成员（删除原来的，重新加入）
func (hook *MailHook) FlushedNotifyMem(to ...string) error {
	hook.mail.To.Flushall()
	return hook.AddNotifyMem(to...)
}

// Levels returns the available logging levels.
func (hook *MailHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

// Fire is called when a log event is fired.
func (hook *MailHook) Fire(entry *logrus.Entry) error {
	title := fmt.Sprintf("[%s]:[%s] - %s", strings.ToUpper(entry.Level.String()), hook.AppName, entry.Message)
	stack := make(map[string]string, 1)
	var errcarrystack bool
	if e, ok := entry.Data["error"]; ok {
		if err, ok := e.(error); ok {
			errmsg := fmt.Sprintf("%+v", err)
			lines := strings.Split(errmsg, "\n")
			// err 有堆栈信息
			if len(lines) > 1 {
				errcarrystack = true
				for i := 0; i < len(lines)-1; i = i + 2 {
					stack[lines[i]] = lines[i+1]
				}
			}
		}
	}
	if !errcarrystack && entry.HasCaller() {
		path := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
		fun := entry.Caller.Function
		stack[path] = fun
	}
	dt := NewDefaultTemplate(title, entry.Message, entry.Time, stack, entry.Data)
	return hook.mail.Send(title, dt)
}
