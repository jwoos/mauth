# Go Auth
This is a command line program that is meant to serve as a drop in replacement for OTP apps such as Google Authenticator or Authy. HOTP (RFC 4226) and TOTP (RFC 6238) are supported, each supporting base32 encoded secrets as well as differing length of the passwords.

## Why?
Have you ever sat at your computer and tried to log in to a site and remember that you need your phone for the two factor authentication? This is a common occurrence for me and I thought I might implement a CLI program since I am always using the commandline.

### Note
Multifactor authentication is a tricky topic - it always comes down to a tradeoff of convenience and security. Authentication is typically done by one or more parts of who you are, what you have, and what you know. The password is clearly what you know, but where does this OTP fall into play? It's certainly not who we are and I can argue that it's not something we know since it will change depending on time, for TOTP, and on the counter for HOTP. So it comes down to what we have.
Multifactor authentication was made to fulfill the part of what we have. It's supposed to act as a convenient replacement for an actual hardware device by emulating it via software. So having it on multiple devices doesn't really make sense from a security point of view - typically it should be only on one device and the secret should never be shared or stored (this is also debatable e.g. on paper in a secure location, but I won't get into that). However, it's hard to argue that it's not convenient. I'm not sure if the tradeoff is worth it but people do in general opt for the convenience, for example Authy is another client that syncs multifactor accounts on it. This debate of security is one worth having and one that should be kept in mind while using this or any other multifactor program in addition to the one on your phone.

## Build
You should have Go installed and `$GOPATH` defined.

```
go get github.com/jwoos/go_auth
```

This will pull the package, compile and drop the executable `go_auth` in `$GOPATH/bin`

## Usage
```
go_auth [options] ACCOUNT
```
`ACTION`
- `get`: Display an account
- `add`: Add an account
- `edit`: Edit an account
- `delete`: Delete an account

`ACCOUNT`: the exact account name or a search
