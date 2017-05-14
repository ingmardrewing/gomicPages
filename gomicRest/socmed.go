package gomicRest

import (
	"fmt"
	"log"
	"unicode/utf8"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicRest/config"
	"github.com/ingmardrewing/gomicRest/db"
)

func NewSocMedService() *restful.WebService {
	path := "/0.1/gomic/socmed"
	service := new(restful.WebService)
	service.
		Path(path).
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	log.Printf("Creating socmed service at localhost:8080 -- access with http://localhost:8080%s\n", path)

	service.Route(service.POST("/twitter").Filter(basicAuthenticate).To(Tweet))
	return service
}

func Tweet(request *restful.Request, response *restful.Response) {
	cred := oauth1.NewConfig(
		config.GetTwitterConsumerKey(),
		config.GetTwitterConsumerSecret())

	token := oauth1.NewToken(
		config.GetTwitterAccessToken(),
		config.GetTwitterAccessTokenSecret())

	httpClient := cred.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)

	/*
		verifyParams := &twitter.AccountVerifyParams{
			SkipStatus:   twitter.Bool(true),
			IncludeEmail: twitter.Bool(true),
		}
		user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
		fmt.Printf("User's ACCOUNT:\n%+v\n", user)

		fmt.Println(getTweetText())
	*/

	// actually tweet
	tweet, _, _ := client.Statuses.Update(getTweetText(), nil)
	fmt.Printf("Posted tweet \n%v\n", tweet)
}

func getTweetText() string {
	pages := db.GetAllPages()
	lastPage := pages[len(pages)-1]

	url := "https://devabo.de" + lastPage.Path
	tweet := lastPage.Title + " " + url

	for _, tag := range config.GetTags() {
		if utf8.RuneCountInString(tweet+" "+tag) > 140 {
			return tweet
		}
		tweet += " " + tag
	}

	return tweet
}

/*
func PostToTumblr() {
	fmt.Println("Post to tumblr")
	cons_key, cons_secret, token, token_secret := config.GetTumblData()
	client := gotumblr.NewTumblrRestClient(cons_key, cons_secret, token, token_secret, "http://localhost/~drewing/cgi-bin/tumblr.pl", "http://api.tumblr.com")

	blogname := "devabo-de.tumblr.com"
	state := "published"
	tags := "comic,webcomic,graphicnovel,drawing,art,narrative,scifi,sci-fi,science-fiction,dystopy,parody,humor,nerd,pulp,geek,blackandwhite"
	photoPostByURL := client.CreatePhoto(
		blogname,
		map[string]string{
			"link":    prodUrl,
			"source":  imgurl,
			"caption": title,
			"tags":    tags,
			"state":   state})
	if photoPostByURL == nil {
		fmt.Println("done")
	} else {
		fmt.Println(photoPostByURL)
	}
}
*/
