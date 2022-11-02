package us_phone_generator

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"math/rand"

	crand "crypto/rand"
)

func randomFromSlice[V numeric | string](slice []V) V {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	return slice[rand.Intn(len(slice))]
}

func removeFromSlice[T comparable](l []T, item T) []T {
	for i, other := range l {
		if other == item {
			return append(l[:i], l[i+1:]...)
		}
	}
	return l
}

func randomInt(min, max int) int {
	var b [8]byte
	_, err := crand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))

	return rand.Intn(max-min+1) + min
}

func hashMD5(line string) string {
	hash := md5.Sum([]byte(line))
	return hex.EncodeToString(hash[:])
}
