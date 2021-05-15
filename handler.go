package commando

type CommandHandler = func(cmd Command) (interface{}, error)
