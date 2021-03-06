package main

import (
	"fmt"
	"os"

	taxreturn "github.com/antklim/tax-return"
	"github.com/spf13/cobra"
)

var (
	recordsFile string // path to the file with the bills records
	noHeader    bool
	year        int

	rootCmd = &cobra.Command{
		Use:   "taxreturn",
		Short: "taxreturn - tax and expenses calculator.",
		Long:  "",
		RunE:  runRoot,
	}
)

func init() {
	rootCmd.Flags().StringVarP(&recordsFile, "records", "r", "", "Path to the bills records file (required)")
	rootCmd.Flags().IntVarP(&year, "year", "y", 0, "Financial year ending in (required)")
	rootCmd.Flags().BoolVarP(&noHeader, "noHeader", "H", false, "Records file does not have header (default: false)")
	rootCmd.MarkFlagRequired("records")
	rootCmd.MarkFlagRequired("year")
}

func runRoot(cmd *cobra.Command, args []string) error {
	file, err := os.Open(recordsFile)
	if err != nil {
		return err
	}

	bills, err := taxreturn.ReadCsv(file, !noHeader)
	if err != nil {
		return err
	}

	financialYear := taxreturn.FinancialYearEnding(year)
	report, err := bills.Report(financialYear)
	if err != nil {
		return err
	}

	fmt.Println(report)

	return nil
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
