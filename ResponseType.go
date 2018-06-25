package kiteconnect

type Response struct {
	Status  string `json:"status"`
	Data    string `json:"data"`
	Message string `json:"message"`
	Type    string `json:"error_type"`
}

type LoginResponse struct {
	Status  string    `json:"status"`
	Data    LoginData `json:"data"`
	Message string    `json:"message"`
	Type    string    `json:"error_type"`
}

type LoginData struct {
	UserId        string   `json:"user_id"`
	UserName      string   `json:"user_name"`
	UserShortName string   `json:"user_shortname"`
	Email         string   `json:"email"`
	UserType      string   `json:"user_type"`
	Broker        string   `json:"broker"`
	Exchanges     []string `json:exchanges`
	Products      []string `json:products`
	OrderTypes    []string `json:"order_types"`
	ApiKey        string   `json:"api_key"`
	AccessToken   string   `json:"access_token"`
	PublicToken   string   `json:"public_token"`
	LoginTime     string   `json:"login_time"`
}

func (err Response) Error() string {
	return err.Message
}

func (err LoginResponse) Error() string {
	return err.Message
}
