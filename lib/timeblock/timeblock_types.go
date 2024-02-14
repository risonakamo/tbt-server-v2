// typings for timeblock lib

package timeblock_lib

import (
	"time"
)

// collection of timeblocks, keyed by their ID
type TimeblocksDict map[string]*Timeblock

type Timeblock struct {
    Id string `json:"id"`

	Title string `json:"title"`

    // timerows should be ordered by start date, backwards. the last time row should be the
    // most latest
	Timerows []Timerow `json:"timerows"`

    // --- derived ---
    // total of all durations of time rows
    // TotalTime time.Duration

    // if the "top" timerow is ongoing
    // Ongoing bool
}

type Timerow struct {
    Id string `json:"id"`

	Start time.Time `json:"start"`
    End time.Time `json:"end"`

    // if ongoing, end time is invalid
    Ongoing bool `json:"ongoing"`

    // --- derived ---
    // Duration time.Duration
}

// get duration of time row. returns 0 duration if it is still ongoing
func (timerow *Timerow) Duration() time.Duration {
    if timerow.Ongoing {
        return time.Duration(0)
    }

    return timerow.End.Sub(timerow.Start)
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