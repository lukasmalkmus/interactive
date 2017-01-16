/*
Package interactive provides an easy to implement shell for simple, interactive
commandline applications. It is build on top of the excellent
https://golang.org/x/crypto/ssh/terminal package and tries to simplify the creation of
small and simple applications which run in shell mode. It isn't very powerful
(yet) but has enough features for basic usage. For example it shuts down
gracefully on Ctrl^C and Ctrl^D.

This small application echos the entered text:
    s := interactive.New(">")
    s.Action = func(c *interactive.Context) error {
        text := c.ReadLine()
        c.WriteLine(text)
        return nil
    }
    s.Run()
Returning an error instead of nil will print the error and close the session
with a return code of 1.
*/
package interactive
