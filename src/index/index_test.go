package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestIndex() *Index {
	index := NewIndex()
	index.AddDocument("hamlet", map[string]string{
		"author": "William Shakespeare",
		"content": `To be, or not to be: that is the question:
Whether 'tis nobler in the mind to suffer
The slings and arrows of outrageous fortune,
Or to take arms against a sea of troubles,
And by opposing end them? To die: to sleep;`,
	})
	index.AddDocument("archery", map[string]string{
		"author": "Horace A. Ford",
		"content": `Mr. Ford was the founder of modern scientific archery. First by example, and then by precept, he changed what before was 'playing at bows and arrows' into a scientific pastime. He held the Champion's medal for eleven years in succession—from 1849 to 1859. He also won it again in 1867. After this time, although he was seen occasionally in the archery field, his powers began to wane. He died in the year 1880. His best scores, whether at public matches or in private practice, have never been surpassed. But, although no one has risen who can claim that on him has fallen the mantle of[vii] Mr. Ford, his work was not in vain. Thanks to the more scientific and rational principles laid down by this great archer, any active lad nowadays can, with a few months' practice, make scores which would have been thought fabulous when George III. was king.`,
	})
	return index
}

func TestGetFieldID(t *testing.T) {
	index := NewIndex()
	assert.Equal(t, fieldID(1), index.getFieldID("a"))
	assert.Equal(t, fieldID(2), index.getFieldID("b"))
	assert.Equal(t, map[fieldID]string{
		fieldID(1): "a",
		fieldID(2): "b",
	}, index.fieldIDToName)
	assert.Equal(t, map[string]fieldID{
		"a": fieldID(1),
		"b": fieldID(2),
	}, index.fieldNameToID)
	assert.Equal(t, fieldID(2), index.biggestFieldID)
}

func TestGetDocID(t *testing.T) {
	index := NewIndex()
	assert.Equal(t, docID(1), index.getDocID("a"))
	assert.Equal(t, docID(2), index.getDocID("b"))
	assert.Equal(t, map[docID]string{
		docID(1): "a",
		docID(2): "b",
	}, index.docIDToName)
	assert.Equal(t, map[string]docID{
		"a": docID(1),
		"b": docID(2),
	}, index.docNameToID)
	assert.Equal(t, docID(2), index.biggestDocID)
}

func TestAddDocument(t *testing.T) {
	index := createTestIndex()
	assert.Equal(t, []location{
		{
			document: docID(1),
			field: fieldID(2),
			offset: uint64(33),
		},
	}, index.index["question"])
	assert.Equal(t, []location{
		{
			document: docID(1),
			field: fieldID(2),
			offset: uint64(100),
		},
		{
			document: docID(2),
			field: fieldID(2),
			offset: uint64(142),
		},
	}, index.index["arrow"])
	assert.Equal(t, []location{
		{
			document: docID(1),
			field: fieldID(1),
			offset: uint64(8),
		},
	}, index.index["shakespear"])
}

func TestQuery(t *testing.T) {
	index := createTestIndex()
	assert.Equal(t,
		[]location{
			{
				document: docID(1),
				field: fieldID(2),
				offset: uint64(100),
			},
			{
				document: docID(2),
				field: fieldID(2),
				offset: uint64(142),
			},
			{
				document: docID(1),
				field: fieldID(2),
				offset: uint64(33),
			},
		},
		*index.Query("arrow questions"),
	)
}
