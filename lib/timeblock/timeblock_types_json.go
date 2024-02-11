package timeblock_lib

// collection of timeblock_jsons, keyed by id
type TimeblocksDict_json map[string]*Timeblock_json

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