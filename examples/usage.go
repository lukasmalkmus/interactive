package main

import "github.com/LukasMa/interactive"

func main() {
	// Create a new session. Use prompt char of your choice.
	s := interactive.New(">")

	// This is where the magic happens. This action will simply echo user input.
	s.Action = func(c *interactive.Context) error {
		text := c.ReadLine()
		c.WriteLine(text)
		return nil
	}

	// This is executed BEFORE the Action function. Great for printing
	// information, help text, etc.
	s.Before = func(c *interactive.Context) error {
		c.WriteLine("Welcome to Echo!")
		return nil
	}

	// This is executed AFTER the Action function and invoked by context.Close().
	s.After = func(c *interactive.Context) error {
		c.WriteLine("Bye!")
		return nil
	}

	// Finally run the application.
	s.Run()
}
