# Go Language Web Base

## Run localhost
`go run main.go`

## Deploying a Golang app to Heroku
`heroku create -b https://github.com/heroku/heroku-buildpack-go.git`

`go mod init github.com/bxcodec/sample` i try use this module for make file go.mod, without this can running because heroku cannot verifying deploy. you can skip this if your project have go.mod

`git add .`

`git commit -m "your massage"`

`git push heroku master`