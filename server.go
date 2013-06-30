package main

import (
	"net/http"
	"net/url"
	"net/http/httputil"
	"log"
	"math/rand"
)

func startServer() {
	proxy := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// Connect to Redis
			c := pool.Get()

			checkError(c.Err(), "Failed to connect to redis")

			// Make sure that the Redis connection is closed
			defer c.Close()

			keys, error := getRedirectIPs(c, req.Host)
			checkError(error, "Failed to retrieve ip from redis")

			if len(keys) > 0 {
				key := keys[rand.Intn(len(keys))] // TODO: Make a plugin system for algorithms and stop using rand
				destination, err := url.Parse(key)
				checkError(err, "Failed to parse URL provided by redis")
				req.URL.Scheme = destination.Scheme
				req.URL.Host = destination.Host
				req.Host = destination.Host
			} else {
				destination, err := url.Parse("https://s3.amazonaws.com/heroku_pages/error.html")
				checkError(err, "Failed to parse error page URL")
				req.URL.Scheme = destination.Scheme
				req.URL.Host = destination.Host
				req.URL.Path = destination.Path
				req.Host = destination.Host
			}
		},
	}

	log.Fatal(http.ListenAndServe(":8080", &proxy))
}

func checkError(err error, description string) {
	if err != nil {
		log.Print(description);
		log.Print(err);
	}
}
