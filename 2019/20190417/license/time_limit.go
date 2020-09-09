package license

import (
	"fmt"
	"github.com/bangwork/bang-api/app/models/organization"
	"time"

	"github.com/bangwork/bang-api/app/models"
	stampModel "github.com/bangwork/bang-api/app/models/stamp"
	// "github.com/bangwork/bang-api/app/utils/log"
	gorp "gopkg.in/gorp.v1"
)

var (
	LicenceExpireSwitch = true            //LicenceExpireSwitch license过期时(LicenceExpireSwitch==false)，设置所有用户token失效。 慎用.false表示所有用户均不能登录。
	systemTimeOffset    *SystemTimeOffset // 定时检测系统时间是否有改动
	reporter            *LogReporter
	timeIntervalError   = 24 * 60 * 60 // 允许的误差时间间隔
	minTime             = 1544511536   // 2018/12/11 14:58:56

)

type SystemTimeOffset struct {
	timeOld time.Time
	Offset  int64
}

func CheckTimeLimit() {
	initSystemTimeOffset()
	go func() {
		for {
			licenseDatas := GetLicenseDatas()
			tInt64 := time.Now().Unix()
			org, _ := organization.GetPrivateDeployOrganization(models.DBM)
			if org != nil {
				if reporter == nil {
					reporter = InitLogReporter(licenseDatas.OrgIdentity, org.UUID, org.Name)
				}
				if int64(org.Scale) > licenseDatas.MaxMemberCount {
					reporter.RecordMsg("organization scale exceed the limit")
					// DisableAllUserLogin()
				}
				realMemberCount, _ := organization.GetOrgMemberCount(models.DBM, org.UUID)
				if int64(realMemberCount) > licenseDatas.MaxMemberCount {
					reporter.RecordMsg("team member count exceed the limit")
					// DisableAllUserLogin()
				}
			} else {
				//允许取不到机构信息，取不到则不校验license
				time.Sleep(time.Second * 60)
				continue
			}

			if tInt64 > licenseDatas.ExpiredTime {
				reporter.RecordMsg("product license has expired")
				// DisableAllUserLogin()
			}

			if tInt64 < int64(minTime) {
				reporter.RecordMsg("invalid system date")
				// DisableAllUserLogin()
			}

			if tInt64 > licenseDatas.ProjectExpireTime {
				reporter.RecordMsg("project license has expired")
				// DisableAllUserLogin()
			}
			if tInt64 > licenseDatas.WikiExpireTime {
				reporter.RecordMsg("wiki license has expired")
				// DisableAllUserLogin()
			}
			if tInt64 > licenseDatas.TestcaseExpireTime {
				reporter.RecordMsg("testcase license has expired")
				// DisableAllUserLogin()
			}
			if tInt64 > licenseDatas.PipelineExpireTime {
				reporter.RecordMsg("pipeline license has expired")
				// DisableAllUserLogin()
			}

			systemTimeOffset.CheckTimeOffset()
			time.Sleep(time.Second * 60)
		}
	}()

	go func() {
		for {
			// 这里不要过于频繁
			checkTimeLimitWithTimeStamp(models.DBM)
			time.Sleep(time.Hour)
		}
	}()
}

func checkTimeLimitWithTimeStamp(src gorp.SqlExecutor) {
	licenseDatas := GetLicenseDatas()
	check := func(max int64) {
		t := time.Now()
		if max > 0 {
			if t.Unix() < max-int64(timeIntervalError) {
				reporter.RecordMsg("invalid system date")
				// 当 用户数据的最大时间戳 - 误差时间间隔  < 当前系统时间时 强制退出
				// 这种情况发生在 用户大幅度修改系统时间的时候
				// DisableAllUserLogin()
			}
			if max > licenseDatas.ExpiredTime {
				reporter.RecordMsg("product license has expired")
				// 用户数据的最大时间戳 不能 超过限制日期， 超过则强制退出
				// DisableAllUserLogin()
			}
		}
	}

	max, _ := stampModel.GetMaxUserAccessTime(src)
	check(max)
	max, _ = stampModel.GetMaxTaskUpdateTime(src)
	check(max / 1e6)
	max, _ = stampModel.GetMaxOrgUpdateTime(src)
	check(max / 1e6)
}

func initSystemTimeOffset() {
	systemTimeOffset = &SystemTimeOffset{
		timeOld: time.Now(),
		Offset:  0,
	}
}
func (this *SystemTimeOffset) CheckTimeOffset() {
	timeNew := time.Now()
	this.timeOld = timeNew
	offset := timeNew.Sub(this.timeOld) - 60*time.Second
	if abs(int64(offset)) > 2*int64(time.Second) {
		this.Offset += int64(offset)
		if abs(this.Offset) > 60*int64(time.Second) {
			updateMinute := int64(this.Offset / (60 * int64(time.Second)))
			fmt.Printf("dangerous operation: system time has been modified,updateMinute=%d min", updateMinute)
			reporter.RecordMsg(fmt.Sprintf("dangerous operation: system time has been modified,updateMinute=%d min", updateMinute))
		}
	}
}

func abs(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}

//DisableAllUserLogin 禁用所有用户登录
func DisableAllUserLogin() {
	LicenceExpireSwitch = false
	go watcher()
}
