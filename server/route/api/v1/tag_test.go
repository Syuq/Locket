package v1

import (
	"testing"
)

func TestFindTagListFromLocketContent(t *testing.T) {
	tests := []struct {
		locketContent string
		want        []string
	}{
		{
			locketContent: "#tag1 ",
			want:        []string{"tag1"},
		},
		{
			locketContent: "#tag1 #tag2 ",
			want:        []string{"tag1", "tag2"},
		},
		{
			locketContent: "#tag1 #tag2 \n#tag3 ",
			want:        []string{"tag1", "tag2", "tag3"},
		},
		{
			locketContent: "#tag1 #tag2 \n#tag3 #tag4 ",
			want:        []string{"tag1", "tag2", "tag3", "tag4"},
		},
		{
			locketContent: "#tag1 #tag2 \n#tag3  #tag4 ",
			want:        []string{"tag1", "tag2", "tag3", "tag4"},
		},
		{
			locketContent: "#tag1 123123#tag2 \n#tag3  #tag4 ",
			want:        []string{"tag1", "tag2", "tag3", "tag4"},
		},
		{
			locketContent: "#tag1 http://123123.com?123123#tag2 \n#tag3  #tag4 http://123123.com?123123#tag2) ",
			want:        []string{"tag1", "tag2", "tag2)", "tag3", "tag4"},
		},
	}
	for _, test := range tests {
		result := findTagListFromLocketContent(test.locketContent)
		if len(result) != len(test.want) {
			t.Errorf("Find tag list %s: got result %v, want %v.", test.locketContent, result, test.want)
		}
	}
}
