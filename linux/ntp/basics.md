# Basics
- NTP uses what is called a clock **stratum** scheme, which provides a layered approach to accessing correct time sources. The stratums are numbered from 0 to 15. The devices at stratum 0 are highly accurate time-keeping hardware devices, such as atomic clocks.

# SNTP
- The Simple Network Time Protocol (SNTP) mentioned earlier in the chapter is a simplified version of NTP. Although NTP can achieve high levels of time accuracy, SNTP is more for applications or systems not dependent on accurate clocks.

# Choose NTP Servers
- if ISP provides a time server, use it
- common name is `http://pool.ntp.org`
- location based `http://ca.pool.ntp.org`

# Leap Second
- Because the earth’s rotation has been slowing down, our actual day is about 0.001 seconds less than 24 hours
- To combat this problem, Google introduced free public time servers that use NTP and smear the leap second over the course of time so that there is no need to issue a leap second announcement.
- This is called **leap-smearing**. The Google leap-smearing NTP servers are `timen.google.com`, where n is set to 1 through 4. If you choose to use a leap-smearing time server on your system, you should not mix in time servers on your NTP client program that do not employ this technique.

# NTP Client (or NTP Daemon)
- `ntpd`
- `chronyd` is newer
