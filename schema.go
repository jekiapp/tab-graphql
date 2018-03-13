package main

import "github.com/graphql-go/graphql"

// example query:
//	{
//		get_shop(shop_id:123){
//			shop_id
//			shop_name
//			products{
//				product_id
//				product_name
//			}
//		}
//	}

// create the root query for the whole node
func createSchema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"get_shop": getShopField,
			},
		}),
	})

	if err != nil {
		panic(err)
	}

	return schema
}

var getShopField = &graphql.Field{
	Name: "get_shop",
	Type: shop,
	Args: graphql.FieldConfigArgument{
		"shop_id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		id := p.Args["shop_id"].(int)
		return getShopByID(id)
	},
}

// we use this struct to represent data from database
type Shop struct {
	ShopID   int
	ShopName string
}

func getShopByID(id int) (Shop, error) {
	return Shop{
		ShopID:   id,
		ShopName: "Tokopedia",
	}, nil
}

var shop = func() *graphql.Object {
	fields := graphql.Fields{}

	fields["shop_id"] = &graphql.Field{
		Name: "ShopID",
		Type: graphql.Int,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			shop := p.Source.(Shop)
			return shop.ShopID, nil
		},
	}

	fields["shop_name"] = &graphql.Field{
		Name: "ShopName",
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			shop := p.Source.(Shop)
			return shop.ShopName, nil
		},
	}

	fields["products"] = &graphql.Field{
		Name: "Products",
		Type: graphql.NewList(product),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			shopId := p.Source.(Shop).ShopID
			return getProductsByShop(shopId)
		},
	}

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Shop",
		Fields: fields,
	})
}()

// this struct can represent data for the field also from database
// we do that with json tag and BindFields method
type Product struct {
	ProductID   int    `json:"product_id"`
	ProductName string `json:"product_name"`
}

func getProductsByShop(shopId int) ([]Product, error) {
	return []Product{
		{
			ProductID:   101,
			ProductName: "Sepatu",
		},
		{
			ProductID:   102,
			ProductName: "Celana",
		},
	}, nil
}

var product = func() *graphql.Object {
	fields := graphql.BindFields(Product{})

	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Product",
		Fields: fields,
	})
}()
