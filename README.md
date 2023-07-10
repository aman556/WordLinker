# WordLinker

WordLinker is a CLI Developed using the Cobra package in Golang. The use case of this CLI is that it will link all the occurrences of the word with the respective link.

WordLinker has the flexibility to pass the configuration file in the command.

## Commands

### link command

Command to link the words.

``` bash
WordLinker link --configFile <configFilePath> <dirPath>
```

- `--configFile` - Path to the config file
- `dirPath` - Path to the dir over which you want to run the WordLinker command.

### version command

To print the version of WordLinker.

```bash
WordLinker version
```

## How to install WordLinker?

For now, to install WordLinker you need to clone the repo and run the below Golang commands.

- `go build`
- `go install`

## How you can skip a file or files with a particular extension?

To do this you can mention this thing in the config file like below.

``` yaml
- excludepath:
  - ".git/"
  - ".gitignore"
  - ".github/"
```

## How you can mention what set of files or files with a particular extension WordLinker should run on?

To do this you can mention this thing in the config file like below.

``` yaml
includePath:
  - ".md"
```

## How to tell the CLI which word should match with which Link?

To do this you can mention this thing in the config file like below.

``` yaml
wordLinks:
  - youtube: https://www.youtube.com/
```

Note: WIP project
