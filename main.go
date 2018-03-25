package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	commandAdd    = "add"
	commandSearch = "search"
	commandRemove = "remove"
	commandHelp   = "help"
)

const (
	helpRoot = `Wordy is a personal collection of words that you find interesting or useful in writing.
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
`
	helpAdd = `Add a new word

Usage: wordy add

Flags:
	-file
		File to save to (default: ~/.wordy)
`
	helpSearch = `Search existing words

Usage: wordy remove

Flags:
	-file
		File to save to (default: ~/.wordy)
`
	helpRemove = `Remove existing word

Usage: wordy remove

Flags:
	-file
		File to save to (default: ~/.wordy)
	-confirm
		Confirm before removing (default: true)
`
)

type add struct {
	filename string
}

func (add *add) run() {
}

type search struct {
	filename string
}

func (search *search) run() {
}

type remove struct {
	filename string
	confirm  bool
}

func (remove *remove) run() {
}

func help(args []string) {
	if len(args) == 0 {
		fmt.Println(helpRoot)
		os.Exit(2)
	}
	switch args[0] {
	case commandAdd:
		fmt.Println(helpAdd)
	case commandSearch:
		fmt.Println(helpSearch)
	case commandRemove:
		fmt.Println(helpRemove)
	case commandHelp:
		fmt.Println(helpRoot)
	default:
		fmt.Println(helpRoot)
		os.Exit(2)
	}
	os.Exit(0)
}

func main() {
	rootFlagSet := flag.CommandLine
	rootFlagSet.SetOutput(ioutil.Discard)
	rootFlagSet.Usage = func() {
		fmt.Println(helpRoot)
	}

	add := new(add)
	addFlagSet := flag.NewFlagSet(commandAdd, flag.ExitOnError)
	addFlagSet.SetOutput(ioutil.Discard)
	addFlagSet.Usage = func() {
		fmt.Println(helpAdd)
	}

	search := new(search)
	searchFlagSet := flag.NewFlagSet(commandSearch, flag.ExitOnError)
	searchFlagSet.SetOutput(ioutil.Discard)
	searchFlagSet.Usage = func() {
		fmt.Println(helpSearch)
	}

	remove := new(remove)
	removeFlagSet := flag.NewFlagSet(commandRemove, flag.ExitOnError)
	removeFlagSet.SetOutput(ioutil.Discard)
	removeFlagSet.Usage = func() {
		fmt.Println(helpRemove)
	}

	if len(os.Args) < 2 {
		fmt.Println(helpRoot)
		os.Exit(2)
	}

	switch os.Args[1] {
	case commandAdd:
		addFlagSet.Parse(os.Args[2:])
		add.run()
	case commandSearch:
		searchFlagSet.Parse(os.Args[2:])
		search.run()
	case commandRemove:
		removeFlagSet.Parse(os.Args[2:])
		remove.run()
	case commandHelp:
		help(os.Args[2:])
	default:
		fmt.Println(helpRoot)
		os.Exit(2)
	}
}
