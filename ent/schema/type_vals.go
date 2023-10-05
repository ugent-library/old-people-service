package schema

type TypeVals map[string][]string

func (ir TypeVals) Add(key string, vals ...string) TypeVals {
	ir[key] = append(ir[key], vals...)
	return ir
}
