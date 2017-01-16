package main

import "github.com/LukasMa/interactive"

func main() {
	// Create a new session. Use prompt char of your choice.
	s := interactive.New(">")

	// This is where the magic happens. This action will simply echo user input.
	s.Action = func(c *interactive.Context) {
		text := c.ReadLine()
		c.WriteLine(text)
	}

	// This is executed BEFORE the Action function. Great for printing
	// information, help text, etc.
	s.Before = func(c *interactive.Context) {
		c.WriteLine("Welcome to Echo!")
	}

	// This is executed AFTER the Action function and invoked by context.Close().
	s.After = func(c *interactive.Context) {
		c.WriteLine("Bye!")
	}

	// Finally run the application.
	s.Run()
}
