package openpgp

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

func (o *FastOpenPGP) Verify(signature, message, publicKey string) (bool, error) {
	return o.verifyBytes(signature, strings.NewReader(message), publicKey)
}

func (o *FastOpenPGP) VerifyFile(signature, input, publicKey string) (bool, error) {
	message, err := os.Open(input)
	if err != nil {
		return false, err
	}
	defer message.Close()
	return o.verifyBytes(signature, message, publicKey)
}

func (o *FastOpenPGP) VerifyBytes(signature string, message []byte, publicKey string) (bool, error) {
	return o.verifyBytes(signature, bytes.NewReader(message), publicKey)
}

func (o *FastOpenPGP) verifyBytes(signature string, message io.Reader, publicKey string) (bool, error) {
	entityList, err := o.readPublicKeys(publicKey)
	if err != nil {
		return false, fmt.Errorf("publicKey error: %w", err)
	}
	if len(entityList) < 1 {
		return false, fmt.Errorf("publicKey error: %w", errors.New("no key found"))
	}

	sig, err := o.readSignature(signature)
	if err != nil {
		return false, fmt.Errorf("signature error: %w", err)
	}

	hash := sig.Hash.New()
	_, err = io.Copy(hash, message)
	if err != nil {
		return false, err
	}

	publicKeyItem := entityList[0].PrimaryKey
	if publicKeyItem == nil {
		return false, fmt.Errorf("publicKey error: %w", errors.New("no publicKey found"))
	}
	err = publicKeyItem.VerifySignature(hash, sig)
	if err != nil {
		return false, err
	}

	return true, nil
}
