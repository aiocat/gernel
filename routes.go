// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RootRoute(ctx *fiber.Ctx) error {
	return ctx.SendFile(".\\web\\index.html")
}

func InjectRoute(ctx *fiber.Ctx) error {
	KRNL.Inject()
	return ctx.SendStatus(204)
}

func ExecuteRoute(ctx *fiber.Ctx) error {
	KRNL.Execute(string(ctx.Body()))
	return ctx.SendStatus(204)
}

func IsInjectedRoute(ctx *fiber.Ctx) error {
	if KRNL.IsInjected() {
		return ctx.Status(200).SendString("1")
	} else {
		return ctx.Status(200).SendString("0")
	}
}

func AllScriptsRoute(ctx *fiber.Ctx) error {
	list, err := os.ReadDir(".\\scripts")

	if err != nil {
		return ctx.Status(200).SendString("")
	}

	listedStr := ""
	for _, file := range list {
		listedStr += (file.Name() + "\n")
	}

	listedStr = strings.TrimSuffix(listedStr, "\n")
	return ctx.Status(200).SendString(listedStr)
}

func ScriptRoute(ctx *fiber.Ctx) error {
	name := string(ctx.Body())
	if name == "" {
		return ctx.Status(200).SendString("")
	}

	bytes, err := os.ReadFile(".\\scripts\\" + name)
	if err != nil {
		return ctx.Status(200).SendString("")
	}

	return ctx.Status(200).SendString(string(bytes))
}
