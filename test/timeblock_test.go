package timeblock_lib

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
	timeblock_lib "time-block-tracker/v2/lib"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		DisableColors: false,
	})

	m.Run()
}

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
	time.Sleep(5*time.Second)

	if !newblock.Ongoing() {
		t.Error("set to ongoing failed")
	}

	timeblock_lib.ToggleTimeblock(blocks,newblock.Id)
	time.Sleep(5*time.Second)

	if newblock.Ongoing() {
		t.Error("set to not ongoing failed")
	}

	timeblock_lib.ToggleTimeblock(blocks,newblock.Id)
	time.Sleep(5*time.Second)

	if len(newblock.Timerows)!=2 {
		t.Error("not correct amount of time rows")
	}

	timeblock_lib.AddTimeblock(blocks)

	spew.Dump(blocks)
}

// test json output of a timeblock
// create a timeblock with some timerows, and print the json
func Test_jsonconvert(t *testing.T) {
	var blocks timeblock_lib.TimeblocksDict=timeblock_lib.TimeblocksDict{}
	var newblock *timeblock_lib.Timeblock=timeblock_lib.AddTimeblock(blocks)

	// start
	timeblock_lib.ToggleTimeblock(blocks,newblock.Id)
	time.Sleep(4*time.Second)

	// stop
	timeblock_lib.ToggleTimeblock(blocks,newblock.Id)
	time.Sleep(3*time.Second)

	// start again
	timeblock_lib.ToggleTimeblock(blocks,newblock.Id)

	var jsonblock timeblock_lib.Timeblock_json=timeblock_lib.ConvertToJsonTimeblock(
		*newblock,
	)

	var res []byte
	var e error
	res,e=json.MarshalIndent(jsonblock,"","  ")

	if e!=nil {
		t.Error(e)
	}

	fmt.Println(string(res))
}