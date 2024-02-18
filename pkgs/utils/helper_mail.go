package utils

import (
	"OnlineStoreBackend/pkgs/logging"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MailData struct {
	Name                       string `json:"name"`
	EmailTo                    string `json:"email_to"`
	EmailFrom                  string `json:"email_from"`
	EmailPretext               string `json:"email_pretext"`
	Company                    string `json:"company"`
	Phone                      string `json:"phone"`
	Subject                    string `json:"subject"`
	SourceChannel              string `json:"source_channel"`
	BodyBlock                  string `json:"body_block"`
	TargetTeam                 string `json:"target_team"`
	FirstName                  string `json:"first_name"`
	CompanyID                  uint64 `json:"company_id"`
	HeaderPosterImageUrl       string `json:"header_poster_image_url"`
	HeaderPosterSloganTitle    string `json:"header_poster_slogan_title"`
	HeaderPosterSloganSubtitle string `json:"header_poster_slogan_subtitle"`
	BodyHeading                string `json:"body_heading"`
	BodyGreeting               string `json:"body_greeting"`
	BodyCtaBtnLabel            string `json:"body_cta_btn_label"`
	BodyCtaBtnLink             string `json:"body_cta_btn_link"`
}

func HelperMail(requestURL string, c echo.Context, mailData MailData) {
	logger, err := logging.NewLogger()
	if err != nil {
		panic(err)
	}

	jsonStr, err := json.Marshal(mailData)
	if err != nil {
		logger.Error(err.Error())
	}

	responseMail, err := http.Post(requestURL, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		logger.Error(err.Error())
	}
	defer responseMail.Body.Close()

	// if we want to check for a specific status code, we can do so here
	// for example, a successful request should return a 200 OK status
	if !(responseMail.StatusCode == http.StatusOK || responseMail.StatusCode == http.StatusCreated) {
		// if the status code is not 200, we should log the status code and the
		// status string, then exit with a fatal error
		errMsg := fmt.Sprintf("Audit service error: %d %s", responseMail.StatusCode, responseMail.Status)
		logger.Error(errMsg)
	}
}
