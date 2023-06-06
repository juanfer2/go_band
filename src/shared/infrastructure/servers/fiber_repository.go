package servers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	group_mate_routes "github.com/juanfer2/whorship_band/src/group_mates/infrastructure/routes"
	"github.com/juanfer2/whorship_band/src/shared/infrastructure/routes"
	"github.com/juanfer2/whorship_band/src/shared/utilities"
	wompi_routes "github.com/juanfer2/whorship_band/src/wompi/infrastructure/routes"
)

func StartServerApp() {
	PORT := getPort()
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format:   "${pid} ${status} - ${method} ${path}\n",
		TimeZone: "America/Bogotá",
		Done: func(c *fiber.Ctx, logString []byte) {
			fmt.Sprintf("%s - %d", c.Request().RequestURI(), c.Response().StatusCode())
		},
	}))
	app.Use(cors.New())
	app.Static("/images", "./tmp/images")

	routes.SetupRoutes(app)
	wompi_routes.SetupRoutes(app)
	group_mate_routes.SetupRoutes(app)
	//products_routes.SetupRoutes(app)
	app.Listen(PORT)
}

func getPort() string {
	if utilities.GetEnv("PORT") != "" {
		return ":" + utilities.GetEnv("PORT")
	} else {
		return ":4000"
	}
}