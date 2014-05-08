Notes
====

## I have used:

- [Render](https://github.com/martini-contrib/render) to render templates
- [Binding](https://github.com/martini-contrib/binding) to bind form fields in a struct
- [Viper](https://github.com/spf13/viper) to use a user password configuration file. The example file config.yaml
includes only one admin user (password 'secret'). If the file is empty, the password is set inside the server.go file.
- [Bcrypt](http://godoc.org/code.google.com/p/go.crypto/bcrypt) to decrypt the has from the configuraiton file.
