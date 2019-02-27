package main

import (
	_ "github.com/BTBurke/caddy-jwt"
	"github.com/mholt/caddy/caddy/caddymain"
	_ "github.com/xadereq/loginsrv/caddy"
)

func main() {
	caddymain.Run()
}
