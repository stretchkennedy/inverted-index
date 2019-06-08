package index

import (
	"strings"

	"github.com/kljensen/snowball"
)

type fieldID uint8
type docID uint32

type location struct {
	document docID
	field fieldID
	offset uint64
}

type Index struct {
	fieldIDToName map[fieldID]string
	fieldNameToID map[string]fieldID
	biggestFieldID fieldID

	docIDToName map[docID]string
	docNameToID map[string]docID
	biggestDocID docID

	index map[string][]location
}

func NewIndex() *Index {
	return &Index{
		fieldIDToName: map[fieldID]string{},
		fieldNameToID: map[string]fieldID{},

		docIDToName: map[docID]string{},
		docNameToID: map[string]docID{},

		index: map[string][]location{},
	}
}

func (i *Index) Query(terms string) {

}

func (i *Index) AddDocument(docName string, fields map[string]string) {
	dID := i.getDocID(docName)
	for fieldName, content := range fields {
		fID := i.getFieldID(fieldName)
		// tokenise word by word
		scanner := newScanner(&content)
		for scanner.Scan() {
			i.indexWord(
				dID, fID, uint64(scanner.Start()),
				scanner.Text(),
			)
		}
	}
}

func (i *Index) indexWord(
	dID docID,
	fID fieldID,
	offset uint64,
	word string,
) {
	stem, err := stemWord(word)
	if err != nil {
		panic(err)
	}
	if stem == "" {
		return
	}
	i.index[stem] = append(i.index[stem], location{
		field: fID,
		document: dID,
		offset: offset,
	})
}

func (i *Index) getFieldID(name string) fieldID {
	id, ok := i.fieldNameToID[name]
	if ok {
		return id
	} else {
		i.biggestFieldID += 1
		i.fieldNameToID[name] = i.biggestFieldID
		i.fieldIDToName[i.biggestFieldID] = name
		return i.biggestFieldID
	}
}

func (i *Index) getDocID(name string) docID {
	id, ok := i.docNameToID[name]
	if ok {
		return id
	} else {
		i.biggestDocID += 1
		i.docNameToID[name] = i.biggestDocID
		i.docIDToName[i.biggestDocID] = name
		return i.biggestDocID
	}
}

func stemWord(word string) (string, error) {
	lower := strings.ToLower(word)
	switch lower {
	case "a", "about", "above", "after", "again", "against", "all", "am", "an",
		"and", "any", "are", "as", "at", "be", "because", "been", "before",
		"being", "below", "between", "both", "but", "by", "can", "did", "do",
		"does", "doing", "don", "down", "during", "each", "few", "for", "from",
		"further", "had", "has", "have", "having", "he", "her", "here", "hers",
		"herself", "him", "himself", "his", "how", "i", "if", "in", "into", "is",
		"it", "its", "itself", "just", "me", "more", "most", "my", "myself",
		"no", "nor", "not", "now", "of", "off", "on", "once", "only", "or",
		"other", "our", "ours", "ourselves", "out", "over", "own", "s", "same",
		"she", "should", "so", "some", "such", "t", "than", "that", "the", "their",
		"theirs", "them", "themselves", "then", "there", "these", "they",
		"this", "those", "through", "to", "too", "under", "until", "up",
		"very", "was", "we", "were", "what", "when", "where", "which", "while",
		"who", "whom", "why", "will", "with", "you", "your", "yours", "yourself",
		"yourselves":
		return "", nil
	}
	return snowball.Stem(lower, "english", false)
}
