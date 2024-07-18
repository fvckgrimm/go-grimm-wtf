package routes

import (
	"goravel/app/http/controllers"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
        isMobile, ok := ctx.Value("isMobile").(bool)
        if ok && isMobile {
            // Render mobile.tmpl for mobile devices
            return ctx.Response().View().Make("mobile.tmpl", map[string]interface{}{
                "version": support.Version,
            })
        }

        // Otherwise, render home.tmpl for non-mobile devices
        return ctx.Response().View().Make("home.tmpl", map[string]interface{}{
            "version": support.Version,
        })
    })

    facades.Route().Get("/blog", controllers.NewBlogController().Index)
    facades.Route().Get("/blog/{slug}", controllers.NewBlogController().Show) 

    facades.Route().Static("/css", "./resources/css")
    facades.Route().Static("/javascript", "./resources/javascript")
    facades.Route().Static("/assets", "./resources/assets")

    rssController := controllers.NewRSSController()
    facades.Route().Get("/rss", rssController.Index)

}
