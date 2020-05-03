package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

const (
	DIR_PATH = "动物农场"
	URL = "https://www.kanunu8.com/book3/6879/"
)


type chapterContent struct {
	title string
	content string
}

var workChapter = make(chan string) // 章节地址
var workContent = make(chan chapterContent) // 章节内容

var wg sync.WaitGroup

func main() {
	// 初始化目录
	_, err := os.Stat(DIR_PATH)
	if err != nil {
		if os.IsNotExist(err){
			log.Println(fmt.Sprintf("创建%s目录", DIR_PATH))
			err = os.Mkdir(DIR_PATH, os.ModePerm)
			if err != nil {
				log.Panicf("创建%s目录失败，系统退出\n", DIR_PATH)
			}
		}else {
			log.Panicf("打开%s目录失败，系统退出\n", DIR_PATH)
		}

	}

	log.Println("开始采集...")


	wg.Add(3)

	// 爬取章节地址
	go getChapterUrl()

	// 获取章节内容
	go getChapterContent()


	// 将内容写入到文件里
	go saveChapterContent()


	wg.Wait()

	log.Println("采集结束")

}

func getChapterUrl(){

	content, err := httpGet(URL)
	if err != nil {
		log.Panicln(err)
	}

	var dom *goquery.Document
	dom, err = goquery.NewDocumentFromReader(strings.NewReader(convertToString(string(content), "gbk", "utf-8")))
	if err != nil {
		log.Panicln(err)
	}

	dom.Find("tbody tr[bgcolor]").Each(func(i int, selection *goquery.Selection) {
		selection.Find("td > a").Each(func(i int, selection *goquery.Selection) {
			link, ok := selection.Attr("href")
			if ok {
				// 处理相对地址
				link = URL + link
				workChapter <- link
			}
		})
	})
	defer func() {
		close(workChapter)
		defer wg.Done()
	}()

}


func getChapterContent(){
	for link := range workChapter{
		// 根据章节获取章节内容
		content, err := httpGet(link)
		if err != nil {
			log.Panicln(err)
		}

		var dom *goquery.Document

		dom, err = goquery.NewDocumentFromReader(strings.NewReader(convertToString(string(content), "gbk", "utf-8")))
		if err != nil {
			log.Panicln(err)
		}


		title := dom.Find("font[size]").Text()

		dom.Find("td > p").Each(func(i int, selection *goquery.Selection) {
			workContent <- chapterContent{
				title: title,
				content: selection.Text(),
			}
		})
	}

	defer func() {
		close(workContent)
		wg.Done()
	}()
}


func httpGet(link string) ([]byte, error){
	response, err := http.Get(link)
	if err != nil {
		return nil, err
	}

	defer func() {
		if response.Body != nil {
			_ = response.Body.Close()
		}
	}()

	return ioutil.ReadAll(response.Body)
}

func saveChapterContent(){
	defer wg.Done()
	for content := range workContent{
		// 写入文件
		chapterFile, err := os.OpenFile(DIR_PATH + "/" + content.title + ".txt", os.O_CREATE | os.O_WRONLY, os.ModePerm)
		if err != nil {
			log.Panicf("创建%s文件出错\n", content.title)
		}
		log.Printf("创建%s文件\n", content.title)

		n, err := chapterFile.WriteString(content.content)
		if err != nil {
			log.Panicf("写入%s文件出错\n", content.title)
		}

		log.Printf("写入%s文件成功，成功写入%d个字符", content.title, n)
	}
}

func convertToString(src string, srcCode string, tagCode string) string {

	srcCoder := mahonia.NewDecoder(srcCode)

	srcResult := srcCoder.ConvertString(src)

	tagCoder := mahonia.NewDecoder(tagCode)

	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)

	result := string(cdata)

	return result

}