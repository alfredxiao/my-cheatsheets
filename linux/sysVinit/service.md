# Status
`service httpd status`

# Start
`service httpd start`

# Stop
`service httpd Stop`

# Restart
`service httpd restart`

# Reload (service's own config)
`service httpd reload`

# List all status
`service --status-all`

# Disable a service
`update-rc.d -f SERVICE-NAME remove`, The -f option on the preceding command is needed only if a script exists for the service (SERVICE-NAME) in the /etc/init.d/ directory
