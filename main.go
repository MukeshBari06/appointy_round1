// for time, instead of (dd/mm/yyyy, hh,mn) format... int is used, just assume time is integer ex- start time: 1, endTime: 5

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	//for generating id
	"crypto/rand"
	"encoding/hex"
)

/*
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
*/
/*
type time struct {
	d int `json:"d"`
	m int `json:"m"`
	y int `json:"y"`
	hh int `json:"hh"`
	mm int `json:"mm"`
}
*/

//uri --------- "mongodb+srv://new_user:<PASSWORD>@cluster0.anp21.gcp.mongodb.net/<meetings>?retryWrites=true&w=majority"         <PASSWORD>:5XIeSJBPkyECny53
//url --------- "mongodb://localhost:27017"

//Global variables for storing results for different requests

//var Meetings []Meeting //stores all meetings
var meetingById bson.M
var meetingsInTime bson.M
var meetingsOfParticipant bson.M

func handleRequests() {
	http.HandleFunc("/", homePage) //for testing, homepage
	//http.HandleFunc("/allmeetings", returnAllMeetings)	//allmeetings
	http.HandleFunc("/meeting/", returnMeetingById)
	http.HandleFunc("/meetings", returnMeetingsInTime) //same func will schedule meeting if no querries are given
	http.HandleFunc("/articles", returnMeetinsOfParticipant)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

/* //return all hardcoded meetings
func returnAllMeetings(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllMeetings")
	json.NewEncoder(w).Encode(Meetings)
}*/

func returnMeetingById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnMeetingById")

	id := r.URL.Path[len("/meeting/"):]

	getMeetingById(id)
	json.NewEncoder(w).Encode(meetingById)
}

func returnMeetingsInTime(w http.ResponseWriter, r *http.Request) {

	u, _ := url.Parse(r.URL.String())
	q, _ := url.ParseQuery(u.RawQuery)
	if q.Get("start") == "" && q.Get("end") == "" {
		fmt.Println("Endpoint Hit: schedulingMeeting")
		scheduleMeeting(w, r)
	} else {
		fmt.Println("Endpoint Hit: returnMeetingsInTime")

		startTime, _ := strconv.Atoi(q.Get("start"))
		endTime, _ := strconv.Atoi(q.Get("end"))
		getMeetingsInTime(startTime, endTime)
		json.NewEncoder(w).Encode(meetingsInTime)
	}
}

func returnMeetinsOfParticipant(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnMeetinsOfParticipant")

	u, _ := url.Parse(r.URL.String())
	q, _ := url.ParseQuery(u.RawQuery)

	email := q.Get("participant")

	getMeetingsOfParticipant(email)
	json.NewEncoder(w).Encode(meetingsOfParticipant)
}

func getMeetingById(id1 string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://new_user:5XIeSJBPkyECny53@cluster0.anp21.gcp.mongodb.net/<meetings>?retryWrites=true&w=majority"))

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	meetingCollection := client.Database("db").Collection("meetings")

	// Get a MongoDB document using the FindOne() method
	err = meetingCollection.FindOne(context.TODO(), bson.D{{"ID", id1}}).Decode(&meetingById)
	if err != nil {
		fmt.Println("FindOne() No meeting found:", err)
		os.Exit(1)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}

func newMeetPossible(email1 string, startTime int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://new_user:5XIeSJBPkyECny53@cluster0.anp21.gcp.mongodb.net/<meetings>?retryWrites=true&w=majority"))

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	meetingCollection := client.Database("db").Collection("meetings")

	cursor, err := meetingCollection.Find(context.TODO(), bson.D{{"StartTime", startTime}, {"RSVP", "yes"}, {"Participants", bson.D{{"$elemMatch", bson.D{{"Email", email1}}}}}})

	var possible bool = true

	// Find() method raised an error
	if err != nil {
		fmt.Println("checking for scheduling meeting  ERROR:", err)
		defer cursor.Close(ctx)

		// If the API call was a success
	} else {
		// iterate over docs using Next()

		for cursor.Next(ctx) {

			err := cursor.Decode(&meetingsOfParticipant)

			// If there is a cursor.Decode error
			if err != nil {
				fmt.Println("cursor.Next() checking for scheduling meeting error:", err)
				os.Exit(1)
			} else {
				possible = false
				break
			}

		}
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	return possible
}

func getMeetingsOfParticipant(email1 string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://new_user:5XIeSJBPkyECny53@cluster0.anp21.gcp.mongodb.net/<meetings>?retryWrites=true&w=majority"))

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	meetingCollection := client.Database("db").Collection("meetings")

	cursor, err := meetingCollection.Find(context.TODO(), bson.D{{"Participants", bson.D{{"$elemMatch", bson.D{{"Email", email1}}}}}})

	// Find() method raised an error
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)

		// If the API call was a success
	} else {
		// iterate over docs using Next()
		for cursor.Next(ctx) {

			err := cursor.Decode(&meetingsOfParticipant)

			// If there is a cursor.Decode error
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			}
		}
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}
func getMeetingsInTime(startTime1 int, endTime1 int) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://new_user:5XIeSJBPkyECny53@cluster0.anp21.gcp.mongodb.net/<meetings>?retryWrites=true&w=majority"))

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	meetingCollection := client.Database("db").Collection("meetings")

	cursor, err := meetingCollection.Find(context.TODO(), bson.D{{"StartTime", bson.D{{"$gt", startTime1}, {"$lt", endTime1}}}, {"EndTime", bson.D{{"$gt", startTime1}, {"$lt", endTime1}}}}) //$ giving error

	// Find() method raised an error
	if err != nil {
		fmt.Println("Finding all documents ERROR:", err)
		defer cursor.Close(ctx)

		// If the API call was a success
	} else {
		// iterate over docs using Next()
		for cursor.Next(ctx) {

			err := cursor.Decode(&meetingsInTime)

			// If there is a cursor.Decode error
			if err != nil {
				fmt.Println("cursor.Next() error:", err)
				os.Exit(1)
			}
		}
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}

/*
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
*/

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func scheduleMeeting(w http.ResponseWriter, r *http.Request) {
	id, _ := randomHex(6)
	
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
	
	meetingTitle := "$_POST['meetingTitle']"
	participantName := "$_POST['participantName']"
	email := "$_POST['email']"
	rsvp := "$_POST['rsvp']"
	startTime, _ := strconv.Atoi("$_POST['startTime']")
	endTime, _ := strconv.Atoi("$_POST['endTime']")

	var possible bool = newMeetPossible(email, startTime)
	if possible == true {
		insertDocument(id, meetingTitle, participantName, email, rsvp, startTime, endTime)
	} else {
		fmt.Fprintf(w, "busy schedule, cant schedule meeting")
	}

}

func insertDocument(id string, title string, name string, email string, rsvp string, startTime int, endTime int) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://new_user:5XIeSJBPkyECny53@cluster0.anp21.gcp.mongodb.net/<meetings>?retryWrites=true&w=majority"))

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())

	meetingCollection := client.Database("db").Collection("meetings")

	insertResult, err := meetingCollection.InsertOne(
		ctx,
		bson.D{
			{"ID", id},
			{"Title", title},
			{"Participants", bson.A{
				bson.D{
					{"Name", name},
					{"Email", email},
					{"RSVP", rsvp}}}},
			{"StartTime", startTime},
			{"EndTime", endTime},
			{"CreationTimestamp", "xyz"}})

	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID, "inserted")

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

}

func main() {
	//fewHardCodedMeetings()
	handleRequests()
}
