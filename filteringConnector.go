package radio

import "errors"

// NewFilteringConnector constructs new filtering connector
func NewFilteringConnector(users []string, next Connector) (*FilteringConnector, error) {
	m := map[User]bool{}

	for _, u := range users {
		if obj, err := ParseUser(u); err != nil {
			return nil, err
		} else {
			m[*obj] = true
		}
	}

	return &FilteringConnector{
		Next:    next,
		Allowed: m,
	}, nil
}

// FilteringConnector is an adapter over other connector, that filters users
type FilteringConnector struct {
	Next    Connector
	Allowed map[User]bool
}

// Invoke checks that user is in allowed list and invokes command
func (f FilteringConnector) Invoke(context Context) error {
	if len(f.Allowed) > 0 {
		if a, b := f.Allowed[context.GetUser()]; a && b {
			return f.Next.Invoke(context)
		}
	}

	err := errors.New("access forbidden for " + context.GetUser().String())
	context.SendMessage(err)
	return err
}
