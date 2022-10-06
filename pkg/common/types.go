package common

type Hash [32]byte

func (h *Hash) SetBytes(data []byte) {
	if len(h) == len(data) {
		_ = append(h[:0], data...)
	}
}

func (h *Hash) Bytes() []byte {
	return h[:]
}

type Address [20]byte

func (a *Address) SetBytes(data []byte) {
	if len(a) == len(data) {
		_ = append(a[:0], data...)
	}
}

func (a *Address) Bytes() []byte {
	return a[:]
}

type Bloom [256]byte

func (b *Bloom) SetBytes(data []byte) {
	if len(b) == len(data) {
		_ = append(b[:0], data...)
	}
}

func (b *Bloom) Bytes() []byte {
	return b[:]
}

type BlockNonce [8]byte

func (n *BlockNonce) SetBytes(data []byte) {
	if len(n) == len(data) {
		_ = append(n[:0], data...)
	}
}

func (n *BlockNonce) Bytes() []byte {
	return n[:]
}
