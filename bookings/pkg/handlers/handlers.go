package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JusticeMuch/bookings/pkg/config"
	"github.com/JusticeMuch/bookings/pkg/models"
	"github.com/JusticeMuch/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "This is the home page")
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

//ABout is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "world"
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Gudetama is the gudetama page handler
func (m *Repository) Gudetama(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "world"
	render.RenderTemplate(w, r, "gudetama.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Gintama is the gintama page handler
func (m *Repository) Gintama(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "gintama.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Gintama is the gintama page handler
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "reservation.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool `json:"ok"`
	Message string 	`json:"message"`
}

//handler for request , sends json response
func (m *Repository) ReservationJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse {
		OK: true,
		Message : "Available!",
	}

	out,err := json.MarshalIndent(resp, "", "   ");
	if err != nil{
		log.Println(err)
	}

	w.Header().Set("Content-type", "application/json");
	w.Write(out)
}

//Gintama is the gintama page handler
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

//Gintama is the gintama page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

func addValues(x, y int) int {
	return x + y
}

// func Divide(w http.ResponseWriter, r *http.Request){
// 	f, err := divideValues(100.0, 0.0)
// 	if err != nil{
// 		fmt.Fprintf(w, err.Error())
// 		return
// 	}

// 	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is  %f", 100.0 ,0.0 ,f))
// }

// func divideValues(x, y float32) (float32, error){
// 	var err error
// 	if y == 0{
// 		err = errors.New("cannot divide by zero")
// 	}

// 	result := x/y

// 	return result, err

// }
