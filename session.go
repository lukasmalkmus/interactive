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
	// Action is the actual application logic that is looped until the
	// application gets terminated.
	Action ActionFunc
	// After is run AFTER the action function, BEFORE the session is closed.
	// It is invoked by context.Close().
	After AfterFunc
	// Before is run BEFORE the action function.
	Before BeforeFunc

	context *Context
	fd      int
	shell   *Shell
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
	shell := &Shell{
		r: os.Stdin,
		w: os.Stdout,
	}

	// Create new terminal with desired prompt sign.
	term := terminal.NewTerminal(shell, strings.Trim(prompt, " ")+" ")

	// Finally create the session.
	s := &Session{
		Action: dummyAction,
		fd:     fd,
		shell:  shell,
		state:  oldState,
		term:   term,
	}
	s.context = &Context{session: s}

	// Set up Ctrl^C listener.
	term.AutoCompleteCallback = func(line string, pos int, key rune) (newLine string, newPos int, ok bool) {
		if key == '\x03' {
			s.close(0)
		}
		return "", 0, false
	}

	return s
}

// Run is a blocking method that executes the actual logic.
func (s *Session) Run() {
	if s.Action == nil {
		panic(`Dear developer. If you see this, you fucked up big time. You
ignored every bit of documentation and manually set the root
Action to NIL. How is this supposed to work?!`)
	}

	// Run Before function if present. Close session if an error occurs.
	if s.Before != nil {
		if err := s.Before(s.context); err != nil {
			s.write(err.Error() + "\n")
			s.close(1)
		}
	}

	// Loop root action. Close session if an error occurs.
	for {
		if err := s.Action(s.context); err != nil {
			s.write(err.Error() + "\n")
			s.close(1)
		}
	}
}

// Shell returns the sessions underlying shell (ReadWriter).
func (s *Session) Shell() *Shell {
	return s.shell
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

func dummyAction(c *Context) error {
	c.Println("No Action defined!")
	c.Close()
	return nil
}
