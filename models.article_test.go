package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	
	// get the all articles
	alist := getAllArticles()

	// check the length
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// check content one by one
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
		  v.ID != articleList[i].ID ||
		  v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}
