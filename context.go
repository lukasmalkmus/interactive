package interactive

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

// ReadLine reads a line of text from the terminal.
func (c *Context) ReadLine() string {
	return c.session.readLine()
}

// WriteLine writes a line of text to the terminal.
func (c *Context) WriteLine(text string) {
	c.session.writeLine(text)
}
