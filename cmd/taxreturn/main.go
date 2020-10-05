package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	recordsFile string // path to the file with the bills records

	rootCmd = &cobra.Command{
		Use:   "taxreturn",
		Short: "taxreturn - tax and expenses calculator.",
		Long:  "",
		RunE:  taxreturn,
	}
)

func init() {
	rootCmd.Flags().StringVarP(&recordsFile, "records", "r", "", "Path to the bills records file (required)")
	rootCmd.MarkFlagRequired("records")
}

func taxreturn(ccmd *cobra.Command, args []string) error {
	recordFile, err := os.Open(recordsFile)
	if err != nil {
		return err
	}

	reader := csv.NewReader(recordFile)
	records, err := reader.ReadAll()

	fmt.Println(records)

	return nil
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Can't execute command:", err)
		os.Exit(1)
	}
}
