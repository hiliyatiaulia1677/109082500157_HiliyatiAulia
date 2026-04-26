package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	for {
		fmt.Scanf("%c", &ch)
		if ch == '\n' || ch == ' ' {
			continue
		}
		if ch == '.' || *n >= NMAX {
			break
		}
		t[*n] = ch
		*n++
	}
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks  : ")
	isiArray(&tab, &m)

	for i := 0; i < m; i++ {
		fmt.Printf("%c ", tab[i])
	}
	fmt.Println()

	// palindrom
	if palindrom(tab, m) {
		fmt.Println("Palindrom ? true")
	} else {
		fmt.Println("Palindrom ? false")
	}

	// reverse
	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	for i := 0; i < m; i++ {
		fmt.Printf("%c ", tab[i])
	}
	fmt.Println()
}
