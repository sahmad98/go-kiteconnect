package kiteconnect

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	// "io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type KiteConnect struct {
	ApiKey       string
	ApiSecret    string
	CheckSum     string
	RequestToken string
	AccessToken  string
	LoginUrl     string
	ApiBaseUrl   string
	ErrorMsg     string
	LoginResp    LoginResponse
}

func NewKiteConnect(api_key string) *KiteConnect {
	kite := new(KiteConnect)
	kite.ApiKey = api_key
	kite.LoginUrl = "https://kite.trade/connect/login?v=3&api_key=" + kite.ApiKey
	kite.ApiBaseUrl = "https://api.kite.trade"
	return kite
}

func (k *KiteConnect) GetLoginUrl() string {
	return k.LoginUrl
}

func (k *KiteConnect) GenerateSession(request_token, api_secret string) error {
	k.RequestToken = request_token
	k.ApiSecret = api_secret
	hex_checksum := sha256.Sum256([]byte(k.ApiKey + k.RequestToken + k.ApiSecret))
	k.CheckSum = hex.EncodeToString(hex_checksum[0:32])
	log.Printf("Checksum %s, len %d, %x", k.CheckSum, len(k.CheckSum), hex_checksum)

	session := k.ApiBaseUrl + "/session/token"
	data := url.Values{}
	data.Set("api_key", k.ApiKey)
	data.Set("request_token", k.RequestToken)
	data.Set("checksum", k.CheckSum)

	client := &http.Client{}
	request, _ := http.NewRequest("POST", session, strings.NewReader(data.Encode()))
	request.Header.Add("X-Kite-Version", "3")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(request)
	if err != nil {
		log.Printf("Exception: %s", err.Error())
		return err
	}
	log.Println(resp.Status)
	// resp_data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// log.Println(string(resp_data))
	response := LoginResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		log.Println("Unable to decode")
		return err
	}

	if response.Status == "error" {
		log.Println("Parsed Error")
		log.Println(response.Message)
		log.Println(response.Type)
		return response
	}

	k.LoginResp = response
	log.Printf("Access Token: %s", k.LoginResp.Data.AccessToken)
	return nil
}

func (k *KiteConnect) GetUserProfile() LoginResponse {
}
