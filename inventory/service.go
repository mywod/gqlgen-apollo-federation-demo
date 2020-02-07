// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package inventory

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/plugin/federation/fedruntime"
)

func (ec *executionContext) __resolve__service(ctx context.Context) (fedruntime.Service, error) {
	if ec.DisableIntrospection {
		return fedruntime.Service{}, errors.New("federated introspection disabled")
	}
	return fedruntime.Service{
		SDL: `type Product @key(fields: "upc") @extends {
	upc: String! @external
	weight: Int @external
	price: Int @external
	inStock: Boolean
	shippingEstimate: Int @requires(fields: "price weight")
}
`,
	}, nil
}

func (ec *executionContext) __resolve_entities(ctx context.Context, representations []map[string]interface{}) ([]fedruntime.Entity, error) {
	list := []fedruntime.Entity{}
	for _, rep := range representations {
		typeName, ok := rep["__typename"].(string)
		if !ok {
			return nil, errors.New("__typename must be an existing string")
		}
		switch typeName {

		case "Product":
			id0, err := ec.unmarshalNString2string(ctx, rep["upc"])
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Field %s undefined in schema.", "upc"))
			}

			entity, err := ec.resolvers.Entity().FindProductByUpc(ctx,
				id0)
			if err != nil {
				return nil, err
			}

			entity.Price, err = ec.unmarshalOInt2ᚖint(ctx, rep["price"])
			if err != nil {
				return nil, err
			}

			entity.Weight, err = ec.unmarshalOInt2ᚖint(ctx, rep["weight"])
			if err != nil {
				return nil, err
			}

			list = append(list, entity)

		default:
			return nil, errors.New("unknown type: " + typeName)
		}
	}
	return list, nil
}
