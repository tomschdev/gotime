package main

import "log"

func main() {
	// Create a NotificationBuilder and use it to set properties
	bldr := newNotificationBuilder()
	// Use the builder to set some properties
	bldr.SetTitle("New Notification")
	bldr.SetIcon("icon.png")
	bldr.SetSubTitle("This is the subtitle for my icon")
	bldr.SetImage("image.jpg")
	bldr.SetPriority(5)
	bldr.SetMessage("This is a basic notification with text.")
	bldr.SetType("alert")
	// Use the Build function to create a finished object
	notif, err := bldr.Build()
	if err != nil {
		log.Fatalf("error creating the notification: ", err)
	} else {
		log.Printf("new notification created: %+v\n", notif)
	}

}
