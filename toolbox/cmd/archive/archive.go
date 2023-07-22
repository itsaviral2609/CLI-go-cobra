/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package archive

import (
	"io"
	"archive/zip"
	"os"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
)

func archiveFile(sourceFilePath, destinationFilePath string) error {
	// Create a new ZIP archive
	zipFile, err := os.Create(destinationFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Create a new ZIP writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk through the source directory and add each file to the ZIP archive
	err = filepath.Walk(sourceFilePath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			relPath, err := filepath.Rel(sourceFilePath, filePath)
			if err != nil {
				return err
			}

			zipEntry, err := zipWriter.Create(relPath)
			if err != nil {
				return err
			}

			fileToArchive, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer fileToArchive.Close()

			_, err = io.Copy(zipEntry, fileToArchive)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	fmt.Println("Files archived successfully!")
	return nil
}

// archiveCmd represents the archive command
var ArchiveCmd = &cobra.Command{
	Use:   "archive",
	Short: "Archive lets you create an archive of any file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sourceFilePath := "/home/hp/toolbox"
		destinationFilePath := "/home/hp/toolbox/archive.zip"

		err := archiveFile(sourceFilePath, destinationFilePath)
		if err != nil {
			fmt.Println("Error archiving the files:", err)
		} else {
			fmt.Println("Files archived successfully!")
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// archiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// archiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
