package handlers

// Login initiates the login process. A link will be emailed to the user to complete the login.
func Login() {
	// look up user by email.
	// if found, generate a one time code and email it to the user.
	// else, return not found.
}

// Authenticate completes the login process. The user will be authenticated and a token will be returned.
func Authenticate() {
	// get one time code from request.
	// look up user by one time code.
	// return JWT.
}
