# battlesnake-go

A simple [BattleSnake AI](http://battlesnake.io) written in Go.

To get started you'll need a working Go development environment.

If you haven't setup a Go development environment before, read [how to get started with Go](https://golang.org/doc/install).

You'll also need the [Godep](https://github.com/tools/godep) dependency management tool

You can also [![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

### Running the AI locally

Fork and clone this repo:
```
cd $GOPATH/github.com/sendwithus/battlesnake-go
git clone git@github.com:sendwithus/battlesnake-go.git
cd battlesnake-go
```

Load dependencies:
```
godep restore
```

Compile the battlesnake-go server:
```
go build
```
This will create a `battlesnake-go` executable.

Run the server:
```
./battlesnake-go
```

Test the client in your browser: [http://127.0.0.1:9000](http://127.0.0.1:9000)


### Deploying to Heroku

Save godep dependencies:
```
godep save
git add .
git commit -m "save godependencies"
```

Create a new Go Heroku app:
```
heroku create [APP_NAME] --buildpack https://github.com/kr/heroku-buildpack-go
```

Push code to Heroku servers:
```
git push heroku master
```

Open Heroku app in browser:
```
heroku open
```
Or go directly via http://APP_NAME.herokuapp.com

View/stream server logs:
```
heroku logs --tail
```
