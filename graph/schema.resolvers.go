package graph

import (
	"context"
	"fmt"

	"graphql-books/graph/model"
)

func (r *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	r.Resolver.mu.Lock()
	defer r.Resolver.mu.Unlock()
	return r.Resolver.Books, nil
}

func (r *queryResolver) Book(ctx context.Context, id string) (*model.Book, error) {
	r.Resolver.mu.Lock()
	defer r.Resolver.mu.Unlock()
	for _, b := range r.Resolver.Books {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, nil
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	r.Resolver.mu.Lock()
	defer r.Resolver.mu.Unlock()
	return r.Resolver.Authors, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model.NewBook) (*model.Book, error) {
	r.Resolver.mu.Lock()
	defer r.Resolver.mu.Unlock()

	var author *model.Author
	for _, a := range r.Resolver.Authors {
		if a.ID == input.AuthorID {
			author = a
			break
		}
	}
	if author == nil {
		return nil, fmt.Errorf("author not found: %s", input.AuthorID)
	}

	r.Resolver.counter++
	book := &model.Book{
		ID:     fmt.Sprintf("book-%d", r.Resolver.counter),
		Title:  input.Title,
		Year:   input.Year,
		Author: author,
	}
	r.Resolver.Books = append(r.Resolver.Books, book)
	author.Books = append(author.Books, book)

	for _, ch := range r.Resolver.subscribers {
		ch <- book
	}

	return book, nil
}

func (r *subscriptionResolver) BookAdded(ctx context.Context) (<-chan *model.Book, error) {
	r.Resolver.mu.Lock()
	ch := make(chan *model.Book, 1)
	r.Resolver.subscribers = append(r.Resolver.subscribers, ch)
	r.Resolver.mu.Unlock()

	go func() {
		<-ctx.Done()
		r.Resolver.mu.Lock()
		for i, sub := range r.Resolver.subscribers {
			if sub == ch {
				r.Resolver.subscribers = append(r.Resolver.subscribers[:i], r.Resolver.subscribers[i+1:]...)
				break
			}
		}
		r.Resolver.mu.Unlock()
	}()

	return ch, nil
}

type (
	queryResolver        struct{ *Resolver }
	mutationResolver     struct{ *Resolver }
	subscriptionResolver struct{ *Resolver }
)
