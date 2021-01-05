# Go Language Web Base

## Run localhost
`go run main.go`

## Deploying a Golang app to Heroku
`heroku create -b https://github.com/heroku/heroku-buildpack-go.git`

`go mod init github.com/bxcodec/sample` i try use this command for make file go.mod from other repo, without this can't running because heroku can't verifying deploy. you can skip this if your project have `go.mod`

`git add .`

`git commit -m "your massage"`

`git push heroku master`