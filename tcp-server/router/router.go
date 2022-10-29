package router

import "math"

type Router interface {
	Handler(msg string, msgChan chan string)
}

type AddRouter struct{}

func (a *AddRouter) Handler(msg string, msgChan chan string) {
	size := len(msg)
	i := 0
	for ; i < size && msg[i] != '+'; i++ {
	}
	if i >= size-1 {
		msgChan <- "0"
		return
	}
	s1 := msg[:i]
	s2 := msg[i+1:]
	l1 := len(s1)
	l2 := len(s2)
	l := int(math.Max(float64(l1), float64(l2))) + 1
	p := 0
	ans := make([]uint8, l)
	for l1 > 0 && l2 > 0 {
		l1--
		l2--
		l--
		if s1[l1] > '9' || s1[l1] < '0' || s2[l2] > '9' || s2[l2] < '0' {
			msgChan <- "0"
			return
		}
		x := int(s1[l1] + s2[l2] - '0'*2)
		x += p
		p = x / 10
		ans[l] = uint8(x%10) + '0'
	}
	for l1 > 0 {
		l1--
		l--
		if s1[l1] > '9' || s1[l1] < '0' {
			msgChan <- "0"
			return
		}
		x := int(s1[l1] - '0')
		x += p
		p = x / 10
		ans[l] = uint8(x%10) + '0'
	}
	for l2 > 0 {
		l2--
		l--
		if s2[l2] > '9' || s2[l2] < '0' {
			msgChan <- "0"
			return
		}
		x := int(s2[l2] - '0')
		x += p
		p = x / 10
		ans[l] = uint8(x%10) + '0'
	}
	if p == 1 {
		ans[0] = '1'
		msgChan <- string(ans)
	} else {
		msgChan <- string(ans[1:])
	}
}
