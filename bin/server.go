package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
	timeblock_api "time-block-tracker/v2/lib/api"
	timeblock_lib "time-block-tracker/v2/lib/timeblock"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

// auto save timeblock file every N seconds
const AUTOSAVE_INTERVAL int=30

func main() {
    // --- initial setup ---
    var HERE string
    HERE,_=os.Executable()
    HERE=filepath.Dir(HERE)

    var timeblockFilePath string=filepath.Join(HERE,"../data/timeblock.json")

	logrus.SetFormatter(&logrus.TextFormatter {
		ForceColors: true,
		DisableColors: false,
	})

	var app *fiber.App=fiber.New(fiber.Config {
        CaseSensitive: true,
        EnablePrintRoutes: false,
    })

    app.Use(cors.New())

    // retrieving saved timeblocks
    var blocks timeblock_lib.TimeblocksDict
    os.MkdirAll(filepath.Dir(timeblockFilePath),os.ModePerm)
    blocks,_=timeblock_lib.LoadTimeblockFile(timeblockFilePath)

    var blockMtx sync.Mutex



    // ---- workers ---
    go saveBlocksWorker(blocks,timeblockFilePath,&blockMtx)


    // ---- routes ----
    // gets TimeblocksDict in json form
    app.Get("/get-timeblocks",func(c *fiber.Ctx) error {
        return c.JSON(timeblock_lib.ConvertToJsonTimeblockDict(blocks))
    })

    // takes title change request obj. returns string if successful
    app.Post("/set-title",func(c *fiber.Ctx) error {
        var setTitleReq timeblock_api.TitleChangeReq
        var e error=c.BodyParser(&setTitleReq)

        if e!=nil {
            return e
        }

        blockMtx.Lock()
        e=timeblock_lib.ChangeTimeblockTitle(
            blocks,
            setTitleReq.BlockId,
            setTitleReq.NewTitle,
        )
        blockMtx.Unlock()

        if e!=nil {
            return e
        }

        return c.SendString("successfully changed title")
    })

    // takes nothing, returns string if successful
    app.Post("/new-timeblock",func(c *fiber.Ctx) error {
        blockMtx.Lock()
        var newtimeblock *timeblock_lib.Timeblock=timeblock_lib.AddTimeblock(
            blocks,
        )
        blockMtx.Unlock()

        var e error=timeblock_lib.SaveTimeblockFile(blocks,timeblockFilePath)

        if e!=nil {
            fmt.Println("error while trying to save timeblock file")
            fmt.Println(e)
        }

        return c.SendString(fmt.Sprintf("created timeblock %v",newtimeblock.Id))
    })

    // give id of time block in url. returns string on success.
    // also the timeblock json.
    app.Post("/toggle-timeblock/:id",func(c *fiber.Ctx) error {
        var id string=c.Params("id")

        blockMtx.Lock()
        var e error=timeblock_lib.ToggleTimeblock(
            blocks,
            id,
        )
        blockMtx.Unlock()

        if e!=nil {
            return e
        }

        e=timeblock_lib.SaveTimeblockFile(blocks,timeblockFilePath)

        if e!=nil {
            fmt.Println("error while trying to save timeblock file")
            fmt.Println(e)
        }

        return c.SendString("successfully toggled timeblock")
    })

    app.Listen(":4201")
}

// worker that saves the current blocks once in a while to the timeblock file
func saveBlocksWorker(
    blocks timeblock_lib.TimeblocksDict,
    timeblockFilePath string,
    blockMtx *sync.Mutex,
) {
    var ticker *time.Ticker=time.NewTicker(time.Duration(AUTOSAVE_INTERVAL)*time.Second)

    for range ticker.C {
        blockMtx.Lock()
        var e error=timeblock_lib.SaveTimeblockFile(blocks,timeblockFilePath)

        if e!=nil {
            fmt.Println("error while trying to save timeblock file")
            fmt.Println(e)
        }
        blockMtx.Unlock()
    }
}