# Api

Olivia's REST Api 

## Installation

Clone the project :

```
$ git clone https://github.com/OliviaBot/Api.git
```

Get a key on [OpenWeathermap](https://www.openweathermap.org/) and set it in the environment variable

```
WEATHER_KEY = xxx
```

## Usage

Run the `main.go`: 

```
$ go run main.go
```

Then send a POST request at `localhost:8080/api/response` with `sentence` and `authorId` parameters in the request Body

The api respond with this:

```json
{
  "content": "Plutôt pas mal et toi ?"
}
```

## Contributing

1. Fork it (https://github.com/OliviaBot/Api/fork)
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Contributors

- [ananagame](https://github.com/ananagame) - creator, maintainer
