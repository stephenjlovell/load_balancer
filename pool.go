
package main

type Pool []*Worker // a heap implemented as a priority queue of pointers to worker objects.

// satisfy the container#heap interface...

func (p Pool) Len() int { return len(p) }

func (p Pool) Less(i, j int) bool { return p[i].pending < p[j].pending }

func (p Pool) Swap(i, j int) {
  p[i], p[j] = p[j], p[i]
  p[i].index, p[j].index = j, i
}

func (p *Pool) Push(w interface{}) {
  n := len(*p)
  item := w.(*Worker)
  item.index = n
  *p = append(*p, item)
}

func (p *Pool) Pop() interface{} {
  old := *p
  n := len(old)
  item := old[n-1]
  *p = old[0 : n-1]
  return item
}
