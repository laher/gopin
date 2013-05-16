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

Warnings
--------
This is very experimental.
 * I have only used it on a very simple project which only imports standard library packages.
 * I haven't got this working for vanity URLs yet - only the 'static' sites supported currently (wouldn't take long).

To do
-----
The idea is that gopinned projects can have a manifest file, e.g. package.json, which would define dependencies.
It'd be pretty easy to add 2 'commands'. 
One to add dependencies, and a second to actually pull in these definitions.

1. 'pin' task. This will simply create/update a manifest file, defining a pinned dependency.


        gopin pin -repo bitbucket.org/laher/mkdo -tag v0.0.7 github.com/laher/mkdo

2. 'getdeps' task. This will parse the manifest and install all the pinned dependency definitions. (can't think of a better name right now)


        gopin getdeps
        

3. Recursively locate manifest files.
