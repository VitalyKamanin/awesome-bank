package integrations

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/beevik/etree"
)

const cbrSoapURL = "https://www.cbr.ru/DailyInfoWebServ/DailyInfo.asmx"

func BuildSOAPRequest() string {
	fromDate := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	toDate := time.Now().Format("2006-01-02")

	return fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
<soap12:Envelope xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
  <soap12:Body>
    <KeyRate xmlns="http://web.cbr.ru/">
      <fromDate>%s</fromDate>
      <ToDate>%s</ToDate>
    </KeyRate>
  </soap12:Body>
</soap12:Envelope>`, fromDate, toDate)
}

func SendSOAPRequest(soapXML string) ([]byte, error) {
	req, err := http.NewRequest("POST", cbrSoapURL, bytes.NewBuffer([]byte(soapXML)))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Set("SOAPAction", `"http://web.cbr.ru/KeyRate"`)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	return body, nil
}

func ParseXMLResponse(xmlData []byte) (float64, error) {
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(xmlData); err != nil {
		return 0, fmt.Errorf("failed to parse XML: %v", err)
	}

	rateElement := doc.FindElement("//Rate")
	if rateElement == nil {
		return 0, errors.New("tag 'Rate' not found in response")
	}

	var rate float64
	_, err := fmt.Sscanf(rateElement.Text(), "%f", &rate)
	if err != nil {
		return 0, fmt.Errorf("failed to convert rate to float: %v", err)
	}

	return rate, nil
}

func GetCentralBankRate() (float64, error) {
	soapRequest := BuildSOAPRequest()
	rawBody, err := SendSOAPRequest(soapRequest)
	if err != nil {
		return 0, err
	}

	rate, err := ParseXMLResponse(rawBody)
	if err != nil {
		return 0, err
	}

	rate += 5.0

	return rate, nil
}
