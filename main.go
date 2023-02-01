package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	port   = "8080" // The default PORT to serve this
	albums = []album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99, Poster: "https://m.media-amazon.com/images/I/71aSnR4uTNL._AC_SX385_.jpg", Sample: "https://storage.highresaudio.com/library/audio/c_190000/197842/format110516.mp3"},
		{ID: "2", Title: "Briefcase Full Of Blues", Artist: "Blues Brothers", Price: 17.99, Poster: "https://ia801806.us.archive.org/18/items/lp_briefcase-full-of-blues_the-blues-brothers/lp_briefcase-full-of-blues_the-blues-brothers_itemimage.png", Sample: "https://ia801806.us.archive.org/18/items/lp_briefcase-full-of-blues_the-blues-brothers/disc1/02.03.%20Soul%20Man_sample.mp3"},
		{ID: "3", Title: "Cowboy Bebop OST", Artist: "Various", Price: 39.99, Poster: "https://ia902305.us.archive.org/35/items/Cowboy-Bebop-OST-Collection/810VrvZrQsL._SL1442_.jpg?cnt=0", Sample: "https://ia802305.us.archive.org/35/items/Cowboy-Bebop-OST-Collection/Cowboy%20Bebop%20OST%20Collection/Cowboy%20Bebop%20CD-BOX%20Original%20Sound%20Track%20Limited%20Edition/Disc%201/02.%20Seatbelts%20-%20Tank%21%20%28TV%20Edit%29.mp3"}, // The initial albums, can be expaneded using post methods.
	}
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", getHome)

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Static("/js", "./js")
	router.Static("/css", "./css")

	router.Run(fmt.Sprintf("0.0.0.0:%s", port))
	fmt.Printf("Serving on %s", port)
}

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Poster string  `json:"poster"`
	Sample string  `json:"sample"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Adds a new album struct marshalled from the POST body.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// The handler for the root path.
func getHome(c *gin.Context) {
	// Loads the index template and adds the initial information.
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title":    "The Vinyl Co.",
		"Subtitle": "Enjoy the best curated collection of Vinyl albums. The Vinyl Co. now offers a wide array of Vinyl albums at the reach of your fingertips, come and enjoy some classics with us.",
		"Body":     "Just tap on the album you want to hear, and let the music set you free.",
		"albums":   albums,
	})
}

// Handles POST requests
// Bind the content body to a JSON marshalled from the album struct.
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		fmt.Printf("Encountered Error from POST albums %s", err)
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
