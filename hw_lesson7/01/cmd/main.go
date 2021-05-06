package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zn11ch/GO01-onl-work-hard/hw_lesson7/01/internal"
	"github.com/zn11ch/GO01-onl-work-hard/hw_lesson7/01/internal/workers"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/production.json", "path to config file")
}

func main() {

	flag.Parse()

	config := internal.NewConfig()
	configFile, _ := ioutil.ReadFile(configPath)
	err := json.Unmarshal([]byte(configFile), &config)

	fmt.Println("Config:", config)
	if err != nil {
		panic(err)
	}

	absPath, _ := filepath.Abs("./test.txt")
	file, err := os.Open(absPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := internal.NewLinesChannel()
	words := internal.NewWordsChannel()
	quit := make(chan bool)
	sm := sync.Map{}
	once := sync.Once{}
	wg := sync.WaitGroup{}
	fileReaderWg := sync.WaitGroup{}
	countWordsWg := sync.WaitGroup{}
	checkWg := sync.WaitGroup{}

	w := workers.Worker{false, &wg, config}
	for i := 0; i < config.Workers.FileRead; i++ {
		wg.Add(1)
		fileReaderWg.Add(1)
		r := workers.NewFileReader(w, lines, quit, scanner, &fileReaderWg)
		go r.EventLoop()

	}

	for i := 0; i < config.Workers.CountWords; i++ {
		wg.Add(1)
		countWordsWg.Add(1)
		c := workers.NewCountWords(w, lines, words, quit, &countWordsWg)
		go c.EventLoop()
	}

	for i := 0; i < config.Workers.Check; i++ {
		wg.Add(1)
		checkWg.Add(1)
		d := workers.NewCheckWorker(w, words, quit, &sm, &once, &checkWg)
		go d.EventLoop()
	}

	wg.Wait()

	sm.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
