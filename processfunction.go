package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/dymm/gorchestrator/pkg/messaging"
)

type dataType struct {
	Name  string
	Value int
}

func deserializeDataType(values map[string]string) (dataType, error) {

	serialized, found := values["data"]

	var data dataType
	if !found {
		return data, errors.New("No data found")
	}

	err := json.Unmarshal([]byte(serialized), &data)
	return data, err
}

func createValueProducer(queue messaging.Queue, outgoing string) {
	time.Sleep(3 * time.Second)
	counter := 0
	for {
		counter = counter + 1
		newValue := dataType{
			Name:  fmt.Sprintf("Value %d", counter),
			Value: rand.Intn(100),
		}
		fmt.Printf("%s : Producing the value %d\n", newValue.Name, newValue.Value)
		serialized, _ := json.Marshal(newValue)
		newWorkItem := messaging.NewWorkItem(map[string]string{"data": string(serialized)})

		if err := queue.Send(outgoing, newWorkItem); err != nil {
			fmt.Println("Error while sending the message. ", err)
			os.Exit(0)
		}
		if counter%222 == 0 {
			time.Sleep(30 * time.Second)
		} else {
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func addConstToValue(queue messaging.Queue, outgoing string, valueToAdd int) {
	fmt.Println(("Starting addConstToValue"))
	defer fmt.Println(("Stoping addConstToValue"))

	for {
		workItem, err := queue.Receive()

		var data dataType
		if err == nil {
			data, err = deserializeDataType(workItem.GetValues())
		}
		if err != nil {
			fmt.Println("addConstToValue : error while reading the message. ", err)
			os.Exit(0)
		}

		data.Value = data.Value + valueToAdd
		serializedValue, _ := json.Marshal(data)
		workItem.GetValues()["data"] = string(serializedValue)

		if data.Value >= 0 && data.Value <= 100 {
			fmt.Printf("%s addConstToValue : loosing the value\n", data.Name)
			continue //Lose the message for a timeout
		}

		err = queue.Send(outgoing, workItem)
		if err != nil {
			fmt.Println("addConstToValue : error while sending the message. ", err)
			os.Exit(0)
		}
	}
}

func subConstToValue(queue messaging.Queue, outgoing string, valueToSub int) {
	fmt.Println(("Starting subConstToValue"))
	defer fmt.Println(("Stoping subConstToValue"))

	for {
		workItem, err := queue.Receive()

		var data dataType
		if err == nil {
			data, err = deserializeDataType(workItem.GetValues())
		}
		if err != nil {
			fmt.Println("subConstToValue : error while reading the message. ", err)
			os.Exit(0)
		}

		data.Value = data.Value - valueToSub
		serializedValue, _ := json.Marshal(data)
		workItem.GetValues()["data"] = string(serializedValue)

		err = queue.Send(outgoing, workItem)
		if err != nil {
			fmt.Println("subConstToValue : error while sending the message. ", err)
			os.Exit(0)
		}
	}
}

func printTheValue(queue messaging.Queue, outgoing string) {
	fmt.Println(("Starting printTheValue"))
	defer fmt.Println(("Stoping printTheValue"))

	for {
		workItem, err := queue.Receive()
		var data dataType
		if err == nil {
			data, err = deserializeDataType(workItem.GetValues())
		}
		if err != nil {
			fmt.Println("printTheValue : error while reading the message. ", err)
			os.Exit(0)
		}

		fmt.Printf("%s : value is %d\n", data.Name, data.Value)

		err = queue.Send(outgoing, workItem)
		if err != nil {
			fmt.Println("printTheValue : error while sending the message. ", err)
			os.Exit(0)
		}
	}
}

func handleError(queue messaging.Queue, outgoing string) {
	fmt.Println(("Starting handleError"))
	defer fmt.Println(("Stoping handleError"))

	for {
		workItem, err := queue.Receive()
		var data dataType
		if err == nil {
			data, err = deserializeDataType(workItem.GetValues())
		}
		if err != nil {
			fmt.Println("handleError : error while reading the message. ", err)
		} else {
			fmt.Printf("%s : An error occured -> %v\n", data.Name, workItem.GetValues()["error"])
		}

		err = queue.Send(outgoing, workItem)
		if err != nil {
			fmt.Println("handleError : error while sending the message. ", err)
			os.Exit(0)
		}
	}
}
