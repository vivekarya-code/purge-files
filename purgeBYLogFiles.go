package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

var (
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

// Constants
const DirectoryPath = "D:/logs"
const PurgeDays = 30
const LogFilePath = "D:/logs/purge-files.txt"

// Initialize...
func init() {
	file, err := os.OpenFile(LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(file, "\r\nINFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "\r\nWARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "\r\nERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {

	InfoLogger.Println("\n =====================Starting the application=====================")

	purgeFiles(DirectoryPath, PurgeDays)

	InfoLogger.Println("                 Exiting the application                   ")

}

func purgeFiles(path string, Days int) {

	file, err := os.Open(path)
	if err != nil {
		ErrorLogger.Println(err)
		panic(err)
	}
	defer file.Close()

	today := time.Now()

	filesPurgeCount := 0
	list, _ := file.Readdirnames(0)
	for _, name := range list {
		fileStat, err := os.Stat(path + "/" + name)
		if err != nil {
			ErrorLogger.Println(err)
			panic(err)
		}
		lastModified := fileStat.ModTime()
		daysOld := int(today.Sub(lastModified).Hours() / 24)

		if daysOld > PurgeDays {
			err := os.Remove(path + "/" + name)
			if err != nil {
				ErrorLogger.Panicln(err)
				panic(err)
			}
			filesPurgeCount += 1
		}
	}
	InfoLogger.Println("Number of files purged: " + strconv.Itoa(filesPurgeCount))

}
