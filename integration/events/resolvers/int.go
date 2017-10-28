package resolvers

type intResolver struct {

}

func (d intResolver) Resolve(value interface{}) interface{}{
	return value
}
