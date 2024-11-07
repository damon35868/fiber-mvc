package schedule

import (
	"fiber-mvc/app/service"

	"github.com/gofiber/fiber/v2/log"
	"github.com/robfig/cron/v3"
)

func tokenScheduleRun(s *service.Service) {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("@every 30m", func() {
		go (func() {
			log.Info("----------定时刷新token任务开始-----------")
			// s.Common.RefreshAccessToken()
		})()
	})
	c.Start()
}
