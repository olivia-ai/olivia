package embeddings

import (
	"strings"

	"github.com/olivia-ai/olivia/data"
)

func tokenize(sentence string) (tokens []string) {
	tokens = strings.Fields(
		strings.ToLower(sentence),
	)

	for i, token := range tokens {
		tokens[i] = stem(token)
	}

	return
}

func stem(word string) string {
	return word
}

func appendToVocabulary(tokenBase *[]string, words ...string) {
	// TODO
}

// EstablishVocabulary takes a slice of Conversation structs and establish the vocabulary set.
func EstablishVocabulary(conversations []data.Conversation) (vocabulary []string) {
	for _, conversation := range conversations {
		// Iterate through the answer and question to avoid code duplication
		for _, sentence := range []string{conversation.Answer, conversation.Question} {
			appendToVocabulary(&vocabulary, tokenize(sentence)...)
		}
	}

	return
}

func getEmbedding(vocabulary []string, word string) (embedding []float64) {
	for _, token := range vocabulary {
		var value float64 = 0
		if token == word {
			value = 1
		}

		embedding = append(embedding, value)
	}
	
	return
}