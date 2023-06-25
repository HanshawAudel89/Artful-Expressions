package gallery

import (
	"fmt"
	"net/http"
	"time"
)

// Gallery holds information about an art gallery and studio.
type Gallery struct {
	Name string
	Images []string
	Location string
	HoursOfOperation time.Time
	SocialMedia string
	Events []string
}

// NewGallery creates an instance of Gallery.
func NewGallery(name, location, socialMedia string, images []string, hoursOfOperation time.Time, events []string) *Gallery {
	return &Gallery{
		Name: name,
		Location: location,
		SocialMedia: socialMedia,
		Images: images,
		HoursOfOperation: hoursOfOperation,
		Events: events,
	}
}

// GetInfo returns information about the gallery.
func (g *Gallery) GetInfo() string {
	return fmt.Sprintf("Welcome to %v! We are located at %v and are open from %v to %v. Follow us on %v for updates about our events: %v", 
		g.Name, g.Location, g.HoursOfOperation.InputFormat("03:04 PM"), g.HoursOfOperation.Format("03:04 PM"), g.SocialMedia, g.Events)
}

// RequestHandler is the HTTP handler for requests to the gallery.
func (g *Gallery) RequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, g.GetInfo())
	case "POST":
		if len(g.Images) > 0 {
			fmt.Fprintf(w, g.Images[0])
		}
	case "PUT":
		if len(g.Events) > 0 {
			fmt.Fprintf(w, g.Events[0])
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}