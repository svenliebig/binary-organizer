# 👻 bo(o) - Binary Organizer

What is 👻 bo(o)? A small CLI Tool to organize different versions of your binaries in you $PATH.

## Why?

Modern development is evolving fast and so are the tools we use. Sometimes we have the need to use three or more different versions of the same binary, across different projects (talking to you NodeJS). Despite that there are tools like [`nvm`](https://github.com/nvm-sh/nvm), [`pyenv`](https://github.com/pyenv/pyenv) or other tools for other languages, binaries and tools, I prefer to have actually *less* installed on my system, so I have more control over the behavior of my system. That's what `👻 bo(o)` is trying to solve.

## Installation

1) Download the latest release from the [releases page](...)
2) Extract the archive
3) Create an alias to the `boo.sh` file in your shell profile (read more about 'why alias' [here](#why-alias)):
  * e.g. `alias boo='. /path/to/boo.sh'`
4) Run `boo init` to:
  * create the configuration file
5) Add `boo load` to your shell profile (e.g. `.bashrc`, `.zshrc`, etc.)

## Usage

```bash
boo node # interactively select a version
boo node 14.17.0 # sets the node version to 14.17.0
boo node 14 # sets the node version to the latest 14.x version
boo node -v # prints the current node version

boo mvn 3.8.1 # sets the maven version to 3.8.1
```

## Configuration

👻 bo(o) will lookup the following locations for a configuration file, if not specified with the `-c` flag:

```
~/.config/boo/config.toml
~/.boo/config.toml
~/.config/boo.toml
~/.boo.toml
```

## FAQ

### Why Alias?

The reason why we need an alias for the execution, is the way shell scripts get executed. If you run a shell script, it will be executed in a subshell, which means that the script can't change the environment of the parent shell. That's why we need to source the script with the `.` command, which will execute the script in the current shell, and therefore can change the environment of the current shell. That's also the reason why we need a shell script in the first place, instead of using the go binary directly.
