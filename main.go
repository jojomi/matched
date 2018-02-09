package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	defaultCharSet = `()[]{}<>""''`

	charSet   string
	errLogger *log.Logger
)

func main() {
	errLogger = log.New(os.Stderr, "", 0)

	var rootCmd = &cobra.Command{
		Use: "matched test-string",
		Run: cmdRoot,
	}
	rootCmd.PersistentFlags().StringVarP(&charSet, "char-set", "c", defaultCharSet, "pairs of chars to be checked")

	if err := rootCmd.Execute(); err != nil {
		errLogger.Println(err)
		os.Exit(1)
	}
}

func cmdRoot(cmd *cobra.Command, args []string) {
	matches := splitCharSet(charSet)
	if len(args) < 1 {
		errLogger.Println("No input given.")
		os.Exit(2)
	}
	input := args[0]

	var (
		err   error
		valid bool
	)
	for _, match := range matches {
		err = match.ValidateForString(input)
		if err != nil {
			errLogger.Println(err)
			valid = false
		}
	}

	if !valid {
		os.Exit(3)
	}
}
