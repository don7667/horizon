package horizon

import (
	"github.com/openbankit/horizon/simplepath"
)

func initPathFinding(app *App) {
	app.paths = &simplepath.Finder{app.CoreQ()}
}

func init() {
	appInit.Add("path-finder", initPathFinding, "app-context", "log", "core-db")
}
