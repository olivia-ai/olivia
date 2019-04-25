<h1 align="center">
  <br>
  <img src="https://i.imgur.com/Xz0DUXf.png" alt="Cute character" width="400">
  <br>
</h1>

<h4 align="center">Your new best friend</h4>

<p align="center">
  <a href="https://travis-ci.org/olivia-ai/olivia"><img src="https://travis-ci.org/olivia-ai/olivia.svg?branch=master"></a>
</p>

<p align="center">
  <a href="#how-to-use">How To Use</a> •
  <a href="#license">License</a>
</p>

## How To Use

The requests must be POST requests and point to `/api/response` with the following parameters :
```
sentence=hey
authorId=192003943049
```

Choose your API :

- [Use your API version](#setup-the-api)
- [Use official API](#official-api)

### Setup the API
Clone the API:
```
$ git clone https://github.com/olivia-ai/olivia.git
```

Then run it: 
```
$ go run main.go
```

### Official API

Just point your request to [https://olivia-api.herokuapp.com/](https://olivia-api.herokuapp.com/)


## License

MIT