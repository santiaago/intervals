package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var (
	file              = flag.String("f", "file.log", "File name to the log file to analyse.")
	start             = flag.String("s", "start task", "Determines the begining of the task to measure.")
	end               = flag.String("e", "end task", "Determines the end of the task to measure.")
	timeForm          = flag.String("t", "2006-01-02 15:04:05.000", "Time format present in the log file.")
	startTimePosition = flag.Int("ts", 0, "Start position of the time stamp in the log file.")
	endTimePosition   = flag.Int("te", 23, "End position of the time stamp in the log file.")
)

var usage = `Usage: intervals [options...]
Options:
  -f   File name to the log file to analyse.
  -s   Determines the beginning of the task to measure.
  -e   Determines the end of the task to measure.
  -t   Time format present in the log file.
  -ts  Start position of the time stamp in the log file.
  -te  End position of the time stamp in the log file.
`

func main() {

	flag.Parse()

	if data, err := ioutil.ReadFile(*file); err != nil {
		log.Fatalln(err)
	} else {
		m := measure{
			*start,
			*end,
			*timeForm,
			*startTimePosition,
			*endTimePosition,
			strings.NewReader(fmt.Sprintf("%s", data)),
		}
		avg, c := m.Intervals()

		fmt.Println("number of intervals: ", c)
		fmt.Println("avg duration: ", avg)
	}
}
