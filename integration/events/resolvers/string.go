package resolvers


type stringResolver struct {

}

func (d stringResolver) Resolve(value interface{}) interface{}{
	return value
}


