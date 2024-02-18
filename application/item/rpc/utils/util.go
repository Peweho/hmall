package utils

import "hmall/application/item/rpc/types/service"

type ItemsArr []*service.Items

func (m ItemsArr) Len() int { return len(m) }

func (m ItemsArr) Less(i, j int) bool { return m[i].Id < m[j].Id }

func (m ItemsArr) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
