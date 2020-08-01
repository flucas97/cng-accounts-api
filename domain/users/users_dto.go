package users

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"first_name"`
	LastName string `json:"last_name"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

/*

CREATE TABLE users (
	ID INT PRIMARY KEY NOT NULL,
	NAME TEXT NOT NULL,
	EMAIL TEXT NOT NULL
	PASSWORD TEXT NOT NULL,
	ACCOUNT_REPOSITORY_ID INT,
	ROLE TEXT NOT NULL,
	ADDRESS TEXT NOT NULL
);

*/
