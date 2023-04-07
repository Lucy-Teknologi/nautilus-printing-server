package zpl

import (
	"fmt"
)

func RetouchingZPLString(printable Printable) string {
	basket_str := fmt.Sprint(printable.Basket)
	loin_str := fmt.Sprint(printable.Loin)
	weight_str := fmt.Sprint(printable.Weight)
	return fmt.Sprintf(`
		^XA
		^PW406
		^FO120,40
		^BQ,2,6
		^FDQA,%s^FS
		^FO125,216^A0,20^FD%s^FS
		^FO115,240^A0,20^FDCutting %s^FS
		^FO110,260^A0,20^FDRetouch %s^FS
		^FO135,308^A0,20^FDBasket: %s^FS
		^FO135,328^A0,20^FDLoin: %s^FS
		^FO135,348^A0,20^FDBerat: %s^FS
		^FO135,368^A0,20^FDGrade: %s^FS
		^FS
		^XZ
	`,
		printable.ILC,
		printable.CommonILC,
		printable.CuttingDate,
		printable.RetouchingDate,
		padStart(basket_str, BasketSpace-len(basket_str), " "),
		padStart(loin_str, LoinSpace-len(loin_str), " "),
		padStart(weight_str, WeightSpace-len(weight_str), " "),
		padStart(printable.Grade, GradeSpace-len(printable.Grade), " "),
	)
}

func CuttingZPLString(printable Printable) string {
	return fmt.Sprintf(`
		^XA
		^PW406
		^FO120,40
		^BQ,2,6
		^FDQA,%s
		^FS
		^FO125,216^A0,20^FD%s^FS
		^FO115,240^A0,20^FDCutting %s^FS
		^FO135,288^A0,20^FDBasket: %d^FS
		^FO135,308^A0,20^FDLoin: %d^FS
		^FO135,328^A0,20^FDBerat: %.3f^FS
		^FO135,348^A0,20^FDGrade: %s^FS
		^XZ
		`,
		printable.ILC,
		printable.CommonILC,
		printable.CuttingDate,
		printable.Basket,
		printable.Loin,
		printable.Weight,
		printable.Grade,
	)
}
