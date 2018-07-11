package main

import (
	"fmt"

	"path/filepath"

	"github.com/otiai10/gosseract"
	"github.com/spf13/cobra"
)

func main() {
	var cmdExtract = &cobra.Command{
		Use:   "extract [path to image]",
		Short: "Extract recognizable characters from a given image.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			abs, err := filepath.Abs(args[0])
			if err != nil {
				panic(err)
			}

			client := gosseract.NewClient()
			defer client.Close()
			client.SetImage(abs)
			text, err := client.Text()
			if err != nil {
				panic(err)
			}
			fmt.Println(text)
		},
	}

	var rootCmd = &cobra.Command{Use: "ocr"}
	rootCmd.AddCommand(cmdExtract)
	rootCmd.Execute()
}
