package requestid

import (
	"github.com/WedgeServer/wedge"
	"github.com/WedgeServer/wedge/caddyhttp/httpserver"
)

func init() {
	caddy.RegisterPlugin("request_id", caddy.Plugin{
		ServerType: "http",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	for c.Next() {
		if c.NextArg() {
			return c.ArgErr() //no arg expected.
		}
	}

	httpserver.GetConfig(c).AddMiddleware(func(next httpserver.Handler) httpserver.Handler {
		return Handler{Next: next}
	})

	return nil
}
