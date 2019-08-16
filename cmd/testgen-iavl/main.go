package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	iavlproofs "github.com/confio/proofs-iavl"
	"github.com/confio/proofs-iavl/helpers"

	"github.com/tendermint/iavl"
	cmn "github.com/tendermint/tendermint/libs/common"
	db "github.com/tendermint/tm-db"
)

/**
testgen-iavl will generate a json struct on stdout (meant to be saved to file for testdata).
this will be an auto-generated existence proof in the form:

{
	"root": "<hex encoded root hash of tree>",
	"existence": "<hex encoded protobuf marshaling of an existence proof>"
}
**/

func main() {
	tree, keys := helpers.BuildTree(400)
	root := tree.WorkingHash()

	// TODO: allow exist/nonexist, left/right/center
	key := keys[87]

	proof, err := iavlproofs.CreateMembershipProof(tree, key)
	if err != nil {
		fmt.Printf("Error: create proof: %+v\n", err)
		os.Exit(1)
	}

	binary, err := proof.Marshal()
	if err != nil {
		fmt.Printf("Error: protobuf marshal: %+v\n", err)
		os.Exit(1)
	}

	res := map[string]interface{}{
		"root":      hex.EncodeToString(root),
		"existence": hex.EncodeToString(binary),
	}
	out, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Printf("Error: json encoding: %+v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
