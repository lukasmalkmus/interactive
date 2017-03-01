package interactive

import "fmt"

// A Context is bound to a session and provides methods to interact with the
// session. Don't be misled by the name of this type! It hasn't anything in
// common with the official context package/type. This type only exists so the
// full session isn't exposed in the functions (see funcs.go).
type Context struct {
	session *Session
}

// Close will stop the session and restore the terminal.
func (c *Context) Close() {
	c.session.close(0)
}

// Print formats using the default formats for its operands and writes to the
// sessions output. Spaces are added between operands when neither is a string.
// See fmt.Sprint and fmt.Print for more information.
func (c *Context) Print(a ...interface{}) {
	c.session.write(fmt.Sprint(a...))
}

// Printf formats according to a format specifier and writes to the sessions
// output. See fmt.Sprintf and fmt.Printf for more information.
func (c *Context) Printf(format string, a ...interface{}) {
	c.session.write(fmt.Sprintf(format, a...))
}

// Println formats using the default formats for its operands and writes to the
// sessions output. Spaces are always added between operands and a newline is
// appended. See fmt.Sprintln and fmt.Println for more information.
func (c *Context) Println(a ...interface{}) {
	c.session.write(fmt.Sprintln(a...))
}

// Scan reads user input from the session. It does NOT behave like fmt.Scan().
func (c *Context) Scan() (string, error) {
	return c.session.read()
}
