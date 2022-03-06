package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/JusticeMuch/bookings/internal/config"
	"github.com/JusticeMuch/bookings/internal/forms"
	"github.com/JusticeMuch/bookings/internal/models"
	"github.com/JusticeMuch/bookings/internal/render"
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
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w, r, "reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

//Post handler for post method on reservation page
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email", "phone")
	form.MinLength("first_name", 3, r)
	form.IsEmail("email")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
	}

	m.App.Session.Put(r.Context(), "reservation", reservation)

	http.Redirect(w,r, "/reservation-summary", http.StatusSeeOther);
}

func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

//handler for request , sends json response
func (m *Repository) ReservationJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "   ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(out)
}

//Gintama is the gintama page handler
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
	})
}

//Gintama is the gintama page handler
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

//reservation-summary handler
func (m *Repository) ReservationSummary(w http.ResponseWriter, r *http.Request) {
	reservation,ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok{
		log.Println("Cannot get item from session");
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		m.App.Session.Put(r.Context(), "error", "Cannot get reservation from session")
		return 
	}

	m.App.Session.Remove(r.Context(), "reservation")

	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
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
