package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"github.com/tdewolff/minify/v2/json"
	"github.com/tdewolff/minify/v2/svg"
	"github.com/tdewolff/minify/v2/xml"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: program <source_directory> <target_directory>")
		os.Exit(1)
	}

	sourceDir := os.Args[1]

	targetDir := os.Args[2]

	sourceAbs, err := filepath.Abs(sourceDir)
	if err != nil {
		fmt.Printf("Error getting absolute path for source directory: %v\n", err)
		os.Exit(1)
	}

	targetAbs, err := filepath.Abs(targetDir)
	if err != nil {
		fmt.Printf("Error getting absolute path for target directory: %v\n", err)

		os.Exit(1)
	}

	err = Process(sourceAbs, targetAbs)
	if err != nil {
		fmt.Printf("Error processing files: %v\n", err)
		os.Exit(1)

	}

	fmt.Println("File processing completed successfully.")
}

func Process(sourceDir, targetDir string) error {
	if err := os.MkdirAll(targetDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating target directory: %w", err)
	}

	if err := processFiles(sourceDir, targetDir); err != nil {
		return err
	}

	return processDirectories(sourceDir, targetDir)
}

func processFiles(sourceDir, targetDir string) error {
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("error reading source directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			sourcePath := filepath.Join(sourceDir, entry.Name())
			targetPath := filepath.Join(targetDir, entry.Name())

			if err := processFile(sourcePath, targetPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func processFile(sourcePath, targetPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer sourceFile.Close()

	// Get the file extension
	ext := strings.TrimPrefix(filepath.Ext(sourcePath), ".")

	// Process the content using the  minifier
	processedReader, err := minifier(sourceFile, ext)
	if err != nil {
		return fmt.Errorf("error processing file content: %w", err)
	}

	// Write the processed content to the target file
	targetFile, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("error creating target file: %w", err)
	}
	defer targetFile.Close()

	_, err = io.Copy(targetFile, processedReader)
	if err != nil {
		return fmt.Errorf("error writing to target file: %w", err)
	}

	return nil
}

func processDirectories(sourceDir, targetDir string) error {
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("error reading source directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			sourcePath := filepath.Join(sourceDir, entry.Name())
			targetPath := filepath.Join(targetDir, entry.Name())

			if err := os.MkdirAll(targetPath, os.ModePerm); err != nil {
				return fmt.Errorf("error creating target subdirectory: %w", err)
			}

			if err := Process(sourcePath, targetPath); err != nil {
				return err
			}
		}
	}

	return nil
}

func minifier(input io.Reader, fileExt string) (io.Reader, error) {
	content, err := io.ReadAll(input)
	if err != nil {
		return nil, fmt.Errorf("error reading input: %w", err)
	}

	m := minify.New()
	var b bytes.Buffer

	switch strings.ToLower(fileExt) {
	case "html":
		m.AddFunc("text/html", html.Minify)
		if err := m.Minify("text/html", &b, bytes.NewReader(content)); err != nil {
			return nil, err
		}
	case "css":
		m.AddFunc("text/css", css.Minify)
		if err := m.Minify("text/css", &b, bytes.NewReader(content)); err != nil {
			return nil, err
		}
	case "js":
		m.AddFunc("text/javascript", js.Minify)
		if err := m.Minify("text/javascript", &b, bytes.NewReader(content)); err != nil {
			return nil, err
		}
	case "svg":
		m.AddFunc("image/svg+xml", svg.Minify)
		if err := m.Minify("image/svg+xml", &b, bytes.NewReader(content)); err != nil {
			return nil, err
		}
	case "json":
		m.AddFunc("application/json", json.Minify)
		if err := m.Minify("application/json", &b, bytes.NewReader(content)); err != nil {
			return nil, err
		}
	case "xml":
		m.AddFunc("text/xml", xml.Minify)
		if err := m.Minify("text/xml", &b, bytes.NewReader(content)); err != nil {
			return nil, err
		}
	default:
		return bytes.NewReader(content), nil
	}

	return bytes.NewReader(b.Bytes()), nil
}
