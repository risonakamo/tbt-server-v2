package timeblock_lib

import (
	"testing"
	timeblock_lib "time-block-tracker/v2/lib"

	"github.com/davecgh/go-spew/spew"
)

// general test of timeblock lib usage
func Test1(t *testing.T) {
	var blocks timeblock_lib.TimeblocksDict=timeblock_lib.TimeblocksDict{}

	var newblock *timeblock_lib.Timeblock=timeblock_lib.AddTimeblock(blocks)

	timeblock_lib.ChangeTimeblockTitle(
		blocks,
		newblock.Id,
		"hello",
	)

	if blocks[newblock.Id].Title!="hello" {
		t.Error("name change failed")
	}

	timeblock_lib.ToggleTimeblock(blocks,newblock.Id)

	if !newblock.Ongoing() {
		t.Error("set to ongoing failed")
	}

	timeblock_lib.ToggleTimeblock(blocks,newblock.Id)

	if newblock.Ongoing() {
		t.Error("set to not ongoing failed")
	}

	timeblock_lib.ToggleTimeblock(blocks,newblock.Id)

	if len(newblock.Timerows)!=2 {
		t.Error("not correct amount of time rows")
	}

	timeblock_lib.AddTimeblock(blocks)

	spew.Dump(blocks)
}