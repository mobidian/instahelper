package config

import (
	"time"
)

// InstahelperConfig is information about the Instahelper package as a whole
type InstahelperConfig struct {
	// AESKey used to encrypt password and Account.CachedInsta
	// Don't change
	AESKey []byte

	ID int `storm:"id"`

	// Port to run the application on - defaults to :3333
	Port int

	// Username for Instahelper
	Username string

	// Password for Instahelper
	Password string

	// Analytics enabled?
	Analytics bool
}

// Account is an Instagram Account
type Account struct {
	ID int `storm:"id,increment" json:"id"`

	Username string `storm:"unique" json:"username"`
	Password string `json:"-"`

	// INFO
	FullName string `json:"full_name"`
	Bio      string `json:"bio"`
	// Profile pic url
	ProfilePic string `json:"profile_pic"`
	Private    bool   `json:"private"`

	// Cached GoInsta object
	CachedInsta []byte `json:"-"`

	// Settings is not inline to be able to copy over settings between accounts
	Settings *Settings `json:"settings"`

	// AddedAt is when the user added this account
	AddedAt time.Time `storm:"index" json:"added_at"`
}

// Settings for a given account
type Settings struct {
	FollowsPerDay   int `json:"follows_per_day,omitempty"`
	CommentsPerDay  int `json:"comments_per_day,omitempty"`
	LikesPerDay     int `json:"likes_per_day,omitempty"`
	UnfollowsPerDay int `json:"unfollows_per_day,omitempty"`

	// Proxy to make requests with
	Proxy string `json:"proxy,omitempty"`

	// UnfollowAt is the number of follows when the bot should start unfollowing
	UnfollowAt int `json:"unfollow_at,omitempty"`
	// UnfollowNonFollowers will decide if we unfollow those who do not follow after one day
	UnfollowNonFollowers bool `json:"unfollow_non_followers,omitempty"`

	// Tags to follow, comment, or like
	Tags []string `json:"tags,omitempty"`
	// CommentList is the list of comments to choose from when commenting
	CommentList []string `json:"comment_list,omitempty"`

	// Blacklist is a list of accounts to avoid following, commenting, or liking
	Blacklist []string `json:"blacklist,omitempty"`
	// Whitelist is the list of users to only follow, comment, and like on
	Whitelist []string `json:"whitelist,omitempty"`

	// FollowPrivate will decide if we follow private accounts
	FollowPrivate bool `json:"follow_private,omitempty"`
}

// A Notification that would be shown in the top navbar.
type Notification struct {
	ID   int    `storm:"id,increment" json:"id"`
	Link string `json:"link"`
	Text string `json:"text,omitempty"`
}

// USE SAVE OVER UPDATE
// https://github.com/asdine/storm/issues/160

// Update the model using the values provided - helper function
func (m *InstahelperConfig) Update() error {
	return DB.Save(m)
}

// Update the model using the values provided - helper function
func (m *Account) Update() error {
	return DB.Save(m)
}

// Models for boltdb
var Models = []interface{}{
	&Account{}, &InstahelperConfig{}, &Notification{},
}

// Migrate will reindex all fields
func Migrate() error {

	for _, m := range Models {

		if err := DB.Init(m); err != nil {
			return err
		}

		if err := DB.ReIndex(m); err != nil {
			return err
		}
	}
	return nil
}
