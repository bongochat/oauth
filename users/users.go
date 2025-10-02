package users

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/bongochat/utils/resterrors"
	"github.com/go-resty/resty/v2"
)

type User struct {
	Id            int64  `json:"id"`
	AccountNumber int64  `json:"account_number"`
	CountryId     int8   `json:"country_id"`
	CountryCode   string `json:"country_code"`
	PhoneNumber   string `json:"phone_number"`
	DateCreated   string `json:"date_created"`
	IsActive      bool   `json:"is_active"`
}

type UserRegistrationRequest struct {
	CountryId   int8    `json:"country_id"`
	PhoneNumber string  `json:"phone_number"`
	Password    string  `json:"password"`
	DeviceId    string  `json:"device_id"`
	DeviceType  string  `json:"device_type"`
	DeviceModel string  `json:"device_model"`
	IPAddress   net.IP  `json:"ip_address"`
	AppVersion  string  `json:"app_version"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type UserRegistrationResponse struct {
	Message string `json:"message"`
	Result  User   `json:"result"`
	Status  int    `json:"status"`
}

type UserLoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

var (
	userHostUrl            = os.Getenv("USER_HOST_URL")
	userRegistrationApiUrl = os.Getenv("USER_REGISTRATION_API_URL")
	userLoginApiUrl        = os.Getenv("USER_LOGIN_API_URL")
	userClient             = resty.New().SetBaseURL(userHostUrl)
)

func RegisterUser(country_id int8, phone_number string, password string, deviceId string, deviceType string, deviceModel string, ipAddress net.IP, appVersion string, latitude float64, longitude float64) (*User, resterrors.RestError) {
	request := UserRegistrationRequest{
		CountryId:   country_id,
		PhoneNumber: phone_number,
		Password:    password,
		DeviceId:    deviceId,
		DeviceType:  deviceType,
		DeviceModel: deviceModel,
		IPAddress:   ipAddress,
		AppVersion:  appVersion,
		Latitude:    latitude,
		Longitude:   longitude,
	}

	response, err := userClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(userRegistrationApiUrl)

	if err != nil {
		return nil, resterrors.NewInternalServerError("Could not register", "", err)
	}

	if response == nil {
		return nil, resterrors.NewInternalServerError("Invalid client response", "", nil)
	}
	fmt.Println("TEST LOG", string(response.Body()))
	if response.StatusCode() != http.StatusCreated {
		apiErr, err := resterrors.NewRestErrorFromBytes(response.Body())
		if err != nil {
			return nil, resterrors.NewInternalServerError("Invalid rest client error interface", "", err)
		}
		return nil, apiErr
	}

	var apiResponse UserRegistrationResponse
	if err := json.Unmarshal(response.Body(), &apiResponse); err != nil {
		return nil, resterrors.NewInternalServerError("Error unmarshalling response", "", err)
	}

	// var user User
	// if err := json.Unmarshal(&apiResponse.Result, &user); err != nil {
	// 	return nil, resterrors.NewInternalServerError("error unmarshalling", "", err)
	// }
	fmt.Println("API DATA:.............", &apiResponse.Result)
	return &apiResponse.Result, nil
}

func LoginUser(phone_number string, password string) (*User, resterrors.RestError) {
	request := UserLoginRequest{
		PhoneNumber: phone_number,
		Password:    password,
	}

	response, err := userClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		Post(userLoginApiUrl)

	if err != nil {
		return nil, resterrors.NewInternalServerError("Could not login", "", err)
	}

	if response == nil {
		return nil, resterrors.NewInternalServerError("Invalid client response", "", nil)
	}

	if response.StatusCode() != http.StatusOK {
		apiErr, err := resterrors.NewRestErrorFromBytes(response.Body())
		if err != nil {
			return nil, resterrors.NewInternalServerError("Invalid rest client error interface", "", err)
		}
		return nil, apiErr
	}

	var user User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, resterrors.NewInternalServerError("error unmarshalling", "", err)
	}
	return &user, nil
}
