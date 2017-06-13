/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"encoding/hex"
	"crypto/sha256"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("--- Init001 ---")
	fmt.Printf("Init called, initializing chaincode")
	// ID, Name, Date, Detail
//	var A, B string    // Entities
//	var Aval, Bval int // Asset holdings
	var err error

//	var Id, Name, Date, Detail string
//	var IdVal, NameVal, DateVal, DetailVal string
	
	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
	}

	// Initialize the chaincode
//	A = args[0]
//	Aval, err = strconv.Atoi(args[1])
//	if err != nil {
//		return nil, errors.New("Expecting integer value for asset holding")
//	}
//	B = args[2]
//	Bval, err = strconv.Atoi(args[3])
//	if err != nil {
//		return nil, errors.New("Expecting integer value for asset holding")
//	}
//	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	err = stub.PutState("Id",  []byte(args[0]))
	if err != nil {
		return nil, errors.New(" Id: ERROR! ")
	}
	fmt.Printf("@--- set Id 001 ---")

	err = stub.PutState("Name",  []byte(args[1]))
	if err != nil {
		return nil, errors.New(" Name: ERROR! ")
	}
	fmt.Printf("@--- set Name 001 ---")

	err = stub.PutState("Date",  []byte(args[2]))
	if err != nil {
		return nil, errors.New(" Date: ERROR! ")
	}
	fmt.Printf("@--- set Date 001 ---")

	err = stub.PutState("Detail",  []byte(args[3]))
	if err != nil {
		return nil, errors.New(" Detail: ERROR! ")
	}
	fmt.Printf("@--- set Detail 001 ---")

	return nil, nil
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("--- invoke001 ---")
	fmt.Printf("Running invoke")
	
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X int          // Transaction value
	var err error

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Avalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	if Bvalbytes == nil {
		return nil, errors.New("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	Aval = Aval - X
	Bval = Bval + X + 1 + 1
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return nil, err
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Printf("--- delete001 ---")
	fmt.Printf("Running delete")
	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return nil, errors.New("Failed to delete state")
	}

	return nil, nil
}

// Invoke callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("--- Invoke002 ---")
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "invoke" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

func (t* SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("--- Run001 ---")
	fmt.Printf("Run called, passing through to Invoke (same function)")
	
	// Handle different functions
	if function == "invoke" {
		// Transaction makes payment of X units from A to B
		fmt.Printf("Function is invoke")
		return t.invoke(stub, args)
	} else if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		fmt.Printf("Function is delete")
		return t.delete(stub, args)
	}

	return nil, errors.New("Received unknown function invocation")
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("--- Query001 ---")
	fmt.Printf("Query called, determining function")
    	converted := sha256.Sum256([]byte("123ABC456"))
	fmt.Printf("@@@ hash: " + hex.EncodeToString(converted[:]))

	if function != "query" {
		fmt.Printf("Function is query")
		return nil, errors.New("1113 Invalid query function name. Expecting \"query\"")
	}

	key := args[0]
//	Avalbytes, err := stub.GetState(key)
//	if err != nil {
//		jsonResp := "{\"Error\":\"Failed to get state for " + key + "\"}"
//		return nil, errors.New(jsonResp)
//	}
//	if Avalbytes == nil {
//		jsonResp := "{\"Error\":\"Nil amount for " + key + "\"}"
//		return nil, errors.New(jsonResp)
//	}
//
//	jsonResp := "{\"Name\":\"" + key + "\",\"Value\":\"" + string(Avalbytes) + "\"}"
//	fmt.Printf("Query Response:%s\n", jsonResp)
	
	fmt.Printf(key)
	return nil, nil
//	return Avalbytes, nil
}

func main() {
	fmt.Printf("--- main001 ---")
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
