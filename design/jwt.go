package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var JWT = JWTSecurity("jwt", func() {
	Header("Authorization")
	Scope("api:access", "API access")
})
