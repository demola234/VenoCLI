package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-isatty"
	"github.com/muesli/reflow/indent"
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

	p := tea.NewProgram(newProgram(), opts...)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}

	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting Bubble Tea program:", err)
		os.Exit(1)
	}
}

func newModel() model {
	const showLastResults = 5

	sp := spinner.New()
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("206"))

	return model{
		spinner: sp,
		results: make([]result, showLastResults),
	}
}

func newModel() model {
	const showLastResults = 5

	sp := spinner.New()
	sp.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("206"))

	return model{
		spinner: sp,
		results: make([]result, showLastResults),
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

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case processFinishedMsg:
		d := time.Duration(msg)
		res := result{emoji: randomEmoji(), duration: d}
		log.Printf("%s Job finished in %s", res.emoji, res.duration)
		m.results = append(m.results[1:], res)
		return m, runPretendProcess
	default:
		return m, nil
	}
}

func (m model) View() string {
	s := "\n" +
		m.spinner.View() + " Doing some work...\n\n"

	for _, res := range m.results {
		if res.duration == 0 {
			s += "........................\n"

		} else {
			s += fmt.Sprintf("%s Job finished in %s\n", res.emoji, res.duration)
		}
	}
	s += helpStyle("\nPress any key to exit\n")

	if m.quitting {
		s += "\n"
	}

	return indent.String(s, 1)
}

// processFinishedMsg is sent when a pretend process completes.
type processFinishedMsg time.Duration

// pretendProcess simulates a long-running process.
func runPretendProcess() tea.Msg {
	pause := time.Duration(rand.Int63n(899)+100) * time.Millisecond // nolint:gosec
	time.Sleep(pause)
	return processFinishedMsg(pause)
}

func randomEmoji() string {
	emojis := []rune("🍦🧋🍡🤠👾😭🦊🐯🦆🥨🎏🍔🍒🍥🎮📦🦁🐶🐸🍕🥐🧲🚒🥇🏆🌽")
	return string(emojis[rand.Intn(len(emojis))]) // nolint:gosec
}
