package main

import "fmt"

func main() {
	var h1, m1, s1, h2, m2, s2 int
	fmt.Scanf("%d %d %d\n%d %d %d", &h1, &m1, &s1, &h2, &m2, &s2)
	if h2 < h1 || h2 == h1 && m2 < m1 || h2 == h1 && m2 == m1 && s2 < s1 {
		h2 += 24
	}
	h3 := h2 - h1
	m3 := m2 - m1
	s3 := s2 - s1
	if s3 < 0 {
		s3 = 60 + s3
		m3--
	}
	if m3 < 0 {
		m3 = 60 + m3
		h3--
	}
	fmt.Printf("%02d %02d %02d\n", h3, m3, s3)
}
