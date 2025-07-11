package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rizome-dev/opun/pkg/plugin"
)

type Config struct {
	MaxComplexity int  `json:"max_complexity"`
	IgnoreTests   bool `json:"ignore_tests"`
}

type LineCount struct {
	TotalLines   int `json:"total_lines"`
	CodeLines    int `json:"code_lines"`
	CommentLines int `json:"comment_lines"`
	BlankLines   int `json:"blank_lines"`
}

func main() {
	p := plugin.New("code-analyzer", "1.0.0")

	// Register command handler
	p.RegisterCommand("analyze", func(args map[string]interface{}) (interface{}, error) {
		file, ok := args["file"].(string)
		if !ok {
			return nil, fmt.Errorf("file argument required")
		}

		result, err := analyzeFile(file)
		if err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"message": fmt.Sprintf("Analysis complete for %s", file),
			"result":  result,
		}, nil
	})

	// Register tool handler
	p.RegisterTool("count-lines", func(input map[string]interface{}) (interface{}, error) {
		filePath, ok := input["file_path"].(string)
		if !ok {
			return nil, fmt.Errorf("file_path required")
		}

		return countLines(filePath)
	})

	// Run the plugin
	if err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Plugin error: %v\n", err)
		os.Exit(1)
	}
}

func analyzeFile(filePath string) (map[string]interface{}, error) {
	counts, err := countLines(filePath)
	if err != nil {
		return nil, err
	}

	// Simple complexity calculation (lines of code / 10)
	complexity := counts.CodeLines / 10
	if complexity < 1 {
		complexity = 1
	}

	return map[string]interface{}{
		"file":       filePath,
		"lines":      counts,
		"complexity": complexity,
		"rating":     getComplexityRating(complexity),
	}, nil
}

func countLines(filePath string) (*LineCount, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	counts := &LineCount{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		counts.TotalLines++

		if line == "" {
			counts.BlankLines++
		} else if strings.HasPrefix(line, "//") || strings.HasPrefix(line, "/*") || strings.HasPrefix(line, "*") {
			counts.CommentLines++
		} else {
			counts.CodeLines++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return counts, nil
}

func getComplexityRating(complexity int) string {
	switch {
	case complexity <= 5:
		return "Simple"
	case complexity <= 10:
		return "Moderate"
	case complexity <= 20:
		return "Complex"
	default:
		return "Very Complex"
	}
}