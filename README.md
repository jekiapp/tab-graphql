# Tokopedia Tech a Break : Graphql In Go

Example of basic serving graphql using Go

## Run

	go build -o gql-server && ./gql-server 


## Test

	curl -d '{ get_shop(shop_id:123){ shop_id, shop_name, products{ product_id, product_name } } }' localhost:8080/graphql