### Local Development
#### Add the following to `go.mod` if you need direct access to local changes made for those libraries
`replace github.com/jaydamon/moneymakergocloak => ../moneymakergocloak`
`replace github.com/jaydamon/moneymakerplaid => ../moneymakerplaid`
`replace github.com/jaydamon/moneymakerrabbit => ../moneymakerrabbit`
### Update version
#### Get latest commit
`go get github.com/someone/some_module@af044c0995fe`