package main

// Package -
// This API serves a web hook for TYPE FORM
// and transform the input JSON to Zen Desk equivalent
// request along side Weather API Input Data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	//Type form Data Type
	BOOLEAN_FEILD = "boolean"
	TEXT_FEILD    = "text"
	DATE_FEILD    = "date"
	FILE_UPLOAD   = "file_upload"

	//HTML tags
	PARA_START = "<p>"
	PARA_END   = "</p><hr>"

	//TV FIELD IDs
	TV_CLAIM          = "TV claim"
	TV_INCIDENT_DATE  = "63241330"
	TV_PURCHASE_PRICE = "63391165"
	TV_CRIME_REF      = "AZbSqcXTKlED"
	TV_MODEL_NO       = "lGCVB9tse6Re"
	TV_MAKE           = "WwzqHPb0K9Wv"
	TV_SERIAL_NO      = "u2dzBoYFjbRA"
	TV_DAMAGE_IMAGE   = "63241151"
	TV_RECIEPT        = "63241244"

	//Storm Surge Field IDs
	STORM_SURGE_CLAIM     = "Storm surge claim"
	STORM_INCIDENT_DATE   = "nAz5fZvtiuLO"
	STORM_INCIDENT_PLACE  = "64049754"
	STORM_DAMAGE_IMAGE_1  = "63907299"
	STORM_DAMAGE_IMAGE_2  = "j79cNctIvogK"
	STORM_REPAIR_ESTIMATE = "wFpTHm7AZVNO"
	STORM_SHORT_VIDEO     = "63907004"
)

// func main() {
// 	println("staritng app..")
// 	http.HandleFunc("/", Handler)
// 	http.ListenAndServe(":8083", nil)
// }

