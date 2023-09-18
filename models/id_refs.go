package models

type IdRefs map[string][]string

func (ir IdRefs) Add(key string, vals ...string) IdRefs {
	ir[key] = append(ir[key], vals...)
	return ir
}

func (ir IdRefs) Clear() IdRefs {
	for key := range ir {
		delete(ir, key)
	}
	return ir
}

func (ir IdRefs) Dup() IdRefs {
	copyIr := make(map[string][]string, len(ir))
	for key, vals := range ir {
		copyIr[key] = append(copyIr[key], vals...)
	}
	return copyIr
}
