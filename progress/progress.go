package progress

import "github.com/sirupsen/logrus"

var DefaultChan chan *Progress

// DefaultChanBufferSize defines the number of progress events which can be buffered.
// If the buffer is full the next progress event will be dropped
const DefaultChanBufferSize = 50000

func init() {
	DefaultChan = make(chan *Progress, DefaultChanBufferSize)
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
		logrus.Error("buffer is full. Progress event has been dropped")
	}
}

func (p *Progress) Percent() float64 {
	return float64(p.Current) / float64(p.Total) * 100
}