// Handler method transform incoming request payload
// to Zendesk Specific Format (JSON to JSON conveion)
func Handler(w http.ResponseWriter, r *http.Request) {

	// Marhsal TYPE FORM DATA to TypeFormData struct
	var typeFormdata TypeFormData
	err := json.NewDecoder(r.Body).Decode(&typeFormdata)
	if err == io.EOF || err != nil { // create error response on empty payload or malform JSON
		createErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	// transform TypeForm JSON to hold Zendesk Specific JSON Data
	transformedData := transformData(typeFormdata)
	transformedData.Status = 200
	transformedDataJSON, err := json.Marshal(&transformedData)
	fmt.Println("Type form data after transformatin -----> ", string(transformedDataJSON))
	if err != nil { // create error response on error
		createErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	//write transformed JSON data
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(string(transformedDataJSON)))
}

// transformData - Map TYPEFORM data to Zendesk JSON Schema,
func transformData(typeFormdata TypeFormData) TranformedData {
	//populate Fresh Desk Specific Struct
	var ticketDetails = TicketDetails{}
	var transformedData = TranformedData{}

	var claimType = typeFormdata.FormResponse.Definition.Title

	//update Ticket details
	ticketDetails.Ticket.EventID = typeFormdata.EventID
	ticketDetails.Ticket.Token = typeFormdata.FormResponse.Token
	ticketDetails.Ticket.SubmittedAt = typeFormdata.FormResponse.SubmittedAt

	ticketDetails.Ticket.Subject = claimType
	ticketDetails.Ticket.Requester.Email = typeFormdata.FormResponse.Hidden.Email
	ticketDetails.Ticket.Requester.Name = typeFormdata.FormResponse.Hidden.Name
	ticketDetails.Ticket.Requester.Phone = typeFormdata.FormResponse.Hidden.Phone
	ticketDetails.Ticket.Requester.PolicyNumber = typeFormdata.FormResponse.Hidden.Policy
	ticketDetails.Ticket.Requester.LocaleID = 1

	// default ticket data
	ticketDetails.Ticket.Status = "new"
	ticketDetails.Ticket.Priority = "normal"
	ticketDetails.Ticket.Type = "incident"
	ticketBody := ""

	//populate Ticket Descripton by adding all Question and Answers
	for i, field := range typeFormdata.FormResponse.Definition.Fields {
		answer := typeFormdata.FormResponse.Answers[i]
		//add all data to ticket body
		switch field.Type {
		case BOOLEAN_FEILD:
			ticketBody += PARA_START + "<b>" + field.Title + "</b>"
			ticketBody += strconv.FormatBool(answer.Boolean) + PARA_END
		case DATE_FEILD:
			ticketBody += PARA_START + "<b>" + field.Title + "</b>"
			ticketBody += " : " + answer.Date + PARA_END
		case FILE_UPLOAD:
			//ticketBody += PARA_START + "<b>" + field.Title + "</b>"
			// ticketBody += " <a href='" + answer.FileURL + "'>" + answer.FileURL + "</a>" + PARA_END
		default:
			ticketBody += PARA_START + "<b>" + field.Title + "</b>"
			ticketBody += " : " + answer.Text + PARA_END
		}
		ticketDetails.Ticket.Comment.HTMLBody = ticketBody
		//claim specific important data to be added as custom fields
		if strings.Compare(claimType, STORM_SURGE_CLAIM) == 0 {
			//for Storm Surge Claims
			switch field.ID {
			case STORM_INCIDENT_DATE:
				transformedData.WeatherAPIInput.Date = strings.Replace(answer.Date, "-", "", 2)
				transformedData.StromClaimData.IncidentDate = answer.Date
			case STORM_INCIDENT_PLACE:
				transformedData.WeatherAPIInput.City = answer.Text
				transformedData.StromClaimData.IncidentPlace = answer.Text
			case STORM_DAMAGE_IMAGE_1:
				transformedData.StromClaimData.DamageImageURL1 = answer.FileURL
			case STORM_DAMAGE_IMAGE_2:
				transformedData.StromClaimData.DamageImageURL2 = answer.FileURL
			case STORM_REPAIR_ESTIMATE:
				transformedData.StromClaimData.RepairEstimateImage = answer.FileURL
			case STORM_SHORT_VIDEO:
				transformedData.StromClaimData.DamageVideoURL = answer.FileURL
			default:
			}
		} else {
			//assuming TV claim request
			//for Storm Surge Claims
			switch field.ID {
			case TV_INCIDENT_DATE:
				transformedData.WeatherAPIInput.Date = strings.Replace(answer.Date, "-", "", 2)
				transformedData.TVClaimData.IncidentDate = answer.Date
			case TV_PURCHASE_PRICE:
				transformedData.TVClaimData.TVPrice = answer.Number
			case TV_CRIME_REF:
				transformedData.TVClaimData.CrimeRef = answer.Text
			case TV_MODEL_NO:
				transformedData.TVClaimData.TVModelNo = answer.Text
			case TV_MAKE:
				transformedData.TVClaimData.TVMake = answer.Text
			case TV_SERIAL_NO:
				transformedData.TVClaimData.TVSerialNo = answer.Text
			case TV_DAMAGE_IMAGE:
				transformedData.TVClaimData.DamageImageURL1 = answer.FileURL
			case TV_RECIEPT:
				transformedData.TVClaimData.TVReceiptImage = answer.FileURL
			default:
			}
		}
		//update ticket specific data
		transformedData.TicketDetails = ticketDetails
	}
	return transformedData
}

// createErrorResponse - this function forms a error reposne with
// error message and http code
func createErrorResponse(w http.ResponseWriter, message string, status int) {
	errorJSON, _ := json.Marshal(&Error{
		Status:  status,
		Message: message})
	//Send custom error message to caller
	w.WriteHeader(status)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(errorJSON))
}

