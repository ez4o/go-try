package try

func ThrowOnError(err any) {
	if err != nil {
		panic(err)
	}
}
