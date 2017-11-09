package progress

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
	"github.com/sirupsen/logrus"
)

type Percent struct {
	states map[string]map[string]float64
}

func NewPercent() *Percent {
	perc := &Percent{}
	perc.states = make(map[string]map[string]float64)

	return perc
}

func (perc *Percent) CatchProgress(pChan chan *Progress) {

	for p := range pChan {
		if _, ok := perc.states[p.Step]; !ok {
			perc.states[p.Step] = make(map[string]float64)
		}
		perc.states[p.Step][p.Name] = p.Percent()
	}
}

func (perc *Percent) ShowLogs() {
	tracking := make(map[string]float64)
	var max int
	for k, v := range perc.states {
		var value float64
		for _, v2 := range v {
			value += v2
		}
		tracking[k] = value
		if len(v) > max {
			max = len(v)
		}
	}
	for k, v := range tracking {
		tracking[k] = v / float64(max)
	}

	log := logrus.NewEntry(logrus.StandardLogger())
	for k, v := range tracking {
		log = log.WithField(k, fmt.Sprintf("%0.03f%%", v))
	}
	log.Info("Progress")
}

func (perc *Percent) Start() {
	gocron.Every(1).Seconds().Do(perc.ShowLogs)
	<-gocron.Start()
}
