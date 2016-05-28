# goimserver
GoIMServer - a simple XMPP server

Base on the work of https://github.com/tam7t/xmpp.

Status: WIP

## Test suit result

Coming soon...

## Goals

### Primary goals

* Core xmpp server 
  * [RFC 6120 - Extensible Messaging and Presence Protocol (XMPP): Core](https://datatracker.ietf.org/doc/rfc6120/)
  * [RFC 6121 - Extensible Messaging and Presence Protocol (XMPP): Instant Messaging and Presence](https://datatracker.ietf.org/doc/rfc6121/)
  * [RFC 7622 - Extensible Messaging and Presence Protocol (XMPP): Address Format](https://datatracker.ietf.org/doc/rfc7622/)
  * [RFC 7395 - An Extensible Messaging and Presence Protocol (XMPP) Subprotocol for WebSocket](https://datatracker.ietf.org/doc/rfc7395/)
* Standard plugin interfaces for XEPs (internal, DLL, ... ?)
* Multi-User-Chat (MUC)
  * [XEP-0045: Multi-User Chat](http://www.xmpp.org/extensions/xep-0045.html)
* PubSub
  * [XEP-0060: Publish-Subscribe](http://www.xmpp.org/extensions/xep-0060.html)
  * [XEP-0163: Personal Eventing Protocol](http://www.xmpp.org/extensions/xep-0163.html)
* BOSH
  * [XEP-0124: Bidirectional-streams Over Synchronous HTTP](http://www.xmpp.org/extensions/xep-0124.html)
  * [XEP-0206: XMPP Over BOSH](http://www.xmpp.org/extensions/xep-0206.html)
* Deep systemd integration (optional feature)

### secondary goals

* More XEP support - [XEPs](http://xmpp.org/extensions/index.html)
* ...

## Tips and Hints

https://cjones.org/hg/go-xmpp2/file/tip/xmpp