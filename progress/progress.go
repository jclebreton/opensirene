package progress

import "github.com/sirupsen/logrus"

var DefaultProgressChan chan *Progress

func init() {
	DefaultProgressChan = make(chan *Progress)
}

type Progress struct {
	Name         string
	Step         string
	Total        uint64
	Current      uint64
	ProgressChan chan *Progress
}

func (p *Progress) Add(n int64) {
	p.Current += uint64(n)
	select {
	case p.ProgressChan <- p:
	default:
		logrus.Error("Unable to send progress")
	}
}

func (p *Progress) Percent() float64 {
	return float64(p.Current) / float64(p.Total) * 100
}
