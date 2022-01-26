package models

//TemplateData holds data sento to handlers from templates
type TemplateData struct {
	StringMap map[string]string
	InfoMap   map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Error     string
}