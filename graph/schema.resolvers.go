package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	test "example.com/m/Test"
	"example.com/m/graph/generated"
	"example.com/m/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Upsertstudent(ctx context.Context, input *model.Studententer) (*model.StudentResponse, error) {
	//	panic(fmt.Errorf("not implemented"))
	result, err := test.TableInsert(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Addition(ctx context.Context, input *model.Sum) (*model.AdditionResponse, error) {
	//	panic(fmt.Errorf("not implemented"I
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.AdditionResponse
	num3 = num1 + num2
	mod = model.AdditionResponse{Number3: &num3}
	return &mod, nil
}

func (r *queryResolver) Substraction(ctx context.Context, input *model.Sub) (*model.SubResponse, error) {
	//	panic(fmt.Errorf("not implemented"))
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.SubResponse
	num3 = num1 - num2
	mod = model.SubResponse{Number3: &num3}
	return &mod, nil
}

func (r *queryResolver) Multiplication(ctx context.Context, input *model.Product) (*model.MulResponse, error) {
	//panic(fmt.Errorf("not implemented"))
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.MulResponse
	num3 = num1 * num2
	mod = model.MulResponse{Number3: &num3}
	return &mod, nil
}

func (r *queryResolver) Division(ctx context.Context, input *model.Div) (*model.DivResponse, error) {
	//panic(fmt.Errorf("not implemented"))
	var num1 int
	var num2 int
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	var mod model.DivResponse
	num3 = num1 / num2
	mod = model.DivResponse{Number3: &num3}
	return &mod, nil
}

func (r *queryResolver) Calculator(ctx context.Context, input *model.Calculation) (*model.CalculateResponse, error) {
	//panic(fmt.Errorf("not implemented"))
	var num1 int
	var num2 int
	var operate string
	var num3 int
	num1 = *input.Number1
	num2 = *input.Number2
	operate = *input.Opertor
	var mod model.CalculateResponse
	if operate == "+" {
		num3 = num1 + num2
	}
	if operate == "-" {
		num3 = num1 - num2
	}
	if operate == "*" {
		num3 = num1 * num2
	}
	if operate == "*" {
		num3 = num1 * num2
	}
	if operate == "/" {
		if num2 == 0 {
			return nil, errors.New("can't divide by zero")
		}
		num3 = num1 / num2
	}

	mod = model.CalculateResponse{Number3: &num3}
	return &mod, nil
}

func (r *queryResolver) Studentdata(ctx context.Context, input *model.DataFetch) ([]*model.Studentoutput, error) {
	//	panic(fmt.Errorf("not implemented"))
	result, err := test.Fetch(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (r *queryResolver) Exceldata(ctx context.Context) (*model.Exeloutput, error) {
	//panic(fmt.Errorf("not implemented"))
	result, err := test.Exfetch(ctx)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *queryResolver) Dbexcel(ctx context.Context, input *model.DbInput) (*model.Response, error) {
	//panic(fmt.Errorf("not implemented"))
	result, err := test.Insertexcel(ctx, input)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *queryResolver) Chartexcel(ctx context.Context, input *model.Chartenter) (*model.ChartResponse, error) {
	//panic(fmt.Errorf("not implemented"))
	result, err := test.Enterexcel(ctx, input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *queryResolver) Piedata(ctx context.Context) (*model.PieResponse, error) {
	//	panic(fmt.Errorf("not implemented"))
	result, err := test.Enterpie(ctx)
	if err != nil {
		return nil, err
	}
	return result, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
