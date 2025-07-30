***GopherSMTP***

What is GopherSMTP?

GopherSMTP is a SMTP phishing tool written in golang that (at it's core) automates sending emails. GopherSMTP also takes advantage of smtp to sms gateways to create convincing sms messages.
Basically it's a terminal tool that lets you send scary emails and texts. 

Here's what you can do:

in the form of either an email or sms message you can:

choose to send your message to a single person or list of people.

use templates for your message (You can write your own templates)

Or you can write your own message in the terminal if you like

I'd like to be able to spin up new emails from my own mail server but those would be traced to your ip 
and any mail sent from that server would just end up in the spam folder.

we may add an option for google gemini to generate a phishing message. and flood the targets inbox with multiple emails.

Who made this?

I made this! Me, M.E

Why would you use this?

I don't know. To scare your friends.

Where can I use this software?

You can use this on damn near anything with a golang compiler and an internet connection.

When is this useful?
This isn't useful unless you want to scare someone and know what you're doing. An example of when to use it is that my friend got a needle at that hospital. 
I know what hospital he goes to, so I make a gmail in that hospitals name and use smtpToSmS to tell him his shot was contaminated. It's awesome when it works.
Pranks are fun.

How does it work?

GopherSMTP works by sending emails. Any emails you want to use are stored in a text file called "AttackerEmails.txt", these are the emails you send from. Email is the first line, and your app password should be on the second line.
The email selecor reads the file line by line and expects the gmail to be on the odd lines, and the password for the subsequent email to be on the line below it. So email1@gmail.com goes on line 1, and the password for that goes on line 2.
If you can't figure out those instructions, you're the problem.
