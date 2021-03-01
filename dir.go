package main

import (
    "metricstream.com/metricstream/modules"
    "fmt"
)

func main() {

	dir := mdir.UserHomeDir()
        fmt.Println(dir)
}
