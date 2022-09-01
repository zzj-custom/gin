// Package logger
package logger

//type SubscribeMap map[logrus.Level][]*email.Receiver
//type SubscribeHook struct {
//	subMap SubscribeMap
//}
//
//// Levels 此处可以自实现hook 目前使用三方hook
//func (h *SubscribeHook) Levels() []logrus.Level {
//	return logrus.AllLevels
//}
//
//func (h *SubscribeHook) Fire(entry *logrus.Entry) error {
//	for level, receivers := range h.subMap {
//		//命中 准备消费
//		if level == entry.Level {
//			if len(receivers) > 0 {
//				email.SendEmail(receivers, fmt.Sprintf("%s:[系统日志警报]", entry.Level.String()),
//					fmt.Sprintf("错误内容: %s", entry.Message))
//			}
//		}
//	}
//	return nil
//}
//func NewSubscribeMap(level logrus.Level, receiverStr string) SubscribeMap {
//	subMap := SubscribeMap{}
//	addressList := strings.Split(receiverStr, ";")
//	var receivers []*email.Receiver
//	for _, address := range addressList {
//		receivers = append(receivers, &email.Receiver{Email: address})
//	}
//	subMap[level] = receivers
//	return subMap
//}
//func newSubScribeHook(subMap SubscribeMap) *SubscribeHook {
//	return &SubscribeHook{subMap}
//}
