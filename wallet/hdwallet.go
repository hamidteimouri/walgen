package wallet

import (
	"fmt"
	"strings"
)

type HDWallet struct {
	seed       []byte
	btcChainId int
	ethChainId int
}

func NewHDWallet(mnemonic, password string, btcChainId int, ethChainId int) (*HDWallet, error) {
	mnemonic = strings.ReplaceAll(mnemonic, "\n", "")
	mnemonic = strings.ReplaceAll(mnemonic, "\r", "")

	seed, err := NewSeedFromMnemonic(mnemonic, password)
	if err != nil {
		return nil, err
	}
	return &HDWallet{seed: seed, btcChainId: btcChainId, ethChainId: ethChainId}, nil
}

func (h *HDWallet) NewWallet(symbol string, accountIndex, changeType, index int) (Wallet, error) {
	path, err := MakeBip44Path(symbol, h.btcChainId, accountIndex, changeType, index)
	if err != nil {
		return nil, err
	}

	return h.NewWalletByPath(symbol, path, SegWitNone)
}

func (h *HDWallet) NewSegWitWallet(accountIndex, changeType, index int) (Wallet, error) {
	path, err := MakeBip49Path(SymbolBtc, h.btcChainId, accountIndex, changeType, index)
	if err != nil {
		return nil, err
	}
	return h.NewWalletByPath(SymbolBtc, path, SegWitScript)
}

func (h *HDWallet) NewNativeSegWitWallet(accountIndex, changeType, index int) (Wallet, error) {
	path, err := MakeBip84Path(SymbolBtc, h.btcChainId, accountIndex, changeType, index)
	if err != nil {
		return nil, err
	}
	return h.NewWalletByPath(SymbolBtc, path, SegWitNative)
}

func (h *HDWallet) NewWalletByPath(symbol string, path string, segWitType SegWitType) (Wallet, error) {
	var w Wallet
	var err error

	switch symbol {
	case SymbolBtc:
		w, err = NewBtcWalletByPath(path, h.seed, h.btcChainId, segWitType)
	case SymbolEth:
		w, err = NewEthWalletByPath(path, h.seed, h.ethChainId)
	case SymbolBnb:
		w, err = NewEthWalletByPath(path, h.seed, ChainBsc)
	case SymbolTrx:
		w, err = NewTrxWalletByPath(path, h.seed)
	default:
		err = fmt.Errorf("symbol not supported: %s", symbol)
	}

	if err != nil {
		return nil, err
	}
	return w, nil
}
