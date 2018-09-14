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
 * @file shh-post.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */

package dto

import "math/big"

/* type SHHPostParameters struct {
	From     string   `json:"from"`
	To       string   `json:"to"`
	Topics   []string `json:"topics"`
	Payload  string   `json:"payload"`
	Priority string   `json:"priority"`
	TTL      *big.Int `json:"ttl"`
} */

type SHHPostParameters struct {
	SymKeyID   string   `json:"symKeyID"`
	PubKey     string   `json:"pubKey"`
	Sig        string   `json:"sig"`
	TTL        *big.Int `json:"ttl"`
	Topic      string   `json:"topic"`
	Payload    string   `json:"payload"`
	Padding    string   `json:"padding"`
	POWTime    *big.Int `json:"powTime"`
	POWTarget  *big.Int `json:"powTarget"`
	TargetPeer string   `json:"targetPeer"`
}

// SHHSubscribeParam contains the required parameters of shh_subscribe
type SHHSubscribeParam struct {
	SymKeyID     string   `json:"symKeyId"`
	PrivateKeyID string   `json:"privateKeyId"`
	Sig          string   `json:"sig"`
	MinPOW       int      `json:"minPow"`
	Topics       []string `json:"topics"`
	AllowP2P     bool     `json:"allowP2P"`
}

type WhisperMsg struct {
	Sig                string `json:"sig"` // public key who signed this message
	TTL                int    `json:"ttl"`
	Timestamp          int    `json:"timestamp"`
	Topic              string `json:"topic"`
	Payload            string `json:"payload"` // Decrypted payload
	Padding            string `json:"padding"` // Optional padding (byte array of arbitrary length).
	POW                int    `json:"pow"`
	Hash               string `json:"hash"`
	RecipientPublicKey string `json:"recipientPublicKey"`
}
