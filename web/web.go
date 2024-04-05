package web

import (
	"embed"
	"github.com/labstack/echo/v4"
)

/*
	import(
	"embed"

"github.com/gin-gonic/gin"
)
*/
var (
	dist      embed.FS
	indexHTML embed.FS

	distDirFS     = echo.MustSubFS(dist, "dist")
	distIndexHTML = echo.MustSubFS(indexHTML, "dist")
)

func RegisterHandler(e *echo.Echo) {
	e.FileFS("/", "index.html", distIndexHTML)
}
