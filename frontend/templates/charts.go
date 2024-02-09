package templates

import (
	"context"
	"github.com/a-h/templ"
	"io"
)

type Renderable interface {
	Render(w io.Writer) error
}

func ConvertChartToTemplComponent(chart Renderable) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return chart.Render(w)
	})
}
