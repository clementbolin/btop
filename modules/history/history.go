package history

import (
	"fmt"
	"os"
)

// History :
type History struct {
	cmd []string
}

// GetHistoryCmd :
func (h *History) GetHistoryCmd() {
	shell := os.Getenv("SHELL")
	fmt.Println(shell)
}
