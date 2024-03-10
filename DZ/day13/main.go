package main

import "fmt"

type alif struct {
	Payment  int
	Balanse  int
	Transfer int
}

type humo struct {
	Payment  int
	Balanse  int
	Transfer int
}

type ib struct {
	Payment  int
	Balanse  int
	Transfer int
}

type ds struct {
	Payment  int
	Balanse  int
	Transfer int
}
type qrs struct{
	Payment  int
}

type bancs interface{
	balansebanc
	transferbanc
	paymentbanc

}

type balansebanc interface{
	balanse()
}

type transferbanc interface{
	transfer()
}

type paymentbanc interface{
	payment()
}



type qr interface {
	plusBalanse(int)
	makePayment(int)
	//trTransfer(int)
}

func qrConstructor(q qr) {
	q.plusBalanse(1000)
	q.makePayment(200)
	//q.Transfer(154)
}
func main() {

	var userhumo *humo
	userhumo = new(humo)
	qrConstructor(userhumo)
	fmt.Printf(" Humo Банк : %+v\n", userhumo)

	var useralif *alif
	useralif = new(alif)
	qrConstructor(useralif)
	fmt.Printf(" Alif Банк: %+v\n", useralif)

	var userib *ib
	userib = new(ib)
	qrConstructor(userib)
	fmt.Printf(" IB Банк: %+v\n", userib)

	var userds *ds
	userds = new(ds)
	qrConstructor(userds)
	fmt.Printf(" DS Банк: %+v\n", userds)
	
	var user *bancs
	user = new(bancs)
	fmt.Println(" QR bancs: 12341", user)



}

type user struct{
	
}


func (h *humo) makePayment(payment int) {
	h.Payment = h.Balanse - payment
	fmt.Print(h.Balanse)

}

func (h *humo) plusBalanse(balanse int) {
	h.Balanse = h.Payment + balanse
	fmt.Print(h.Balanse)
}

/*func (h *humo) trTransfer(transfer int) {
	h.Transfer = h.Transfer - 
}*/

func (d *ds) makePayment(payment int) {
	d.Payment = (d.Balanse - d.Payment) - payment
	fmt.Printf("Баланс: %+v\n", d.Balanse)

}

func (d *ds) plusBalanse(balanse int) {
	d.Balanse = d.Payment + balanse
	fmt.Printf("Баланс: %+v\n", d.Balanse)
}

func (i *ib) makePayment(payment int) {
	i.Payment = i.Balanse -  payment
	fmt.Print(i.Balanse)
}

func (i *ib) plusBalanse(balanse int) {
	i.Balanse = i.Balanse + balanse
	fmt.Print(i.Balanse)
}

func (a *alif) makePayment(payment int) {
	a.Payment = a.Balanse - payment
	fmt.Print(a.Balanse)
}

func (a *alif) plusBalanse(balanse int) {
	a.Balanse = a.Payment + balanse
	fmt.Print(a.Balanse)
}
