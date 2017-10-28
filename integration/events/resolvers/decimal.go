package resolvers

type decimalResolver struct {

}

func (d decimalResolver) Resolve(value interface{}) interface{} {
	if value != nil {
		return string(value.([]byte))
	}
	return value
}

