package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	elastic "gopkg.in/olivere/elastic.v5"
	"os"
	"golang.org/x/net/context"
	"net/http"
	"fmt"
	"time"
	"encoding/json"
	"reflect"
)

var (
	client *elastic.Client
	errConn error
	ctx context.Context
)

const (
	Index = "articles"
	Type = "article"
	DateForm = "2006-01-02"
	DateShortForm = "20060102"
	DatePrintForm = "2006-01-02"
)

type ArticleView struct {
	Id    string    `json:"id,omitempty"`
	Title string    `json:"title,omitempty"`
	Date  string 	`json:"date,omitempty"`
	Body  string    `json:"body,omitempty"`
	Tags  []string  `json:"tags,omitempty"`
}

type ArticleSanitized struct {
	Id    string
	Title string
	Date  time.Time
	Body  string
	Tags  []string
}

type Tag struct {
	Tag         string   `json:"tag"`
	Count       int      `json:"count"`
	Articles    []string `json:"article,omitempty"`
	RelatedTags []string `json:"related_tags,omitempty""`
}


func insertArticle(article ArticleSanitized) (error) {
	_, err := client.Index().
		Index(Index).
		Type(Type).
		Id(article.Id).
		BodyJson(article).
		Do(ctx)


	return err
}

func fetchArticle(id string) (int, ArticleSanitized) {

	result, err := client.Get().
		Index(Index).
		Type(Type).
		Id(id).
		FetchSource(true).
		Do(ctx)

	if err != nil {
		return 500, ArticleSanitized{}
	}

	if result.Found {
		var article ArticleSanitized

		err = json.Unmarshal(*result.Source, &article)

		if err == nil {
			return 200, article
		} else {
			fmt.Println(err.Error())
			return 500, ArticleSanitized{}
		}

	} else {
		return 404, ArticleSanitized{}
	}
}

func fetchTag(tag string, timestamp time.Time) (int, Tag) {
	query := elastic.NewBoolQuery().Must(elastic.NewTermQuery("Date", timestamp)).Must(elastic.NewTermQuery("Tags", tag))
	searchResult, err := client.Search().
		Index(Index).
		Query(query).
		Do(ctx)

	if err != nil {
		return 500, Tag{}
	}

	var article ArticleSanitized

	count := 0
	var ids []string
	var relatedTags []string
	for _, item := range searchResult.Each(reflect.TypeOf(article)) {
		if a, ok := item.(ArticleSanitized); ok {
			fmt.Printf("Title: %s Id: %s Tags: %v\n", a.Title, a.Id, a.Tags);
			for _, t := range a.Tags {
				fmt.Println(t + " " + tag)

				if !stringInSlice(t, relatedTags) && t != tag {
					fmt.Println("APPENDING")
					relatedTags = append(relatedTags, t)
				}
			}

			ids = append(ids, a.Id)
			count++
		}
	}

	return 200, Tag{
		Tag: tag,
		Count:count,
		Articles: ids,
		RelatedTags: relatedTags,
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func articleCreate(c *gin.Context) {
	var json ArticleView

	if c.BindJSON(&json) == nil {

		tc, _ := time.Parse(DateForm, json.Date)

		article := ArticleSanitized{
			Id: json.Id,
			Title: json.Title,
			Date: tc,
			Tags: json.Tags,
		}

		err := insertArticle(article)

		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})

			return
		} else {
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println("Incorrect params given")
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status": "error",
	})

}

func articleShow(c *gin.Context) {
	id := c.Param("id")

	status, result := fetchArticle(id)

	tc := result.Date.Format(DatePrintForm)

	article := ArticleView{
		Id: result.Id,
		Title: result.Title,
		Date: tc,
		Tags: result.Tags,
	}

	c.JSON(status, article)

}

func tagShow(c *gin.Context) {
	tag := c.Param("tagName")
	tc, _ := time.Parse(DateShortForm, c.Param("date"))

	status, result := fetchTag(tag, tc)

	c.JSON(status, result)
}

func main() {
	ctx = context.Background()

	username := os.Getenv("ESUSER")
	password := os.Getenv("ESPASS")
	host := os.Getenv("ESHOST")
	port := os.Getenv("ESPORT")

	// initialize elasticsearch client
	client, errConn = elastic.NewClient(
		elastic.SetBasicAuth(username, password),
		elastic.SetURL("http://"+host+":"+port))

	if errConn != nil {
		panic(errConn)
	}

	exists, err := client.IndexExists("articles").Do(context.Background())

	if err != nil {
		panic(err)
	}

	if !exists {
		// Creating the index
		_, err := client.CreateIndex("articles").Do(ctx)
		if err != nil {
			panic(err)
		}
	}

	// setting up routes
	router := gin.Default()
	router.POST("/articles", articleCreate)
	router.GET("/articles/:id", articleShow)
	router.GET("/tags/:tagName/:date", tagShow)

	router.Run(":8080")
}
