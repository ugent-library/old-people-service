package models

type IdRefs map[string][]string

func (ir IdRefs) Add(key string, vals ...string) IdRefs {
	ir[key] = append(ir[key], vals...)
	return ir
}
