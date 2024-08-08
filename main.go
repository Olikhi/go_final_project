package main

func main() {
	initDb()
	startServer()
	defer db.Close()
}
