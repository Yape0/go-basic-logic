package main

import (
	"fmt"
	"github.com/Yape0/go-basic-logic/logic3"
	"github.com/Yape0/go-basic-logic/utils"
)

func main() {

	//slice := []int{1, 2, 3}
	//go_print_slice.PrintSlice(logic1.Satusatu(10))
	//go_print_slice.PrintSlice(SatuDua(10))
	//go_print_slice.PrintSlice(SatuTiga(10))
	//go_print_slice.PrintSlice(SatuEmpat(10))
	//go_print_slice.PrintSlice(SatuLima(10))
	//go_print_slice.PrintSlice(SatuEnam(10))
	//go_print_slice.PrintSlice(SatuTujuh(10))   //bisa 10 / 11
	//go_print_slice.PrintSlice(SatuDelapan(11)) //bisa 10 / 11
	//go_print_slice.PrintSlice(SatuSembilan(11))
	//go_print_slice.PrintSlice(SatuDuabelas(12))

	fmt.Println("\n\n")

	//go_print_slice.PrintSlice2(DuaSatu(9))
	//go_print_slice.PrintSlice2(logic2.DuaSatu(9))
	//go_print_slice.PrintSlice2(DuaTiga(9))
	//go_print_slice.PrintSlice2(DuaEnam(9))
	//go_print_slice.PrintSlice2(DuaEmpat(9)) //masih perlu adjustment
	//go_print_slice.PrintSlice2(TigaTujuh(9))
	utils.PrintSlice2(logic3.TigaTigabelas(9))

}
