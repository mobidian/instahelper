// Package insta provides a wrapper over goinsta
package insta

import (
	"github.com/ahmdrz/goinsta"
	"github.com/socialplanner/instahelper/app/config"
)

// CachedInsta logs into Instagram and returns []bytes needed to Import this object
func CachedInsta(username, password, proxy string) ([]byte, error) {

	ig, err := Login(username, password, proxy)

	if err != nil {
		return []byte{}, err
	}

	c, err := config.Config()

	if err != nil {
		return []byte{}, err
	}

	// Export the cached instagram object using the configs AESKey.
	return ig.Export(c.AESKey)
}

// ExportCached will export the cached instagram object using the configs AESKey.
func ExportCached(ig *goinsta.Instagram) ([]byte, error) {
	c, err := config.Config()

	if err != nil {
		return []byte{}, err
	}
	return ig.Export(c.AESKey)
}

// Login will connect to Instagram through a proxy if one is passed
func Login(username, password, proxy string) (*goinsta.Instagram, error) {
	var ig *goinsta.Instagram

	// If proxy passed create a goinsta connection using that proxy
	if proxy == "" {
		ig = goinsta.New(username, password)
	} else {
		ig = goinsta.NewViaProxy(username, password, password)
	}

	err := ig.Login()

	if err != nil {
		return nil, err
	}

	return ig, nil
}

// Import an account from it's cached bytes
func Import(b []byte) (*goinsta.Instagram, error) {
	c, err := config.Config()

	if err != nil {
		return nil, err
	}

	return goinsta.Import(b, c.AESKey)
}
