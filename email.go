
// Sending Email Using Smtp in Golang
package main
 
import (
    "fmt"
    "net/smtp"
    "os"
    "os/exec"
    "bufio"
    "time"
    "html/template"
    "bytes"
    "log"
)

// EmailData holds the data we want to inject into the email template.
type EmailData struct {
	A string
	B string
}

func main () {


	var penis string
	//attackers email
	from := "EMAIL HERE"
   	password := "PASS KEY HERE"


	loop: for {
		Clear()
		fmt.Println(Figlet("GopherSMTP", "slant"))
		fmt.Println("SMTP attack tool")
		fmt.Println("List:")
		fmt.Println("0: quit \n1: send single email \n2: Phish from a template \n3: send to multiple emails \n4: smsOverSmtp\n5: Send smsOverSmtp to multiple numbers")

		fmt.Scanln(&penis)

		switch penis {

		case "0":
			break loop
		case "1":
			emailPhish(from, password)
		case "2":
                        phishFromATemplate(from, password)

		case "3":
			multiSender(from, password)
		case "4":
			textOverSmtp(from, password)
		case "5":
			multiTextOverSmtp(from, password)
		}
	}
}
func phishFromATemplate (from, password string) {

	//var from string
	//var password string
	var to string
	var subject string
	var file string
	var injectA string
	var injectB string

	fmt.Println("Paste email here:")
	fmt.Scanln(&to)

	fmt.Println("subject")
	subject = getInput("")

	fmt.Println("file here:")
	fmt.Scanln(&file)

	fmt.Println("INJECT A:")
	fmt.Scanln(&injectA)
	fmt.Println("INJECT B:")
        fmt.Scanln(&injectB)

	sendEmailFromTemplate(from, password, to, subject, file, injectA, injectB)

}

func sendEmailFromTemplate(from string, password string, to string, subject string, file string, injectA string, injectB string) {

	data := EmailData{
        	A: injectA,
        	B: injectB,
        }
        err := sendEmailTemplate(from, password, to, subject, file, data)
        if err != nil {
        	log.Fatal("Error sending email:", err)
        }
        fmt.Println("Email sent successfully!")
}

//sends email from template
func sendEmailTemplate(from string, password string, to string, subject string, file string, data EmailData) error {
	// Step 1: Parse the template
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		return err
	}

	// Step 2: Render the template to a buffer
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return err
	}

	// Step 3: Set up the email headers and body
	//from := "youremail@example.com"
	//password := "yourpassword"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Email headers
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body.String()

	// Step 4: Send the email
	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}

func emailPhish(from, password string) {
	var to, subject, choice string

	fmt.Println("Type victim email address:")
	fmt.Scanln(&to)

	fmt.Println("Subject:")
	subject = getInput("")

	fmt.Println("Do you want to use a template? (Y/N)")
	fmt.Scanln(&choice)

	if choice == "Y" || choice == "y" {
		var templateFile, name, link string

		fmt.Println("Template file name (e.g. email_template.txt):")
		fmt.Scanln(&templateFile)

		fmt.Println("Inject a name into the template:")
		fmt.Scanln(&name)

		fmt.Println("Paste your phishing link:")
		fmt.Scanln(&link)

		data := EmailData{
			A: name,
			B: link,
		}

		err := sendEmailTemplate(from, password, to, subject, templateFile, data)
		if err != nil {
			log.Fatal("Error sending email:", err)
		}
		fmt.Println("Email sent successfully!")

	} else {
		fmt.Println("Write your email message (press Enter when done):")
		body := getInput("")

		fullMessage := "From: " + from + "\n" +
			"To: " + to + "\n" +
			"Subject: " + subject + "\n\n" +
			body

		auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
		err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(fullMessage))
		if err != nil {
			log.Fatal("Error sending custom email:", err)
		}
		fmt.Println("Custom email sent successfully!")
	}
}


func getInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // Reads input until newline

	return input[:len(input)-1] // Removes the newline character at the end
}

