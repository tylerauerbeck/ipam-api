package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

// IPBlockType returns IPBlockTypeResolver implementation.
func (r *Resolver) IPBlockType() IPBlockTypeResolver { return &iPBlockTypeResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type iPBlockTypeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
