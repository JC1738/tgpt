package helper

import (
	"github.com/charmbracelet/glamour"
)

var markdownRenderer *glamour.TermRenderer

// InitMarkdownRenderer initializes the Glamour renderer
func InitMarkdownRenderer() error {
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),  // Auto-detect terminal theme
		glamour.WithWordWrap(-1), // Use terminal width
	)
	if err != nil {
		return err
	}
	markdownRenderer = r
	return nil
}

// RenderMarkdown renders markdown content using Glamour
func RenderMarkdown(content string) (string, error) {
	if markdownRenderer == nil {
		if err := InitMarkdownRenderer(); err != nil {
			return content, err // Fallback to raw content
		}
	}

	rendered, err := markdownRenderer.Render(content)
	if err != nil {
		return content, err // Fallback to raw content
	}

	return rendered, nil
}
