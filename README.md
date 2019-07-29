<h1 align="center">ipdiscover</h1>
<h4 align="center">🔍 A simple tool to obtain long lists of ips from domains using goroutines</h4>
<p align="center">
<a href="https://travis-ci.org/JavierOlmedo/ipdiscover"><img src="https://api.travis-ci.org/JavierOlmedo/ipdiscover.svg?branch=master"></a>
<a href="https://www.python.org/"><img src="https://img.shields.io/badge/Golang-blue.svg"></a>
<a href="https://raw.githubusercontent.com/JavierOlmedo/ipdiscover/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-blue.svg"></a>
<a href="https://hackpuntes.com"><img src="https://img.shields.io/badge/website-hackpuntes.com-blue.svg"></a>
<a href="https://twitter.com/jjavierolmedo"><img src="https://img.shields.io/badge/twitter-@jjavierolmedo-blue.svg"></a>
</p>

## About ipdiscover
ipdiscover is a golang tool designed to obtain long lists of IPs from domains. It helps penetration tester and bug bounty hunters to quickly collect IPs to work with other tools, for example, nmap.

## Install

```
▶ go get -u github.com/JavierOlmedo/ipdiscover
```

## Usage

ipdiscover accepts line-delimited domains on `stdin`:

```
▶ cat domains.txt
google.es
google.net
google.com
google.edu
▶ cat domains.txt | ipdiscover
google.es;172.217.16.227
google.net;216.58.211.36
google.com;172.217.17.14
google.edu;Unknown
```

Only one domain:
```
▶ ipdiscover google.es
google.es;172.217.16.227
```

## Concurrency

You can set the concurrency level with the `-t` flag and specifying a number (default 23 concurrencies):

```
▶ cat domains.txt | ipdiscover -t 99
```

## All IPs Nslookup

You can get all ips of a domain that solves nslookup using the `-a` flag:

```
▶ cat domains.txt | httprobe -a
```

## Credit
This tool was inspired by @tomnomnom [scripts](https://github.com/tomnomnom?utf8=%E2%9C%93&tab=repositories&q=&type=&language=go). Thanks to them I learned to program in Go!

> *Made with ❤️ in Spain*