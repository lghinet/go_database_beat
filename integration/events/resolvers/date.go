package resolvers

type dateResolver struct {

}

func (d dateResolver) Resolve(value interface{}) interface{}{
	return value
}
