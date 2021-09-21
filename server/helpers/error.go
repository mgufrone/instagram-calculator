package helpers

type FuncOrErr func() error
func TryOrError(funcs ...FuncOrErr) (err error) {
	for _, c := range funcs {
		if err = c(); err != nil {
			return
		}
	}
	return
}
