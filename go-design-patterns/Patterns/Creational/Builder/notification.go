package main

// This is the finished product that is created by the builder
//-NOTE: All fields are lowercase i.e., they are not exported therefore making them immutable
//-unless, you create methods for updating fields
type Notification struct {
	title    string
	subtitle string
	message  string
	image    string
	icon     string
	priority int
	notType  string
}
