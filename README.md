<h1 align="center">
  <br>
  <img src="https://olivia-ai.org/img/icons/olivia-with-text.png" alt="Olivia's character" width="300">
  <br>
</h1>

<h4 align="center">💁‍♀️ Your new best friend</h4>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/olivia-ai/olivia"><img src="https://goreportcard.com/badge/github.com/olivia-ai/olivia"></a>
  <a href="https://godoc.org/github.com/olivia-ai/olivia"><img src="https://godoc.org/github.com/olivia-ai/olivia?status.svg" alt="GoDoc"></a>
  <a href="https://app.fossa.io/projects/git%2Bgithub.com%2Folivia-ai%2Folivia?ref=badge_shield"><img src="https://app.fossa.io/api/projects/git%2Bgithub.com%2Folivia-ai%2Folivia.svg?type=shield"></a>
  <a href="https://codecov.io/gh/olivia-ai/olivia"><img src="https://codecov.io/gh/olivia-ai/olivia/branch/master/graph/badge.svg" /></a>
  <br>
  <img src="https://github.com/olivia-ai/olivia/workflows/Code%20coverage/badge.svg">
  <img src="https://github.com/olivia-ai/olivia/workflows/Docker%20CI/badge.svg">
  <img src="https://github.com/olivia-ai/olivia/workflows/Format%20checker/badge.svg">
</p>

<p align="center">
  <a href="https://twitter.com/oliv_ai"><img alt="Twitter Follow" src="https://img.shields.io/twitter/follow/oliv_ai"></a>
  <a href="https://discord.gg/wXDwTdy"><img src="https://img.shields.io/discord/699567909235720224?label=Discord&style=social"></a>
</p>

<p align="center">
  <a href="https://www.youtube.com/watch?v=JRSNnW05suo"><img width="250" src="https://i.imgur.com/kEKJjJn.png"></a>
</p>

<p align="center">
  <a href="https://olivia-ai.org">Website</a> —
  <a href="https://docs.olivia-ai.org">Documentation</a> —
  <a href="#getting-started">Getting started</a> —
  <a href="#introduction">Introduction</a> —
  <a href="#translations">Translations</a> —
  <a href="#contributors">Contributors</a> —
  <a href="#license">License</a>
</p>

<p align="center">
  ⚠️ Please check the <strong><a href="https://github.com/olivia-ai/olivia/issues">Call for contributors</a></strong>
</p>

## Introduction
<p align="center">
  <img alt="introduction" height="100" src="https://i.imgur.com/Ygm9CMc.png">
</p>

### Description
Olivia is an open-source chatbot built in Golang using Machine Learning technologies.
Its goal is to provide a free and open-source alternative to big services like DialogFlow. 

You can chat with her by speaking (STT) or writing, she replies with a text message but you can enable her voice (TTS).

You can clone the project and customize it as you want using [GitHub](https://github.com/olivia-ai/olivia)
Try it on [her website!](https://olivia-ai.org)

### Why Olivia?
- The only chatbot project in Go that could be modulable and customizable.
- Using daily a privacy-friendly chatbot is great.
- The Website is a Progressive Web Application, which means you can add it to your phone and it seems like a native app!


## Getting started
### Installation 
#### Login to Github 

To get a personal access token from Github go to `Setings > Developer settings > Personal Access Tokens`

Click on Generate new Token and name it you MUST have read and write packages ticked on.
Then click Generate new token

Replace `TOKEN` with the Token that you just made.
```bash
$ export PAT=TOKEN
```

Login to Github (Note: change USERNAME to Gthub username)
```bash
$ echo $PAT | docker login docker.pkg.github.com -u USERNAME --password-stdin
```

#### Docker

<p align="center">
  <img alt="docker installation" height="100" src="https://i.imgur.com/5NDCfF3.png">
</p>

Pull the image from GitHub Packages
```bash
$ docker pull docker.pkg.github.com/olivia-ai/olivia/olivia:latest
```

Then start it
```bash
$ docker run -d -e PORT=8080 -p 8080:8080 docker.pkg.github.com/olivia-ai/olivia/olivia:latest
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

The app will automatically check for `res/datasets/training.json` file which contains the save of the neural network.
By default when you clone the repository from Github you have a stable save.
If you want to train a new model just delete this file and rerun the app.

#### GitHub
<p align="center">
  <img height="100" src="https://i.imgur.com/RRPoP69.png">
</p>

Clone the project via GitHub:

```bash 
$ git clone git@github.com:olivia-ai/olivia.git
```

Then download the dependencies
```bash
$ go mod download
```

And run it
```bash
$ go run main.go
```

### Frontend and Backend
To install the frontend and the backend together, please use the `docker-compose.yml` file:

```bash
$ docker-compose up
```

And all done!

## Architecture
<p align="center">
  <img alt="architecture" height="85" src="https://i.imgur.com/95h8WIU.png">
  <br>
  <img src="https://i.imgur.com/G9BYf4Y.png">
</p>

## Translations

<p align="center">
  <img alt="introduction" height="130" src="https://i.imgur.com/MDKbP0R.png">
</p>

### Languages supported
- <img src="https://i.imgur.com/URqxsb0.png" width="25"> English
- <img src="https://i.imgur.com/Oo5BNk0.png" width="25"> Spanish
- <img src="https://i.imgur.com/2DWxeF9.png" width="25"> Catalan
- <img src="https://i.imgur.com/0dVqbjf.png" width="25"> French
- <img src="https://i.imgur.com/sXLQp8e.png" width="25"> German
- <img src="https://i.imgur.com/DGNcrRF.png" width="25"> Italian
- <img src="https://i.imgur.com/kB0RoFZ.png" width="25"> Brazilian portuguese - not completed

### Coverage
The coverage of the translations is given [here](https://olivia-ai.org/dashboard/language).
To add a language please read [the documentation for that](https://docs.olivia-ai.org/translations.html).

## Contributors

<p align="center">
  <img alt="docker installation" height="85" src="https://i.imgur.com/6xr2zdp.png">
</p>
  
### Contributing
Please refer to the [contributing file](.github/CONTRIBUTING.md)
  
### Code Contributors
Thanks to the people who contribute to Olivia. 

[Contribute](.github/CONTRIBUTING.md)
<a href="https://github.com/olivia-ai/olivia/graphs/contributors"><img src="https://opencollective.com/olivia-ai/contributors.svg?width=950&button=false" /></a>

### Financial Contributors
Become a financial contributor and help Olivia growth. 

Contribute on the GitHub page of [hugolgst](https://github.com/sponsors/hugolgst) ❤️

## License

<p align="center">
  <img src="https://i.imgur.com/9Xxtchv.png" height="90">
</p>

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Folivia-ai%2Folivia.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Folivia-ai%2Folivia?ref=badge_large)

<p align="center">
  <img width="60" src="https://olivia-ai.org/img/icons/olivia.png">
<p>

<p align="center">
  Made with ❤️ by <a href="https://github.com/hugolgst">Hugo Lageneste</a>
</p>

![Olivia's wave](https://olivia-ai.org/img/background-olivia.png)
