package main

import (
    "net/http"
    "os"
    "strconv"
    "time"

    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*")
    r.Static("/static", "./static")

    // Route for the main page
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })

    // Route for the results page
    r.GET("/results", func(c *gin.Context) {
        killsAchieved, _ := strconv.Atoi(c.Query("killsAchieved"))
        timeTaken, _ := strconv.ParseFloat(c.Query("timeTaken"), 64)
        averageTimePerKill, _ := strconv.ParseFloat(c.Query("averageTimePerKill"), 64)

        c.HTML(http.StatusOK, "results.html", gin.H{
            "KillsAchieved":          killsAchieved,
            "TimeTaken":              formatDuration(time.Duration(timeTaken) * time.Second),
            "AverageTimePerKill":     formatDuration(time.Duration(averageTimePerKill) * time.Second),
            "AverageTimePerKillRaw":  averageTimePerKill,
        })
    })

    // Get the PORT from the environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Start the server on the specified port
    r.Run(":" + port)
}

// Helper function to format duration in a more readable form
func formatDuration(d time.Duration) string {
    hours := int(d.Hours())
    minutes := int(d.Minutes()) % 60
    seconds := int(d.Seconds()) % 60
    return strconv.Itoa(hours) + "h " + strconv.Itoa(minutes) + "m " + strconv.Itoa(seconds) + "s"
}
