## battlesnake-go

A simple [Battlesnake AI](http://battlesnake.io) written in Go.

Visit [https://github.com/battlesnakeio/community/blob/master/starter-snakes.md](https://github.com/battlesnakeio/community/blob/master/starter-snakes.md) for API documentation and instructions for running your AI.

To get started, you'll need:
  1. A working Go development environment ([guide](https://golang.org/doc/install)).
  2. Read [Heroku's guide to deploying Go apps](https://devcenter.heroku.com/articles/getting-started-with-go#introduction)

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

### Running the AI locally

1) [Fork this repo](https://github.com/sendwithus/battlesnake-go/fork).

2) Clone repo to your development environment:
```
git clone git@github.com:USERNAME/battlesnake-go.git $GOPATH/github.com/USERNAME/battlesnake-go
cd $GOPATH/github.com/USERNAME/battlesnake-go
```

3) Compile and run the server with:

```
Make run
```

4) Test the client in your browser: [http://127.0.0.1:9000/start](http://127.0.0.1:9000/start)

### Running tests locally

```
Make test
```

Note: if you're missing any packages, use `Make get`.

### Deploying to Heroku

1) Create a new Go Heroku app using Go buildpack.
```
heroku create
```

2) Push code to Heroku servers.
```
git push heroku master
```

3) Open Heroku app in browser.
```
heroku open
```
Or go directly via http://APP_NAME.herokuapp.com

4) View/stream server logs.
```
heroku logs --tail
```

### Questions?

[Email](mailto:battlesnake@sendwithus.com), [Twitter](http://twitter.com/send_with_us)
