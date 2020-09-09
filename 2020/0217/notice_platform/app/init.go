package app

import (
	"fmt"
	"path/filepath"

	"github.com/bangwork/ones-plugin/notice_platform/app/models/db"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/common/i18n"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/distributor"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/filter"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/payload"
	"github.com/bangwork/ones-plugin/notice_platform/app/services/webhook"
	"github.com/bangwork/ones-plugin/notice_platform/app/utils/config"
	"github.com/ngaut/log"
)

func init() {
	if err := log.SetOutputByName("./notice.log"); err != nil {
		panic(fmt.Sprintf("Error at onStart: %s\n", err))
	}
	onStart(config.InitConfigs)
	onStart(db.InitDB)
	onStart(db.InitRedis)
	onStart(initI18N)
	onStart(webhook.InitWebHookMap)
	onStart(filter.InitTaskFilterMap)
	onStart(webhook.InitTunnelMap)
	onStart(filter.InitTeamIssueType)

	go payload.TimeTicker()
	go distributor.Sender()

}

func onStart(fn func() error) {
	if err := fn(); err != nil {
		panic(fmt.Sprintf("Error at onStart: %s\n", err))
	}
}

func initI18N() error {
	return i18n.LoadLocales(BasePath(), "locales", "en")
}

var basePath = func() string {
	dir, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	return dir
}()

func BasePath() string {
	return basePath
}
