# Go Language Web Base

## Run localhost
`go run main.go`

## Deploying a Golang app to Heroku
`heroku create -b https://github.com/heroku/heroku-buildpack-go.git`
`go mod init github.com/bxcodec/sample`
`git add .`
`git commit -m "your massage"`
`git push heroku master`