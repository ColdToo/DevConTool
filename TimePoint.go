package DevConTool

import (
	"fmt"
	"time"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	Second     = 1
	HalfMinute = 30 * Second
	Minute     = 60 * Second
	Hour       = 60 * Minute
	Day        = 24 * Hour
	Week       = 7 * Day
)

var (
	ErrTimeAxisEof     = fmt.Errorf("reach end of the timeTool axis")
	errTimeStampGoBack = fmt.Errorf("timestamp go back")
)

func alignUnixTimeStamp(t int64, a int64) int64 {
	_, offset := time.Now().Zone()
	return (t+int64(offset))/a*a - int64(offset)
}

// timeTool axis for construct timeTool series
type TimeAxis struct {
	start int64
	end   int64
	step  int64
	index int64
	point int64
}

func NewTimeAxis(start int64, end int64, step int64) *TimeAxis {
	t := &TimeAxis{start: alignUnixTimeStamp(start, step), end: alignUnixTimeStamp(end, step), step: step}
	t.point = (t.end-t.start)/t.step + 1
	return t
}

func (t *TimeAxis) Start() int64                    { return t.start }
func (t *TimeAxis) End() int64                      { return t.end }
func (t *TimeAxis) Step() int64                     { return t.step }
func (t *TimeAxis) PointNum() int64                 { return t.point }
func (t *TimeAxis) Reset()                          { t.index = 0 }
func (t *TimeAxis) getPointTimeStamp(i int64) int64 { return t.start + i*t.step }
func (t *TimeAxis) Next(i int64) int64              { return t.start + i*t.step }

func (t *TimeAxis) IndexOf(timestamp int64) (int64, error) {
	for {
		if t.index >= t.point {
			return 0, ErrTimeAxisEof
		} else if timestamp >= t.getPointTimeStamp(t.index) && timestamp < t.getPointTimeStamp(t.index+1) {
			return t.index, nil
		} else if timestamp >= t.getPointTimeStamp(t.index+1) {
			t.index++
		} else {
			return 0, errTimeStampGoBack
		}
	}
}

func TimeUnix(str string) (int64, error) {
	t, err := time.ParseInLocation(TimeFormat, str, time.Local)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

type TimeRange struct {
	Start string
	End   string
}

func resolveTime(seconds int) (day, hour, minute, second int) {
	day = seconds / (24 * 3600)
	hour = (seconds - day*3600*24) / 3600
	minute = (seconds - day*24*3600 - hour*3600) / 60
	second = seconds - day*24*3600 - hour*3600 - minute*60

	return
}

func ResolveTime(seconds int) string {
	var day, hour, minute, second string
	d, h, m, s := resolveTime(seconds)
	if d > 0 {
		day = fmt.Sprintf("%d天", d)
	}
	if h > 0 {
		hour = fmt.Sprintf("%d小时", h)
	}
	if m > 0 {
		minute = fmt.Sprintf("%d分钟", m)
	}
	if s > 0 {
		second = fmt.Sprintf("%d秒", s)
	}

	return day + hour + minute + second
}
