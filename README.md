<h1 align="center">
  <br>
  <img src="https://i.imgur.com/Xz0DUXf.png" alt="Olivia's character" width="400">
  <br>
</h1>

<h4 align="center">💁‍♀️ Your new best friend built with an artificial neural network</h4>
<h5 align="center">Inspired by <a href="https://github.com/leon-ai/leon">leon-ai/leon</a> :)</h5>

<p align="center">
  <a href="https://travis-ci.org/olivia-ai/olivia"><img src="https://travis-ci.org/olivia-ai/olivia.svg?branch=master"></a>
  <a href="https://app.fossa.io/projects/git%2Bgithub.com%2Folivia-ai%2Folivia?ref=badge_shield"><img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Folivia-ai%2Folivia.svg?type=shield"></a>
</p>

<p align="center">
  <a href="#getting-started">Getting started</a> •
  <a href="https://docs.olivia-ai.org">Documentation</a> •
  <a href="https://github.com/orgs/olivia-ai/projects">Projects</a> •
  <a href="https://www.youtube.com/watch?v=JmJZi9gmKvI">Video</a> •
  <a href="#contributors">Contributors</a> •
  <a href="#license">License</a>
</p>

## Getting started
### Installation
Clone Olivia's REST Api from the master branch of Github repository

```bash
git clone https://github.com/olivia-ai/olivia.git
```

Then go inside the project and install the dependencies

```bash
cd olivia

# Install the dependencies with dep (https://github.com/golang/dep)
dep ensure
```

And run the application

```bash
go run main.go
```

The REST Api is now listening on the port `8080`, to change it just set it inside the environment variable `PORT`

The app will automatically check for `res/training.json` file which contains the save of the neural network.
By default when you clone the repository from Github you have a stable save.
If you want to train a new model just delete this file and rerun the app.

### How to use
To use the REST Api you must establish `POST` request to `/api/response` with two parameters:
- `sentence` which is the message you want to send to Olivia
- `authorId` which is an arbitrary ID to identify the user for having a contextual chat

The latest release is online at `https://olivia-api.herokuapp.com`

#### Example with curl
```bash
curl -X POST 'https://olivia-api.herokuapp.com/api/response' --data "sentence=Hello" --data "authorId=81278329032"
```

The response arrives in this format

```json
{
  "content": "Good morning!",
  "tag": "hello"
}
```

## Donate

### Backers
<a href="https://opencollective.com/olivia-ai#backer" target="_blank"><img src="https://opencollective.com/olivia-ai/tiers/backer.svg?width=890"></a>

### Sponsors
<a href="https://opencollective.com/olivia-ai/sponsor/0/website" target="_blank"><img src="https://opencollective.com/olivia-ai/sponsor/0/avatar.svg"></a>

## Contributors
- [ananagame](https://github.com/ananagame) - creator, maintainer

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Folivia-ai%2Folivia.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Folivia-ai%2Folivia?ref=badge_large)

