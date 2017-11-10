package progress

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type Percent struct {
	states map[string]map[string]float64
	stop   chan interface{}
}

func NewPercent(steps ...string) *Percent {
	perc := &Percent{}
	perc.states = make(map[string]map[string]float64)
	for _, step := range steps {
		perc.states[step] = make(map[string]float64)
	}
	perc.stop = make(chan interface{}, 1)
	return perc
}

func (perc *Percent) CatchProgress(pChan chan *Progress) {
	for p := range pChan {
		if _, ok := perc.states[p.Step]; !ok {
			logrus.Error("unknown step progression")
			continue
		}
		perc.states[p.Step][p.Name] = p.Percent()
	}
}

func (perc *Percent) ShowLogs() {
	tracking := make(map[string]float64)

	for k, v := range perc.states {
		var value float64
		for _, v2 := range v {
			value += v2
		}
		if n := len(v); n == 0 {
			tracking[k] = value
		} else {
			tracking[k] = value / float64(n)
		}
	}

	log := logrus.NewEntry(logrus.StandardLogger())
	for k, v := range tracking {
		log = log.WithField(k, fmt.Sprintf("%0.02f%%", v))
	}
	log.Info("Progress")
}

func (perc *Percent) Start() {
	tick := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-perc.stop:
			tick.Stop()
			close(perc.stop)
			return
		case <-tick.C:
			perc.ShowLogs()
		}
	}
}

func (perc *Percent) Stop() {
	perc.stop <- true
}
