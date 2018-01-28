package liveclock

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/bketelsen/cog"
	"github.com/gobuffalo/packr"
)

var _ cog.Cog = (*LiveClock)(nil)
var cogType reflect.Type
var box packr.Box

type LiveClock struct {
	cog.UXCog
	ticker *time.Ticker
}

func NewLiveClock() *LiveClock {
	liveClock := &LiveClock{}
	liveClock.SetCogType(cogType)
	liveClock.SetCleanupFunc(liveClock.Cleanup)
	liveClock.UXCog.Box = box
	return liveClock
}

func (lc *LiveClock) Cleanup() {
	lc.ticker.Stop()
}

func (lc *LiveClock) Start() error {

	const layout = time.RFC1123
	var location *time.Location

	if lc.Props["timezonename"] != nil && lc.Props["timezoneoffset"] != nil {
		tzo, err := strconv.Atoi(lc.Props["timezoneoffset"].(string))
		if err != nil {
			return errors.New("The timezonename and timezoneoffset props need to be set!")
		}
		location = time.FixedZone(lc.Props["timezonename"].(string), tzo)
	} else {
		return errors.New("The timezonename and timezoneoffset props need to be set!")
	}

	lc.ticker = time.NewTicker(time.Millisecond * 1000)

	go func() {
		for t := range lc.ticker.C {
			lc.SetProp("currentTime", t.In(location).Format(layout))
		}
	}()

	err := lc.Render()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	cogType = reflect.TypeOf(LiveClock{})
	box = packr.NewBox("../assets")
}
