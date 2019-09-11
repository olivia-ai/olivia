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
  <img src="https://circleci.com/gh/olivia-ai/olivia/tree/master.svg?style=svg">
</p>

<p align="center">
  <a href="https://olivia-ai.org">Website</a> •
  <a href="#getting-started">Getting started</a> •
  <a href="https://docs.olivia-ai.org">Documentation</a> •
  <a href="https://github.com/orgs/olivia-ai/projects">Projects</a> •
  <a href="https://www.youtube.com/watch?v=JmJZi9gmKvI">Video</a> •
  <a href="#license">License</a>
</p>

## Getting started
### Installation
Clone Olivia from the master branch of Github repository

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

The Websocket is now listening on the port `8080`, to change it just set it inside the environment variable `PORT`

The app will automatically check for `res/training.json` file which contains the save of the neural network.
By default when you clone the repository from Github you have a stable save.
If you want to train a new model just delete this file and rerun the app.

### How to use
Connect to `wss://olivia-api.herokuapp.com/` and send a JSON message like this

```json
{
  "content": "Hello!",
  "authorid": "129390230"
}
```

and the websocket will respond you with 
```json
{
  "content": "Good morning!",
  "tag": "hello"
}
```

## Contributors

### Code Contributors

This project exists thanks to all the people who contribute. [[Contribute](CONTRIBUTING.md)].
<a href="https://github.com/olivia-ai/olivia/graphs/contributors"><img src="https://opencollective.com/olivia-ai/contributors.svg?width=890&button=false" /></a>

### Financial Contributors

Become a financial contributor and help us sustain our community. [[Contribute](https://opencollective.com/olivia-ai/contribute)]

#### Individuals

<a href="https://opencollective.com/olivia-ai"><img src="https://opencollective.com/olivia-ai/individuals.svg?width=890"></a>

#### Organizations

Support this project with your organization. Your logo will show up here with a link to your website. [[Contribute](https://opencollective.com/olivia-ai/contribute)]

<a href="https://opencollective.com/olivia-ai/organization/0/website"><img src="https://opencollective.com/olivia-ai/organization/0/avatar.svg"></a>
<a href="https://opencollective.com/olivia-ai/organization/1/website"><img src="https://opencollective.com/olivia-ai/organization/1/avatar.svg"></a>
<a href="https://opencollective.com/olivia-ai/organization/2/website"><img src="https://opencollective.com/olivia-ai/organization/2/avatar.svg"></a>
<a href="https://opencollective.com/olivia-ai/organization/3/website"><img src="https://opencollective.com/olivia-ai/organization/3/avatar.svg"></a>
<a href="https://opencollective.com/olivia-ai/organization/4/website"><img src="https://opencollective.com/olivia-ai/organization/4/avatar.svg"></a>
<a href="https://opencollective.com/olivia-ai/organization/5/website"><img src="https://opencollective.com/olivia-ai/organization/5/avatar.svg"></a>
<a href="https://opencollective.com/olivia-ai/organization/6/website"><img src="https://opencollective.com/olivia-ai/organization/6/avatar.svg"></a>
<a href="https://opencollective.com/olivia-ai/organization/7/website"><img src="https://opencollective.com/olivia-ai/organization/7/avatar.svg"></a>
<a href="https://opencollective.com/olivia-ai/organization/8/website"><img src="https://opencollective.com/olivia-ai/organization/8/avatar.svg"></a>
<a href="https://opencollective.com/olivia-ai/organization/9/website"><img src="https://opencollective.com/olivia-ai/organization/9/avatar.svg"></a>

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Folivia-ai%2Folivia.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Folivia-ai%2Folivia?ref=badge_large)

<p align="center">
  Made with ❤️ by <a href="https://github.com/ananagame">Hugo Lageneste</a>
</p>

![Olivia's wave](https://olivia-ai.org/img/background-olivia.png)
