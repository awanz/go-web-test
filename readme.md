# Go Language Web Base

## Run localhost
`go run main.go`

## Deploying a Golang app to Heroku
`heroku create -b https://github.com/heroku/heroku-buildpack-go.git`
`git push heroku master`
`glide create`
`glide install`