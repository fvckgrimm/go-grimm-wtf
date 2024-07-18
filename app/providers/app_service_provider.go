package providers

import (
	"github.com/goravel/framework/contracts/foundation"
    "html/template"
    "github.com/goravel/framework/facades"
    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/html"
    "github.com/gomarkdown/markdown/parser"
)

type AppServiceProvider struct {
}

func (receiver *AppServiceProvider) Register(app foundation.Application) {

}

func (receiver *AppServiceProvider) Boot(app foundation.Application) {
    facades.View().Share("markdown", func(input string) template.HTML {
        extensions := parser.CommonExtensions | parser.AutoHeadingIDs
        p := parser.NewWithExtensions(extensions)
        doc := p.Parse([]byte(input))

        htmlFlags := html.CommonFlags | html.HrefTargetBlank
        opts := html.RendererOptions{Flags: htmlFlags}
        renderer := html.NewRenderer(opts)

        return template.HTML(markdown.Render(doc, renderer))
    })
}