func textOverSmtp(from, password string) {
	var number string
	var text string
	var A string
	var gate string

	gateway := []string{"@msg.koodomobile.com", "@txt.att.net", "@cingularme.com", "@txt.bellmobility.com", "@sms.myboostmobile.com", "@text.scsouth1.com", "@vmobl.com", "@vtext.com", "@msg.telus.com", "@tms.suncom.com", "@messaging.sprintpcs.com", "@email.uscc.net", "@sms.sasktel.com", "@text.republicwireless.com", "@qwestmp.com", "@vtext.com", "@page.nextel.com", "@cspire1.com", "@sms.cricketwireless.net", "@mailmymobile.net", "@cwemail.com"}

	fmt.Println("what is the victims phone number?")
	fmt.Scanln(&number)
	
	fmt.Println("what is the text you want to send?")
	text = input()
	
	fmt.Println("Do you know the cell provider gateway? Y/N")
	fmt.Scanln(&A)
	if A == "n" || A == "N" {
		for i := 0; i <= 10; i++ {
		//vic := number + "@msg.koodomobile.com"
	
		vic := number + gateway[i]
		fmt.Println(vic)
		email(from, password, text, vic)
		fmt.Println(vic)
	
		}
	
	} else {
	
	fmt.Println("what is the sms gateway?")
        fmt.Scanln(&gate)

	vic := number + gate

	email(from, password, text, vic)
        fmt.Println(vic)
	
	}
	
	time.Sleep(2 * time.Second)

}

func multiTextOverSmtp(from, password string) {
	var listFile string
	var knowsGateway string
	var gateway string
	var useTemplate string
	var text string
	var templateFile string
	var name string
	var link string

	gateways := []string{
		"@msg.koodomobile.com", "@txt.att.net", "@cingularme.com", "@txt.bellmobility.com",
		"@sms.myboostmobile.com", "@text.scsouth1.com", "@vmobl.com", "@vtext.com",
		"@msg.telus.com", "@tms.suncom.com", "@messaging.sprintpcs.com", "@email.uscc.net",
		"@sms.sasktel.com", "@text.republicwireless.com", "@qwestmp.com", "@page.nextel.com",
		"@cspire1.com", "@sms.cricketwireless.net", "@mailmymobile.net", "@cwemail.com",
	}

	fmt.Println("Enter filename of phone number list:")
	fmt.Scanln(&listFile)
	numbers := fileLines(listFile)

	fmt.Println("Do you know the SMS gateway? (Y/N)")
	fmt.Scanln(&knowsGateway)

	fmt.Println("Do you want to use a template? (Y/N)")
	fmt.Scanln(&useTemplate)

	if useTemplate == "Y" || useTemplate == "y" {
		fmt.Println("Enter template filename (e.g. 'sms_template.txt'):")
		fmt.Scanln(&templateFile)

		fmt.Println("Inject variable A (e.g. name):")
		fmt.Scanln(&name)

		fmt.Println("Inject variable B (e.g. link):")
		fmt.Scanln(&link)

		data := EmailData{A: name, B: link}

		body := renderTemplate(templateFile, data)
		if body == "" {
			log.Fatal("Could not parse template.")
		}

		if knowsGateway == "Y" || knowsGateway == "y" {
			fmt.Println("Enter the SMS gateway (e.g. @vtext.com):")
			fmt.Scanln(&gateway)
			for _, num := range numbers {
				vic := num + gateway
				email(from, password, body, vic)
			}
		} else {
			for _, num := range numbers {
				for i := 0; i <= 10 && i < len(gateways); i++ {
					vic := num + gateways[i]
					email(from, password, body, vic)
				}
			}
		}
	} else {
		fmt.Println("Enter your message:")
		text = getInput("")

		if knowsGateway == "Y" || knowsGateway == "y" {
			fmt.Println("Enter the SMS gateway (e.g. @vtext.com):")
			fmt.Scanln(&gateway)
			for _, num := range numbers {
				vic := num + gateway
				email(from, password, text, vic)
			}
		} else {
			for _, num := range numbers {
				for i := 0; i <= 10 && i < len(gateways); i++ {
					vic := num + gateways[i]
					email(from, password, text, vic)
				}
			}
		}
	}
}

func renderTemplate(file string, data EmailData) string {
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		fmt.Println("Template parsing error:", err)
		return ""
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		fmt.Println("Template execution error:", err)
		return ""
	}
	return buf.String()
}



func input() string {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter a string with spaces: ")
	input, _ := reader.ReadString('\n')
	//fmt.Println("You entered:", input)
	return input
}

