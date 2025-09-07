package grpc

import (
	"context"
	categorypb "golang-crud/proto/categorypb"
	"golang-crud/repository"
)

type CategoryServer struct {
    categorypb.UnimplementedCategoryServiceServer
}

func (s *CategoryServer) GetAllCategories(ctx context.Context, req *categorypb.Empty) (*categorypb.CategoryList, error) {
    repo := repository.CategoryRepository{}
    categories, _ := repo.GetAll()
    var list []*categorypb.Category
    for _, c := range categories {
        list = append(list, &categorypb.Category{
            Id:   int32(c.Id),
            Name: c.Name,
        })
    }
    return &categorypb.CategoryList{Categories: list}, nil
}
