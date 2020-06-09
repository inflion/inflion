package flow

type actionResults struct {
	results []ActionResult
}

func (a actionResults) isSuccess() bool {
	for _, ar := range a.results {
		if ar.ExitStatus == false {
			return false
		}
	}
	return true
}

func (a actionResults) getExitMessage() string {
	if a.isSuccess() {
		return "success"
	} else {
		return "fail"
	}
}

func (a actionResults) append(result ActionResult) actionResults {
	a.results = append(a.results, result)
	return a
}
