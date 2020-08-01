package accounts

type Account struct {
	ID                   int64    `json:"id"`
	Name                 string   `json:"name"`
	CannabisRepositoryID int64    `belongs_to:"cannabis_repository_id" json:"cannabis_repository_id"`
	Status               string   `json:"status"`
	AvaliableFeatures    []string `json:"avaliable_features"`
	Language             string   `json:"language"`
	UsersRepositoryID    int64    `belongs_to:"users_repository_id" json:"user_repository_id"`
}

/*

CREATE TABLE accounts (
	ID INT PRIMARY KEY NOT NULL,
	NAME TEXT NOT NULL,
	CANNABIS_REPOSITORY_ID INT,
	STATUS TEXT NOT NULL,
	AVALIABLE_FEATURES TEXT,
	LANGUAGE TEXT NOT NULL,
	USERS_REPOSITORY_ID INT
);

*/
