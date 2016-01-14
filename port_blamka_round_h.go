package argon2

func fBlaMka(x, y uint64) uint64 {
	var (
		m  = uint64(0xFFFFFFFF)
		xy = (x & m) * (y & m)
	)
	return x + y + 2*xy
}

func rotr64(w, c uint64) uint64 {
	return (w >> c) | (w << (64 - c))
}

func bkG(a, b, c, d uint64) (uint64, uint64, uint64, uint64) {
	a = fBlaMka(a, b)
	d = rotr64(d^a, 32)
	c = fBlaMka(c, d)
	b = rotr64(b^c, 24)
	a = fBlaMka(a, b)
	d = rotr64(d^a, 16)
	c = fBlaMka(c, d)
	b = rotr64(b^c, 63)

	return a, b, c, d
}

func blake2_round_nomsg(v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15 uint64) (uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) {
	v0, v4, v8, v12 = bkG(v0, v4, v8, v12)
	v1, v5, v9, v13 = bkG(v1, v5, v9, v13)
	v2, v6, v10, v14 = bkG(v2, v6, v10, v14)
	v3, v7, v11, v15 = bkG(v3, v7, v11, v15)
	v0, v5, v10, v15 = bkG(v0, v5, v10, v15)
	v1, v6, v11, v12 = bkG(v1, v6, v11, v12)
	v2, v7, v8, v13 = bkG(v2, v7, v8, v13)
	v3, v4, v9, v14 = bkG(v3, v4, v9, v14)

	return v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15
}