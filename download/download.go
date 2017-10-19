package download

import (
	"net/http"
	"sort"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

const (
	url     = "http://files.data.gouv.fr/sirene/"
	datefmt = "2006-01-02  15:04:05"
)

func gather(n *html.Node) bool {
	return n.DataAtom == atom.Tr && scrape.Attr(n, "class") == "item type-application type-zip"
}

func link(n *html.Node) bool {
	return n.DataAtom == atom.A && scrape.Attr(n, "href") != ""
}

// NewFileFromRow takes an HTML node and returns the structured file
func NewFileFromRow(n *html.Node) (*File, error) {
	var err error
	var t time.Time
	var h, date *html.Node
	var ok bool

	if h, ok = scrape.Find(n, link); !ok {
		return nil, errors.New("no link for this entry")
	}
	if date, ok = scrape.Find(n, scrape.ByClass("coldate")); !ok {
		return nil, errors.New("no date for this entry")
	}
	if t, err = time.Parse(datefmt, scrape.Text(date)); err != nil {
		return nil, errors.Wrap(err, "couldn't parse date for this entry")
	}
	f := File{
		Name: scrape.Text(h),
		Link: url + scrape.Attr(h, "href"),
		Date: t,
	}
	return &f, nil
}

func Start() error {
	var err error
	var f *File
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	ff := Files{}
	rows := scrape.FindAllNested(root, gather)
	for _, row := range rows {
		if f, err = NewFileFromRow(row); err != nil {
			continue
		}
		ff = append(ff, f)
	}
	sort.Sort(ff)
	for _, file := range ff {
		spew.Dump(file)
	}
	return nil
}
