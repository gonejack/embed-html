package embedhtml

import (
	"path/filepath"

	"github.com/alecthomas/kong"
)

type Options struct {
	Verbose bool `short:"v" help:"Verbose printing."`
	About   bool `help:"About."`

	HTML []string `name:".html" arg:"" optional:"" help:"list of .html files"`
}

func MustParseOption() (opt Options) {
	kong.Parse(&opt,
		kong.Name("embed-html"),
		kong.Description("This command line converts .html file into single complete html."),
		kong.UsageOnError(),
	)
	if len(opt.HTML) == 0 || opt.HTML[0] == "*.html" {
		opt.HTML, _ = filepath.Glob("*.html")
	}
	return
}
