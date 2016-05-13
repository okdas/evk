# evk

`evk` (environment variable keychain) is a simple command line tool to help manage secret environment variables and store them in OS X Keychain.

Note: under development. Things may won't work until this line will be removed!

Inspired by [mmfa](https://github.com/thbishop/mmfa) and [envchain](https://github.com/sorah/envchain) (which is almost the same, but a little bit different way to work with tool).

## Install

...

## Quick Start

Bucket — a bunch of variables, useful for separating different environments with same secret environment variables (e.g. Prod and Dev AWS accounts). If bucket will not be specified, `main` bucket will be used to store and get data.

### Add a variables

That is how you can add environment variable to `main` bucket:

```
$ evk add ENVIRONMENT_VARIABLE
```

You also can specify bucket:

```
$ evk add -b aws_production ENVIRONMENT_VARIABLE
```

### Get secret environment variable

You can set variables from `main` bucket to your current shell session:

```
$ eval $(evk get)
```

You can also specify bucket to get environment variables from:

```
$ eval $(evk get -b aws_production)
```

<!--
### Remove variable and bucket

You need to specify bucket to remove:

```
$ evk remove -b aws_production
```

If you will delete `main` bucket, all environment variables will be deleted from keychain.

### List Buckets

You can list existing buckets:

```
$ evk list
```
-->

## TODO
* Commands execution with EnVars overriding (e.g. `evk run -b openstack_bucket -c "swift list"`)
* Rewrite search with C instead of os.Exec (dependency [go-keychain](https://github.com/okdas/go-keychain))

## Contribute
* Fork the project
* Make your feature addition or bug fix (with tests and docs) in a topic branch
* Send a pull request and I'll get it integrated
