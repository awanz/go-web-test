# Go Language Web Base

## Run localhost
`go run main.go`

## Deploying a Golang app to Heroku

`git init` initialization git repository

`heroku create -b https://github.com/heroku/heroku-buildpack-go.git` create heroku and add buildpack go lang

`go mod init github.com/bxcodec/sample` i try use this command for make file go.mod from other repo, without this can't running because heroku can't verifying deploy. you can skip this if your project have `go.mod`

`git add .` add all file to stage

`git commit -m "your massage"` make tag or massage to your file

`git push heroku master` add push to heroku server