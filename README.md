# adventofcode/2025
This is a repo for [advent of code 2025](https://adventofcode.com/2025) solutions in Go.

Utilises [cobra](https://github.com/spf13/cobra) to create a cli.

## Structure
Solutions and tests are in ```/challenge/aoc2025/<day>/``` with
* ```<day>.go```
* ```<day>_test.go```
* ```<day>_cmd.go```

Input files are in ```/input/aoc2025/<day>.txt```

### Setup
```/main.go``` uses ```/cmd/cmd.go``` to create the cli and include each year (e.g. ```/challenge/aoc2025/aoc2025.go```) as a command.

The year command adds each day (e.g. ```/challenge/aoc2025/day01/day01_cmd.go```) as a subcommand.

### Prep for a new day
A code generator is in ```/gen/day``` with templates as ```/gen/*.tmpl```.

To use:
```go run ./gen/day -year=<year> -day=<day>```

2025 is the default year so to create for day01 use:

```go run ./gen/day -day=day01```
```go run ./gen/pull_input```

This will generate:
- ```/challenge/aoc2025/```
    - ```day01/```
        - ```day01.go```
        - ```day01_test.go```
    - ```aoc2025.go```
- ```/cmd/aoc/2025/```
    - ```day01/```
        - ```day01_cmd.go```
- ```/input/aoc2025```
    - ```day01.txt```

Note: if the challenge directory exists then the generator will exit without creating/overwriting anything:
- ```/challenge/aoc<year>/<day>/```

### Pulling input
AOC request that input files are not stored in repositories. See [here](https://adventofcode.com/2024/about)

A little util to pull all of the input files can be used as
```go run ./gen/pull_input```
The ```./input``` is git ignored.

This reads the session token from a file called ```./input/token```. You can find your token as the ```session``` from under cookies from the firefox dev tools.

## Usage example
### Single solution
```go run main.go <year> <day>```

```go run main.go 2025 day01```
### All solutions for a year
```go run main.go <year>```

```go run main.go 2025```

### All solutions
```go run main.go```
