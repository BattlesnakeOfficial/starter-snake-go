## battlesnake-go

A simple [BattleSnake AI](http://battlesnake.io) written in Go.

Visit [battlesnake.io/readme](http://battlesnake.io/readme) for API documentation and instructions for running your AI.

To get started, you'll need:
  1. A working Go development environment ([guide](https://golang.org/doc/install)).
  2. Experience [deploying Go apps to Heroku](https://devcenter.heroku.com/articles/getting-started-with-go#introduction)
  3. Your Snake ID from http://www.battlesnake.io/team

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

### Running the AI locally

1) [Fork this repo](https://github.com/sendwithus/battlesnake-go/fork).

2) Clone repo to your development environment:
```
git clone git@github.com:USERNAME/battlesnake-go.git $GOPATH/github.com/USERNAME/battlesnake-go
cd $GOPATH/github.com/USERNAME/battlesnake-go
```

3) Compile the battlesnake-go server.
```
go build
```
This will create a `battlesnake-go` executable.

4) Set your snake ID as an environment variable.
```
export SNAKE_ID=ABCDEF1234
```
This will allow your snake to locate itself during the game.

5) Run the server.
```
./battlesnake-go
```

6) Test the client in your browser: [http://127.0.0.1:9000](http://127.0.0.1:9000)


### Deploying to Heroku

1) Create a new Go Heroku app using Go buildpack.
```
heroku create [APP_NAME] --buildpack https://github.com/kr/heroku-buildpack-go
```

2) Push code to Heroku servers.
```
git push heroku master
```

3) Set the SNAKE_ID environment variable in your Heroku app.
```
heroku config:set SNAKE_ID=ABCDEF1234
```

4) Open Heroku app in browser.
```
heroku open
```
Or go directly via http://APP_NAME.herokuapp.com

5) View/stream server logs.
```
heroku logs --tail
```

### Questions?

[Email](mailto:battlesnake@sendwithus.com), [Twitter](http://twitter.com/send_with_us)
