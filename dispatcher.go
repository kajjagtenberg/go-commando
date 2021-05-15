package commando

type CommandDispatcher struct {
	handlers map[string]CommandHandler
}

func (r *CommandDispatcher) Register(name string, shorthand string, handler CommandHandler) error {
	if r.handlers[name] != nil {
		return ErrNameAlreadyUsed
	}

	if r.handlers[shorthand] != nil {
		return ErrShorthandNameAlreadyUsed
	}

	r.handlers[name] = handler
	r.handlers[shorthand] = handler

	return nil
}

func (r *CommandDispatcher) Handle(cmd Command) (interface{}, error) {
	handler := r.handlers[cmd.Name]
	if handler == nil {
		return nil, ErrUnknownCommand
	}

	return handler(cmd)
}

func NewCommandDispatcher() *CommandDispatcher {
	return &CommandDispatcher{
		handlers: make(map[string]CommandHandler),
	}
}
