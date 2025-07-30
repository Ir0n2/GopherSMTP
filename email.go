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
    "strconv"
)

// EmailData holds the data we want to inject into the email template.
type EmailData struct {
	A string
	B string
}

func main () {

	//penis is the user input var
	var penis string
	var emailArray []string
	emailArray = readFileLineByLine("AttackerEmails.txt")
	var emailName int
	var pass int
	
	emailName = 0
	pass = 1

	loop: for {

	from := emailArray[emailName]
   	password := emailArray[pass]


	
		Clear()
		fmt.Println(Figlet("GopherSMTP", "slant"))
		fmt.Println("Current email:", from)
		fmt.Println("SMTP attack tool")
		fmt.Println("List:")
		fmt.Println("0: quit \n1: send single email\n2: send to multiple emails \n3: smsOverSmtp\n4: Send smsOverSmtp to multiple numbers\n5: Set email\n6: Spam emails\n7. Spam smsOverSmtp")

		fmt.Scanln(&penis)

		switch penis {

		case "0":
			break loop
		case "1":
			emailPhish(from, password)
		case "2":
			multiSender(from, password)
		case "3":
			textOverSmtp(from, password)
		case "4":
			multiTextOverSmtp(from, password)
		case "5":
			emailName, pass = changeEmail()
		case "6":
			spamEmails(from, password)
		case "7":
			spamTextOverSmtp(from, password)
		}
	}
}

func spamEmails(from, password string) {

	var emailArray []string
        emailArray = readFileLineByLine("AttackerEmails.txt")
	var limit int
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
	var choice string	
	fmt.Println("Do you want to use the experimental mode?\n This mode uses all emails in list. (Y/N)?")
        fmt.Scanln(&choice)

        if choice == "Y" || choice == "y" {

	        fmt.Println("Warning! This tool sends from every email in AttackEmails.txt\nIf you have 2 emails in the list and enter 100, 200 emails should be sent.\nEnter limit here:")
        	fmt.Scanln(&limit)
        	for i := 0; i <= limit; i++ {

                	for i := 0; i <= (len(emailArray) - 1); i++ { 
                		subject2 := subject + strconv.Itoa(i)
                		sendEmailFromTemplate(emailArray[i], emailArray[i + 1], to, subject2, file, injectA, injectB)

                	}
        	}


        } else {

	fmt.Println("Enter limit here:")
	fmt.Scanln(&limit)
		for i := 0; i <= limit; i++ {

			subject2 := subject + strconv.Itoa(i)
			sendEmailFromTemplate(from, password, to, subject2, file, injectA, injectB)
	
		}
	}
}

func spamTextOverSmtp(from, password string) {
        var number string
        var text string
        var A string
        var gate string

        gateway := []string{"@msg.koodomobile.com", "@txt.att.net", "@cingularme.com", "@txt.bellmobility.com", "@sms.myboostmobile.com", "@text.scsouth1.com", "@vmobl.com", "@vtext.com", "@msg.telus.com", "@tms.suncom.com", "@messaging.sprintpcs.com", "@email.uscc.net", "@sms.sasktel.com", "@text.republicwireless.com", "@qwestmp.com", "@vtext.com", "@page.nextel.com", "@cspire1.com", "@sms.cricketwireless.net", "@mailmymobile.net", "@cwemail.com"}

        fmt.Println("what is the victims phone number?")
        fmt.Scanln(&number)

        fmt.Println("what is the text you want to send?")
        text = input()
	var limit int
	fmt.Println("How many do you want to send?")
	fmt.Scanln(&limit)

        fmt.Println("Do you know the cell provider gateway? Y/N")
        fmt.Scanln(&A)
        if A == "n" || A == "N" {
		for i := 0; i <= limit; i++ {
                	for i := 0; i <= len(gateway); i++ {

                		vic := number + gateway[i]
                		fmt.Println(vic)
                		email(from, password, text, vic)
                		fmt.Println(vic)

                	}
		}
        } else {

        	fmt.Println("what is the sms gateway?")
        	fmt.Scanln(&gate)

       		vic := number + gate
		for i := 0; i <= limit; i++ {
        		email(from, password, text, vic)
		}

        }

        time.Sleep(2 * time.Second)

}

func displayEmailList() {//filepath string) {

	var emailArray []string
        emailArray = readFileLineByLine("AttackerEmails.txt")
	
	for i := 0; i <= (len(emailArray) - 1); i++{
		
		if i % 2 == 0 {
			fmt.Println("\033[32mEmail", (i+1), " is:\033[0m", emailArray[i])
		} else {
			fmt.Println("\033[34mPass", i, " is:\033[0m", emailArray[i])
		}
	}

}

func changeEmail() (int, int) {
	//print selections
	displayEmailList()
	fmt.Println("Enter 0 to exit")
	
	var nig int
	var nog int
	var input int
	fmt.Println("Please enter the corresponding integer:")
	fmt.Scanln(&input)

	if input == 0 {
		
	} else {
		nig = input - 1
		nog = input
	}
	return nig, nog
}

func phishFromATemplate (from, password string) {

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
//This function injects the A and B variables in the template. Those can really be whatever you want. I need to work on dynamic naming.
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
	var choice string
	fmt.Println("Do you want to use a template? (Y/N)")
	fmt.Scanln(&choice)

	if choice == "Y" || choice == "y" {
		phishFromATemplate (from, password)

	} else {

		var to, subject string

        	fmt.Println("Type victim email address:")
        	fmt.Scanln(&to)

        	fmt.Println("Subject:")
	        subject = getInput("")

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

func readFileLineByLine(filePath string) []string {

	var array []string

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		array = append(array, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return array
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
	
	input, _ := reader.ReadString('\n')
	
	return input
}

//This function is for sendinng emails to multiple users from a file list. list file? idfk
func multiSender(from, password string) {
	var file string
        var sub string
        var template string
        var name string
        var link string

	fmt.Println("Paste your list of emails here, it should just be emails in a list with no interuptions.")
	fmt.Scanln(&file)
	
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
	//mess is the template and the subject getting sent to the email	
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
 
    // toList is list of email address that email is to be sent.
    toList := []string{victim}
    //email server host 
    host := "smtp.gmail.com"
    //email server port
    port := "587"
    //message body
    body := []byte(msg)
    //auth data
    auth := smtp.PlainAuth("", from, password, host)
 
    err := smtp.SendMail(host+":"+port, auth, from, toList, body)
 
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
//this function calls figlet to print my logo
func Figlet(text string, font string) (string) {
	// Run `figlet` with the specified text and font
	cmd := exec.Command("figlet", "-f", font, text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	return string(output)
}

func check (err error) {

	if err != nil {
        	fmt.Println(err)
        	os.Exit(1)
    	}
}

