package forms

// Fields
const (
	Generic         = "generic"
	Title           = "title"
	Content         = "content"
	Username        = "username"
	Email           = "email"
	Password        = "password"
	CsrfToken       = "csrf_token"
	CurrentPassword = "current_password"
	NewPassword     = "new_password"
	ConfirmPassword = "confirm_password"
	SnipId          = "id"
)

// Errors
const (
	ErrBlankField             = "This field cannot be blank"
	ErrInvalidField           = "This field is invalid"
	ErrInvalidEmailOrPassword = "Email or password is incorrect"
	ErrExistingUsername       = "Username is already taken"
	ErrExistingEmail          = "Email already in use"
	ErrInvalidEmail           = "Invalid email"
	ErrIncorrectPassword      = "Incorrect password"
	ErrMismatchedPasswords    = "Confirmation does not match"
	ErrWhitespace             = "This field must not contain spaces"
)
