# Advent of Code Solutions

Here be answers/solutions to Advent of Code challenges for various years.

Answers are divided into individual subdirectories for each year and day.

`cmd/aoc` contains a small CLI tool for automatically generating subdirectory structure, pre-populating each day with files to be completed.

`cmd/aoc-announce`, and the associated Docker image, `ghcr.io/arylatt/advent-of-code/aoc-announce:latest`, is a tool for watching a private leaderboard and sending notifications to a Discord channel via a webhook when players complete puzzles.

`elves` is my helper package. It contains functionality to:

* Enable easier testing of the samples provided as part of the challenge descriptions
* Automatically pull the full inputs when running the test case (uses session cookie placed in `AOC_SESSION_COOKIE` env var)
* Automatically submit answer from test case to AoC (uses `AOC_SESSION_COOKIE` and is enabled by setting `AOC_SUBMIT_ANSWERS` to any non-empty value)

## Friends

**[rmhyde](https://github.com/rmhyde)**

[AoC 2021](https://github.com/rmhyde/advent-of-code-2021), [AoC 2022](https://github.com/rmhyde/advent-of-code-2022), [AoC 2023](https://github.com/rmhyde/advent-of-code-2023), [AoC 2024](https://github.com/rmhyde/advent-of-code-2024)

**[CallumLRT](https://github.com/CallumLRT)**

[AoC 2023, 2024](https://github.com/CallumLRT/advent-of-code)

**[dantdj](https://github.com/dantdj)**

[AoC 2020, 2023, 2024](https://github.com/dantdj/AdventOfCode)

![ooh, friend!](https://media.tenor.com/Y_EhxEaS4MEAAAAC/friend.gif)
