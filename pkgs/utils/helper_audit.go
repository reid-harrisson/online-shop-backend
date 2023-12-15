package utils

import (
	"OnlineStoreBackend/pkgs/logging"
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuditData struct {
	Description string `json:"description"`
	Code        uint   `json:"code"`
	Category    string `json:"category"`
	Type        string `json:"type"`
	Action      string `json:"action"`
	Event       string `json:"event"`
	IpAddress   string `json:"ip_address"`
	HttpRequest string `json:"http_request"`
	HttpHeader  string `json:"http_header"`
	ConsumerID  uint   `json:"consumer_id"`
	UserAgent   string `json:"user_agent"`
	RemoteHost  string `json:"remote_host"`
}

func HelperAudit(requestUrl string, c echo.Context, auditData AuditData) {

	logger, err := logging.NewLogger()
	if err != nil {
		panic(err)
	}

	request := c.Request()

	auditData.IpAddress = c.RealIP()
	auditData.UserAgent = request.UserAgent()

	remoteHost, err := net.LookupAddr(c.RealIP())
	if err != nil {
		fmt.Println("Unable to resolve hostname:", err)
		return
	}
	auditData.RemoteHost = strings.Join(remoteHost, " ")

	var protocol string
	if strings.Contains(request.Proto, "HTTPS") {
		protocol = "https"
	} else {
		protocol = "http"
	}
	auditData.HttpRequest = protocol + "://" + request.Host + request.RequestURI

	auditData.HttpHeader = "PROTOCOL: " + protocol + " / METHOD: " + request.Method + " / HTTP COOKIE: " + request.Header.Get("Cookie") + " / REFERRER: " + request.Header.Get("Referer")

	jsonStr, err := json.Marshal(auditData)
	if err != nil {
		logger.Error(err.Error())
	}

	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		logger.Error(err.Error())
	}
	defer resp.Body.Close()

	// if we want to check for a specific status code, we can do so here
	// for example, a successful request should return a 200 OK status
	if !(resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusCreated) {
		// if the status code is not 200, we should log the status code and the
		// status string, then exit with a fatal error
		errMsg := fmt.Sprintf("Audit service error: %d %s", resp.StatusCode, resp.Status)
		logger.Error(errMsg)
	}
}
