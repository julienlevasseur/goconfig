package crontab

import (
	"fmt"

	"github.com/julienlevasseur/goconfig/pkg/file"
	"github.com/julienlevasseur/goconfig/pkg/user"
)

// * * * * * command to be executed
// | | | | |
// | | | | |
// | | | | |
// | | | | |_______________ Day of the Week (0 - 6)(Sunday to Saturday)
// | | | |
// | | | |_______________ Month of the Year (1 - 12)
// | | |
// | | |_______________ Day of the Month (1 - 31)
// | |
// | |_______________ Hour (0 - 23)
// |
// |_______________ Minute (0 - 59)

func Cron(
	minute,
	hour,
	dayOfMonth,
	monthOfYear,
	dayOfWeek,
	command string,
) error {
	username, err := user.Username()
	if err != nil {
		return err
	}

	// Check if entry already exists
	found, err := file.LineIsPresent(
		fmt.Sprintf("/var/spool/cron/crontabs/%s", username),
		fmt.Sprintf(
			"%s %s %s %s %s %s",
			minute,
			hour,
			dayOfMonth,
			monthOfYear,
			dayOfWeek,
			command,
		),
	)
	if err != nil {
		return err
	}

	fmt.Println("[Crontab] Add new crontab entry")
	file.Append(
		fmt.Sprintf("/var/spool/cron/crontabs/%s", username),
		fmt.Sprintf(
			"%s %s %s %s %s %s",
			minute,
			hour,
			dayOfMonth,
			monthOfYear,
			dayOfWeek,
			command,
		),
		&found,
	)

	return nil
}
