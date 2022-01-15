# Basics
## Understanding output
- Lines beginning with `;` are comments not part of the information.
- Sections of the output:
  - `->>HEADER<<-`
  - `QUESTION SECTION:`
  - `ANSWER SECTION:`
- `IN` means this is an Internet lookup (in the Internet class).

# Short Answer Mode
- `dig google.com +short`
- `dig google.com MX +short`
- `dig +short google.com MX`

# For a record type
- `dig hostname MX`

# Specify a DNS Server to query
- `dig @8.8.8.8 yahoo.com`

# Find TTL
- `dig google.com A`
```
;; ANSWER SECTION:
google.com.		190	IN	A	142.250.70.174
```
- `190` is ttl in seconds
## to find the raw TTL value
- `dig google.com NS`
- which outputs `ns1.google.com.`
- then `dig @ns1.google.com google.com`
- which outputs
```
google.com.		300	IN	A	142.250.70.238
```
- where `300` is the raw TTL

# Answer Only
* `dig yahoo.com +nocomments +noquestion +noauthority +noadditional +nostats`

# All DNS Records
- `dig yahoo.com ANY +noall +answer`

# Reverse Lookup
* `dig -x 123.123.123.123`
