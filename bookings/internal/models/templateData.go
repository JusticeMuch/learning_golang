package models

import (
	"github.com/JusticeMuch/bookings/internal/forms"
)

//TemplateData holds data sento to handlers from templates
type TemplateData struct {
	StringMap map[string]string
	InfoMap   map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Warning   string
	Flash     string
	Error     string
	Form      *forms.Form
}
