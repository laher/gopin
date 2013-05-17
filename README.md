gopin
=====

experimental go-get fork with support for tags and alternative repos.

The aim is to provide a mechanism for repeatable builds.

Installation
------------

    go get github.com/laher/gopin

Usage
-----
The following will install a *specific (tagged) version* of a *fork* of 'mkdo'.

    gopin get -u -repo bitbucket.org/laher/mkdo -tag v0.0.7 github.com/laher/mkdo

So, the package is downloaded from bitbucket.org, but keeps the original github.com import paths.

In this way, gopin also provides a straightforward way to manage package forks, without having to edit import paths.

Hurray!

Also, you can list available tags for a particular import.

     gopin tags github.com/laher/mkdo


Warnings
--------
This is very experimental.
 * I have only used it on a very simple project which only imports standard library packages.
 * I haven't got this working for vanity URLs yet - only the 'static' sites supported currently (wouldn't take long).

To NOT Do
---------

I have decided NOT to make this a tool which deals with dependency heirarchies.

If I do make a config file then I'll make it system-level (or potentially user-level or GOPATH-level).

To do (maybe oneday)
--------------------

1. 'pin' task. This will simply create/update a config file, defining a pinned dependency for the system.

        gopin pin -repo bitbucket.org/laher/mkdo -tag v0.0.7 github.com/laher/mkdo

2. 'getdeps' task. This will parse the manifest and install all the pinned dependency definitions. (can't think of a better name right now)

        gopin getdeps


3. Version range support? e.g. 'v1.*' or similar.
