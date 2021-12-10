package types

type Chunk struct {
	Closing string
}

type Stack struct {
	chunks []Chunk
}

func NewStack() *Stack {
	return &Stack{make([]Chunk, 0, 10)}
}

func (s *Stack) Push(v Chunk) {
	s.chunks = append(s.chunks, v)
}

func (s *Stack) Pop() Chunk {
	l := len(s.chunks)
	top := s.chunks[l-1]
	s.chunks = s.chunks[:l-1]
	return top
}

func (s *Stack) Len() int {
	return len(s.chunks)
}
