/********************************************************************************
   This file is part of go-web3.
   go-web3 is free software: you can redistribute it and/or modify
   it under the terms of the GNU Lesser General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   go-web3 is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Lesser General Public License for more details.
   You should have received a copy of the GNU Lesser General Public License
   along with go-web3.  If not, see <http://www.gnu.org/licenses/>.
*********************************************************************************/

/**
 * @file shh.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */

package shh

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/enkhalifapro/go-web3/dto"
	"github.com/enkhalifapro/go-web3/providers"
)

// SHH - The Net Module
type SHH struct {
	provider providers.ProviderInterface
}

// NewSHH - Net Module constructor to set the default provider
func NewSHH(provider providers.ProviderInterface) *SHH {
	shh := new(SHH)
	shh.provider = provider
	return shh
}

// GetVersion - Returns the current whisper protocol version.
// Reference: https://github.com/ethereum/wiki/wiki/JSON-RPC#shh_version
// Parameters:
//    - none
// Returns:
// 	  - String - The current whisper protocol version
func (shh *SHH) GetVersion() (string, error) {

	pointer := &dto.RequestResult{}

	err := shh.provider.SendRequest(pointer, "shh_version", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// AsymPost - Sends a whisper message Asymetric
func (shh *SHH) AsymPost(asymKeyID string, recipientPubKey string, topic string, payload string, ttl *big.Int) (bool, error) {

	params := make([]dto.SHHPostParameters, 1)
	params[0].TTL = ttl // utils.IntToHex(ttl)
	params[0].Topic = topic
	params[0].PubKey = recipientPubKey
	params[0].POWTarget = big.NewInt(2)
	params[0].POWTime = big.NewInt(100)
	params[0].Payload = hexutil.Encode([]byte(payload))
	params[0].Sig = asymKeyID

	pointer := &dto.RequestResult{}

	err := shh.provider.SendRequest(pointer, "shh_post", params)

	if err != nil {
		return false, err
	}

	return pointer.ToBoolean()

}

// SymPost - Sends a whisper message Asymetric
func (shh *SHH) SymPost(symKeyID string, recipientPubKey string, topic string, payload string, ttl *big.Int) (bool, error) {

	params := make([]dto.SHHPostParameters, 1)
	params[0].TTL = ttl // utils.IntToHex(ttl)
	params[0].Topic = topic
	params[0].PubKey = recipientPubKey
	params[0].POWTarget = big.NewInt(2)
	params[0].POWTime = big.NewInt(100)
	params[0].Payload = hexutil.Encode([]byte(payload))
	params[0].SymKeyID = symKeyID

	pointer := &dto.RequestResult{}

	err := shh.provider.SendRequest(pointer, "shh_post", params)

	if err != nil {
		return false, err
	}

	return pointer.ToBoolean()

}

// NewKeyPair - Generates a new public and private key pair for message decryption and encryption.
// Reference: https://github.com/ethereum/go-ethereum/wiki/Whisper-v5-RPC-API#shh_newkeypair
// Parameters:
// - none
// Returns:
// 	  - String - returns Key ID on success and an error on failure.
func (shh *SHH) NewKeyPair() (string, error) {

	pointer := &dto.RequestResult{}

	err := shh.provider.SendRequest(pointer, "shh_newKeyPair", nil)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// NewMsgFilter filters messages based on critirias like keyID and topic
func (shh *SHH) NewMsgFilter(subscribeParams *dto.SHHSubscribeParam) (string, error) {
	params := make([]dto.SHHSubscribeParam, 0)
	params = append(params, *subscribeParams)

	pointer := &dto.RequestResult{}

	err := shh.provider.SendRequest(pointer, "shh_newMessageFilter", params)

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// GetPublicKey by privateKey
func (shh *SHH) GetPublicKey(keyID string) (string, error) {

	pointer := &dto.RequestResult{}

	err := shh.provider.SendRequest(pointer, "shh_getPublicKey", []string{keyID})

	if err != nil {
		return "", err
	}

	return pointer.ToString()

}

// GetFilterMsgs Gets all messages based on provided filterID
func (shh *SHH) GetFilterMsgs(filterID string) []*dto.WhisperMsg {

	pointer := &dto.RequestResult{}

	err := shh.provider.SendRequest(pointer, "shh_getFilterMessages", []string{filterID})

	if err != nil {
		return nil
	}
	return pointer.ToWhisperMsgs()
}

// GenerateSymKeyFromPassword Generates the key from password, stores it, and returns its ID.
func (shh *SHH) GenerateSymKeyFromPassword(pwd string) (string, error) {

	pointer := &dto.RequestResult{}

	err := shh.provider.SendRequest(pointer, "shh_generateSymKeyFromPassword", []string{pwd})

	if err != nil {
		return "", err
	}
	return pointer.ToString()
}
