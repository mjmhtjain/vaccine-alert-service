package customerror

type NoRecordExists struct {
	Msg string
}

func (re *NoRecordExists) Error() string {
	return re.Msg
}

type DBQueryScanError struct {
	Msg string
}

func (re *DBQueryScanError) Error() string {
	return re.Msg
}
