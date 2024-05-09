package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Event struct {
	Time   string
	ID     int
	Name   string
	Number int
}

type Client struct {
	Name      string
	Table     int
	Arrival   string
	Departure string
}

func formatTime(time string) string {
	return time
}

// Функция для конвертации времени в формат int
func convertTimeToInt(timeStr string) int {
	var timeInt int
	fmt.Sscanf(timeStr, "%d", &timeInt)
	return timeInt
}

// Функция для получения времени, на которое стол был занят
func getOccupiedTime(clients map[string]Client, table int) int {
	var totalOccupiedTime int
	for _, client := range clients {
		if client.Table == table && client.Departure != "" {
			arrivalTime := convertTimeToInt(client.Arrival)
			departureTime := convertTimeToInt(client.Departure)
			duration := (departureTime - arrivalTime + 100) / 100
			totalOccupiedTime += duration
		}
	}
	return totalOccupiedTime
}

func main() {
	var tables int
	var openTime, closeTime string
	var hourCost int
	var clients = make(map[string]Client)
	var waitList []string

	file, err := os.Open("input_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if len(line) == 1 {
			tables, _ = strconv.Atoi(line[0])
		} else if len(line) == 2 {
			openTime = line[0]
			closeTime = line[1]
			fmt.Println(openTime)
		} else if len(line) == 1 {
			hourCost, _ = strconv.Atoi(line[0])
		} else {
			eventTime := line[0]
			eventID, _ := strconv.Atoi(line[1])
			eventName := line[len(line)-1]

			switch eventID {
			case 1:
				if _, ok := clients[eventName]; ok {
					fmt.Println(eventTime, "YouShallNotPass")
					continue
				}
				if eventTime < openTime || eventTime > closeTime {
					fmt.Println(eventTime, "NotOpenYet")
					continue
				}
				clients[eventName] = Client{Name: eventName, Arrival: eventTime}
				fmt.Println(eventTime, eventID, eventName)
			case 2:
				client, ok := clients[eventName]
				if !ok {
					fmt.Println(eventTime, "ClientUnknown")
					continue
				}
				table, _ := strconv.Atoi(line[3])
				if client.Table == table {
					continue
				}
				if client.Table != 0 {
					fmt.Println(eventTime, "PlaceIsBusy")
					continue
				}
				client.Table = table
				clients[eventName] = client
				fmt.Println(eventTime, eventID, eventName, table)
			case 3:
				if len(clients) < tables {
					fmt.Println(eventTime, "ICanWaitNoLonger!")
					continue
				}
				waitList = append(waitList, eventName)
				fmt.Println(eventTime, eventID, eventName)
			case 4:
				client, ok := clients[eventName]
				if !ok {
					fmt.Println(eventTime, "ClientUnknown")
					continue
				}
				client.Departure = eventTime
				delete(clients, eventName)
				fmt.Println(eventTime, eventID, eventName)
				if len(waitList) > 0 {
					firstClient := waitList[0]
					waitList = waitList[1:]
					table, _ := strconv.Atoi(line[3])
					client.Table = table
					clients[firstClient] = client
					fmt.Printf("%s %d %s %d\n", eventTime, 12, firstClient, table)
				}
			}
		}
	}

	for i := 1; i <= tables; i++ {
		var revenue, timeOccupied int
		for _, client := range clients {
			if client.Table == i && client.Departure != "" {
				arrivalMinute, _ := strconv.Atoi(strings.Replace(client.Arrival, ":", "", 1))
				departureMinute, _ := strconv.Atoi(strings.Replace(client.Departure, ":", "", 1))
				duration := (departureMinute - arrivalMinute + 59) / 60
				revenue += duration * hourCost
				timeOccupied += duration
			}
		}
		if timeOccupied > 0 {
			fmt.Printf("%d %d %02d:%02d\n", i, revenue, timeOccupied/60, timeOccupied%60)
		}
	}
	fmt.Println(closeTime)
}
