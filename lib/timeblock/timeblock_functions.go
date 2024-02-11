// functions to work with timeblocks

package timeblock_lib

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// add a new timeblock to a timeblocks dict.
// MUTATES the dict.
// returns the timeblock that was just made
func AddTimeblock(blocks TimeblocksDict) *Timeblock {
    var newblock Timeblock=newTimeblock()

    blocks[newblock.Id]=&newblock

    return &newblock
}

// starts or stops a timeblock inside of a dict. returns error if it
// doesn't exist in the dict. MUTATES the dict.
func ToggleTimeblock(blocks TimeblocksDict,blockId string) error {
    var block *Timeblock
    var found bool

    block,found=blocks[blockId]

    if !found {
        logrus.Error("failed to find block with id: ",blockId)
        return errors.New("failed to find block id")
    }

    toggleTimeblock(block)

    logrus.Infof("toggled timeblock [%v] to %v\n",blockId,block.Ongoing())

    return nil
}

// changes title of a block in timeblocks dict. returns error if failed.
// MUTATES dict.
func ChangeTimeblockTitle(
    blocks TimeblocksDict,
    blockId string,
    newTitle string,
) error {
    var block *Timeblock
    var found bool

    block,found=blocks[blockId]

    if !found {
        logrus.Error("failed to find block with id: ",blockId)
        return errors.New("failed to find block id")
    }

    block.Title=newTitle

    return nil
}

// convert a timeblock to the json version
func ConvertToJsonTimeblock(block Timeblock) Timeblock_json {
    return Timeblock_json {
        Timeblock: block,

        Timerows: convertTimerowsToJson(block.Timerows),

        TotalTime: block.TotalTime().Minutes(),
        Ongoing: block.Ongoing(),
    }
}

// convert all blocks in timeblock dict to json version
func ConvertToJsonTimeblockDict(blocks TimeblocksDict) TimeblocksDict_json {
    var newdict TimeblocksDict_json=TimeblocksDict_json{}

    var id string
    var block *Timeblock
    for id,block = range blocks {
        var conblock Timeblock_json=ConvertToJsonTimeblock(*block)
        newdict[id]=&conblock
    }

    return newdict
}





// ----- private -----
// create a new timeblock
func newTimeblock() Timeblock {
    return Timeblock {
        Id: genUUid(),
        Timerows: []Timerow{},
    }
}

// create new timerow, setting start time as now, and it as ongoing
func newTimerow() Timerow {
    return Timerow {
        Id: genUUid(),
        Start: time.Now(),
        Ongoing: true,
    }
}

// return a short uuid
func genUUid() string {
    return uuid.New().String()[0:6]
}

// perform "toggle" action on time block.  this results in adding a timerow to the
// time block if it is not "ongoing", or "closing" the last time row in the timeblock,
// if it is ongoing.
func toggleTimeblock(block *Timeblock) {
    if !block.Ongoing() {
        // if not ongoing, add a new timerow to the end of the timerow list
        block.Timerows=append(block.Timerows,newTimerow())
    } else {
        // otherwise, close the last row in the timeblock
        closeTimerow(&block.Timerows[len(block.Timerows)-1])
    }
}

// ends ongoing status of time row by setting the end date and setting
// ongoing to false. MUTATES time row.
// if the timerow is not ongoing, does nothing
func closeTimerow(timerow *Timerow) {
    if !timerow.Ongoing {
        return
    }

    timerow.Ongoing=false

    timerow.End=time.Now()
}

// convert array of time rows into time rows json
func convertTimerowsToJson(rows []Timerow) []Timerow_json {
    var newrows []Timerow_json=[]Timerow_json{}

    for i := range rows {
        newrows=append(newrows,Timerow_json {
            Timerow: rows[i],
            Duration: rows[i].Duration().Minutes(),
        })
    }

    return newrows
}