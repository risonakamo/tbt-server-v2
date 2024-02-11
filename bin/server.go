package main

import (
	"fmt"
	timeblock_api "time-block-tracker/v2/lib/api"
	timeblock_lib "time-block-tracker/v2/lib/timeblock"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	var app *fiber.App=fiber.New(fiber.Config {
        CaseSensitive: true,
        EnablePrintRoutes: false,
    })

    app.Use(cors.New())

    var blocks timeblock_lib.TimeblocksDict=timeblock_lib.TimeblocksDict{}

    app.Get("/get-timeblocks",func(c *fiber.Ctx) error {
        return c.JSON(blocks)
    })

    app.Post("/set-title",func(c *fiber.Ctx) error {
        var setTitleReq timeblock_api.TitleChangeReq
        var e error=c.BodyParser(&setTitleReq)

        if e!=nil {
            return e
        }

        e=timeblock_lib.ChangeTimeblockTitle(
            blocks,
            setTitleReq.BlockId,
            setTitleReq.NewTitle,
        )

        if e!=nil {
            return e
        }

        return c.SendString("successfully changed title")
    })

    app.Post("/new-timeblock",func(c *fiber.Ctx) error {
        var newtimeblock *timeblock_lib.Timeblock=timeblock_lib.AddTimeblock(
            blocks,
        )

        return c.SendString(fmt.Sprintf("created timeblock %v",newtimeblock.Id))
    })

    app.Post("/toggle-timeblock/:id",func(c *fiber.Ctx) error {
        var id string=c.Params("id")

        var e error=timeblock_lib.ToggleTimeblock(
            blocks,
            id,
        )

        if e!=nil {
            return e
        }

        return c.SendString("successfully toggled timeblock")
    })

    app.Listen(":4201")
}