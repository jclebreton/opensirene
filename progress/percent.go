package progress

import (
	"fmt"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type states struct {
	sync.RWMutex
	m map[string]map[string]float64
}

type Percent struct {
	states states
	stop   chan interface{}
}

func NewPercent(steps ...string) *Percent {
	perc := &Percent{}

	perc.states.m = make(map[string]map[string]float64)
	for _, step := range steps {
		perc.states.m[step] = make(map[string]float64)
	}
	perc.stop = make(chan interface{}, 1)
	return perc
}

func (perc *Percent) CatchProgress(pChan chan *Progress) {
	for p := range pChan {
		func() {
			perc.states.Lock()
			defer perc.states.Unlock()
			if _, ok := perc.states.m[p.Step]; !ok {
				logrus.Error("unknown step progression")
				return
			}
			perc.states.m[p.Step][p.Name] = p.Percent()
		}()
	}
}

func (perc *Percent) ShowLogs() {
	tracking := make(map[string]float64)
	perc.states.RLock()
	defer perc.states.RUnlock()
	for k, v := range perc.states.m {
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
