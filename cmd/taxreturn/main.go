package main

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "taxreturn",
	Short: "taxreturn - tax and expenses calculator.",
	Long:  "",
}

func main() {
	rootCmd.Execute()
}
