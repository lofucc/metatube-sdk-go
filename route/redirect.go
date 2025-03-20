package route

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/lofucc/metatube-sdk-go/common/parser"
	"github.com/lofucc/metatube-sdk-go/engine"
	"github.com/lofucc/metatube-sdk-go/model"
	mt "github.com/lofucc/metatube-sdk-go/provider"
)

func redirect(app *engine.Engine) gin.HandlerFunc {
	const (
		separator = ":"
		queryKey  = "redirect"
	)
	return func(c *gin.Context) {
		if redir := c.Query(queryKey); redir != "" {
			provider, id, found := strings.Cut(
				parser.ParseProviderID(redir),
				separator)
			if !found || id == "" {
				abortWithStatusMessage(c, http.StatusBadRequest, "invalid provider id")
				return
			}

			var (
				info any
				err  error
			)
			if id, err = url.QueryUnescape(id); err != nil {
				abortWithError(c, err)
				return
			}

			switch {
			case app.IsActorProvider(provider):
				info, err = app.GetActorInfoByProviderID(provider, id, true)
			case app.IsMovieProvider(provider):
				info, err = app.GetMovieInfoByProviderID(provider, id, true)
			default:
				abortWithError(c, mt.ErrProviderNotFound)
				return
			}
			if err != nil {
				abortWithError(c, err)
				return
			}

			var homepage string
			switch v := info.(type) {
			case *model.ActorInfo:
				homepage = v.Homepage
			case *model.MovieInfo:
				homepage = v.Homepage
			}
			c.Redirect(http.StatusTemporaryRedirect, homepage)

			c.Abort() // abort pending middlewares
			return
		}
		c.Next()
	}
}
