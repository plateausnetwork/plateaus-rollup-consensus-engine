package merkletree

import (
	"golang.org/x/crypto/sha3"
)

var Sha256Func = func(bytes []byte) ([]byte, error) {
	hash := sha3.New256()
	hash.Write(bytes)

	return hash.Sum(nil), nil
}

//TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type TestContent struct {
	X string
}

//Serialize hashes the values of a DataBlock
func (t TestContent) Serialize() ([]byte, error) {
	return []byte(t.X), nil
}

func (t TestContent) Hash() ([]byte, error) {
	data, _ := t.Serialize()

	return Sha256Func(data)
}
