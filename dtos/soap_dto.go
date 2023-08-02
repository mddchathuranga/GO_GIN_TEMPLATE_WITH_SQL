package dtos

import "encoding/xml"

type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapEnv string   `xml:"xmlns:soapenv,attr"`
	Tem     string   `xml:"xmlns:tem,attr"`
	Body    SOAPBody
}

type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Add     Add
}

type Add struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Add"`
	IntA    int      `xml:"intA"`
	IntB    int      `xml:"intB"`
}

type AddJSONResponse struct {
	Result int `json:"result"`
}