// Error - error object
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// Output Schema
type TranformedData struct {
	Status          int             `json:"status,omitempty"`
	TicketDetails   TicketDetails   `json:"ticket_details,omitempty"`
	WeatherAPIInput WeatherAPIInput `json:"weather_api_input,omitempty"`
	TVClaimData     TVClaimData     `json:"tv_claim_data,omitempty"`
	StromClaimData  StromClaimData  `json:"storm_claim_data,omitempty"`
}

type TVClaimData struct {
	TVPrice         int    `json:"tv_price,omitempty"`
	CrimeRef        string `json:"crime_ref,omitempty"`
	IncidentDate    string `json:"incident_date,omitempty"`
	TVModelNo       string `json:"tv_model_no,omitempty"`
	TVMake          string `json:"tv_make,omitempty"`
	TVSerialNo      string `json:"tv_serial_no,omitempty"`
	DamageImageURL1 string `json:"damage_image_url_1,omitempty"`
	DamageImageURL2 string `json:"damage_image_url_2,omitempty"`
	TVReceiptImage  string `json:"tv_reciept_image_url"`
	DamageVideoURL  string `json:"damage_video_url,omitempty"`
}

type StromClaimData struct {
	IncidentPlace       string `json:"incident_place,omitempty"`
	IncidentDate        string `json:"incident_date,omitempty"`
	DamageImageURL1     string `json:"damage_image_url_1,omitempty"`
	DamageImageURL2     string `json:"damage_image_url_2,omitempty"`
	RepairEstimateImage string `json:"estimate_image_url,omitempty"`
	DamageVideoURL      string `json:"damage_video_url,omitempty"`
}

//Model for TYPE FORM

type TypeFormData struct {
	EventID      string       `json:"event_id"`
	EventType    string       `json:"event_type"`
	FormResponse FormResponse `json:"form_response"`
}

type FormResponse struct {
	FormID      string     `json:"form_id"`
	Token       string     `json:"token"`
	SubmittedAt time.Time  `json:"submitted_at"`
	Hidden      Hidden     `json:"hidden"`
	Definition  Definition `json:"definition"`
	Answers     []Answers  `json:"answers"`
}

type Hidden struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Policy string `json:"policy"`
}

type Definition struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Fields []Fields `json:"fields"`
}

type Fields struct {
	ID                      string `json:"id"`
	Title                   string `json:"title"`
	Type                    string `json:"type"`
	AllowMultipleSelections bool   `json:"allow_multiple_selections"`
	AllowOtherChoice        bool   `json:"allow_other_choice"`
}

type Answers struct {
	Type    string `json:"type"`
	Text    string `json:"text,omitempty"`
	Field   Field  `json:"field"`
	FileURL string `json:"file_url,omitempty"`
	Date    string `json:"date,omitempty"`
	Boolean bool   `json:"boolean,omitempty"`
	Number  int    `json:"number,omitempty"`
}

type Field struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

//Model for ZEN DESK

type TicketDetails struct {
	Ticket struct {
		Type     string `json:"type"`
		Subject  string `json:"subject"`
		Priority string `json:"priority"`
		Status   string `json:"status"`
		Comment  struct {
			HTMLBody string   `json:"html_body"`
			Uploads  []string `json:"uploads,omitempty"`
		} `json:"comment"`
		CustomFields []CustomFields `json:"custom_fields,omitempty"`
		Requester    struct {
			LocaleID     int    `json:"locale_id"`
			Name         string `json:"name"`
			Email        string `json:"email"`
			Phone        string `json:"phone"`
			PolicyNumber string `json:"policy_number"`
		} `json:"requester"`
		TicketFormID int64     `json:"ticket_form_id"`
		EventID      string    `json:"event_id"`
		Token        string    `json:"token"`
		SubmittedAt  time.Time `json:"submitted_at"`
	} `json:"ticket"`
}

type CustomFields struct {
	ID    int64  `json:"id"`
	Value string `json:"value"`
}

// WeatherAPIInput weather specific data
type WeatherAPIInput struct {
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Date    string `json:"date,omitempty"` //YYYYMMDD
}
