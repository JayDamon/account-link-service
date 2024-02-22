module github.com/factotum/moneymaker/account-link-service

go 1.19

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/go-chi/cors v1.2.1
	github.com/jaydamon/http-toolbox v0.0.0-20230114132444-809dfa8092f7
)

require (
	github.com/jaydamon/moneymakergocloak v0.0.0-20230916210526-12136784735d
	github.com/jaydamon/moneymakerplaid v0.0.0-20230221115648-a8aa3efc6a1c
	github.com/jaydamon/moneymakerrabbit v0.0.0-20231018224209-6a93251ce145
	github.com/joho/godotenv v1.5.1
	github.com/plaid/plaid-go v1.10.0
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/Nerzal/gocloak/v12 v12.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-resty/resty/v2 v2.10.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rabbitmq/amqp091-go v1.9.0 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/oauth2 v0.12.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/jaydamon/moneymakergocloak => ../moneymakergocloak

replace github.com/jaydamon/moneymakerrabbit => ../moneymakerrabbit
