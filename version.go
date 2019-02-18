package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Provide version information",
	Long:  `Provide version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(logo)
<<<<<<< HEAD
		fmt.Println("Docgen version: v2.1")
=======
		fmt.Println("Docgen version: v2.1 - testing edit")
>>>>>>> 258ed8768a14ed29f6b6d6617125d5691d4d74ca
		fmt.Println("Support postman collection version > 2.1")
	},
}
