// typings for timeblock lib

package timeblock_lib

import (
	"fmt"
	"time"
)

// collection of timeblocks, keyed by their ID
type TimeblocksDict map[string]*Timeblock

type Timeblock struct {
    Id string

	Title string

    // timerows should be ordered by start date, backwards. the last time row should be the
    // most latest
	Timerows []Timerow

    // --- derived ---
    // total of all durations of time rows
    // TotalTime time.Duration

    // if the "top" timerow is ongoing
    // Ongoing bool
}

type Timerow struct {
    Id string

	Start time.Time
    End time.Time

    // if ongoing, end time is invalid
    Ongoing bool

    // --- derived ---
    // Duration time.Duration
}

// get duration of time row. returns -1 duration if it is still ongoing
func (timerow *Timerow) Duration() time.Duration {
    if timerow.Ongoing {
        fmt.Println("error: tried to get duration of ongoing time row")
        return time.Duration(-1)
    }

    return timerow.Start.Sub(timerow.End)
}

// return total time duration of all non-ongoing time rows
func (timeblock *Timeblock) TotalTime() time.Duration {
    var totaltime time.Duration

    for i := range timeblock.Timerows {
        totaltime=totaltime+timeblock.Timerows[i].Duration()
    }

    return totaltime
}

// return if the timeblock is ongoing
func (timeblock *Timeblock) Ongoing() bool {
    if len(timeblock.Timerows)==0 {
        return false
    }

    return timeblock.Timerows[len(timeblock.Timerows)-1].Ongoing
}