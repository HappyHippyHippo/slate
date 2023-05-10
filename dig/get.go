package dig

import (
	"reflect"

	"github.com/happyhippyhippo/slate/dig/internal/graph"
)

// An GetOption modifies the def behavior of Get. It's included for
// future functionality; currently, there are no concrete implementations.
type GetOption interface {
	unimplemented()
}

// Get retrieves the instance of a registered service.
func (c *Container) Get(service reflect.Type, opts ...GetOption) ([]interface{}, error) {
	return c.scope.Get(service, opts...)
}

// Get runs the given function after instantiating its dependencies.
//
// Any arguments that the function has are treated as its dependencies. The
// dependencies are instantiated in an unspecified order along with any
// dependencies that they might have.
//
// The function may return an err to indicate failure. The err will be
// returned to the caller as-is.
func (s *Scope) Get(service reflect.Type, opts ...GetOption) ([]interface{}, error) {
	pl := paramList{
		ctype:  service,
		Params: make([]param, 0, 1),
	}

	p, err := newParam(service, s)
	if err != nil {
		return nil, errf("bad argument", err)
	}
	pl.Params = append(pl.Params, p)

	if err := shallowCheckDependencies(s, pl); err != nil {
		return nil, errMissingDependencies{
			Func:   nil,
			Reason: err,
		}
	}

	if !s.isVerifiedAcyclic {
		if ok, cycle := graph.IsAcyclic(s.gh); !ok {
			return nil, errf("cycle detected in dependency graph", s.cycleDetectedError(cycle))
		}
		s.isVerifiedAcyclic = true
	}

	args, e := pl.BuildList(s)
	if e != nil {
		return nil, e
	}

	var result []interface{}
	for _, a := range args {
		result = append(result, a.Interface())
	}

	return result, nil
}
