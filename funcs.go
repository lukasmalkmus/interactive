package interactive

// BeforeFunc is the function to execute BEFORE the Action function is looped.
type BeforeFunc func(*Context)

// ActionFunc is the function that should contain the logic. In most cases this
// is reading input, parsing input, doing stuff and printing output.
type ActionFunc func(*Context)

// AfterFunc is the function to execute BEFORE the session is closed. It is
// invoked by context.Close().
type AfterFunc func(*Context)
