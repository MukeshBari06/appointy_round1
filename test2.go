//............for testing how meetings will be encoded in json.........

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Meeting struct {
	ID                string        `json:"id"`
	Title             string        `json:"title"`
	Participants      []Participant `json:"participants"`
	StartTime         int           `json:"startTime"`
	EndTime           int           `json:"endTime"`
	CreationTimestamp string        `json:"creationTimestamp"`
}
type Participant struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	RSVP  string `json:"rsvp"`
}

var Meetings []Meeting //stores all meetings

func handleRequests() {
	http.HandleFunc("/", homePage)                  //for testing, homepage
	http.HandleFunc("/meetings", returnAllMeetings) //allmeetings
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//return all hardcoded meetings
func returnAllMeetings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllMeetings")
	json.NewEncoder(w).Encode(Meetings)
}

func fewHardCodedMeetings() {
	Meetings = []Meeting{
		Meeting{ID: "id1", Title: "title1", Participants: []Participant{Participant{Name: "name", Email: "email", RSVP: "rsvp"}, Participant{Name: "name", Email: "email", RSVP: "rsvp"}}, StartTime: 0, EndTime: 1, CreationTimestamp: "xyz"},
		Meeting{ID: "id2", Title: "title2", Participants: []Participant{Participant{Name: "name", Email: "email", RSVP: "rsvp"}, Participant{Name: "name", Email: "email", RSVP: "rsvp"}}, StartTime: 0, EndTime: 1, CreationTimestamp: "xyz"},
		Meeting{ID: "id3", Title: "title3", Participants: []Participant{Participant{Name: "name", Email: "email", RSVP: "rsvp"}, Participant{Name: "name", Email: "email", RSVP: "rsvp"}}, StartTime: 3, EndTime: 4, CreationTimestamp: "xyz"},
		Meeting{ID: "id4", Title: "title4", Participants: []Participant{Participant{Name: "name", Email: "email", RSVP: "rsvp"}, Participant{Name: "name", Email: "email", RSVP: "rsvp"}}, StartTime: 4, EndTime: 5, CreationTimestamp: "xyz"},
		Meeting{ID: "id5", Title: "title5", Participants: []Participant{Participant{Name: "name", Email: "email", RSVP: "rsvp"}, Participant{Name: "name", Email: "email", RSVP: "rsvp"}}, StartTime: 5, EndTime: 6, CreationTimestamp: "xyz"},
		Meeting{ID: "id6", Title: "title6", Participants: []Participant{Participant{Name: "name", Email: "email", RSVP: "rsvp"}, Participant{Name: "name", Email: "email", RSVP: "rsvp"}}, StartTime: 6, EndTime: 7, CreationTimestamp: "xyz"},
		Meeting{ID: "id7", Title: "title7", Participants: []Participant{Participant{Name: "name", Email: "email", RSVP: "rsvp"}, Participant{Name: "name", Email: "email", RSVP: "rsvp"}}, StartTime: 7, EndTime: 8, CreationTimestamp: "xyz"},
		Meeting{ID: "id8", Title: "title8", Participants: []Participant{Participant{Name: "name", Email: "email", RSVP: "rsvp"}, Participant{Name: "name", Email: "email", RSVP: "rsvp"}}, StartTime: 8, EndTime: 9, CreationTimestamp: "xyz"},
	}
}
func main() {
	fewHardCodedMeetings()
	handleRequests()
}
