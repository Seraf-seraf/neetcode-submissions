func isValid(parentheses string) bool {
	set := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	mem := []byte{}
	for i := 0; i < len(parentheses); i++ {
		s := parentheses[i]

		if openedBracket, isClosed := set[s]; isClosed {
			if len(mem) == 0 {
				return false
			}

			opened := mem[len(mem)-1]
			if opened != openedBracket {
				return false
			}

			mem = mem[:len(mem)-1]
			continue
		}

		mem = append(mem, s)
	}

	return len(mem) == 0
}
