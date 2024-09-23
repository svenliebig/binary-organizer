# 👻 bo(o) - Binary Organizer

What is 👻 bo(o)? A small CLI Tool to organize different versions of your binaries in you $PATH.

## Why?

Modern development is evolving fast and so are the tools we use. Sometimes we have the need to use three or more different versions of the same binary, across different projects (talking to you NodeJS). Despite that there are tools like [`nvm`](https://github.com/nvm-sh/nvm), [`pyenv`](https://github.com/pyenv/pyenv) or other tools for other languages, binaries and tools, I prefer to have actually _less_ installed on my system, so I have more control over the behavior of my system. That's what `👻 bo(o)` is trying to solve.

## Installation

1. Make sure to have [Go](https://go.dev/dl/) >= 1.22 installed
2. Clone the repository: `git clone https://github.com/svenliebig/binary-organizer.git`
3. Create an alias to the `boo.sh` file in your shell profile (read more about 'why alias' [here](#why-alias)):

- e.g. `alias boo='. /path/to/binary-organizer/boo.sh'`

4. 🚀 You're ready to go!

## Required Structure

👻 bo(o) will look into the [Configuration](#configuration) file to get the path where it expects the binaries to be installed. Considering the example configuration, the folder structure should look like this:

```
/usr/local/software
├── node
│   ├── node-v16.20.2-darwin-arm64
│   ├── node-v18.20.3-darwin-arm64
│   ├── node-v20.15.0-darwin-arm64
├── maven
│   ├── apache-maven-2.2.0
│   ├── apache-maven-3.6.3
├── go
├── python
├── java
```

`node` is here the binary name, and the subfolders are the versions of the binary as they are named by [NodeJS](https://nodejs.org/en/download/prebuilt-binaries).

## Usage

```bash
boo node 14.17.0 # sets the node version to 14.17.0
boo node 14      # sets the node version to the latest 14.x version
boo node 22      # sets the node version to the latest 22.x version
boo node list    # prints the installed node versions

boo maven 3.8 # sets the maven version to 3.8.x
```

## Supported Binaries

- [x] NodeJS
- [x] Maven
- [ ] Python
- [ ] Java
- [ ] Go

## Configuration

👻 bo(o) will lookup the following locations for a configuration file:

```
~/.config/boo.toml
```

If you don't have a configuration file, 👻 bo(o) will create a default configuration file for you the first time you run it. You can also create the configuration file manually, if you want to customize the configuration.

```toml
# the path where the binaries will be installed
path = "/usr/local/software"

# the default version for the binaries
[defaults]
# the default version to select for node
node = "14.17.0"
```

## FAQ

### Why Alias?

The reason why we need an alias for the execution, is the way shell scripts get executed. If you run a shell script, it will be executed in a subshell, which means that the script can't change the environment of the parent shell. That's why we need to source the script with the `.` command, which will execute the script in the current shell, and therefore can change the environment of the current shell. That's also the reason why we need a shell script in the first place, instead of using the go binary directly.

## Roadmap

- [ ] Add for boo configuration files in the cwd, possibly also with merge
- [ ] Add a boo init command to setup the configuration file
- [ ] Add a boo config subcommand to manage the configuration file
- [ ] Add a boo check subcommand to verify the configured directories in the $PATH environment variable
- [ ] Add `-s` flag for silent mode (refactor the boo output to be centralized), this will be used mainly for the boo init command
