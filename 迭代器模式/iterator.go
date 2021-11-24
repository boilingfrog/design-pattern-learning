package 迭代器模式

type Iterator interface {
	HasNext() bool
	Next() string
}

type names []string

func (na names) NewIterator() *NameRepository {
	return &NameRepository{
		index: 0,
		names: na,
	}
}

type NameRepository struct {
	index int
	names names
}

func (nr *NameRepository) HasNext() bool {
	if nr.index < len(nr.names) {
		return true
	}
	return false
}

func (nr *NameRepository) Next() string {
	if nr.HasNext() {
		name := nr.names[nr.index]
		nr.index++
		return name
	}

	return ""
}
