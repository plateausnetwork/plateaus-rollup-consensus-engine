package merkletree

import (
	"github.com/rhizomplatform/plateaus-rollup-consensus-engine/internal/hash"
	"github.com/txaty/go-merkletree"
	"log"
)

type HashGenerator struct{}

func NewFactory() hash.Generator {
	return &HashGenerator{}
}

var config = &merkletree.Config{
	HashFunc:     Sha256Func,
	NoDuplicates: false,
}

func (t *HashGenerator) GenerateByCollection(elements *[]string) (hash.Hash, error) {
	var list []merkletree.DataBlock

	for _, v := range *elements {
		// TODO: decouple TestContent
		list = append(list, TestContent{X: v})
	}

	if len(list) == 1 {
		list = append(list, TestContent{X: ""})
	}

	tree, err := merkletree.New(config, list)

	if err != nil {
		log.Printf("could not create a merkle tree based on txs: %s", err)
		return nil, err
	}

	return tree.Root, nil
}

func (t *HashGenerator) GenerateByMap(elements *map[string]string) (*map[string]string, error) {
	newElements := *elements

	for i, v := range newElements {
		var hashContent hash.Hash
		var err error

		hashContent, err = TestContent{X: v}.Hash()

		if err != nil {
			log.Printf("could not calculate hash by Tx.RawLog: %s", err)
			return nil, err
		}

		newElements[i] = hashContent.String()
	}

	return &newElements, nil
}

func (t *HashGenerator) generate(elements string) (hash.Hash, error) {
	return TestContent{X: elements}.Serialize()
}
