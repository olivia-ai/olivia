[![Build Status](https://travis-ci.org/olivia-ai/Api.svg?branch=master)](https://travis-ci.org/olivia-ai/Api)

# Api

Olivia's REST Api 

## Installation

Clone the project :

```
$ git clone https://github.com/olivia-ai/Api.git
```


You can change the port with the PORT environment variable, by default it is 8080


## Usage

Note:

A public version is enabled at https://olivia-api.herokuapp.com/

Run the `main.go`: 

```
$ go run main.go
```

Then send a POST request at `localhost:8080/api/response` with `sentence` and `authorId` parameters in the request Body

The API responds with this:

```json
{
  "content": "Plutôt pas mal et toi ?"
}
```

## Contributing

1. Fork it (https://github.com/olivia-ai/Api/fork)
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Contributors

- [ananagame](https://github.com/ananagame) - creator, maintainer
