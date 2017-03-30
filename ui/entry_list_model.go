package ui

import (
	"github.com/therecipe/qt/core"
	"github.com/untoldwind/trustless/api"
)

const entityIDRole = core.Qt__UserRole

type entryListModel struct {
	*core.QAbstractItemModel

	entries []*api.SecretEntry
}

func newEntryListModel() *entryListModel {
	m := &entryListModel{
		QAbstractItemModel: core.NewQAbstractItemModel(nil),
	}

	m.ConnectColumnCount(m.columnCount)
	m.ConnectRowCount(m.rowCount)
	m.ConnectData(m.data)
	m.ConnectIndex(m.index)

	return m
}

func (m *entryListModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (m *entryListModel) rowCount(parent *core.QModelIndex) int {
	return len(m.entries)
}

func (m *entryListModel) index(row int, column int, parent *core.QModelIndex) *core.QModelIndex {
	return m.CreateIndex2(row, column, uintptr(row))
}

func (m *entryListModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() || index.Row() < 0 || index.Row() > len(m.entries) {
		return core.NewQVariant()
	}
	entry := m.entries[index.Row()]
	switch core.Qt__ItemDataRole(role) {
	case core.Qt__DisplayRole:
		return core.NewQVariant14(entry.Name)
	case entityIDRole:
		return core.NewQVariant14(entry.ID)
	}
	return core.NewQVariant()
}

func (m *entryListModel) updateEntries(entries []*api.SecretEntry) {
	if entriesUnchanged(m.entries, entries) {
		return
	}
	m.BeginResetModel()
	m.entries = entries
	m.EndResetModel()
}

func (m *entryListModel) indexOf(lookup *api.SecretEntry) *core.QModelIndex {
	for i, entry := range m.entries {
		if entry == lookup {
			return m.CreateIndex2(i, 0, uintptr(i))
		}
	}
	return core.NewQModelIndex()
}

func entriesUnchanged(first, second []*api.SecretEntry) bool {
	if len(first) != len(second) {
		return false
	}
	for i, entry := range first {
		if second[i] != entry {
			return false
		}
	}
	return true
}
