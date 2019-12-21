package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)


func IndexHandler(context *gin.Context){
	param := context.DefaultQuery("s", "gexingtx")
	log.Println(param)
	resourceItems := GetResource(param)
	context.HTML(http.StatusOK, "index.html", gin.H{
		"resources" : resourceItems,
	})
}


func GetResource(param string) (resourceItems []map[string]interface{}){
	url := fmt.Sprintf("https://m.qqtn.com/tx/%s", param)
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
		log.Fatalln(err.Error())
	}

	urlRegexp, err := regexp.Compile(`href="([^"]+)`)
	if err != nil {
		log.Fatalln(err.Error())
	}
	matchImages := imageRegexp.FindAllStringSubmatch(bodyString, -1)
	matchUrls := urlRegexp.FindAllStringSubmatch(bodyString, -1)
	//log.Println(imageRegexp.FindAllStringSubmatch(bodyString, -1))
	//log.Println(urlRegexp.FindAllStringSubmatch(bodyString, -1))
	for key, value := range matchImages {
		// 索引28起才是图片的对应的链接
		imageUrlIndex := key + 28
		resourceItems = append(resourceItems, map[string]interface{}{
			"url": matchUrls[imageUrlIndex][1],
			"image": value[1],
		})
	}
	return resourceItems
}