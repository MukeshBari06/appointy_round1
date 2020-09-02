//...............for testing urls.............

// for time, instead of (dd/mm/yyyy, hh,mn) format... int is used, just assume time is integer ex- start time: 1, endTime: 5

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	//for generating id
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

/*
type time struct {
	d int `json:"d"`
	m int `json:"m"`
	y int `json:"y"`
	hh int `json:"hh"`
	mm int `json:"mm"`
}
*/

//Global variables for storing results for different requests

var Meetings []Meeting //stores all meetings
var meetingById Meeting
var meetingsInTime bson.M
var meetingsOfParticipant bson.M

func handleRequests() {
	http.HandleFunc("/", homePage)                     //for testing, homepage
	http.HandleFunc("/allmeetings", returnAllMeetings) //allmeetings
	http.HandleFunc("/meeting/", returnMeetingById)
	http.HandleFunc("/meetings", returnMeetingsInTime) //same func will schedule meeting if no querries are given
	http.HandleFunc("/articles", returnMeetinsOfParticipant)
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

func returnMeetingById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnMeetingById")

	id := r.URL.Path[len("/meeting/"):]
	fmt.Fprintf(w, "<h1>id: %s</h1>", id)
}

func returnMeetingsInTime(w http.ResponseWriter, r *http.Request) {

	u, _ := url.Parse(r.URL.String())
	q, _ := url.ParseQuery(u.RawQuery)
	if q.Get("start") == "" && q.Get("end") == "" {
		fmt.Println("Endpoint Hit: schedulingMeeting")
		//fmt.Fprintf(w, "<h1>Schedule Meeting:</h1>")
		id := "vfdssv" //will be generated
		fmt.Fprintf(w, "<b>Schedule Meeting</b>"+
			"<form action=\"/meeting/%s\">"+
			"<label for=\"participantName\">Name:</label><br>"+
			"<input type=\"text\" id=\"participantName\" name=\"participantName\"><br>"+
			"<label for=\"email\">Email:</label><br>"+
			"<input type=\"text\" id=\"email\" name=\"email\"><br>"+
			"<label for=\"meetingTitle\">Title:</label><br>"+
			"<input type=\"text\" id=\"meetingTitle\" name=\"meetingTitle\"><br>"+
			"<label for=\"rsvp\">RSVP:</label><br>"+
			"<input type=\"text\" id=\"rsvp\" name=\"rsvp\"><br>"+
			"<label for=\"startTime\">Start Time:</label><br>"+
			"<input type=\"text\" id=\"startTime\" name=\"startTime\"><br>"+
			"<label for=\"endTime\">End Time:</label><br>"+
			"<input type=\"text\" id=\"endTime\" name=\"endTime\"><br><br>"+
			"<input type=\"submit\" value=\"Submit\"><br>"+
			"</form>", id)

	} else {
		fmt.Println("Endpoint Hit: returnMeetingsInTime")
		fmt.Fprintf(w, "<h1>startTime: %s</h1>", q.Get("start"))
		fmt.Fprintf(w, "<h1>endTime: %s</h1>", q.Get("end"))
	}
}

func returnMeetinsOfParticipant(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnMeetinsOfParticipant")

	u, _ := url.Parse(r.URL.String())
	q, _ := url.ParseQuery(u.RawQuery)
	fmt.Fprintf(w, "<h1>email: %s</h1>", q.Get("participant"))
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
