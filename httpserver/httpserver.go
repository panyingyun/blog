package httpserver

import (
	"strings"

	"blog/assets/public"
	"blog/assets/templates"

	log "github.com/Sirupsen/logrus"
	"github.com/go-macaron/bindata"
	"gopkg.in/macaron.v1"
)

func StartHttpServer(addr string) {
	macaron.Env = macaron.PROD
	m := macaron.New()
	m.Use(macaron.Logger())
	m.Use(macaron.Recovery())
	m.Use(macaron.Static("public",
		macaron.StaticOptions{
			FileSystem: bindata.Static(bindata.Options{
				Asset:      public.Asset,
				AssetDir:   public.AssetDir,
				AssetNames: public.AssetNames,
				AssetInfo:  public.AssetInfo,
				Prefix:     "",
			}),
		},
	))
	m.Use(macaron.Renderer(macaron.RenderOptions{
		TemplateFileSystem: bindata.Templates(bindata.Options{
			Asset:      templates.Asset,
			AssetDir:   templates.AssetDir,
			AssetNames: templates.AssetNames,
			AssetInfo:  templates.AssetInfo,
			Prefix:     "templates",
		}),
	}))

	//Test Web
	//m.Get("/:name", appserver)
	m.RunAddr(addr)
}

//michaelapp blog
func appserver(ctx *macaron.Context) {
	name := ctx.Params("name")
	if len(name) == 0 {
		name = "index"
	} else {
		name = strings.Replace(name, ".html", "", 1)
	}
	log.Info("michaelapp blog first web name = ", name)
	ctx.HTML(200, name)
}

//Just For Chrome or FireFox Browser Test
func favicon(ctx *macaron.Context) {
	ctx.ServeFileContent("public/favicon.ico")
}
