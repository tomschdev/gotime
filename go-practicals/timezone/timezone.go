package timezone

import (
	"fmt"
	"time"
)

func tzConvert(t, from, to string) (string, error) {
	zoneOne, err := time.LoadLocation(from)
	if err != nil {
		return "location load failed", err
	}
	zoneTwo, err := time.LoadLocation(to)
	if err != nil {
		return "location load failed", err
	}
	const timeFormat = "1999-05-10T10:00"
	fromTime, err := time.ParseInLocation(timeFormat, t, zoneOne)
	if err != nil {
		return "time parse failed", err
	}
	toTime := fromTime.In(zoneTwo)
	return toTime.Format(timeFormat), nil
}

func main() {
	ts := "1999-05-10T10:00"
	out, err := tzConvert(ts, "America/Los_Angeles", "Asia/Jerusalem")
	if err != nil {
		fmt.Errorf("error: %q", err)
		return
	}
	fmt.Println(out)
}
