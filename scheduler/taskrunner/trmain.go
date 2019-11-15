/*
@desc : Created by San on 2019/10/28 20:53
*/
package taskrunner

import "time"

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval*time.Second),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}