# todonotion

A golang rewrite of the following project [todo.sh-notion](https://github.com/Ronald-TR/todo.sh-notion), this version is much faster than the previous one, because uses the official **notion api** instead of a session-based crawler.

# Visit the documentation

Please, visit the official [todonotion](https://ronald-tr.github.io/todonotion/) documentation to know how to setup and use this utility!

## Installation 

Visit the [releases](https://github.com/Ronald-TR/todonotion/releases) page and download the right version for your platform, then add the binary into your PATH directory.

### Homebrew

```
# first of all, tap the repository
brew tap Ronald-TR/todonotion git@github.com:Ronald-TR/todonotion.git # you can try change with https clone method if you are facing any trouble

brew install todonotion
```

to check if the installation goes well, do:

	❯ todonotion version
	v0.2.0

## Configuration

You will need to add in your PATH the following environment variables:

    NOTION_DATABASE_ID='The id of your Task Board'
    NOTION_KEY='Your Notion API key'


## Usage

Following the built-in usage message:

```
❯ ./todonotion -h
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
You can use `todonotion <command> -h` to see the help message for each specific command.

For example, the help message for `new` command:
```
❯ ./todonotion new -h
Create a card on the Task board

Usage:
  todonotion new [flags]

Flags:
  -h, --help          help for new
  -l, --lane string   Lane to create the card. (default "To Do")
  -p, --prop string   Property Select value that you want to tag. (default "todonotion")

```
