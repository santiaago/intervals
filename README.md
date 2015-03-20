# Intervals

Intervals is a tiny program that measures the average duration of a repetitive task from the times present in a log file. The file should be a log file that has a time stamp on each line been logged.

# Installation:

    go get github.com/santiaago/intervals

# Usage

    Usage: intervals [options...]
    Options:
      -f   File name to the log file to analyse.
      -s   Determines the beginning of the task to measure.
      -e   Determines the end of the task to measure.
      -t   Time format present in the log file.
      -ts  Start position of the time stamp in the log file.
      -te  End position of the time stamp in the log file.


This is what happens when you run Intervals:

    .\intervals -f ".\file.log" -s "Starting A" -e "Finished A"
    number of intervals:  855
    avg duration:  21.748794152s

This is what your log can look like:

    2015-03-06 16:13:00.000 +00:00 [Information] Starting A.
    2015-03-06 16:13:20.000 +00:00 [Information] Finished A.
    2015-03-06 16:14:00.000 +00:00 [Information] Starting A.
    2015-03-06 16:14:40.000 +00:00 [Information] Finished A.
