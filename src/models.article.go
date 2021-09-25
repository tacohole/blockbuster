package main

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	{ID: 1, Title: "article1", Content: "some content"},
	{ID: 2, Title: "article2", Content: "some more content"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleById(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

func createNewArticle(title, content string) (*article, error) {
	a := article{ID: (len(articleList) + 1), Title: title, Content: content}

	articleList = append(articleList, a)

	return &a, nil
}
