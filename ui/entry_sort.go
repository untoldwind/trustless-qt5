package ui

import "github.com/untoldwind/trustless/api"

type entryStoreNameAsc []*api.SecretEntry

func (p entryStoreNameAsc) Len() int           { return len(p) }
func (p entryStoreNameAsc) Less(i, j int) bool { return p[i].Name < p[j].Name }
func (p entryStoreNameAsc) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
