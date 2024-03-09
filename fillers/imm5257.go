package fillers

import (
	"fmt"
	"os/exec"
	"time"

	"gocandoit.xyz/model"
)

func FillForm(applicant *model.Applicant) {
	cmd := exec.Command("open", "-a", "Adobe Acrobat Reader", "./forms/IMM5257.pdf")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error Adobe Acrobat Reader", err)
		return
	}

	time.Sleep(5 * time.Second)

	PressTab(6)
	Type(applicant.LastName)
	PressTab(1)
	Type(applicant.FirstName)
	PressTab(1)
	PressKey("space")
	PressTab(3)
	SelectSex(applicant.Sex)
	PressTab(1)
	Type(fmt.Sprintf("%d", applicant.Dob.Year()), 0, 1)
	Type(applicant.Dob.Format("01"), 0, 1)
	Type(applicant.Dob.Format("02"), 0, 1)
	fmt.Println("It was: " + fmt.Sprintf("%d", applicant.Dob.Month()))

}
