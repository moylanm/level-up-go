package main

import (
	"fmt"
	"log"
)

// the number of attendees we need to serve lunch to
const consumerCount = 300

// foodCourses represents the types of resources to pass to the consumers
var foodCourses = []string{
	"Caprese Salad",
	"Spaghetti Carbonara",
	"Vanilla Panna Cotta",
}

// takeLunch is the consumer function for the lunch simulation
// Change the signature of this function as required
func takeLunch(name string, in []chan string, done chan<- struct{}) {
	for _, ch := range in {
		log.Printf("%s eats %s.\n", name, <-ch)
	}
	done <- struct{}{}
}

// serveLunch is the producer function for the lunch simulation.
// Change the signature of this function as required
func serveLunch(course string, out chan<- string, done <-chan struct{}) {
	for {
		select {
		case out <- course:
		case <-done:
			return
		}
	}
}

func main() {
	log.Printf("Welcome to the conference lunch! Serving %d attendees.\n",
		consumerCount)

	var courses []chan string
	doneEating := make(chan struct{})
	doneServing := make(chan struct{})

	for _, c := range foodCourses {
		ch := make(chan string)
		courses = append(courses, ch)
		go serveLunch(c, ch, doneServing)
	}

	for i := range consumerCount {
		name := fmt.Sprintf("Attendee %d", i)
		go takeLunch(name, courses, doneEating)
	}

	for range consumerCount {
		<-doneEating
	}

	close(doneServing)
}
