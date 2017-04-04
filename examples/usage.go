package main

import "github.com/lukasmalkmus/interactive"

func main() {
	// Create a new session. Use prompt char of your choice.
	s := interactive.New(">")

	// This is where the magic happens. This action will simply echo user input.
	s.Action = func(c *interactive.Context) error {
		text, _ := c.Scan()
		c.Println(text)
		return nil
	}

	// This is executed BEFORE the Action function. Great for printing
	// information, help text, etc.
	s.Before = func(c *interactive.Context) error {
		c.Println("Welcome to Echo!")
		return nil
	}

	// This is executed AFTER the Action function and invoked by context.Close().
	s.After = func(c *interactive.Context) error {
		c.Println("Bye!")
		return nil
	}

	// Finally run the application.
	s.Run()
}
