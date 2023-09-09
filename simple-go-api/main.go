package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type article struct {
	ID        int    `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created"`
	UpdatedAt string `json:"updated"`
	Likes     int    `json:"likes"`
}

var articles = []article{
	{ID: 1, Author: "A J", Title: "Post broj 1", Content: "Ovo je neki sadrzaj_1", CreatedAt: "01-08-2022", UpdatedAt: "", Likes: 111},
	{ID: 2, Author: "A J", Title: "Post broj 2", Content: "Ovo je neki sadrzaj_2", CreatedAt: "10-08-2022", UpdatedAt: "13-08-2022", Likes: 171},
	{ID: 3, Author: "A J", Title: "Post broj 3", Content: "Ovo je neki sadrzaj_3", CreatedAt: "21-08-2022", UpdatedAt: "", Likes: 223},
	{ID: 4, Author: "A J", Title: "Post broj 4", Content: "Ovo je neki sadrzaj_4", CreatedAt: "07-09-2022", UpdatedAt: "", Likes: 118},
	{ID: 5, Author: "A J", Title: "Post broj 5", Content: "Ovo je neki sadrzaj_5", CreatedAt: "04-10-2022", UpdatedAt: "05-10-2022", Likes: 197},
	{ID: 6, Author: "A J", Title: "Post broj 6", Content: "Ovo je neki sadrzaj_6", CreatedAt: "11-11-2022", UpdatedAt: "", Likes: 217},
	{ID: 7, Author: "A J", Title: "Post broj 7", Content: "Ovo je neki sadrzaj_7", CreatedAt: "03-01-2023", UpdatedAt: "", Likes: 212},
	{ID: 8, Author: "A J", Title: "Post broj 8", Content: "Ovo je neki sadrzaj_8", CreatedAt: "21-01-2023", UpdatedAt: "01-02-2023", Likes: 119},
	{ID: 9, Author: "A J", Title: "Post broj 9", Content: "Ovo je neki sadrzaj_9", CreatedAt: "04-02-2023", UpdatedAt: "", Likes: 203},
	{ID: 10, Author: "A J", Title: "Post broj 10", Content: "Ovo je neki sadrzaj_10", CreatedAt: "28-02-2023", UpdatedAt: "09-03-2023", Likes: 299},
}

func getArticles(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, articles)
}

func addNewArticle(c *gin.Context) {
	var newArticle article

	if err := c.BindJSON(&newArticle); err != nil {
		return
	}
	articles = append(articles, newArticle)
	c.IndentedJSON(http.StatusCreated, newArticle)
}

func getArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id")) // ovde sam konvertovao string u int posto je u objektu ID definisan kao int, err sam stavio u _ da se ne koristi, mada mogu i da ispitam da li ima greske. Ne znam da li moze bez ove konverzije, koliko kapiram context.Param() vraca string.
	article, err := getArticleByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, article)
}

func updateLikes(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := getArticleByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
		return
	}
	article.Likes++
	c.IndentedJSON(http.StatusOK, article)
}

func getArticleByID(id int) (*article, error) {
	for i, a := range articles {
		if a.ID == id {
			return &articles[i], nil
		}
	}
	return nil, errors.New("article not found")
}

func main() {
	router := gin.Default()
	fmt.Println("Server running on port 8080 ...")
	router.GET("/articles", getArticles)
	router.GET("/articles/:id", getArticle)
	router.PATCH("/articles/:id", updateLikes)
	router.POST("/articles", addNewArticle)
	router.Run("localhost:8080")
}
