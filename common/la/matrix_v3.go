package la

import (
	"bfw/common/lang"
	"errors"
	"strconv"
	"strings"
)

const (
	sparseMatrixNoSize           int = 0
	sparseMatrixIndexOutOfBound  int = -1
	sparseMatrixElementZeroValue int = 0.0
)

var (
	sparseMatrixInValidError         = errors.New("sparse matrix is invalid")
	sparseMatrixIndexOutOfBoundError = errors.New("sparse matrix index is out of bound")
	sparseMatrixInvalidKeyError      = errors.New("sparse matrix key is invalid")
	NullSparseMatrix                 = &SparseMatrix{}
)

type SparseMatrix struct {
	tripleMap map[string]float64
	size      int
	rowSize   int
	lineSize  int
}

func (sm *SparseMatrix) validate() bool {
	if sm.size == sparseMatrixNoSize ||
		sm.tripleMap == nil ||
		len(sm.tripleMap) != sm.size {
		return false
	}
	return true
}

func (sm *SparseMatrix) validateOneIndex(index int) bool {
	if index < 0 ||
		index >= sm.rowSize ||
		index >= sm.lineSize {
		return false
	}
	return true
}

func (sm *SparseMatrix) validateIndex(index ...int) bool {
	if indexLen := len(index); indexLen > 0 {
		if indexLen >= 1 {
			if index[0] < 0 || index[0] >= sm.rowSize {
				return false
			}
		}
		if indexLen >= 2 {
			if index[1] < 0 || index[1] >= sm.lineSize {
				return false
			}
		}
		if indexLen >= 3 {
			for indexIdx := 2; indexIdx < indexLen; indexIdx++ {
				if !sm.validateOneIndex(index[indexIdx]) {
					return false
				}
			}
		}
	}
	return true
}

func (sm *SparseMatrix) assign(size, rowSize, lineSize int) {
	sm.setSize(size, rowSize, lineSize)
	sm.tripleMap = make(map[string]float64, size)
}

func (sm *SparseMatrix) setSelf(smt *SparseMatrix) {
	sm.setValues(smt.tripleMap, smt.size, smt.rowSize, smt.lineSize)
}

func (sm *SparseMatrix) setValues(tm map[string]float64, size, rowSize, lineSize int) {
	sm.setSize(size, rowSize, lineSize)
	sm.setTripleMap(tm)
}

func (sm *SparseMatrix) setSize(size, rowSize, lineSize int) {
	sm.size = size
	sm.rowSize = rowSize
	sm.lineSize = lineSize
}

func (sm *SparseMatrix) setTripleMap(tm map[string]float64) {
	sm.tripleMap = tm
}

func (sm *SparseMatrix) getKey(rowIndex, lineIndex int) string {
	if !sm.validateIndex(rowIndex, lineIndex) {
		panic(sparseMatrixIndexOutOfBoundError)
	}
	return lang.ConcatIntToString("_", rowIndex, lineIndex)
}

func (sm *SparseMatrix) makeCopy() *SparseMatrix {
	smCopy := &SparseMatrix{}
	smCopy.assign(sm.size, sm.rowSize, sm.lineSize)
	for key, value := range sm.tripleMap {
		smCopy.setKV(key, value)
	}
	return smCopy
}

func (sm *SparseMatrix) set(rowIndex, lineIndex int, value float64) {
	if !sm.validateIndex(rowIndex, lineIndex) {
		panic(sparseMatrixIndexOutOfBoundError)
	}
	sm.setKV(sm.getKey(rowIndex, lineIndex), value)
}

func (sm *SparseMatrix) setKV(key string, value float64) {
	sm.tripleMap[key] = value
}

func (sm *SparseMatrix) get(rowIndex, lineIndex int) float64 {
	if !sm.validateIndex(rowIndex, lineIndex) {
		panic(sparseMatrixIndexOutOfBoundError)
	}
	return sm.getValue(sm.getKey(rowIndex, lineIndex))
}

func (sm *SparseMatrix) getValue(key string) float64 {
	return sm.tripleMap[key]
}

func (sm *SparseMatrix) del(rowIndex, lineIndex int) {
	if !sm.validateIndex(rowIndex, lineIndex) {
		panic(sparseMatrixIndexOutOfBoundError)
	}
	sm.delKV(sm.getKey(rowIndex, lineIndex))
}

func (sm *SparseMatrix) delKV(key string) {
	delete(sm.tripleMap, key)
}

func (sm *SparseMatrix) validateKey(key string) (int, int) {
	indexStrArr := strings.Split(key, "_")
	if indexStrArr == nil ||
		len(indexStrArr) != 2 {
		panic(sparseMatrixInvalidKeyError)
	}
	rowIndex, rowIndexErr := strconv.Atoi(indexStrArr[0])
	if rowIndexErr != nil {
		panic(rowIndexErr)
	}
	lineIndex, lineIndexErr := strconv.Atoi(indexStrArr[1])
	if lineIndexErr != nil {
		panic(lineIndexErr)
	}
	return rowIndex, lineIndex
}

func (sm *SparseMatrix) swapKeyRowLineIndex(key string) string {
	rowIndex, lineIndex := sm.validateKey(key)
	if !sm.validateIndex(rowIndex, lineIndex) {
		panic(sparseMatrixIndexOutOfBound)
	}
	return sm.getKey(lineIndex, rowIndex)
}

func (sm *SparseMatrix) convertToMatrix() *Matrix {
	return &Matrix{}
}

func (sm *SparseMatrix) Matrix() *Matrix {
	return sm.convertToMatrix()
}

func (sm *SparseMatrix) transpose() *SparseMatrix {
	for key, value := range sm.tripleMap {
		sm.delKV(key)
		key = sm.swapKeyRowLineIndex(key)
		sm.setKV(key, value)
	}
	return &SparseMatrix{}
}

func (sm *SparseMatrix) GetTranspose() *SparseMatrix {
	smCopy := sm.makeCopy()
	return smCopy.transpose()
}
