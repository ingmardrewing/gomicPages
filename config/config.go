package config

import (
	"fmt"
	"os"
)

func GetDsn() string {
	user := os.Getenv("DB_GOMIC_USER")
	pass := os.Getenv("DB_GOMIC_PASS")
	name := os.Getenv("DB_GOMIC_NAME")
	host := os.Getenv("DB_GOMIC_HOST")
	return fmt.Sprintf("%s:%s@%s/%s", user, pass, host, name)
}

func GetPasswordHashForUser(user string) string {
	// TODO impl. user dependent hash retrieval
	return os.Getenv("GOMIC_BASIC_AUTH_PASS_HASH")
}

func GetTwitterConsumerKey() string {
	s := os.Getenv("TWITTER_ORIGIN_CONSUMER_KEY")
	fmt.Println(s)
	return s
}

func GetTlsPaths() (string, string) {
	cert := os.Getenv("TLS_CERT_PATH")
	key := os.Getenv("TLS_KEY_PATH")
	return cert, key
}

func GetTwitterConsumerSecret() string {
	s := os.Getenv("TWITTER_ORIGIN_CONSUMER_SECRET")
	fmt.Println(s)
	return s
}

func GetTwitterAccessToken() string {
	s := os.Getenv("TWITTER_ORIGIN_ACCESS_TOKEN")
	fmt.Println(s)
	return s
}

func GetTwitterAccessTokenSecret() string {
	s := os.Getenv("TWITTER_ORIGIN_ACCESS_TOKEN_SECRET")
	fmt.Println(s)
	return s
}

func GetTumblrConsumerKey() string {
	return os.Getenv("GOMIC_TUMBLR_CONSUMER_KEY")
}

func GetTumblrConsumerSecret() string {
	return os.Getenv("GOMIC_TUMBLR_CONSUMER_SECRET")
}

func GetTumblrToken() string {
	return os.Getenv("GOMIC_TUMBLR_TOKEN")
}

func GetTumblrTokenSecret() string {
	return os.Getenv("GOMIC_TUMBLR_TOKEN_SECRET")
}

func GetTags() []string {
	return []string{"#inked", "#inking", "#illustration", "#drawing", "#drawings", "#art", "#artwork", "#draw", "#painting", "#sketch", "#sketchbook", "#artist", "#comics", "#comicart", "#comic", "#graphicnovel", "#design", "#concept", "#conceptart", "#create", "#creative", "#image", "#imagination"}
}
