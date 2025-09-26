package database

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ISBN            int    `json:"isbn"`
	LibID           int    `json:"libId"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Publisher       string `json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     int    `json:"totalCopies"`
	AvailableCopies int    `json:"availableCopies"`
}

type Library struct {
	gorm.Model
	Name string `valid:"required,min=3,max=50"  json:"name" gorm:"uniqueIndex"`
}

type User struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	ContactNumber int    `json:"contact_number"` // Use snake case
	Role          string `json:"role"`
	LibId         int    `json:"lib_id"` // Use snake case
	Password      string `json:"password"`

	Library Library `valid:"required" gorm:"foreignKey:lib_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE,default:null;"   json:"library"`
}
type LibraryAdmin struct {
	ID   int    `json:"Id"`
	Name string `json:"name"`
}
type IssueRegistery struct {
	IssueID            int    `json:"issueId"`
	ISBN               int    `json:"isbn"`            //fk book
	ReaderID           int    `json:"readerId"`        //fk users
	IssueApproverID    int    `json:"issueApproverId"` //fk admin
	IssueStatus        string `json:"issueStatus"`
	IssueDate          string `json:"issueDate"`
	ExpectedReturnDate string `json:"expectedReturnDate"`
	ReturnDate         string `json:"returnDate"`
	ReturnApproverID   int    `json:"returnApproverId"` //fk admin

}
type RequestEvents struct {
	ReqID        int       `json:"reqId" gorm:"primaryKey"`
	BookID       int       `json:"bookId"`   //fk book
	ReaderID     int       `json:"readerId"` //fk user
	RequestDate  time.Time `json:"requestDate"`
	ApprovalDate time.Time `json:"approvalDate"`
	RejectDate   time.Time `json:"rejectDate"`
	ApproverID   int       `json:"approverId"` //fk admin
	RejectID     int       `json:"rejectId"`   //fk admin
	RequestType  string    `json:"requestType"`
	Status       string    `valid:"required,oneof=pending approved rejected" json:"status"`
}
