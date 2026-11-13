package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	//? Mailpit using for local testing
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"
		formattedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\n%s\r\n", recipient.Email, "Just testing mail..")
		msg := []byte(formattedMsg)

		fmt.Printf("Worker %d:sending email to %s \n", id, recipient.Email)

		err := smtp.SendMail(smtpHost+":"+smtpPort, nil, "om.prakashchauhan2929@gmail.com", []string{recipient.Email}, msg)

		if err != nil {
			log.Printf("Worker %d: failed to send email to %s: %v", id, recipient.Email, err)
			continue
		}
		time.Sleep(50 * time.Millisecond)
		fmt.Printf("Completed task && Worker %d:sending email to %s \n", id, recipient.Email)
	}
}
