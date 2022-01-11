package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rexposadas/locationhistory/server"
)

func main() {

	expiry := setExpiry(os.Getenv("LOCATION_HISTORY_TTL_SECONDS"))
	s := server.Setup(expiry)

	port := "8080"
	if p := os.Getenv("HISTORY_SERVER_LISTEN_ADDR"); p != "" {
		if _, err := strconv.Atoi(p); err == nil {
			port = p
		}
	}

	fmt.Printf("running on port %s", port)
	s.Run(":" + port)
}

// Default to never expiring.
// TODO:
//	We should probably display an error if we come across an invalid expiry time.
//	But, that's for another time.
func setExpiry(e string) int {
	expiry := -1 // default to never expire

	if e == "" {
		return expiry
	}

	ex, err := strconv.Atoi(e)
	if err != nil {
		return expiry
	}

	return ex
}
