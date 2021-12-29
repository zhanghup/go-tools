package sys

import "fmt"

type SizeNumber uint64

func (this SizeNumber) String() string {
	v := uint64(this)
	if v < 1<<10 {
		return fmt.Sprintf("%dB", this)
	} else if v < 1<<20 {
		return fmt.Sprintf("%.2fKB", float64(this)/float64(1<<10))
	} else if v < 1<<30 {
		return fmt.Sprintf("%.2fMB", float64(this)/float64(1<<20))
	} else if v < 1<<40 {
		return fmt.Sprintf("%.2fGB", float64(this)/float64(1<<30))
	} else {
		return fmt.Sprintf("%.2fPB", float64(this)/float64(1<<40))
	}
}
func (this SizeNumber) B() uint64 {
	return uint64(this)
}
func (this SizeNumber) KB() float64 {
	return float64(this) / float64(2<<10)
}
func (this SizeNumber) MB() float64 {
	return float64(this) / float64(2<<20)
}
func (this SizeNumber) GB() float64 {
	return float64(this) / float64(2<<30)
}
func (this SizeNumber) PB() float64 {
	return float64(this) / float64(2<<40)
}
