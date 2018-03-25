# Wordy
Word bank for the command line.

## Introduction
Wordy is a personal collection of words that you find interesting or useful in writing.
It helps you keep track of the words you add, intelligently filter through the existing
words, and also tag the words to seperate them into different categories.

## Installation
You can download the binaries from the [releases](https://github.com/ravernkoh/wordy/releases)
page.

Alternatively, you can build it from source by installing the package
```bash
$ go get -u github.com/ravernkoh/wordy

# Also requires these packages
$ go get -u github.com/mitchellh/go-homedir
$ go get -u github.com/texttheater/golang-levenshtein/levenshtein
```

Pre-defined word banks can be downloaded from this repository as well *(coming soon)*.


## Usage
Add a new word
```bash
$ wordy add
Word: Tumultuous
Description: Messy; Disorganised
Tag(s): 

Tumultuous
	Messy; Disorganised
	adjective, negative

Added!
```

Search existing words (uses Leveshtein's distance)
```bash
$ wordy search
Search: Ttous
2 result(s)

Tumultuous
	Messy; Disorganised
	adjective, negative

Tetanus
	Infection that causes muscle spasms
	noun, negative, infection

Found!
```

Remove existing word
```bash
$ wordy remove
Word: Tumultuous
Are you sure? (y/n)

Tumultuous
	Messy; Disorganised
	adjective, negative

Removed!
```

List possible options
```bash
$ wordy help
Wordy is a personal collection of words that you find interesting or useful in writing.
It helps you keep track of the words you add, intelligently filter through the existing
words, and also tag the words to seperate them into different categories.

Usage: wordy [command]

Commands:
	add
		Add a new word
	search
		Search existing words
	remove
		Remove existing word
	help
		Displays help


$ wordy help add
Add a new word

Usage: wordy add

Flags:
	-file
		File to save to (default: ~/.wordy)


$ wordy help search
Search existing words

Usage: wordy search

Flags:
	-file
		File to save to (default: ~/.wordy)


$ wordy help remove
Remove existing word

Usage: wordy remove

Flags:
	-file
		File to save to (default: ~/.wordy)
	-confirm
		Confirm before removing (default: true)

```
