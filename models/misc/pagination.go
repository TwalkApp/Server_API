package misc

import (
	"math"
)

type Pagination struct {
	Current	int	`json:"page"`
	Size	int	`json:"per_page"`
	Prev	int	`json:"prev"`
	Next	int	`json:"next"`
	Last	int	`json:"last"`
}

func (p Pagination) IsSet() bool {
	return p.Size != -1 && p.Current != -1
}

func (p *Pagination) SetInformations(count int) {
	p.Last = int(math.Ceil(float64(count) / float64(p.Size)))
	p.Prev = p.GetPrev()
	p.Next = p.GetNext()
}

func (p Pagination) GetFrom() int {
	return p.Size * (p.Current - 1)
}

func (p Pagination) GetTo() int {
	return p.GetFrom() + p.Size -1
}

func (p Pagination) GetPrev() int {
	if p.Current == 1 || p.Current == -1 {
		return p.Current
	}
	return p.Current - 1
}

func (p Pagination) GetNext() int {
	if p.Current == p.Last || p.Current == -1 {
		return p.Current
	}
	return p.Current + 1
}
