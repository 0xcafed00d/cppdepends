package cppdep

type fileInfo struct {
	name string
	deps []*fileInfo
}

type CppDep struct {
	includePaths []string
}

func (i *CppDep) SetRoot(path string) {

}

func (i *CppDep) AddImportPath(importPath string) error {

	return nil
}

func (i *CppDep) AddCppFile(cppFile string) error {

	return nil
}

func (i *CppDep) GetCppFiles() []string {

	return nil
}

func (i *CppDep) GetAllDependencies(file string) []string {
	return nil
}

func (i *CppDep) GetFirstLevelDependencies(file string) []string {
	return nil
}

func (i *CppDep) GetAllDependants(file string) []string {
	return nil
}

func (i *CppDep) GetFirstLevelDependants(file string) []string {
	return nil
}
