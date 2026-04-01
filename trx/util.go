package trx

import (
	"github.com/hamidteimouri/walgen/wallet"
	"github.com/lizc2003/gotron-sdk/pkg/address"
	"github.com/lizc2003/gotron-sdk/pkg/common"
	"math"
)

func DecodeAddress(addr string) ([]byte, error) {
	return common.DecodeCheck(addr)
}

func EncodeAddress(a []byte) string {
	return address.Address(a).String()
}

func TrxToSun(v float64) int64 {
	return int64(math.Round(v * wallet.SunPerTrx))
}

func SunToTrx(v int64) float64 {
	return float64(v) / wallet.SunPerTrx
}
