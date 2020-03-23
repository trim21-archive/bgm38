package web

import (
	"fmt"
	"io"
	"mime"
	"path"

	"github.com/gofiber/fiber"
	"github.com/gofiber/recover"
	"github.com/gofiber/requestid"
	"github.com/markbates/pkger"
	"github.com/sirupsen/logrus"

	"bgm38/pkg/web/bgmtv"
)

func Start() error {
	return CreateApp().Listen(3000)
}

func CreateApp() *fiber.App {
	app := fiber.New()
	app.Use(requestid.New())

	app.Use(recover.New(recover.Config{
		Handler: func(c *fiber.Ctx, err error) {
			fmt.Println(err)
			c.SendString(err.Error())
			c.SendStatus(500)
		},
	}))

	setupSwagger(app)
	app.Get("/asserts/web/*", func(c *fiber.Ctx) {
		filepath := c.Params("*")
		f, err := pkger.Open(path.Join("/asserts/web/", filepath))
		if err != nil {
			c.SendStatus(404)
			return
		}
		defer f.Close()
		mintType := mime.TypeByExtension(path.Ext(filepath))
		c.Set("content-type", mintType)
		_, err = io.Copy(c.Fasthttp.Response.BodyWriter(), f)
		if err != nil {
			logrus.Errorln(err)
		}
	})
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})
	bgmtv.Group(app)
	rootRouter(app)
	app.Use(func(c *fiber.Ctx) {
		c.Status(404).SendString(`{}`)
		// => 404 "Not Found"
	})

	return app
}
