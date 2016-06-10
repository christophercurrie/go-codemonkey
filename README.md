A personal collection of small tools written in go.

[![Build Status](https://travis-ci.org/christophercurrie/go-codemonkey.svg?branch=master)](https://travis-ci.org/christophercurrie/go-codemonkey)

# yaml2json

A bare-bones YAML to JSON converter. Written because I tried several
others written using a variety of tools and none of them met my needs.
Supports zero customization or JSON pretty-printing; I recommend
[jq](https://stedolan.github.io/jq/) for post processing if your
needs are more complex.

# json2yaml

A bare-bones JSON to YAML converter. Written because it was a trivial
mirror of `yaml2json`.
