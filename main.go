package main

import (
	"crypto/md5"
	"encoding/base64"
	"time"
)

func main() {
	starting := 10
	ending := 10000000
	for starting <= ending {
		checkCollisions(starting)
		starting *= 10
	}
}

func checkCollisions(toGenerate int) {
	println()
	currTime := time.Now()
	hashes := make(map[string]bool)
	collisions := 0
	for i := 0; i < toGenerate; i++ {
		currTime = currTime.Add(time.Second)
		hash := md5.New()
		hash.Write([]byte(currTime.String()))
		hashedContent := base64.StdEncoding.EncodeToString(hash.Sum(nil))[0:8]
		if hashes[hashedContent] {
			collisions++
		}

		hashes[hashedContent] = true
	}

	println("Hashes Generated:", toGenerate)
	println("Collisions:", collisions)
}
