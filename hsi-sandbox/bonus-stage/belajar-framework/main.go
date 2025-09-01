package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi Framework Fiber
	app := fiber.New()

	// Middleware Sederhana
	// Fungsi : Untuk Log Request
	app.Use(func(ctx *fiber.Ctx) error {
		log.Print("Middleware Berjalan Sebelum Ke Routing")
		// Melukukan validasi apakah user yang merequest itu mempunyai hak/tidak
		return ctx.Next() // Untuk Meneruskan Request
	})

	// Routing
	// URL http://localhost:3000/api
	// Return Ahlan Wa Sahlan
	app.Get("/api/", func(ctx *fiber.Ctx) error {
		log.Print("Routing ke URL /api")
		return ctx.SendString("Ahlan Wa Sahlan")
	})

	//Parameter URL sampel 1
	app.Get("/api/:p1", func(ctx *fiber.Ctx) error {
		// Untuk mendapatkan Value dari :parameter
		getParam := ctx.Params("p1")
		return ctx.SendString("Parameter Yang Dikirimkan Adalah : " + getParam)
	})

	// Parameter URL sampel 2
	app.Get("/api/:categoryId/:productId", func(ctx *fiber.Ctx) error {
		parameter1 := ctx.Params("categoryId")
		parameter2 := ctx.Params("productId")
		return ctx.SendString("Parameter Yang Dikirimkan Adalah CategoryID: " + parameter1 + " dan ProductID: " + parameter2)
	})

	// Parameter URL sampel 3 : Opsional
	app.Get("/api/opsional/:opsionalData?", func(ctx *fiber.Ctx) error {
		opsionalData := ctx.Params("opsionalData", "Default Data")
		return ctx.SendString("Parameter Yang Dikirimkan Adalah : " + opsionalData)
	})

	type BodySample struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}

	// Handler Post
	app.Post("/api/info", func(ctx *fiber.Ctx) error {
		// inisialisasi object baru
		bodySample := new(BodySample)
		// logika, pengecekan dahulu apakah body nya sudah terformat dengan benar
		if err := ctx.BodyParser(bodySample); err != nil {
			// Jika Error kembalikan pesan kesalahan/eror
			return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		// Jika berhasil di parsing, maka kembalikan pesan dibawah ini
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "body sudah diterima",
			"name":    bodySample.Name,
			"address": bodySample.Address,
		})
	})

	// Reserved PORT
	app.Listen(":3000")
}
