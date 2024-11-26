/*
Copyright © 2024 smcelhose.aws@gmail.com
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/yourbasic/bit"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "eratosthenes max",
	Aliases: []string{"prime", "primes"},
	Args:    cobra.ExactArgs(1),
	Short:   "Calculate primes upto N using sieve of Eratosthenes",
	Long: `Calculate primes upto N using sieve of Eratosthenes.
Toolifying Sieve of Erathosthenes found here for self.  https://yourbasic.org/golang/bitmask-flag-set-clear/

	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "")
			os.Exit(1)
		}

		s := SieveOfEratoshthenes(n)
		primes := make([]int, 0)
		s.Visit(func(n int) bool {
			primes = append(primes, n)
			return false
		})

		bytes, err := json.MarshalIndent(primes, "", "  ")

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		fmt.Fprintln(os.Stdout, string(bytes))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eratosthenes.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func SieveOfEratoshthenes(n int) *bit.Set {
	sieve := bit.New().AddRange(2, n)
	sqrtN := int(math.Sqrt(float64(n)))
	for p := 2; p <= sqrtN; p = sieve.Next(p) {
		for k := p * p; k < n; k += p {
			sieve.Delete(k)
		}
	}
	return sieve
}