//This function is for sendinng emails to multiple users from a file list. list file? idfk
func multiSender(from, password string) {
	var file string
	//var vic string
        var sub string
        var template string
        var name string
        var link string

	//var msg string
	fmt.Println("Paste your list of emails here, it should just be emails in a list with no interuptions.")
	fmt.Scanln(&file)
	
	//fmt.Println("Type victim email here address here:")
        //scan target email address
        //fmt.Scanln(&vic)
	var templateYorN string
	fmt.Println("Use a template? Y/N?")
	fmt.Scanln(&templateYorN)
	
	if templateYorN == "Y" || templateYorN == "y"{

        //you gotta use scanln for some reason. bufio fucks up sending the email. something about CRLF
        fmt.Println("Type subject line here:")
        //scan subject line
        sub = getInput("")



        //Scan template name
        fmt.Println("Paste template filename: example: 'email_template.txt'")
        fmt.Scanln(&template)
        //scan name for the email body
        fmt.Println("Inject a name into the template:")
        fmt.Scanln(&name)
        //scan phishing link
        fmt.Println("paste your phishing link:")
        fmt.Scanln(&link)

        //I tried changing sendername to link and it fucked up, so don't do that
        data := EmailData{
                A:       name,
                B: link,
        }

	emailArray := fileLines(file)
	numOfEmails := len(emailArray)

	for i := 0; i < numOfEmails; i++ {
		//email(from, password, msg, emailArray[i])
		//here we send the auth information, the victims email address, the subject line, email template, and data we want to inject into our email, like a fishing link. Or users names
        	err := sendEmailTemplate(from, password, emailArray[i], sub, template, data)
	        if err != nil {
                	log.Fatal("Error sending email:", err)
        	}
        	fmt.Println("Email sent successfully!")

	}
	} else {
	
		//you gotta use scanln for some reason. bufio fucks up sending the email. something about CRLF
        fmt.Println("Type subject line here:")
        //scan subject line
        sub = getInput("")

        //Scan template name
        fmt.Println("Paste template filename: example: 'email_template.txt'")
        //get the email body
	template = getInput("")
	
	mess := sub + template

        emailArray := fileLines(file)
        numOfEmails := len(emailArray)

        for i := 0; i < numOfEmails; i++ {
                //email(from, password, msg, emailArray[i])
                //here we send the auth information, the victims email address, the subject line, email template, and data we want to inject into our email, like a fishing link. Or users names
                email(from, password, mess,emailArray[i])
                //if err != nil {
                  //      log.Fatal("Error sending email:", err)
                //}
                //fmt.Println("Email sent successfully!")

        }


	}
	time.Sleep(2 * time.Second)
}

// Main function
func email(from, password, msg, victim string) {
 
    // from is senders email address
     
    // we used environment variables to load the
    // email address and the password from the shell
    // you can also directly assign the email address
    // and the password
    //from := "golangbot699@gmail.com"
    //password := "kdxy zlvv unbk rssy"
 
    // toList is list of email address that email is to be sent.
    toList := []string{victim}
 
    // host is address of server that the
    // sender's email address belongs,
    // in this case its gmail.
    // For e.g if your are using yahoo
    // mail change the address as smtp.mail.yahoo.com
    host := "smtp.gmail.com"
 
    // Its the default port of smtp server
    port := "587"
 
    // This is the message to send in the mail
    //msg := "The nukes are cumming!!!"
    //However I commented that out and it's now declared in the func call!
    // We can't send strings directly in mail,
    // strings need to be converted into slice bytes
    body := []byte(msg)
 
    // PlainAuth uses the given username and password to
    // authenticate to host and act as identity.
    // Usually identity should be the empty string, 
    // to act as username.
    auth := smtp.PlainAuth("", from, password, host)
 
    // SendMail uses TLS connection to send the mail
    // The email is sent to all address in the toList,
    // the body should be of type bytes, not strings
    // This returns error if any occurred.
    err := smtp.SendMail(host+":"+port, auth, from, toList, body)
 
    // handling the errors
    check(err)

    fmt.Println("Successfully sent mail to all user in toList")
}
//This is the thingy for reading from list files, files that are just lists of things. yeah.
func fileLines(f string) []string {

        filePath := f
        readFile, err := os.Open(filePath)
        check(err)

        fileScanner := bufio.NewScanner(readFile)
        fileScanner.Split(bufio.ScanLines)
        var fileLines []string


                for fileScanner.Scan() {
                fileLines = append(fileLines, fileScanner.Text())
        }

        readFile.Close()

        return fileLines
}


func Clear() {
        
	cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()


    	//check(err)

}

func Figlet(text string, font string) (string) {
	// Run `figlet` with the specified text and font
	cmd := exec.Command("figlet", "-f", font, text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
		//return "", fmt.Errorf("figlet command failed: %w", err)
	}
	return string(output)
}

func check (err error) {

	if err != nil {
        	fmt.Println(err)
        	os.Exit(1)
    	}
}

