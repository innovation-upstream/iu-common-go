package init

import (
	"os"

	"github.com/bugsnag/bugsnag-go"
)

func InitBugsnag(serviceName string, packages []string) {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          os.Getenv("BUGSNAG_API_KEY"),
		ReleaseStage:    os.Getenv("BUGSNAG_RELEASE_STAGE"),
		ProjectPackages: packages,
	})
	bugsnag.OnBeforeNotify(func(e *bugsnag.Event, c *bugsnag.Configuration) error {
		e.MetaData.Add("Service", "Name", serviceName)
		return nil
	})
}
