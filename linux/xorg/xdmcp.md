XDMCP is a remote desktop protocol. With XDMCP, one computer A running X11 can connecting computer B running X11, and interact with computer B as if one were physically at computer B. XDMCP is integrated into X.org, the default X11 server in Ubuntu. XDMCP also needs to be implemented by the display manager.

Warnings

    XDMCP is inherently insecure as it does not encrypt your traffic. XDMCP is analogous to telnet, and therefore they share the same security issues.
  XDMCP uses a large amount of bandwidth because it uses no compression.
