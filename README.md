# Link-Shortener
Developer tool to shorten links.

## Link Shortening Methodology
The idea that I came up with is to first take the current time and hash it to an MD5 hash. The idea behind this is to reduce the amount of collisions since each timestamp will be unique. This string is then encoded using Base64 to generate a more alphanumeric string. It is important to remember that the entire point of this project is to *shorten* the link. To do this I take the first X characters from the encoded string. This reduces the uniquness of the string and increases the chance of collisions. I wrote the following code to see the ideal number of characters to shorten the Base64 encoded string to to reduce collisions.
```go
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
		hashedContent := base64.StdEncoding.EncodeToString(hash.Sum(nil))[0:6]
		if hashes[hashedContent] {
			collisions++
		}

		hashes[hashedContent] = true
	}

	println("Hashes Generated:", toGenerate)
	println("Collisions:", collisions)
}

```

Using this we can see the amount of collisions based on different lengths of shortened encoded strings. The results are in the table below. As can be seen the collisions for 6 characters are relatively low with collisions only starting to appear at 1000000 hashes while still being relatively short. This can be dealt with by adding or subtractng time then regerating the hash.

|Hashes Generated|Collisions (5 Chars)|Collisions (6 Chars)|Collisions (7 Chars)|Collisions (8 Chars)|
|---|---|---|---|---|
|10|0|0|0|0|
|100|0|0|0|0|
|1000|0|0|0|0|
|10000|0|0|0|0|
|100000|3|0|0|0|
|1000000|460|6|0|0|
|10000000|46166|706|16|1|
