module github.com/lty5240/blue/service-user

go 1.13

replace github.com/golang/lint => golang.org/x/lint v0.0.0-20200130185559-910be7a94367

replace github.com/testcontainers/testcontainer-go => github.com/testcontainers/testcontainers-go v0.2.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/golang/protobuf v1.3.3
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.0.0
	github.com/micro/go-plugins/registry/consul/v2 v2.0.2
	golang.org/x/net v0.0.0-20200202094626-16171245cfb2
)
