# Stic 🍊 - Navigate HN in your terminal

*Stic* is a *maybe-a-little-over-complicated* CLI to navigate Hacker News from the terminal that uses [bubbletea](https://github.com/charmbracelet/bubbletea) for the TUI.

## Motivation
I'm working on some projects that can benefit from an advanced TUI. So this project is something like a "light" contact with [bubbletea](https://github.com/charmbracelet/bubbletea) and its ecosystem.

I also wanted a very simple and easy way to see from the terminal what's new in the [orange site](https://news.ycombinator.com/).

## Installation

For the moment you can just clone this repository and build the project manually or just install it using `go install` as follows:

````shell
go install github.com/daniarlert/stic@latest
````

## Usage
To see the list of available options you can run:
```shell
stic -h

# Or

stic --help
```

This will prompt you something like this:
```text
NAME:
   stic - hn in the terminal

USAGE:
   stic [global options] command [command options] [arguments...]

VERSION:
   v0.0.3

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --category value, -c value  hn category. Available categories are: "top", "news", "best", "ask", "show" and "job" (default: "top")
   --debug, -d                 enables debug mode (default: false)
   --help, -h                  show help (default: false)
   --json                      outputs JSON object (default: false)
   --light                     enables light color scheme (default: false)
   --max value, -m value       max number of items (default: 20)
   --no-limit                  fetches as many stories as possible ignoring the `--max` flag
   --version, -v               print the version (default: false)
```

The two main flags are `--category` or `-c` and `--max` or `-m`. The `--category` flag selects the HN category from which to download the stories, which number will be specified by the `--max`'s flag value.

````shell
stic --category news

stic --category job --max 15
````

## Navigation
You can go to the previous story by using `k` or the `ArrowUp`. As well as the `ArrowLeft` or the `k` key. To go to the next story press `ArrowDown`, `j`, `ArrowRight` or the `l` key.

You can alter between full screen by using the `spacebar`. And to open the URL of the story you want to in the browser press `enter`.

To see a list with all the options use `?`. And to quit `q`.