package try

type Handler struct {
	e *Exception
}

func (h *Handler) Catch(f func(*Exception)) {
	if h == nil {
		return
	}

	f(h.e)
}
