package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Site       []string `json:"Site"`
	SearchText string   `json:"SearchText"`
}

type Response struct {
	FoundAtSite string `json:"FoundAtSite"`
}

func FoundAtSite(r *Request) (string, bool) { // returns website and status

	for i := 0; i < len(r.Site); i++ {
		resp, err := http.Get(r.Site[i]) // request document
		if err != nil {
			continue // skipping the current request, to return appropriate result in case there is a correct website in search list
		}

		defer resp.Body.Close()
		html, err := ioutil.ReadAll(resp.Body) // cast body to string
		if err != nil {
			continue // skipping the current request, to return appropriate result in case there is a correct website in search list
		}
		var siteContent = string(html)

		if strings.Contains(siteContent, r.SearchText) == true {
			return r.Site[i], true
		}
	}
	return "", false
}

func checkText(c *gin.Context) { // function for POST request

	var request Request // json unmarshal
	err := c.BindJSON(&request)
	if err != nil {
		panic(err.Error())
	}

	site, found := FoundAtSite(&request)
	if found == true {

		res := Response{
			FoundAtSite: site,
		}

		c.JSON(200, gin.H{
			"FoundAtSite": res.FoundAtSite,
		})
	} else {
		c.JSON(204, gin.H{})
	}

}

func main() {
	r := gin.Default()
	r.POST("/checkText", checkText)
	r.Run(":8080")
}
