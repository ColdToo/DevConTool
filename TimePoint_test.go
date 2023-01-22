package DevConTool

import (
	"testing"
	"time"
)

func TestNewTimeAxis(t *testing.T) {
	start, _ := time.Parse(TimeFormat, "2022-12-01 00:00:00")
	end, _ := time.Parse(TimeFormat, "2022-12-19 23:59:59")
	t.Log(start.Unix())
	axis := NewTimeAxis(start.Unix(), end.Unix(), 60*60*24)

	for i := 0; i < int(axis.PointNum()); i++ {
		next := axis.Next(int64(i))
		t.Log(time.Unix(next, 0))
	}
	t.Log(end.Unix())
}
