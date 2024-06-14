package event

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

type EventProcessor interface {
	Started(context.Context)
	Stopped(context.Context)
	Process(context.Context)
}

type EventLooper struct {
	Interval  uint64
	Processor EventProcessor

	ctx       context.Context
	cancelCtx context.CancelFunc
	stopSig   chan struct{}
}

// Inits the context and runs the `Started` function if any.
func (l *EventLooper) starting() {
	ctx, cancelFn := context.WithCancel(context.Background())
	l.ctx = ctx
	l.cancelCtx = cancelFn
	l.stopSig = make(chan struct{})
}

// Clears the context and runs the `Stopped` function if any.
func (l *EventLooper) stopping() {
	l.ctx = nil
	l.cancelCtx = nil
	l.stopSig <- struct{}{}
}

// Primary looper of the event. Runs the `Process` function in interval.
func (l *EventLooper) loop() {
	timer := time.NewTimer(0)
	ticker := time.NewTicker(time.Second * time.Duration(l.Interval))
	for {
		select {
		case <-timer.C:
			log.Debug("[EventLooper][loop] call started")
			l.Processor.Started(l.ctx)
			timer.Stop()
		case <-l.ctx.Done():
			log.Debug("[EventLooper][loop] call done")
			ticker.Stop()
			l.stopping()
			l.Processor.Stopped(l.ctx)
			return
		case <-ticker.C:
			log.Debug("[EventLooper][loop] call process")
			l.Processor.Process(l.ctx)
		}
	}
}

// Returns whether the event looper is running.
func (l *EventLooper) Running() bool {
	return l.ctx != nil
}

// Starts the event looper by initializing the new context and call the `loop` function.
// Returns error if the event is already running.
func (l *EventLooper) Start() error {
	if l.Running() {
		return fmt.Errorf("Event is already running")
	}
	l.starting()
	go l.loop()
	return nil
}

// Waits until the event looper stopped.
func (l *EventLooper) Wait() {
	<-l.stopSig
}

// Stops the current looper and waits until the event actual stopped.
// Returns error if the event is not running.
func (l *EventLooper) Stop() error {
	if !l.Running() {
		return fmt.Errorf("Event is not running")
	}
	l.cancelCtx()
	l.Wait()
	l.stopSig = nil
	return nil
}

// Returns a new event looper.
func NewEventLooper(Processor EventProcessor, interval uint64) EventLooper {
	return EventLooper{Processor: Processor, Interval: interval}
}
