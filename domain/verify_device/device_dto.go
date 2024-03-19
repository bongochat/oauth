package verify_device

type VerifyDevice struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id,omitempty"`
	DeviceId    string `json:"device_id"`
	IsVerified  bool   `json:"is_verified"`
}
