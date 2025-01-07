package config

const (
	KickExpiredUserNo  = 0
	KickExpiredUserYes = 1
)

type AD struct {
	KickExpiredUser int `mapstructure:"kick_expired_user"`
}
