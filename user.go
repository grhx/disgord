package discord

import (

)

// client user struct
type ClientUser struct {
	Id       	string
	Username 	string	
	Tag			string	`json:"discriminator"`
	Bot     	bool
	Mfa			bool
	Email		string
	Avatar		string
}

// user struct
type User struct {
	Id               string
	Username         string
	Tag              int
	GlobalName       string // optional
	Avatar           string // optional
	Bot              bool
	System           bool
	MfaEnabled       bool
	Banner           string // optional
	AccentColor      int    // optional
	Locale           string
	Verified         bool
	Email            string // optional
	Flags            []UserFlag
	PremiumType      UserPremium
	PublicFlags      int
	AvatarDecoration string //optional

}

// user manager struct
type userManager struct {
	users map[string]*User
}

func (users userManager) Add(user *User) {
	users.users[user.Id] = user
}
func (users userManager) Get(uid string) *User {
	return users.users[uid]
}
