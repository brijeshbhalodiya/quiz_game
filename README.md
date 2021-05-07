# Quiz game

Simple CLI based Quiz Game which reads the problems from the CSV file & at the end display the result

## Description

A time based quiz game.

Default value of timeout is 30sec & questions are stored in the problems.csv

User can optionally configure the timeout or provide another quiz file.

## Use

To display the different CLI options

```
go run main.go -h
```

To run with the default configuration

```
go run main.go
```

To run with the different problemt statement file

```
go run main.go -csv <filename>
```

To run with different timeout (in seconds)

```
go run main.go -limit <x>
```
