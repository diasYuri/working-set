package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"log"
	"os"
	vm "so.com/page-replacement/virtual-memory"
	"time"
)

type MemoryHit struct {
	ProcessId int
	PageId    int
}

type ConfigMemory struct {
	MemorySize int
	TimeWindow time.Duration
}

type Input struct {
	Config ConfigMemory
	Hits   []MemoryHit
}

func ReadInput() (*Input, error) {
	fileContent, err := os.Open("input.json")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	logrus.Infof("input file successfully read")

	defer func(fileContent *os.File) {
		_ = fileContent.Close()
	}(fileContent)

	byteResult, _ := ioutil.ReadAll(fileContent)

	var input Input

	_ = json.Unmarshal(byteResult, &input)
	return &input, nil
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	input, _ := ReadInput()

	ws := vm.NewWorkingSet(input.Config.TimeWindow)
	m := vm.NewVirtualMemory(ws, input.Config.MemorySize)

	for _, hit := range input.Hits {
		m.AccessPage(hit.ProcessId, hit.PageId)
		time.Sleep(10 * time.Nanosecond)
	}

}
