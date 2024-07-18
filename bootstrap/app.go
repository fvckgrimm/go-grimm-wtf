package bootstrap

import (
	"github.com/goravel/framework/foundation"

	"goravel/config"

    //"github.com/goravel/framework/facades"
    //"github.com/gomarkdown/markdown"
    //"github.com/gomarkdown/markdown/html"
    //"github.com/gomarkdown/markdown/parser"
)

func Boot() {
	app := foundation.NewApplication()

	//Bootstrap the application
	app.Boot()

	//Bootstrap the config.
	config.Boot()

}
