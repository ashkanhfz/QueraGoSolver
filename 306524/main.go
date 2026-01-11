package main

import (
	"fmt"
)

type Event struct {
	costPerPerson int64
	participants  map[string]bool
}

func createEvent(events map[string]*Event, name string, cost int64) {
	if _,ok:=events[name];ok{
		fmt.Println("UNSUCCESSFUL CREATE")
		return
	}
	events[name]=&Event{
			costPerPerson: cost,
			participants: make(map[string]bool)}
	fmt.Println("SUCCESSFUL")
}	

func deleteEvent(events map[string]*Event, name string) {
	if _,ok:=events[name];!ok{
		fmt.Println("INVALID EVENTNAME")
		return
	}
	delete(events, name)
	fmt.Println("SUCCESSFUL")
}

func addUser(events map[string]*Event,user_name string, event_name string) {
	if _,ok:=events[event_name];!ok{
		fmt.Println("INVALID EVENTNAME")
		return
	}
	if _,ok  := events[event_name].participants[user_name];ok{
		fmt.Println("USER ALREADY ADDED TO EVENT")
		return
	}
	events[event_name].participants[user_name] = false
	fmt.Println("SUCCESSFUL")

	
}

func removeUser(events map[string]*Event,user_name string, event_name string) {
	if _,ok:=events[event_name];!ok{
		fmt.Println("INVALID EVENTNAME")
		return
	}
	if _,ok  := events[event_name].participants[user_name];!ok{
		fmt.Println("USER NOT FOUND IN EVENT")
		return
	}

	delete(events[event_name].participants, user_name)
	fmt.Println("SUCCESSFUL")

}

func costEvent(events map[string]*Event,name string) {
	event, ok := events[name]
	if !ok{
		fmt.Println("INVALID EVENTNAME")
		return
	}
	total := int64(len(event.participants)) * event.costPerPerson
	fmt.Println(total)
}



func main() {
	events := make(map[string]*Event)
	for{
		var inp string
		fmt.Scan(&inp)
		switch inp{
			case "CREATE":
				var name string
				var cost int64
				fmt.Scan(&name,&cost)
				createEvent(events, name,cost)
			case "DELETE":
				var name string
				fmt.Scan(&name)
				deleteEvent(events, name)
			case "ADD":
				var user_name, event_name string
				fmt.Scan(&user_name, &event_name)
				addUser(events, user_name, event_name)
			case "REMOVE":
				var user_name, event_name string
				fmt.Scan(&user_name, &event_name)
				removeUser(events, user_name, event_name)
			case "COST":
				var name string
				fmt.Scan(&name)
				costEvent(events, name)
			case "FINISH":
				return
		}
	}
}
