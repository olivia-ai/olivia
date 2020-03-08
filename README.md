<h1 align="center">
  <br>
  <img src="https://olivia-ai.org/img/icons/olivia-with-text.png" alt="Olivia's character" width="300">
  <br>
</h1>

<h4 align="center">💁‍♀️ Your new best friend</h4>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/olivia-ai/olivia"><img src="https://goreportcard.com/badge/github.com/olivia-ai/olivia"></a>
  <a href="https://app.fossa.io/projects/git%2Bgithub.com%2Folivia-ai%2Folivia?ref=badge_shield"><img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Folivia-ai%2Folivia.svg?type=shield"></a>
  <a href="https://travis-ci.org/olivia-ai/olivia"><img src="https://travis-ci.org/olivia-ai/olivia.svg?branch=master"></a>
  <a href="https://circleci.com/gh/olivia-ai/olivia"><img src="https://circleci.com/gh/olivia-ai/olivia/tree/master.svg?style=svg"></a>
</p>

<p align="center">
  <a href="https://olivia-ai.org">Website</a> —
  <a href="https://olivia-ai.org/chat">Chat online</a> —
  <a href="#getting-started">Getting started</a> —
  <a href="https://olivia-ai.org/blog">Blog</a> —
  <a href="https://olivia-ai.org/changelog">Changelog</a> —
  <a href="https://trello.com/b/azB6r2IC/olivia">Trello</a> —
  <a href="#license">License</a>
</p>

## Getting started
### Installation
We will use Docker to setup the project

Pull the image from GitHub Packages
```bash
$ docker pull docker.pkg.github.com/olivia-ai/olivia/olivia:latest
```

Then start it
```bash
$ docker run -d -p 8080:8080 docker.pkg.github.com/olivia-ai/olivia/olivia:latest
```

You can just use the websocket of Olivia now.

To stop it, get the container id:
```bash
$ docker container ls
```
```bash
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS                    NAMES
311b3abb963a        olivia              "./main"            7 minutes ago       Up 7 minutes        0.0.0.0:8080->8080/tcp   quizzical_mayer
```

and stop it
```bash
$ docker container stop 311b3abb963a 
```

The app will automatically check for `res/training.json` file which contains the save of the neural network.
By default when you clone the repository from Github you have a stable save.
If you want to train a new model just delete this file and rerun the app.

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Folivia-ai%2Folivia.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Folivia-ai%2Folivia?ref=badge_large)

<p align="center">
  <img width="60" src="https://olivia-ai.org/img/icons/olivia.png">
<p>

<p align="center">
  Made with ❤️ by <a href="https://github.com/hugolgst">Hugo Lageneste</a>
</p>

![Olivia's wave](https://olivia-ai.org/img/background-olivia.png)
