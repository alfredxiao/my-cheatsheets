# Components
- The mail user agent (`MUA`) is an interface for users to read messages stored in their mailboxes. MUAs do not receive messages; they only display messages that are already in the user’s mailbox.
- The mail delivery agent (`MDA`) is a program that delivers messages to a local user’s inbox.
- The mail transfer agent (`MTA`) sends incoming emails (and outgoing emails being delivered locally) to a mail delivery agent (`MDA`) or local user’s inbox. For outbound messages being transferred to a remote system, the agent establishes a communication link with another `MTA` program on the remote host to transfer the email.
- As you can see, the Linux email server is normally divided into three separate functions: The lines between these three functions are often fuzzy. Some Linux email packages combine functionality for the MTA and MDA functions, whereas others combine the MDA and MUA functions.
