package graph

import (
	"sync"

	"graphql-books/graph/model"
)

type Resolver struct {
	Books       []*model.Book
	Authors     []*model.Author
	mu          sync.Mutex
	counter     int
	subscribers []chan *model.Book
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Subscription() SubscriptionResolver {
	return &subscriptionResolver{r}
}

func (r *Resolver) SeedData() {
	author1 := &model.Author{ID: "author-1", Name: "Alan Donovan"}
	author2 := &model.Author{ID: "author-2", Name: "Katherine Cox-Buday"}

	book1 := &model.Book{ID: "book-1", Title: "The Go Programming Language", Year: 2015, Author: author1}
	book2 := &model.Book{ID: "book-2", Title: "The Practice of Programming", Year: 1999, Author: author1}
	book3 := &model.Book{ID: "book-3", Title: "Concurrency in Go", Year: 2017, Author: author2}

	author1.Books = []*model.Book{book1, book2}
	author2.Books = []*model.Book{book3}

	r.Authors = []*model.Author{author1, author2}
	r.Books = []*model.Book{book1, book2, book3}
	r.counter = 3
}
