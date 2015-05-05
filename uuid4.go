package uuid4

import (
	"crypto/rand"
	"fmt"
)

type UUID struct {
	abytes  []byte
	bbytes  []byte
	cbytes  []byte
	dbytes  []byte
	ebytes  []byte
	storage [16]byte
}

func NewUUID() *UUID {
	uuid := &UUID{}
	uuid.abytes = uuid.storage[0:4]
	uuid.bbytes = uuid.storage[4:6]
	uuid.cbytes = uuid.storage[6:8]
	uuid.dbytes = uuid.storage[8:10]
	uuid.ebytes = uuid.storage[10:16]

	rand.Read(uuid.storage[:])

	uuid.cbytes[0] = uuid.cbytes[0]&0xf | 0x40
	uuid.dbytes[0] = uuid.dbytes[0]&0x3f | 0x80

	return uuid
}

func (u *UUID) DashString() string {
	return fmt.Sprintf("%8x-%4x-%4x-%4x-%12x",
		u.abytes,
		u.bbytes,
		u.cbytes,
		u.dbytes,
		u.ebytes)
}

func (u *UUID) HexString() string {
	return fmt.Sprintf("%8x%4x%4x%4x%12x",
		u.abytes,
		u.bbytes,
		u.cbytes,
		u.dbytes,
		u.ebytes)
}

func (u *UUID) Bytes() []byte {
	return u.storage[:]
}

func DashString() string {
	u := NewUUID()
	return u.DashString()
}
