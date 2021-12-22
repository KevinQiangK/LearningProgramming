package main

import (
	_ "github.com/go-redis/redis/v7"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.Println("Init invoked topst")
}

func main() {
	logrus.Println("Hello, go module mode")
	logrus.Println(uuid.NewString())
}

func init() {
	logrus.Println("Init invoked second")
}
