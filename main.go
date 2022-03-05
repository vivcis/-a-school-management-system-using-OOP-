package main

import (
	"fmt"
	my "github.com/vivcis/CECE-scl-mgt-system"
)

func main() {

	a := my.Applicants{
		Name:   "Olanma",
		Age:    14,
		Gender: "female",
		Class:  "Jss2",
	}

	p := my.Principal{}
	v := p.AdmitApplicant(a)
	fmt.Println(v)

	fmt.Println(my.StudentDB)
	s := my.Students{}
	take := s.TakeACourse("Olanma", "Chemistry")
	fmt.Println(take)

	t := my.Teachers{}
	gradeSt := t.GradeCourses("Franklyn", "Olanma", "Chemistry", 60)
	fmt.Println(gradeSt)
	//fmt.Println(scl.StudentDB)
	//printStudentDB := scl.Students{}
	//printStudentDB.Print()

	//fmt.Println(scl.TeacherDB)
	//printTeacherDB := scl.Teachers{}
	//printTeacherDB.Printer()

	//p.ExpelStudent("Olanma")
	//fmt.Println(scl.StudentDB)
	//p.SackTeacher("Franklyn Omonade")

	//a1 := my.BankAccount{"Funke", 50}
	//a1.Deposit(20)
	//fmt.Println("Balance:", a1.Balance())
	//a2 := my.OverdraftableBankAccount{BankAccount: my.BankAccount{"Femi", 100}, Fee: 20}
	//a2.Deposit(30)
	//fmt.Println("Balance for Femi:", a2.Balance())
	//balance, err := a2.Withdraw(150)
	//if err != nil {
	//	log.Printf("Overdraft incurred: balance is now %.2f", balance)
	//}
	//a2.Deposit(100)
	//fmt.Println("Account 1 Balance:", a1.Balance(), "Account 2 Balance:", a2.Balance())
	//err = my.Transfer(a1, a2, 100)
	//if err != nil {
	//	log.Printf("Could not complete transfer: %v", err)
	//}
	//err = my.Transfer(a2, a1, 100)
	//if err != nil {
	//	log.Printf("Could not complete transfer: %v", err)
	//}
	//fmt.Println("Account 1 Balance:", a1.Balance())
}
