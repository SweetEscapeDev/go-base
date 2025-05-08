package initializer

import (
	"go-base/config"

	"github.com/midtrans/midtrans-go"
)

func InitializeMidtrans(c config.Config) {
	midtrans.ServerKey = c.MidtransServerKey
	midtrans.Environment = midtrans.Sandbox

	if c.ENVIRONMENT == "production" {
		midtrans.Environment = midtrans.Production
	}
}
