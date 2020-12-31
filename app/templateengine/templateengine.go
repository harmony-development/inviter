package templateengine

import (
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/qor/render"
)

var renderer = render.New(&render.Config{
	ViewPaths:     []string{"app/templates"},
	DefaultLayout: "application",
})

type Template struct{}

type dummyByte struct {
	sb strings.Builder
}

func (d dummyByte) Header() http.Header {
	return make(http.Header)
}

func (d dummyByte) WriteHeader(int) {}

func (d *dummyByte) Write(b []byte) (int, error) {
	return d.sb.Write(b)
}

func (d dummyByte) String() string {
	return d.sb.String()
}

func (t Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	d := dummyByte{}
	err := renderer.Execute(name, data, &http.Request{}, &d)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(d.String()))
	return err
}
