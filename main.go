package main

import (
	"fmt"
	"log"
	"os"
	"shop_ping/messenger"
	"shop_ping/watcher"
	"time"
)

// Loops eternally until the text is found (the website being scanned was updated!)
func blockUntilFound(url, text string) {
	hw := watcher.CreateHtmlWatcher(url)
	fmt.Println(fmt.Sprintf("Beginning to search for the phrase: %s \nOn the following page: %s", text, url))

	found := false
	var err error
	for !found {
		log.Println("Checking for text...")
		time.Sleep(20 * time.Second)
		found, err = hw.Find(text)
		if err != nil {
			log.Println("Encountered the following error below: ")
			log.Println(err.Error())
			log.Println("Continuing watch...")
		}
	}
}

// Expects args in the following order:
// shop_ping <www.website.com> <phrase to lookout for>
func main() {
	blockUntilFound(os.Args[1], os.Args[2])

	mg := messenger.CreateTwilioMessenger()
	err := mg.SendAlert("Item is now available!")
	if err != nil {
		log.Panic(err.Error())
	}
}