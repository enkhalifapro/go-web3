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
 * @file personal-unlockaccount_test.go
 * @authors:
 *   Reginaldo Costa <regcostajr@gmail.com>
 * @date 2017
 */
package test

import (
	"errors"
	"testing"

	"github.com/regcostajr/go-web3"
)

func Personal_UnlockAccount(connection *web3.Web3) error {

	accounts, err := listPersonalAccounts(connection)

	if err != nil {
		return err
	}

	result, err := connection.Personal.UnlockAccount(accounts[0], "password", 100)

	if err != nil {
		return err
	}

	if !result {
		return errors.New("Can't unlock account")
	}

	return nil
}

func TestPersonal_UnlockAccount_IPCConnection(t *testing.T) {
	err := Personal_UnlockAccount(IPCConnection())
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestPersonal_UnlockAccount_HTTPConnection(t *testing.T) {
	err := Personal_UnlockAccount(HTTPConnection())
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}