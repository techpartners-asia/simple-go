package simple

import "encoding/json"

type simple struct {
	userName    string
	password    string
	baseUrl     string
	callbackUrl string
	loginObject *SimpleGetTokenResponse
	expire_in   *int64
}

type Simple interface {
	CreateInvoice(input SimpleCreateInvoiceInput) (SimpleCreateInvoiceResponse, error)
	SendInvoiceToNumber(input SimpleSendInvoiceToNumberRequest) (SimpleSendInvoiceToNumberResponse, error)
	GetInvoice(input SimpleGetInvoiceRequest) (SimpleSendInvoiceToNumberResponse, error)
	ChangeInvoiceStatus(input SimpleChangeInvoiceStatusRequest) (SimpleChangeInvoiceStatusResponse, error)
	// GetInvoiceByInvoiceID(invoiceID string) (SimpleInvoiceDetailResponse, error)
	// GetInvoiceByOrderNumber(orderNumber string) (SimpleInvoiceDetailResponse, error)
	Close()
}

func New(userName, password, baseUrl, callbackUrl string) Simple {

	return &simple{
		userName:    userName,
		password:    password,
		baseUrl:     baseUrl,
		callbackUrl: callbackUrl,
		loginObject: nil,
		expire_in:   nil,
	}
}

func (s *simple) CreateInvoice(input SimpleCreateInvoiceInput) (SimpleCreateInvoiceResponse, error) {
	body := SimpleCreateInvoiceRequest{
		OrderID:     input.OrderID,
		Total:       input.Total,
		ExpireDate:  input.ExpireDate,
		CallbackUrl: s.callbackUrl,
	}

	res, err := s.httpRequest(body, SimpleCreateInvoice, "")
	if err != nil {
		return SimpleCreateInvoiceResponse{}, err
	}

	var response SimpleCreateInvoiceResponse
	json.Unmarshal(res, &response)
	return response, nil

}

func (s *simple) SendInvoiceToNumber(input SimpleSendInvoiceToNumberRequest) (SimpleSendInvoiceToNumberResponse, error) {
	res, err := s.httpRequest(input, SimpleSendInvoiceToNumber, "?invoice_uuid="+input.InvoiceUUID+"&mobile="+input.Mobile)
	if err != nil {
		return SimpleSendInvoiceToNumberResponse{}, err
	}

	var response SimpleSendInvoiceToNumberResponse
	json.Unmarshal(res, &response)
	return response, nil
}

func (s *simple) GetInvoice(input SimpleGetInvoiceRequest) (SimpleSendInvoiceToNumberResponse, error) {
	res, err := s.httpRequest(nil, SimpleGetInvoice, "?order_id="+input.OrderID+"&simple_id="+input.SimpleID)
	if err != nil {
		return SimpleSendInvoiceToNumberResponse{}, err
	}

	var response SimpleSendInvoiceToNumberResponse
	json.Unmarshal(res, &response)
	return response, nil
}

func (s *simple) ChangeInvoiceStatus(input SimpleChangeInvoiceStatusRequest) (SimpleChangeInvoiceStatusResponse, error) {
	res, err := s.httpRequest(input, SimpleChangeInvoiceStatus, "")
	if err != nil {
		return SimpleChangeInvoiceStatusResponse{}, err
	}

	var response SimpleChangeInvoiceStatusResponse
	json.Unmarshal(res, &response)
	return response, nil
}

func (s *simple) Close() {
	s.loginObject = nil
	s.expire_in = nil
}
