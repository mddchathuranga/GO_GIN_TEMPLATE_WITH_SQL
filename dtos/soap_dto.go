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

type SOAPResponse struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	SoapEnv string   `xml:"xmlns:soapenv,attr"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Body    SOAPResponseBody
}

type SOAPResponseBody struct {
	XMLName     xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	AddResponse AddResponse
}

type AddResponse struct {
	XMLName   xml.Name `xml:"http://tempuri.org/ AddResponse"`
	AddResult int      `xml:"AddResult"`
}
type AddJSONResponse struct {
	Result int `json:"result"`
}
