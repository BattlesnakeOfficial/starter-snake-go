## battlesnake-go

A simple [BattleSnake AI](http://battlesnake.io) written in Go.

Available at [http://battlesnake-go.herokuapp.com](http://battlesnake-go.herokuapp.com).

To get started, you'll need:
  1. Setup your Go development environment ([guide](https://golang.org/doc/install)).
  2. Install [Godep](https://github.com/tools/godep).

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

### Running the AI locally

1. Fork and clone this repo.
```
cd $GOPATH/github.com/sendwithus/battlesnake-go
git clone git@github.com:sendwithus/battlesnake-go.git
cd battlesnake-go
```

2. Load dependencies.
```
godep restore
```

3. Compile the battlesnake-go server.
```
go build
```
This will create a `battlesnake-go` executable.

4. Run the server.
```
./battlesnake-go
```

5. Test the client in your browser: [http://127.0.0.1:9000](http://127.0.0.1:9000)


### Deploying to Heroku

1. Save godep dependencies:
```
godep save
git add .
git commit -m "save godependencies"
```

2. Create a new Go Heroku app using Go buildpack.
```
heroku create [APP_NAME] --buildpack https://github.com/kr/heroku-buildpack-go
```

3. Push code to Heroku servers.
```
git push heroku master
```

4. Open Heroku app in browser.
```
heroku open
```
Or go directly via http://APP_NAME.herokuapp.com

5. View/stream server logs.
```
heroku logs --tail
```
