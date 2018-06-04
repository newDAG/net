package net

import (
	"time"
)

type WireBlockSignature struct {
	Index     int
	Signature string
}

type WireBody struct {
	Transactions    [][]byte
	BlockSignatures []WireBlockSignature

	SelfParentIndex      int
	OtherParentCreatorID int
	OtherParentIndex     int
	CreatorID            int

	Timestamp time.Time
	Index     int
}

type WireEvent struct {
	Body      WireBody
	Signature string
}

type SyncRequest struct {
	FromID int
	Known  map[int]int
}

type SyncResponse struct {
	FromID    int
	SyncLimit bool
	Events    []WireEvent
	Known     map[int]int
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

type EagerSyncRequest struct {
	FromID int
	Events []WireEvent
}

type EagerSyncResponse struct {
	FromID  int
	Success bool
}
