package cmd

import (
	// "fmt"
	"path/filepath"
	"strings"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/flamego/brotli"
	"github.com/flamego/flamego"
	"github.com/flamego/gzip"
	"github.com/flamego/template"
	"github.com/urfave/cli"

	"github.com/midoks/vez-en/internal/assets/public"
	"github.com/midoks/vez-en/internal/assets/templates"
	"github.com/midoks/vez-en/internal/conf"
	"github.com/midoks/vez-en/internal/context"
	"github.com/midoks/vez-en/internal/router"
	"github.com/midoks/vez-en/internal/tmpl"
	"github.com/midoks/vez-en/internal/tools"
)

var Service = cli.Command{
	Name:        "web",
	Usage:       "This command starts all services",
	Description: `Start Web services`,
	Action:      runWebService,
	Flags: []cli.Flag{
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

func newFlamego() *flamego.Flame {

	f := flamego.New()

	if !conf.Web.DisableRouterLog {
		f.Use(flamego.Logger())
	}

	f.Use(flamego.Recovery())

	if conf.Web.EnableGzip {
		f.Use(gzip.Gzip(gzip.Options{
			CompressionLevel: 9, // 最优压缩
		}))
	}

	f.Use(brotli.Brotli())

	// public
	f.Use(flamego.Static(flamego.StaticOptions{Directory: filepath.Join(conf.CustomDir(), "public")}))

	var publicFs http.FileSystem
	if !conf.Web.LoadAssetsFromDisk {
		publicFs = public.NewFileSystem()
	}

	f.Use(flamego.Static(flamego.StaticOptions{
		Directory:     filepath.Join(conf.WorkDir(), "public"),
		FileSystem:    publicFs,
		EnableLogging: false,
	}))

	// template
	renderOpt := template.Options{
		Directory:         filepath.Join(conf.WorkDir(), "templates"),
		AppendDirectories: []string{filepath.Join(conf.CustomDir(), "templates")},
		FuncMaps:          tmpl.FuncMaps(),
	}

	if !conf.Web.LoadAssetsFromDisk {
		renderOpt.FileSystem = templates.NewTemplateFileSystem("", renderOpt.AppendDirectories[0])
	}

	f.Use(template.Templater(renderOpt))
	f.Use(brotli.Brotli())
	return f
}

func setRouter(f *flamego.Flame) {

	f.Group("", func() {
		f.Get("/", router.Home)
		f.Get("/prev/{pos}", router.Prev)
		f.Get("/next/{pos}", router.Next)
		f.Get("/rand", router.Rand)
		f.Get("/about", router.About)

		f.Get("/so/{kw}.html", router.So)
		f.Get("/so/{kw}/{prevNext}/{pos}.html", router.So)
		f.Get("/html/{pos}.html", router.ContentHtml)
	}, context.Contexter())

}

func runWebService(c *cli.Context) error {

	//check image server
	go func() {
		for {
			pingUrl := conf.Image.Ping

			if pingUrl != "" {
				r, err := tools.GetHttpData(pingUrl)

				// fmt.Println(pingUrl, r, err)
				if err != nil {
					conf.Image.PingStatus = false
				}

				// fmt.Println("..", r, conf.Image.PingResponse, conf.Image.PingStatus)
				if strings.EqualFold(r, conf.Image.PingResponse) {
					conf.Image.PingStatus = true
				}
			}

			// fmt.Println("conf.Image.PingStatus:", conf.Image.PingStatus)
			time.Sleep(time.Second * 25)
		}
	}()

	// go tool pprof -http=:11112 --seconds 30 http://127.0.0.1:11012/debug/pprof/profile
	if conf.App.RunMode != "prod" {
		go func() {
			port := ":" + conf.Debug.Port
			http.ListenAndServe(port, nil)
		}()
	}

	f := newFlamego()
	setRouter(f)
	f.Run(conf.Web.HttpPort)

	return nil
}
