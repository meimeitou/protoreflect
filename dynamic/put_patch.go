package dynamic

import (
	"github.com/jhump/protoreflect/desc"
)

func (m *Message) PutValueField(fd *desc.FieldDescriptor, val interface{}) {
	if err := m.TryPutValueField(fd, val); err != nil {
		panic(err.Error())
	}
}

func (m *Message) TryPutValueField(fd *desc.FieldDescriptor, val interface{}) error {
	if err := m.checkField(fd); err != nil {
		return err
	}
	return m.putValueField(fd, val)
}

func (m *Message) putValueField(fd *desc.FieldDescriptor, val interface{}) error {
	_, ok := m.values[fd.GetNumber()]
	if !ok && m.values !=nil {
		m.values[fd.GetNumber()] = val
	}
	return nil
}
