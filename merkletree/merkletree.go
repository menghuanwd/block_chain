package merkletree

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/cbergoon/merkletree"
)

type TestContent struct {
	x string
}

func (t TestContent) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

func (t TestContent) Equals(other merkletree.Content) (bool, error) {
	return t.x == other.(TestContent).x, nil
}

func CalculateMarkleRoot(hashes []string) string {
	var list []merkletree.Content

	if len(hashes) == 0 {
		return ""
	}

	for _, hash := range hashes {
		list = append(list, TestContent{x: hash})
	}

	t, err := merkletree.NewTree(list)

	if err != nil {
		fmt.Println(err.Error())
	}

	mr := t.MerkleRoot()

	vc, err := t.VerifyContent(t.MerkleRoot(), list[0])

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(vc)
	}

	return string(hex.EncodeToString(mr))
}
