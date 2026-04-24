package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	carservice "github.com/isaquecsilva/graphql/services/car"
)

func NewGraphQLHandler(cs carservice.CarServiceInterface) (*handler.Handler, error) {
	carType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Car",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type:        graphql.Int,
				Description: "Id of the car",
			},
			"brand": &graphql.Field{
				Type:        graphql.String,
				Description: "Brand of the car.",
			},
			"model": &graphql.Field{
				Type:        graphql.String,
				Description: "Car model.",
			},
			"year": &graphql.Field{
				Type:        graphql.Int,
				Description: "Year of the car.",
			},
			"price": &graphql.Field{
				Type:        graphql.Float,
				Description: "Price of the car.",
			},
		},
	})

	fields := graphql.Fields{
		"cars": &graphql.Field{
			Type: &graphql.List{
				OfType: carType,
			},
			Resolve: func(p graphql.ResolveParams) (any, error) {
				return cs.GetAllCars(p.Context)
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	})

	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:     &schema,
		GraphiQL:   true,
		Pretty:     true,
		Playground: true,
	}), nil
}
