package container

type IntVal int

var _ Item = (*IntVal)(nil)

// Compare IntVal 为最小堆
func (i IntVal) Compare(i2 interface{}) CompareRet {
	if i > i2.(IntVal) {
		return CompareGt
	} else if i < i2.(IntVal) {
		return CompareLt
	} else {
		return CompareEqual
	}
}
