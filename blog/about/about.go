package about

import "darklab_blog/blog/common/compon"

type About struct {
	compon.Component
}

const (
	AboutCSSURL string = "about/about.css"
)

func NewAbout() *About {
	return &About{
		compon.NewComponent(
			"about.html",
			AboutT(),
		),
	}
}
