// Copyright (c) 2022 aiocat
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var KRNL *Krnl = NewKrnl()

func main() {
	KRNL.Load()
	defer KRNL.Free()

	app := fiber.New()

	app.Get("/", RootRoute)
	app.Get("/inject", InjectRoute)
	app.Get("/injected", IsInjectedRoute)
	app.Get("/scripts", AllScriptsRoute)
	app.Post("/scripts", ScriptRoute)
	app.Post("/execute", ExecuteRoute)

	app.Static("/", ".\\web\\static")
	app.Static("/scripts", ".\\scripts")

	fmt.Println("Gernel is loaded, visit: http://127.0.0.1:42069 to start using.")
	log.Fatal(app.Listen(":42069"))
}
