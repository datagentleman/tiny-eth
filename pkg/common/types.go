package common

// Hash
type Hash [32]byte

func NewHash(data []byte) *Hash {
	return new(&Hash{}, data)
}

func (h *Hash) SetBytes(data []byte) {
	setBytes(h, data)
}

// Address
type Address [20]byte

func NewAddress(data []byte) *Address {
	return new(&Address{}, data)
}

func (a *Address) SetBytes(data []byte) {
	setBytes(a, data)
}

// Bloom
type Bloom [256]byte

func NewBloom(data []byte) *Bloom {
	return new(&Bloom{}, data)
}

func (b *Bloom) SetBytes(data []byte) {
	setBytes(b, data)
}

// BlockNonce
type BlockNonce [8]byte

func NewBlockNonce(data []byte) *BlockNonce {
	return new(&BlockNonce{}, data)
}

func (n *BlockNonce) SetBytes(data []byte) {
	setBytes(n, data)
}

type Hashes interface {
	*Hash | *Address | *BlockNonce | *Bloom
}

func setBytes[T Hashes](a T, data []byte) {
	if len(data) > len(a) {
		data = data[:len(a)]
	}

	for i := 0; i < len(data); i++ {
		a[i] = data[i]
	}
}

func new[T Hashes](a T, data []byte) T {
	setBytes(a, data)
	return a
}
