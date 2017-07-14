package interactive

// An Action is implemented by the package user and used by the session.
type Action func(*Context) error
