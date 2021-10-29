---
title: "Installation"
date: 2021-10-27T11:28:33-03:00
---

# Installation


## Manual

Go to [releases](https://github.com/Ronald-TR/todonotion/releases) page and choose the right choice to your platform:

```bash
  curl -LO https://github.com/Ronald-TR/todonotion/releases/download/v0.1.0/todonotion-linux-386.tar.gz
  tar -xzvf todonotion-linux-386.tar.gz
  chmod +x todonotion
  sudo ln -s todonotion /usr/local/bin
```

## Using the installation script

Or, you can use the default `install.sh` script (for linux users) that does the installation with one line:


    curl https://raw.githubusercontent.com/Ronald-TR/todonotion/master/install.sh | bash


You can type `todonotion -h` and you will see the help message:

```
A quick way to generate cards that needs to be done without leaving the terminal.
				Full documentation here: https://github.com/ronald-tr/todonotion

Usage:
  todonotion [flags]
  todonotion [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  del         Delete one card from the Task board
  help        Help about any command
  ls          List all cards for a given Lane on the Task board, accepts filter aggregation
  mv          Moves one card between the Task board lanes
  mvp         Moves all cards between the Task board lanes
  new         Create a card on the Task board

Flags:
  -h, --help   help for todonotion

Use "todonotion [command] --help" for more information about a command.
```

yay! it works!