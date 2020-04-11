package language

import "testing"

func TestSearchPatterns(t *testing.T) {
	sentence := "Here are your posts: {posts[0].link}, {user.name}"
	SearchPatterns(sentence)
}
