package license

import (
	"github.com/bangwork/bang-api/app/utils/config"
	"github.com/bangwork/bang-api/app/utils/log"
	"github.com/fsnotify/fsnotify"
)

func watcher() {
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error("watch file err: %v", err)
	}
	defer watch.Close()

	licenseWatchPath := config.String("license_watch_path", "conf/")
	fileName := config.String("license_path", "conf/ones_license.crt")
	err = watch.Add(licenseWatchPath)
	if err != nil {
		log.Error("watch file err: %v", err)
	}
	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					if !LicenceExpireSwitch && fileName == ev.Name && ((ev.Op&fsnotify.Create == fsnotify.Create) || (ev.Op&fsnotify.Write == fsnotify.Write) ||
						(ev.Op == fsnotify.Remove) || (ev.Op == fsnotify.Rename) || (ev.Op == fsnotify.Chmod)) {
						err := InitLicense()
						if err == nil {
							LicenceExpireSwitch = true
							//使用return会跳出for循环
							return
						}
					}
				}
			case err := <-watch.Errors:
				{
					log.Error("watch file err: %v", err)
					return
				}
			}
		}
	}()

	select {}
}
