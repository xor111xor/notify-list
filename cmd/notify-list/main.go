package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

type Task struct {
	Title   string
	Message string
	Urgency string
	Cron    string
}

func (t Task) Notify() {
	path, err := exec.LookPath("notify-send")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Handling special titles
	if t.Title == "@time" {
		t.Title = time.Now().Format("15:00")
	}

	c := exec.Command(path, t.Title, t.Message, "-u", t.Urgency)
	err = c.Run()
	if err != nil {
		fmt.Println("Error send notification, check structure of task")
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Send notification %v, period: %v\n", t.Title, t.Cron)
}

var Tasks []Task

var AppVersion string

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Performing notifications for list of tasks on cron format\n")
		flag.PrintDefaults()
	}
	flagList := flag.String("list", "$HOME/.config/notify-list/list.json", "set path to list of tasks")
	flagVersion := flag.Bool("v", false, "show version")
	flag.Parse()

	if *flagVersion {
		fmt.Println(AppVersion)
		os.Exit(0)
	}

	list_path := fmt.Sprint(os.ExpandEnv(*flagList))

	list, err := os.Open(list_path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = json.NewDecoder(list).Decode(&Tasks)
	if err != nil {
		fmt.Printf("Error parse json: %v", err)
		os.Exit(1)
	}

	// Run
	c := cron.New()

	for i := range Tasks {
		ii := i
		_, err := c.AddFunc(Tasks[ii].Cron, func() { Tasks[ii].Notify() })
		if err != nil {
			panic(err)
		}
	}
	go c.Start()

	// Expect signal for exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
