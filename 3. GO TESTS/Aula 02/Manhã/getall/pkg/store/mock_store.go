package store

type MyMockStore struct {
	CalledRead  int
    CalledWrite int
	Data        interface{}
	MyNil       error
}

func (m *MyMockStore) Read(data interface{}) (interface{}, error) {
	m.CalledRead++
	return m.Data, m.MyNil
}

func (m *MyMockStore) Write(data interface{}) (interface{}, error) {
	m.CalledWrite++
	m.Data = data
	return data, m.MyNil
}