package interactive

import (
	"io"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// A Session is an interactive shell. The New() function should be used to
// obtain a new Session instance.
type Session struct {
	// Action is the actual application logic which is looped until the
	// application gets terminated.
	Action Action

	// After is run AFTER the action function, BEFORE the session is closed.
	After Action

	// Before is run BEFORE the action function.
	Before Action

	context *Context
	fd      int
	prompt  string
	shell   *shell
	state   *terminal.State
	term    *terminal.Terminal
}

// New spawns an interactive session in the current terminal. A prompt character
// needs to be provided which will be printed when user input is awaited.
func New(prompt string) *Session {
	// Save old state and set terminal into raw mode.
	fd := int(os.Stdin.Fd())
	oldState, err := terminal.MakeRaw(fd)
	if err != nil {
		panic(err)
	}

	// Satisfies the ReadWriter interface and serves as I/O for the new terminal.
	shell := &shell{
		Reader: os.Stdin,
		Writer: os.Stdout,
	}

	// Create new terminal with desired prompt sign.
	term := terminal.NewTerminal(shell, strings.TrimSpace(prompt)+" ")

	// Finally create the session.
	s := &Session{
		fd:     fd,
		prompt: prompt,
		shell:  shell,
		state:  oldState,
		term:   term,
	}
	s.context = &Context{Session: s}

	// Setup Ctrl^C listener.
	s.term.AutoCompleteCallback = callback(s)

	return s
}

// Run is a blocking method that executes the actual logic.
func (s *Session) Run() {
	// Abort if no ro0t action is defined.
	if s.Action == nil {
		return
	}

	// Run Before function if present. Print error if present.
	if s.Before != nil {
		if err := s.Before(s.context); err != nil {
			s.write(err.Error() + "\n")
		}
	}

	// Loop root action. Print error if present.
	for {
		if err := s.Action(s.context); err != nil {
			s.write(err.Error() + "\n")
		}
	}
}

func (s *Session) close(exitCode int) {
	// Run After function if present.
	if s.After != nil {
		if err := s.After(s.context); err != nil {
			s.write(err.Error() + "\n")
		}
	}

	// Restore terminal.
	terminal.Restore(s.fd, s.state)
	os.Exit(exitCode)
}

func (s *Session) read() (string, error) {
	text, err := s.term.ReadLine()
	if err != nil {
		// Close session on Ctrl^D.
		if err == io.EOF {
			s.close(0)
		} else {
			return "", err
		}
	}
	return text, nil
}

func (s *Session) write(text string) {
	s.term.Write([]byte(text))
}

func callback(s *Session) func(string, int, rune) (string, int, bool) {
	return func(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
		if key == '\x03' {
			s.close(0)
		}
		return "", 0, false
	}
}
