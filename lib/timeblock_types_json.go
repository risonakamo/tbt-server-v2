package timeblock_lib

// json form of timeblock, with derived fields converted into real fields
type Timeblock_json struct {
    Timeblock

    Timerows []Timerow_json `json:"timerows"`

    // total time in minutes
    TotalTime float64 `json:"totaltime"`
    Ongoing bool `json:"ongoing"`
}

// json version of time row with all derived fields as real fields
type Timerow_json struct {
    Timerow

    // duration in minutes
    Duration float64 `json:"duration"`
}