package main

import (
    "fmt"
    "metricsviews.com/metricsviews/metrics/api/dir"
)

func main() {

	dir := dir.UserHomeDir()
        fmt.Println(dir)
}
