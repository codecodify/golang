package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gammazero/workerpool"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

const (
	CODES = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	CODES_LEN = len(CODES)
	USER_AGENT    = "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0"
	FILE_NAME_LEN = 20
	POOL_MAXSIZE = 16
	PICS_EXT      = ".jpg"     // 图片后缀
	PICS_DIR      = "pics"     // 存放图片文件夹
	RETRY_TIMES = 10
)

// 初始化
func init(){
	createDir(PICS_DIR)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// 如果文件夹不存在，创建文件夹
func createDir(path string){
	_, err := os.Stat(path)
	if err != nil{
		if !os.IsExist(err){
			os.Mkdir(path, os.ModePerm)
			fmt.Println(fmt.Sprintf("创建%s目录", path))
		}
	}
}

// 获取响应内容
func getResponse(url string) (*http.Response, error){
	ref := "https://www.mzitu.com"
	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	request.Header.Set("User-Agent", USER_AGENT)
	request.Header.Set("Referer", ref)
	request.Header.Set("cookie", "Hm_lvt_dbc355aef238b6c32b43eacbbf161c3c=1574513360; Hm_lpvt_dbc355aef238b6c32b43eacbbf161c3c=1574515186")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return response, nil

}

func getDoc(url string) (*goquery.Document,  error){
	response, err := getResponse(url)
	if err != nil {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		panic(err)
	}
	return doc, nil
}

// 进入详情页下载图片
func downloadPics(url string){
	doc, err := getDoc(url)
	if err != nil {
		panic(err)
	}


	// 获取标题，并建立对应标题的子目录
	title := doc.Find("h2.main-title").First().Text()
	saveDir := PICS_DIR + "/" + title
	createDir(saveDir)

	// 最大页数
	maxPage := 1
	doc.Find(".pagenavi a").Each(func(i int, selection *goquery.Selection) {
		page, _ := strconv.Atoi(selection.Text())
		if page > maxPage {
			maxPage = page
		}
	})

	wp := workerpool.New(2)
	for page := 1; page <= maxPage; page ++ {
		imageUrl := fmt.Sprintf("%s/%d", url, page)
		wp.Submit(func() {
			syncSavePic(imageUrl, saveDir)
		})
	}
	wp.StopWait()
}

func syncSavePic(url string, saveDir string){
	doc, err := getDoc(url)
	if err != nil {
		panic(err)
	}
	var localFile *os.File

	// 下载图片
	var imageUrl string
	for retry := 0; retry < RETRY_TIMES; retry++{
		imageUrl, _  = doc.Find(".main-image img").First().Attr("src")
		if len(imageUrl) > 0 {
			break
		}
	}
	if len(imageUrl) == 0 {
		return
	}
	title := doc.Find(".main-title").First().Text()
	var response *http.Response
	for retry := 0; retry < RETRY_TIMES; retry++ {
		response, err = getResponse(imageUrl)
		if err == nil && response.StatusCode == 200 {
			break
		}
	}

	if response == nil || response.StatusCode != 200 {
		fmt.Println("请求失败 " + imageUrl)
		return
	}

	var filename string
	filename = title + PICS_EXT
	//if page == 1{
	//	filename = title + PICS_EXT
	//}else{
	//	filename = fmt.Sprintf("%s(%d)%s", title, page, PICS_EXT)
	//}

	localFile, err = os.Create(path.Join(saveDir, filename))
	if err != nil{
		panic(err)
	}

	fmt.Println("下载图片：", imageUrl)
	if _, err := io.Copy(localFile, response.Body); err != nil {
		fmt.Println(err)
	}

	defer localFile.Close()

}

func main() {
	// 创建 goroutine 池
	wp := workerpool.New(2)
	start := time.Now()
	total := 0

	// 1 - 236
	for page := 1; page <= 1; page++ {
		wp.Submit(func() {
			doc, err := getDoc(fmt.Sprintf("https://www.mzitu.com/page/%d/", page))
			if err != nil {
				panic(err)
			}

			doc.Find("#pins li").Each(func(i int, selection *goquery.Selection) {
				//
				//title, _ := selection.Find("img").First().Attr("alt")
				//
				//fmt.Println(title)
				// 链接
				url, _ := selection.Find("a").First().Attr("href")
				downloadPics(url)
				total++
			})
		})
	}

	// 等待所有任务完成
	wp.StopWait()

	elapsed := time.Since(start)
	fmt.Println("总条数", total)
	fmt.Println("Elapsed: ", elapsed)
}
