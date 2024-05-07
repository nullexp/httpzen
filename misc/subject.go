package misc

type Subject struct {
	AccessMode string `json:"accessMode"`
	UserId     string `json:"userId"`
	SubAccess  string `json:"subAccess"`
}

type SubjectParser interface {
	MustParseSubject(string) Subject
}

type testSubjectParser struct {
	subject Subject
}

func NewTestSubjectParser(s Subject) testSubjectParser {
	return testSubjectParser{subject: s}
}

func (ts testSubjectParser) MustParseSubject(string) Subject {
	return ts.subject
}
