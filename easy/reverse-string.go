func reverseString(s []byte) {
	if len(s) == 0 || len(s) == 1 {
		return
	}
	stidx, edidx := 0, len(s)-1
	for stidx<edidx{
		l, r := s[stidx], s[edidx]
		s[stidx], s[edidx] = r, l
		stidx++
		edidx--
	}
}
