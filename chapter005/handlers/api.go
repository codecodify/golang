package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func ApiHandler(context *gin.Context){
	query := context.Query("s")
	log.Println(query)
	var imageItem []map[string]interface{}
	if len(query) > 0 {
		url := fmt.Sprintf("https://m.qqtn.com/%s", query)
		response, err := getResponse(url)
		if err != nil {
			log.Panicln(err.Error())
		}

		defer func() {
			_ = response.Body.Close()
		}()

		bodyByte, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Panicln(err.Error())
		}

		bodyString := string(bodyByte)

		imageRegexp, err := regexp.Compile(`<img.*?src=["|']?(.*?)["|']?\s.*?>`)
		if err != nil {
			log.Panicln(err.Error())
		}

		images := imageRegexp.FindAllStringSubmatch(bodyString, -1)
		for _, value := range images {
			imageItem = append(imageItem, map[string]interface{}{
				"image": value[1],
			})
		}

		//log.Println(imageItem)
	}
	context.HTML(http.StatusOK, "api.html", gin.H{
		"query": query,
		"images": imageItem,
	})
}