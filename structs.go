package simple

import "time"

type (
	SimpleGetTokenResponse struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    Data   `json:"data"`
	}

	Data struct {
		AccessToken  string      `json:"access_token"`
		ExpiresIn    int64       `json:"expires_in"`
		TokenType    string      `json:"token_type"`
		Scope        interface{} `json:"scope"`
		RefreshToken interface{} `json:"refresh_token"`
		IDToken      interface{} `json:"id_token"`
	}

	SimpleGetTokenBody struct {
		GrantType string `json:"grant_type"`
		DeviceID  string `json:"device_id"`
	}

	SimpleCreateInvoiceRequest struct {
		OrderID     string `json:"order_id"`
		Total       int    `json:"total"`
		ExpireDate  string `json:"expire_date"` // yyyy-MM-dd HH:mm:ss
		CallbackUrl string `json:"callback_url"`
	}

	SimpleCreateInvoiceInput struct {
		OrderID    string `json:"order_id"`
		Total      int    `json:"total"`
		ExpireDate string `json:"expire_date"` // yyyy-MM-dd HH:mm:ss
	}

	SimpleCreateInvoiceResponse struct {
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Data    string      `json:"data"`
		Row     interface{} `json:"row"`
	}

	SimpleSendInvoiceToNumberRequest struct {
		InvoiceUUID string `json:"invoice_uuid"`
		Mobile      string `json:"mobile"`
	}

	SimpleSendInvoiceToNumberResponse struct {
		Code    string                        `json:"code"`
		Message string                        `json:"message"`
		Data    SimpleSendInvoiceToNumberData `json:"data"`
		Row     interface{}                   `json:"row"`
	}

	SimpleSendInvoiceToNumberData struct {
		InvoiceUUID    string      `json:"invoice_uuid"`
		OrderID        string      `json:"order_id"`
		UserMobile     string      `json:"user_mobile"`
		ExpirationDate time.Time   `json:"expiration_date"`
		Condition      interface{} `json:"condition"`
		Total          float64     `json:"total"`
		CbsInfo        interface{} `json:"cbs_info"`
		CbsStatus      string      `json:"cbs_status"`
		DeliveredDate  interface{} `json:"delivered_date"`
		PaidDate       interface{} `json:"paid_date"`
		SourceType     string      `json:"source_type"`
		SimpleID       string      `json:"simple_id"`
		MerchantName   string      `json:"merchant_name"`
		InvoiceStatus  string      `json:"invoice_status"`
		CallbackURL    string      `json:"callback_url"`
	}

	SimpleCallbackBody struct {
		InvoiceUUID string `json:"invoice_uuid"`
		OrderID     string `json:"order_id"`
		Status      string `json:"status"` // PAID, COMPLETED
	}

	SimpleGetInvoiceRequest struct {
		BeginDate     string `json:"begin_date"`     // yyyy-MM-dd HH:mm:ss
		InvoiceStatus string `json:"invoice_status"` // PAID, COMPLETED
		Mobile        string `json:"mobile"`
		OrderID       string `json:"order_id"`
		SimpleID      string `json:"simple_id"`
	}
)
