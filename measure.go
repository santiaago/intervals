package main

import (
	"bufio"
	"errors"
	"log"
	"strings"
	"time"
)

// measure defines the information needed to analyse the data passed
// to the reader.
type measure struct {
	start             string          // defines the start of the task to measure
	end               string          // defines the end of the task to measure
	timeForm          string          // time format to look for
	startPositionTime int             // start position of the time
	endPositionTime   int             // end position of the time
	reader            *strings.Reader // the data to measure
}

// Intervals scans the reader data and computes the avg duration of each
// interval and the number of intervals.
func (m measure) Intervals() (avg time.Duration, count int) {

	s := bufio.NewScanner(m.reader)
	s.Split(bufio.ScanLines)

	var i interval
	var err error
	var ints intervals

	for s.Scan() {

		if i, err = m.line(s.Text(), i); err != nil {
			i = interval{}
		} else if i.isReady() {
			i.duration = i.end.Sub(i.start)
			ints = append(ints, i)
			i = interval{}
		}
	}

	avg = ints.avg()
	count = len(ints)
	return
}

// line scans a line with respect to an interval. It returns an updated
// interval object and an error the line scanned in inconsistent with the
// interval passed by paramenter.
func (m measure) line(line string, i interval) (interval, error) {

	if m.isStart(line) {
		if i.hasStart {
			return i, errors.New("inconsistent interval")
		}
		if err := i.updateStart(m.timeForm, line[m.startPositionTime:m.endPositionTime]); err != nil {
			log.Println(err)
			return i, errors.New("unable to update start time in interval")
		}
		i.endLine = line
	} else if m.isEnd(line) && i.hasStart {
		if i.hasEnd {
			return i, errors.New("inconsistent interval")
		}

		if err := i.updateEnd(m.timeForm, line[m.startPositionTime:m.endPositionTime]); err != nil {
			log.Println(err)
			return i, errors.New("unable to update end time in interval")
		}
		i.endLine = line
	}
	return i, nil
}

// isStart tells you if a line is the begining of a task.
func (m measure) isStart(line string) bool {
	return strings.Contains(line, m.start)
}

// isEnd tells you if a line is the end of a task.
func (m measure) isEnd(line string) bool {
	return strings.Contains(line, m.end)
}
