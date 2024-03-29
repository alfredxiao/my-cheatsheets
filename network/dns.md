# Common Record Types
- `A` (address)
  - Used to map a host name to an IP address. Generally, A records are IP addresses. If a computer consists of multiple IP addresses, adapter cards, or both, it must possess multiple address records.
  - e.g. `142.250.70.174`
- `AAAA`
  - IPV6 address
- `CNAME` (canonical name)
  - Can be used to set an alias for the host name
  - e.g. `ghs.googlehosted.com`
- `MX` (mail exchange)
  - Permits mail to be sent to the right mail servers located in the domain. Other than IP addresses, MX records include fully-qualified domain names.
  - e.g. `ASPMX.L.GOOGLE.COM`
- `NS` (name server)
  - Describes a name server for the domain that permits DNS lookups within several zones. Every primary as well as secondary name server must be reported via this record.
- `PTR` (pointer)
  - Creates a pointer, which maps an IP address to the host name in order to do reverse lookups.
- `SOA` (start of authority)
  - Declares the most authoritative host for the zone. Every zone file should include an SOA record, which is generated automatically when the user adds a zone.
- `TXT` (text record)
  - Permits the insertion of arbitrary text into a DNS record. These records add SPF records into a domain.
  - e.g. `google-site-verification=6tTalLzrBXBO4Gy9700TAbpg2QTKzGYEuZ_Ls69jle8`
