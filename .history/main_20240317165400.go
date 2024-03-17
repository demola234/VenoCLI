package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render

func main() {
	var (
		daemonMode bool
		showHelp   bool
		opts       []tea.ProgramOption
	)

	flag.BoolVar(&daemonMode, "d", false, "run as a daemon")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.Parse()

	if showHelp {
		flag.Usage()
		return
	}

	if daemonMode || !isatty.IsTerminal(os.Stdout.Fd()) {
		opts = []tea.ProgramOption{tea.WithoutRenderer()}
	} else {
		log.SetOutput(io.Discard)
	}

	p := tea.NewProgram(model, opts...)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting Bubble Tea program:", err)
		os.Exit(1)
	}
}

type result struct {
	duration time.Duration
	emoji    string
}

type model struct {
	spinner  spinner.Model
	results  []result
	quitting bool
}

func (m model) Init() tea.Cmd {
	log.Println("Starting work...")
	return tea.Batch(
		m.spinner.Tick,
		runPretendProcess,
	)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m.quitting = true
		return m, tea.Quit

		

}
