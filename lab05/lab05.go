package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"os"
	"fmt"
	"io"
	"encoding/json"
	"github.com/joho/godotenv"
)

// VideoInfo includes the information of a video
type VideoInfo struct {
	Title        string
	Id           string
	ChannelTitle string
	LikeCount    string
	ViewCount    string
	PublishedAt  string
	CommentCount string
}

// YoutubeResponse reflects the JSON response structure
type YoutubeResponse struct {
	Items []struct {
		Snippet struct {
			Title        string `json:"title"`
			ChannelTitle string `json:"channelTitle"`
			PublishedAt  string `json:"publishedAt"`
		} `json:"snippet"`
		Statistics struct {
			ViewCount    string `json:"viewCount"`
			LikeCount    string `json:"likeCount"`
			CommentCount string `json:"commentCount"`
		} `json:"statistics"`
	} `json:"items"`
}

func formatNumber(number string) string {
	if n, err := strconv.Atoi(number); err == nil {
		in := strconv.Itoa(n)
		out := make([]byte, len(in)+(len(in)-1)/3)
		for e, i, j := len(in)-1, len(out)-1, 0; e >= 0; i-- {
			out[i] = in[e]
			if j++; j == 3 && e != 0 {
				i--
				out[i] = ','
				j = 0
			}
			e--
		}
		return string(out)
	}
	return number
}

func formatDate(dateStr string) string {
	t, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return dateStr
	}
	return t.Format("2006年01月02日")
}

func YouTubePage(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	query := r.URL.Query()
	videoID := query.Get("v")
	if videoID == "" {
		http.ServeFile(w, r, "error.html")
		return
	}

	url := fmt.Sprintf("https://www.googleapis.com/youtube/v3/videos?key=%s&id=%s&part=snippet,statistics", apiKey, videoID)
	resp, err := http.Get(url)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		return
	}

	var ytResponse YoutubeResponse
	if err := json.Unmarshal(body, &ytResponse); err != nil {
		http.ServeFile(w, r, "error.html")
		return
	}

	if len(ytResponse.Items) == 0 {
		http.ServeFile(w, r, "error.html")
		return
	}

	video := ytResponse.Items[0]
	videoInfo := VideoInfo{
		Title:        video.Snippet.Title,
		Id:           videoID,
		ChannelTitle: video.Snippet.ChannelTitle,
		LikeCount:    formatNumber(video.Statistics.LikeCount),
		ViewCount:    formatNumber(video.Statistics.ViewCount),
		PublishedAt:  formatDate(video.Snippet.PublishedAt),
		CommentCount: formatNumber(video.Statistics.CommentCount),
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.ServeFile(w, r, "error.html")
		return
	}

	err = tmpl.Execute(w, videoInfo)
	if err != nil {
		http.ServeFile(w, r, "error.html")
		return
	}
}

func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}
