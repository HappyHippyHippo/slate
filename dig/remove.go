package dig

import "reflect"

// An RemoveOption modifies the def behavior of Remove. It's included for
// future functionality; currently, there are no concrete implementations.
type RemoveOption interface {
	unimplemented()
}

// Remove removes the instance of a registered service.
func (c *Container) Remove(ctor interface{}, opts ...RemoveOption) error {
	return c.scope.Remove(ctor, opts...)
}

// Remove removes the instance of a registered service.
func (s *Scope) Remove(ctor interface{}, opts ...RemoveOption) error {
	pCtor := reflect.ValueOf(ctor).Pointer()
	for idx, node := range s.nodes {
		if pCtor == reflect.ValueOf(node.ctor).Pointer() {
			for _, key := range node.resultKeys {
				for nIdx, node := range s.providers[key] {
					if pCtor == reflect.ValueOf(node.ctor).Pointer() {
						s.providers[key] = append(s.providers[key][:nIdx], s.providers[key][nIdx+1:]...)
						break
					}
				}
				delete(s.values, key)
				delete(s.decorators, key)
				delete(s.decoratedValues, key)
				delete(s.groups, key)
			}

			s.nodes = append(s.nodes[:idx], s.nodes[idx+1:]...)
			break
		}
	}
	return nil
}
