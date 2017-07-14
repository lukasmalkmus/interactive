package interactive

import "fmt"

// A Context is bound to a session and provides methods to interact with it. It
// hasn't anything in common with the official context package/type.
type Context struct {
	*Session
}

// Close will close the session with the specified exit code and restore the
//terminal into the previous state.
func (c *Context) Close(exitCode int) {
	c.close(exitCode)
}

// Print formats using the default formats for its operands and writes to the
// sessions output. Spaces are added between operands when neither is a string.
// See fmt.Sprint and fmt.Print for more information.
func (c *Context) Print(a ...interface{}) {
	c.write(fmt.Sprint(a...))
}

// Printf formats according to a format specifier and writes to the sessions
// output. See fmt.Sprintf and fmt.Printf for more information.
func (c *Context) Printf(format string, a ...interface{}) {
	c.write(fmt.Sprintf(format, a...))
}

// Println formats using the default formats for its operands and writes to the
// sessions output. Spaces are always added between operands and a newline is
// appended. See fmt.Sprintln and fmt.Println for more information.
func (c *Context) Println(a ...interface{}) {
	c.write(fmt.Sprintln(a...))
}

// ScanHidden reads without echo.
func (c *Context) ScanHidden() (string, error) {
	return c.term.ReadPassword(c.prompt)
}

// Scan reads user input from the session. It does NOT behave like fmt.Scan().
func (c *Context) Scan() (string, error) {
	return c.read()
}
