package errs

var (
	ErrInvalidCredentials            = new(401, "invalid credentials")
	ErrUserNotFound                  = new(404, "user not found")
	ErrUserWithEmailAlreadyExists    = new(409, "a user with that email already exists.  Please login or use a different email address")
	ErrUserWithGithubIDAlreadyExists = new(409, "a user connected to that github account already exists.  Please login or use a different account")
)
