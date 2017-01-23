/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"bytes"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

var key string ;
var value string ;


   
var counter int = 0;
var stringvalues []string;



func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("abc", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	} else if function == "addLoc" {
		fmt.Println("**** First argument in addLoc:****" + args[0])
		return t.addLoc(stub, args)
	} else if function == "updateLocStatus" {
		fmt.Println("**** First argument in updateLocStatus:****" + args[0])
		return t.updateLocStatus(stub, args)
	} else if function == "uploadBol" {
		fmt.Println("**** First argument in uploadBol:****" + args[0])
		return t.uploadBol(stub, args)
	} else if function == "uploadContract" {
		fmt.Println("**** First argument in uploadContract:****" + args[0])
		return t.uploadContract(stub, args)
	}
	
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation")
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	} else if function == "getLoc" {
	//	i,err := strconv.Atoi(args[0])
	//	fmt.Println(err); 
		return t.getLoc(stub, args);
		 
	} else if function == "getList" {
	
		return t.getLocList(stub, args);
	} else if function == "getNumberOfLocs" {
	
		return t.getNumberOfLocs(stub, args);
	} 
	
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query")
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub ChaincodeStubInterface, args []string) ([]byte, error) {
	
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	
	fmt.Println("saving state for key: " + key);
	
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	
	
	return nil, nil
}

// Adding LOCs 
func (t *SimpleChaincode) addLoc(stub ChaincodeStubInterface, args []string) ([]byte,error){
  var err error;
  var counter1 int;
  
    valAsbytes,err :=stub.GetState(strconv.Itoa(counter))
    s:=string(valAsbytes);
	
     if len(s) != 0 {
	     lastByByte := s[len(s)-1:]
             counter1, err =  strconv.Atoi(lastByByte)
 		if err != nil {
     			return  nil,err
  	         }
	
   	  } else {
             counter1 = 0
    	   }
   
     counter = counter1+1;
    
     counter_s := strconv.Itoa(counter)
     stringvalues = append(args,counter_s)//string array (value)
     s_requester := counter_s //counter value(key)

     stringByte := strings.Join(stringvalues , "|") // x00 = null
     
      err = stub.PutState(s_requester, []byte(stringByte));

      if err != nil {
		return nil, err
	}
	
   
                return nil, nil
}

// Return specific LOC in the system
    func (t *SimpleChaincode) getLoc(stub ChaincodeStubInterface , args []string) ([]byte,error) {
     
    	loc_string,err :=stub.GetState(args[0])
	
    	if err != nil {
		return nil, err
	}
    	
    	s := []string{string(loc_string)};
        
       // s=[]string{string(contract_hash_string),string(bol_hash_string)};
        final_string := strings.Join(s, "|");
    	
    	
    	
        //s := strconv.Itoa(counter) ;
        //ret_s := []byte(s);
        return []byte(final_string), nil;
        
    }


 //Get number of LOCs in the system
    func (t *SimpleChaincode) getNumberOfLocs(stub ChaincodeStubInterface, args []string) ([]byte, error){
        valAsbytes:=strconv.Itoa(counter);
        return []byte(valAsbytes), nil;
    }

    func (t *SimpleChaincode) updateLocStatus(stub *shim.ChaincodeStub, args []string) ([]byte, error){
    	
    	var data string;
    	value , err :=stub.GetState(args[0]);
    		if err != nil {
		return nil, err
	}
	
	for i, data1 := range bytes.Split(value, []byte{0}) { //split by white space
		fmt.Printf("Index%d :  %s\n", i, string(data1));
		 data=string(data1);
	}
	s := strings.Split(data, "|");
	s[4] = args[1];
	stringByte := strings.Join(s, "|") 
	
	err = stub.PutState(args[0], []byte(stringByte));

	  if err != nil {
		return nil, err
		}

	
	return []byte(stringByte), nil;
    }

 //upload Bol document
  func (t *SimpleChaincode) uploadBol(stub ChaincodeStubInterface, args []string) ([]byte, error){
      
       var data string;
	valueAsBytes , err :=stub.GetState(args[0]);

	if err != nil {
		return nil,err
	}
	
	for i, data1 := range bytes.Split(valueAsBytes, []byte{0}) { //split by white space
		fmt.Printf("Index%d :  %s\n", i, string(data1))
		 data=string(data1)
	}
	s := strings.Split(data, "|");
	s[9] = args[1]
	s[4] = args[2];
	stringAsByte := strings.Join(s, "|");
	
	 err = stub.PutState(args[0], []byte(stringAsByte));

      		if err != nil {
		return nil, err
		}

       	return []byte(stringAsByte), nil;
    }
    
    
    // get complete loc list
    
func (t *SimpleChaincode) getLocList(stub ChaincodeStubInterface, args []string) ([]byte, error) {
	var list []string;
	
	for i := 1; i <=counter; i++ {
	 valueAsBytes , err := stub.GetState(strconv.Itoa(i));
	if err != nil {
	 return nil,err	
	}
	  s:=string(valueAsBytes);
	  list =append(list,s);
	}

	stringByte := strings.Join(list, ",");
	
	return []byte(stringByte), nil;
}

 //upload contract document
 func (t *SimpleChaincode) uploadContract(stub ChaincodeStubInterface, args []string) ([]byte, error){
      
       var contract string;
	valueAsBytes , err :=stub.GetState(args[0]);

	if err != nil {
		return nil,err
	}
	
	for i, data := range bytes.Split(valueAsBytes, []byte{0}) { //split by white space
		fmt.Printf("Index%d :  %s\n", i, string(data))
		 contract=string(data)
	}
	s := strings.Split(contract, "|");
	s[7] = args[1]
	s[4] = args[2];
	stringAsByte := strings.Join(s, "|");
	
	 err = stub.PutState(args[0], []byte(stringAsByte));

      		if err != nil {
		return nil, err
		}

       	return []byte(stringAsByte), nil;
    }
    
    
// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	
	fmt.Println("retrieving state for key: " + key);
	
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

        //valAsbytes = []byte(valAsbytes);
	return valAsbytes, nil
}
