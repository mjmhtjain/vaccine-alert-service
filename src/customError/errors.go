package customerror

type RecordExists struct {
	msg string
}

func (re *RecordExists) Error() string {
	return re.msg
}
