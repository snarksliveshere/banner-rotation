package integration_test

import (
	"encoding/json"
	"fmt"
	"time"
)

func (test *notifyTest) notificationSendAddClickBannerMessageMustContainTypeAndBannerAndSlotAndAudience(eventType, banner, slot, audience string) error {
	time.Sleep(3 * time.Second)

	test.messagesMutex.RLock()
	defer test.messagesMutex.RUnlock()

	stat := BannerStatistics{}

	err := json.Unmarshal(test.messages[len(test.messages)-1], &stat)
	panicOnErr(err)
	if stat.Audience != audience ||
		stat.Type != eventType ||
		stat.Slot != slot ||
		stat.Banner != banner {
		return fmt.Errorf("in method notificationSendAddClickBannerMessageMustContainTypeAndBannerAndSlotAndAudience "+
			"data not equal in expected/given format=event:%v/%v,audience:%v/%v,slot:%v/%v,banner:%v/%v\n",
			eventType, stat.Type, audience, stat.Audience, slot, stat.Slot, banner, stat.Banner,
		)
	}
	return nil
}

func (test *notifyTest) errorNotificationSendAddClickBannerMessageMustContainTypeAndBannerAndSlotAndAudience(eventType, banner, slot, audience string) error {
	time.Sleep(3 * time.Second)

	test.messagesMutex.RLock()
	defer test.messagesMutex.RUnlock()

	stat := BannerStatistics{}

	err := json.Unmarshal(test.messages[len(test.messages)-1], &stat)
	panicOnErr(err)

	if stat.Audience != audience ||
		stat.Type != eventType ||
		stat.Slot != slot ||
		stat.Banner != banner {
		test.errorRabbit = "wrong type"
		return nil
	}
	return fmt.Errorf("there is no error in error method:%s\n", "errorNotificationSendAddClickBannerMessageMustContainTypeAndBannerAndSlotAndAudience")
}
