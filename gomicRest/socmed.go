package gomicRest

import (
	"fmt"
	"log"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	restful "github.com/emicklei/go-restful"
	"github.com/ingmardrewing/gomicRest/config"
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

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true)}

	user, _, _ := client.Accounts.VerifyCredentials(verifyParams)
	fmt.Printf("User's Account:\n%+v\n", user)

	/*
		id := request.PathParameter("page-id")
		page := getPage(id)
		response.WriteEntity(page)
	*/
}
