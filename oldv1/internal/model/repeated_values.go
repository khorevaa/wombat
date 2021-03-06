// Copyright 2020 Rogchap. All Rights Reserved.

package model

import (
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/therecipe/qt/core"
)

// The roles for RepeatedValues
const (
	RepeatedValueValue = int(core.Qt__UserRole) + 1<<iota
	RepeatedValueMsgValue
)

// RepeatedValue is a QObject that holds the value or Message value for a repeated field
type RepeatedValue struct {
	core.QObject

	_ string   `property:"value"`
	_ *Message `property:msgValue"`
}

// RepeatedValues is an QAbstractListModel that has the list of repeated values
type RepeatedValues struct {
	core.QAbstractListModel

	ref *desc.MessageDescriptor

	_ func() `constructor:"init"`

	_ map[int]*core.QByteArray `property:"roles"`
	_ []*RepeatedValue         `property:"values"`
	_ int                      `property:"count"`

	_ func()            `slot:"addValue"`
	_ func(int)         `slot:"remove"`
	_ func(int, string) `slot:"editValueAt"`
}

func (m *RepeatedValues) init() {
	m.SetRoles(map[int]*core.QByteArray{
		RepeatedValueValue:    core.NewQByteArray2("value", -1),
		RepeatedValueMsgValue: core.NewQByteArray2("msgValue", -1),
	})
	m.SetCount(0)

	m.ConnectData(m.data)
	m.ConnectRowCount(m.rowCount)
	m.ConnectColumnCount(m.columnCount)
	m.ConnectRoleNames(m.roleNames)
	m.ConnectAddValue(m.addValue)
	m.ConnectRemove(m.remove)
	m.ConnectEditValueAt(m.editValueAt)
}

func (m *RepeatedValues) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}

	if index.Row() >= len(m.Values()) {
		return core.NewQVariant()
	}

	var v = m.Values()[index.Row()]

	switch role {
	case RepeatedValueValue:
		return core.NewQVariant1(v.Value())
	case RepeatedValueMsgValue:
		return core.NewQVariant1(v.MsgValue())

	default:
		return core.NewQVariant()
	}
}

func (m *RepeatedValues) rowCount(parent *core.QModelIndex) int {
	return len(m.Values())
}

func (m *RepeatedValues) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *RepeatedValues) roleNames() map[int]*core.QByteArray {
	return m.Roles()
}

func (m *RepeatedValues) addValue() {
	m.BeginInsertRows(core.NewQModelIndex(), len(m.Values()), len(m.Values()))
	rv := NewRepeatedValue(nil)
	if m.ref != nil {
		rv.SetMsgValue(MapMessage(dynamic.NewMessage(m.ref)))
	}
	m.SetValues(append(m.Values(), rv))
	m.EndInsertRows()
	m.SetCount(m.Count() + 1)
}

func (m *RepeatedValues) remove(row int) {
	m.BeginRemoveRows(core.NewQModelIndex(), row, row)
	m.SetValues(append(m.Values()[:row], m.Values()[row+1:]...))
	m.EndRemoveRows()
	m.SetCount(m.Count() - 1)
}

func (m *RepeatedValues) editValueAt(row int, val string) {
	rv := m.Values()[row]
	if rv.Value() == val {
		return
	}
	rv.SetValue(val)
	idx := m.Index(row, 0, core.NewQModelIndex())
	m.DataChanged(idx, idx, []int{RepeatedValueValue})
}
