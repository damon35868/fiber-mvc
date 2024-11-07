package schedule

import (
	"fiber-mvc/app/common/constant"
	"fiber-mvc/app/service"
	"os"
)

func Boot(s *service.Service) {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == constant.PRO {
		tokenScheduleRun(s)
	}
}
