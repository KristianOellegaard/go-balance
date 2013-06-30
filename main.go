package main

func main() {
	// Create a connection pool for Redis
	createRedisPool()

	// Print the startup message
//	printStartupMessage()

	// Finally start the server
	startServer()
}
