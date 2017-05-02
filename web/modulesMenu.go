package web

type Menu struct {
	ModuleDescription
	Branchres      []Branch         `json:"branches"`
	TargetVariable string           `json:"targetVariable"`
	Descendants    []MenuDescendant `json:"descendants"`
}
type MenuDescendant struct {
	BranchName string `json:"branchName"`
	Descendant Menu   `json:"descendant"`
}
type Branch struct {
	Name      string                `json:"name"`
	Capture   string                `json:"capture"`
	MLChoices []MultiLanguageChoice `json:"mlChoices"`
}
type MultiLanguageChoice struct {
	Language string `json:"language"`
	Capture  string `json:"capture"`
}
