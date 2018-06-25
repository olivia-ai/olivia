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

Install a redis database and set the environment variables, by default it is `localhost:6379` and no password

```
REDIS_ADDRESS = xxx.xxx:xxx
REDIS_PASSWORD = xxx
```

## Usage

Run the `main.go`: 

```
$ go run main.go
```

## Contributing

1. Fork it (https://github.com/OliviaBot/Api/fork)
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Contributors

- [ananagame](https://github.com/ananagame) - creator, maintainer
