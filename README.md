# appointy_round1
<b>API using GoLang and MongoDB</b>


Task 1 | Meetings API<br>

The task is to develop a basic version of meeting scheduling API. You are only required to develop the API for the system. Below are the details.<br>

Meetings should have the following Attributes. All fields are mandatory unless marked optional:<br>
Id<br>
Title<br>
Participants<br>
Start Time<br>
End Time<br>
Creation Timestamp<br>

Participants should have the following Attributes. All fields are mandatory unless marked optional:<br>
Name<br>
Email<br>
RSVP (i.e. Yes/No/MayBe/Not Answered)<br>

You are required to Design and Develop an HTTP JSON API capable of the following operations,<br>
Schedule a meeting<br>
Should be a POST request<br>
Use JSON request body<br>
URL should be ‘/meetings’
<br>Must return the meeting in JSON format

<br>Get a meeting using id
<br>Should be a GET request
<br>Id should be in the url parameter
<br>URL should be ‘/meeting/<id here>’
<br>Must return the meeting in JSON format
  
<br>List all meetings within a time frame
<br>Should be a GET request
<br>URL should be ‘/meetings?start=<start time here>&end=<end time here>’
<br>Must return a an array of meetings in JSON format that are within the time range
  
<br>List all meetings of a participant
<br>Should be a GET request
<br>URL should be ‘/articles?participant=<email id>’
<br>Must return a an array of meetings in JSON format that have the participant received in the email within the time range
