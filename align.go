package bdd

import "strings"

type align struct {
	child *align
	next  *align
	value builder
	colon int // Data separated by a colon. The position of the colon is used to align the data after the colon.
}

// lrPos indented empty brackets
func (n *align) lrPos() {
	b := n
	for x := b; x != nil && x.next != nil; x = x.next {
		if x.child != nil {
			continue
		}
		ss := x.next.value.String()
		if len(ss) == 2 && (ss[1] == ')' || ss[1] == ']' || ss[1] == '}') {
			x.mergeNext(1)
		}
	}
}

func (n *align) spac(i int) {
	for k := 0; k < i; k++ {
		n.value.WriteByte(Space)
	}
}

// mergeNext merge the next node to the current node
func (n *align) mergeNext(max int) {
	n.spac(max - strLen(n.value.String()))
	n.value.WriteString(n.next.value.String())
	putBuilder(n.next.value)
	n.next = n.next.next
}

// tablePos aligns data of array type
func (n *align) tablePos() {
	ms := []int{}
	b := n
	max := 0
	for x := b; x != nil; x = x.next {
		if x.colon > 0 || x.child != nil {
			return
		}
		ll := strLen(x.value.String())
		ms = append(ms, ll)
		if ll > max {
			max = ll
		}
	}

	if max < 10 {
		n.merge(9, ms)
	} else if max < 18 {
		n.merge(5, ms)
	} else if max < 30 {
		n.merge(3, ms)
	}
}

// merge merge to the next node
func (n *align) merge(m int, ms []int) {
	l := len(ms)
	col := 0
	for i := 0; i != m; i++ {
		z := m - i
		if l > z && l%z == 0 {
			col = z
			break
		}
	}
	if col > 1 {
		n.mergeNextSize(col, ms)
	}
}

// mergeNextSize merge to the next node specified length
func (n *align) mergeNextSize(s int, ms []int) {
	lmax := make([]int, s)

	for j := 0; j != s; j++ {
		for i := 0; i*s < len(ms); i++ {
			b := i*s + j
			if ms[b] > lmax[j] {
				lmax[j] = ms[b]
			}
		}
	}

	for i := 1; i < len(lmax); i++ {
		lmax[i] += lmax[i-1]
	}

	for x := n; x != nil; x = x.next {
		for i := 0; i < s-1 && x.next != nil; i++ {
			x.mergeNext(lmax[i])
		}
	}
}

// colonPos align the data after the colon
func (n *align) colonPos() {
	b := n
	for b != nil {
		m := 0
		for x := b; x != nil; x = x.next {
			if x.colon <= 0 {
				continue
			}
			bl := strLen(x.value.String()[:x.colon])
			if bl > m {
				m = bl
			}
			if x.child != nil {
				break
			}
		}
		for x := b; x != nil; x = x.next {

			if x.colon > 0 {
				bl := strLen(x.value.String()[:x.colon])
				if m-bl > 0 {
					t := strings.Replace(x.value.String(), colSym, colSym+spac(m-bl), 1)
					x.value.Reset()
					x.value.WriteString(t)
				}
			}
			b = x.next
			if x.child != nil {
				break
			}
		}
	}
}

func (n *align) put() {
	if n.value != nil {
		putBuilder(n.value)
		n.value = nil
	}

	if n.child != nil {
		n.child.put()
	}

	if n.next != nil {
		n.next.put()
	}
}

