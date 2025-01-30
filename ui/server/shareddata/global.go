package shareddata

const SessionName = "session_token"

var Errors struct {
	InvalidCredentials   string
	InvalidEmail         string
	LongEmail            string
	InvalidUsername      string
	InvalidPassword      string
	UserAlreadyExist     string
	EmailAlreadyExist    string
	UsernameAlreadyExist string
	// other
	ErrorHashingPass string
} = struct {
	InvalidCredentials   string
	InvalidEmail         string
	LongEmail            string
	InvalidUsername      string
	InvalidPassword      string
	UserAlreadyExist     string
	EmailAlreadyExist    string
	UsernameAlreadyExist string
	// other
	ErrorHashingPass string
}{
	InvalidCredentials:   "Invalid Credentials",
	InvalidEmail:         "Invalid Email ex: exmaple@mail.com",
	LongEmail:            "Email must be between 5 and 50 characters.",
	InvalidUsername:      "Username must be between 3 and 15 characters.",
	InvalidPassword:      "Password must be between 6 and 30 characters.",
	UserAlreadyExist:     "Email or Username Already Exist",
	EmailAlreadyExist:    "Email Already Exist",
	UsernameAlreadyExist: "Username Already Exist",
	// other
	ErrorHashingPass: "Error Hashing Password",
}

var UserErrors struct {
	InvalidEmail     string
	InvalidUsername  string
	InvalidPassword  string
	UserAlreadyExist string
	UserNotExist     string
} = struct {
	InvalidEmail     string
	InvalidUsername  string
	InvalidPassword  string
	UserAlreadyExist string
	UserNotExist     string
}{
	InvalidEmail:     "invalid email",
	InvalidUsername:  "invalid username",
	InvalidPassword:  "invalid password",
	UserAlreadyExist: "user already exist",
	UserNotExist:     "user doesn't exist",
}

var PostErrors struct {
	PostNotExist string
	ContentLength string
	TitleLength string
	CategoryDoesntExist string
} = struct{ PostNotExist string; ContentLength string; TitleLength string; CategoryDoesntExist string }{
	PostNotExist: "post doesn't exist",
	ContentLength: "invalid content length",
	TitleLength: "invalid title length",
	CategoryDoesntExist: "categories doesn't exist",
}

var CommentErrors struct {
	InvalidCommentLength string
	InvalidPage          string
} = struct {
	InvalidCommentLength string
	InvalidPage          string
}{
	InvalidCommentLength: "invalid comment length",
	InvalidPage:          "invalid page number",
}
