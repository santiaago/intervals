package main

import "time"

// interval defined
type interval struct {
	start     time.Time     // the start of the interval
	end       time.Time     // the end of the interval
	duration  time.Duration // the duration of the interval
	startLine string        // the line scanned that defines the start of the interval
	endLine   string        // the line scanned that defined the end of the interval
	hasStart  bool          // defined if the current interval has started
	hasEnd    bool          // defines if the current interval has ended
}

type intervals []interval

func (i interval) isReady() bool {
	return i.hasStart && i.hasEnd
}

func (i *interval) updateStart(timeForm, moment string) error {
	var err error
	var t time.Time
	if t, err = time.Parse(timeForm, moment); err != nil {
		return err
	}
	i.start = t
	i.hasStart = true
	return nil
}

func (i *interval) updateEnd(timeForm, moment string) error {
	var err error
	var t time.Time
	if t, err = time.Parse(timeForm, moment); err != nil {
		return err
	}
	i.end = t
	i.hasEnd = true
	return nil
}

// avg computes the average duration from an array of durations.
func (ints *intervals) avg() (d time.Duration) {
	if len(*ints) == 0 {
		return 0
	}
	for _, v := range *ints {
		d += v.duration
	}
	return d / time.Duration(len(*ints))
}
