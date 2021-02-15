// type alias
type hash = map[byte]byte

func isValidSudoku(board [][]byte) bool {
    
    // make row map 
    row :=  make(hash)
    
    // make col maps
    cols := make([]hash, 9)
    for i, _ := range cols {
        cols[i] = make(hash)
    }
    // make box maps
    boxes :=  make([]hash, 9)
    for i, _ := range boxes {
        boxes[i] = make(hash)
    }
    
    for i, _ := range board {
        for j, b := range board[i] {
            if b == '.' {
                continue
            }
            _, dup1 := row[b]
            _, dup2 := cols[j][b]
            _, dup3 := boxes[3*(i/3) + j/3][b]
            if dup1 || dup2 || dup3 {
                return false
            }
            row[b] = b
            cols[j][b] = b
            boxes[3*(i/3) + j/3][b] = b
        }
        row = make(hash)
    }
    
    return true 
}
