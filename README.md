# Awesome Go

<a href="https://awesome-go.com/"><img align="right" src="https://github.com/avelino/awesome-go/raw/main/tmpl/assets/logo.png" alt="awesome-go" title="awesome-go" /></a>

[![Build Status](https://github.com/avelino/awesome-go/actions/workflows/tests.yaml/badge.svg?branch=main)](https://github.com/avelino/awesome-go/actions/workflows/tests.yaml?query=branch%3Amain)
[![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/sindresorhus/awesome)
[![Slack Widget](https://img.shields.io/badge/join-us%20on%20slack-gray.svg?longCache=true&logo=slack&colorB=red)](https://gophers.slack.com/messages/awesome)
[![Netlify Status](https://api.netlify.com/api/v1/badges/83a6dcbe-0da6-433e-b586-f68109286bd5/deploy-status)](https://app.netlify.com/sites/awesome-go/deploys)
[![Track Awesome List](https://www.trackawesomelist.com/badge.svg)](https://www.trackawesomelist.com/avelino/awesome-go/)

We use the _[Golang Bridge](https://github.com/gobridge/about-us/blob/master/README.md)_ community Slack for instant communication, follow the [form here to join](https://invite.slack.golangbridge.org/).

<a href="https://www.producthunt.com/posts/awesome-go?utm_source=badge-featured&utm_medium=badge&utm_souce=badge-awesome-go" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=291535&theme=light" alt="awesome-go - Curated list awesome Go frameworks, libraries and software | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>

**Sponsorships:**

_Special thanks to_

<div align="center">
<table cellpadding="5">
<tbody align="center">
<tr>
<td>
<a href="https://bit.ly/awesome-go-workos">
<img src="https://avelino.run/sponsors/workos-logo-white-bg.svg" width="200" alt="WorkOS"><br/>
<b>Your app, enterprise-ready.</b><br/>
<sub>Start selling to enterprise customers with just a few lines of code.</sub><br/>
<sup>Add Single Sign-On (and more) in minutes instead of months.</sup>
</a>
</td>
<td>
<a href="https://bit.ly/awesome-go-keygen">
<img src="https://avelino.run/sponsors/keygen-logo.png" width="200" alt="keygen"><br/>
<b>An open, source-available software licensing and distribution API.</b><br/>
<sub>Securely license and distribute Go applications with a single API.</sub><br>
<sup>Add auto updates with only a few lines of code.</sup>
</a>
</td>
</tr>
<tr>
<td colspan="2">
<a href="https://bit.ly/awesome-go-digitalocean">
<img src="https://avelino.run/sponsors/do_logo_horizontal_blue-210.png" width="200" alt="Digital Ocean">
</a>
</td>
</tr>
</tbody>
</table>
</div>

**Awesome Go has no monthly fee**_, but we have employees who **work hard** to keep it running. With money raised, we can repay the effort of each person involved! You can see how we calculate our billing and distribution as it is open to the entire community. Want to be a supporter of the project click [here](mailto:avelinorun+oss@gmail.com?subject=awesome-go%3A%20project%20support)._

> A curated list of awesome Go frameworks, libraries, and software. Inspired by [awesome-python](https://github.com/vinta/awesome-python).

**Contributing:**

Please take a quick gander at the [contribution guidelines](https://github.com/avelino/awesome-go/blob/main/CONTRIBUTING.md) first. Thanks to all [contributors](https://github.com/avelino/awesome-go/graphs/contributors); you rock!

> _If you see a package or project here that is no longer maintained or is not a good fit, please submit a pull request to improve this file. Thank you!_

## Contents

- [Awesome Go](#awesome-go)
  - [Contents](#contents)
  - [Artificial Intelligence](#artificial-intelligence)
  - [Audio and Music](#audio-and-music)
  - [Authentication and OAuth](#authentication-and-oauth)
  - [Blockchain](#blockchain)
  - [Bot Building](#bot-building)
  - [Build Automation](#build-automation)
  - [Command Line](#command-line)
    - [Advanced Console UIs](#advanced-console-uis)
    - [Standard CLI](#standard-cli)
  - [Configuration](#configuration)
  - [Continuous Integration](#continuous-integration)
  - [CSS Preprocessors](#css-preprocessors)
  - [Data Structures and Algorithms](#data-structures-and-algorithms)
    - [Bit-packing and Compression](#bit-packing-and-compression)
    - [Bit Sets](#bit-sets)
    - [Bloom and Cuckoo Filters](#bloom-and-cuckoo-filters)
    - [Data Structure and Algorithm Collections](#data-structure-and-algorithm-collections)
    - [Iterators](#iterators)
    - [Maps](#maps)
    - [Miscellaneous Data Structures and Algorithms](#miscellaneous-data-structures-and-algorithms)
    - [Nullable Types](#nullable-types)
    - [Queues](#queues)
    - [Sets](#sets)
    - [Text Analysis](#text-analysis)
    - [Trees](#trees)
    - [Pipes](#pipes)
  - [Database](#database)
    - [Caches](#caches)
    - [Databases Implemented in Go](#databases-implemented-in-go)
    - [Database Schema Migration](#database-schema-migration)
    - [Database Tools](#database-tools)
    - [SQL Query Builders](#sql-query-builders)
  - [Database Drivers](#database-drivers)
    - [Interfaces to Multiple Backends](#interfaces-to-multiple-backends)
    - [Relational Database Drivers](#relational-database-drivers)
    - [NoSQL Database Drivers](#nosql-database-drivers)
    - [Search and Analytic Databases](#search-and-analytic-databases)
  - [Date and Time](#date-and-time)
  - [Distributed Systems](#distributed-systems)
  - [Dynamic DNS](#dynamic-dns)
  - [Email](#email)
  - [Embeddable Scripting Languages](#embeddable-scripting-languages)
  - [Error Handling](#error-handling)
  - [File Handling](#file-handling)
  - [Financial](#financial)
  - [Forms](#forms)
  - [Functional](#functional)
  - [Game Development](#game-development)
  - [Generators](#generators)
  - [Geographic](#geographic)
  - [Go Compilers](#go-compilers)
  - [Goroutines](#goroutines)
  - [GUI](#gui)
  - [Hardware](#hardware)
  - [Images](#images)
  - [IoT (Internet of Things)](#iot-internet-of-things)
  - [Job Scheduler](#job-scheduler)
  - [JSON](#json)
  - [Logging](#logging)
  - [Machine Learning](#machine-learning)
  - [Messaging](#messaging)
  - [Microsoft Office](#microsoft-office)
    - [Microsoft Excel](#microsoft-excel)
  - [Miscellaneous](#miscellaneous)
    - [Dependency Injection](#dependency-injection)
    - [Project Layout](#project-layout)
    - [Strings](#strings)
    - [Uncategorized](#uncategorized)
  - [Natural Language Processing](#natural-language-processing)
    - [Language Detection](#language-detection)
    - [Morphological Analyzers](#morphological-analyzers)
    - [Slugifiers](#slugifiers)
    - [Tokenizers](#tokenizers)
    - [Translation](#translation)
    - [Transliteration](#transliteration)
  - [Networking](#networking)
    - [HTTP Clients](#http-clients)
  - [OpenGL](#opengl)
  - [ORM](#orm)
  - [Package Management](#package-management)
  - [Performance](#performance)
  - [Query Language](#query-language)
  - [Resource Embedding](#resource-embedding)
  - [Science and Data Analysis](#science-and-data-analysis)
  - [Security](#security)
  - [Serialization](#serialization)
  - [Server Applications](#server-applications)
  - [Stream Processing](#stream-processing)
  - [Template Engines](#template-engines)
  - [Testing](#testing)
  - [Text Processing](#text-processing)
    - [Formatters](#formatters)
    - [Markup Languages](#markup-languages)
    - [Parsers/Encoders/Decoders](#parsersencodersdecoders)
    - [Regular Expressions](#regular-expressions)
    - [Sanitation](#sanitation)
    - [Scrapers](#scrapers)
    - [RSS](#rss)
    - [Utility/Miscellaneous](#utilitymiscellaneous)
  - [Third-party APIs](#third-party-apis)
  - [Utilities](#utilities)
  - [UUID](#uuid)
  - [Validation](#validation)
  - [Version Control](#version-control)
  - [Video](#video)
  - [Web Frameworks](#web-frameworks)
    - [Middlewares](#middlewares)
      - [Actual middlewares](#actual-middlewares)
      - [Libraries for creating HTTP middlewares](#libraries-for-creating-http-middlewares)
    - [Routers](#routers)
  - [WebAssembly](#webassembly)
  - [Windows](#windows)
  - [XML](#xml)
  - [Zero Trust](#zero-trust)
  - [Code Analysis](#code-analysis)
  - [Editor Plugins](#editor-plugins)
  - [Go Generate Tools](#go-generate-tools)
  - [Go Tools](#go-tools)
  - [Software Packages](#software-packages)
    - [DevOps Tools](#devops-tools)
    - [Other Software](#other-software)
- [Resources](#resources)
  - [Benchmarks](#benchmarks)
  - [Conferences](#conferences)
  - [E-Books](#e-books)
    - [E-books for purchase](#e-books-for-purchase)
    - [Free e-books](#free-e-books)
  - [Gophers](#gophers)
  - [Meetups](#meetups)
  - [Style Guides](#style-guides)
  - [Social Media](#social-media)
    - [Twitter](#twitter)
    - [Reddit](#reddit)
  - [Websites](#websites)
    - [Tutorials](#tutorials)
    - [Guided Learning](#guided-learning)

**[⬆ back to top](#contents)**

## Artificial Intelligence

_Libraries for building programs that leverage AI._

- [langchaingo(stars: 1860)](https://github.com/tmc/langchaingo) - LangChainGo is a framework for developing applications powered by language models.
- [LocalAI(stars: 14803)](https://github.com/mudler/LocalAI) - Open Source OpenAI alternative, self-host AI models.
- [Ollama(stars: 28712)](https://github.com/jmorganca/ollama) - Run large language models locally.

**[⬆ back to top](#contents)**

## Audio and Music

_Libraries for manipulating audio._

- [flac(stars: 275)](https://github.com/mewkiz/flac) - Native Go FLAC encoder/decoder with support for FLAC streams.
- [gaad(stars: 116)](https://github.com/Comcast/gaad) - Native Go AAC bitstream parser.
- [GoAudio(stars: 307)](https://github.com/DylanMeeus/GoAudio) - Native Go Audio Processing Library.
- [gosamplerate(stars: 31)](https://github.com/dh1tw/gosamplerate) - libsamplerate bindings for go.
- [id3v2(stars: 311)](https://github.com/bogem/id3v2) - ID3 decoding and encoding library for Go.
- [malgo(stars: 241)](https://github.com/gen2brain/malgo) - Mini audio library.
- [minimp3(stars: 116)](https://github.com/tosone/minimp3) - Lightweight MP3 decoder library.
- [Oto(stars: 1411)](https://github.com/hajimehoshi/oto) - A low-level library to play sound on multiple platforms.
- [PortAudio(stars: 644)](https://github.com/gordonklaus/portaudio) - Go bindings for the PortAudio audio I/O library.

**[⬆ back to top](#contents)**

## Authentication and OAuth

_Libraries for implementing authentication schemes._

- [authboss(stars: 3561)](https://github.com/volatiletech/authboss) - Modular authentication system for the web. It tries to remove as much boilerplate and "hard things" as possible so that each time you start a new web project in Go, you can plug it in, configure it, and start building your app without having to build an authentication system each time.
- [branca(stars: 70)](https://github.com/essentialkaos/branca) - branca token [specification implementation](https://github.com/tuupola/branca-spec) for Golang 1.15+.
- [casbin(stars: 16236)](https://github.com/hsluoyz/casbin) - Authorization library that supports access control models like ACL, RBAC, and ABAC.
- [cookiestxt(stars: 12)](https://github.com/mengzhuo/cookiestxt) - provides a parser of cookies.txt file format.
- [go-guardian(stars: 506)](https://github.com/shaj13/go-guardian) - Go-Guardian is a golang library that provides a simple, clean, and idiomatic way to create powerful modern API and web authentication that supports LDAP, Basic, Bearer token, and Certificate based authentication.
- [go-jose(stars: 202)](https://github.com/go-jose/go-jose) - Fairly complete implementation of the JOSE working group's JSON Web Token, JSON Web Signatures, and JSON Web Encryption specs.
- [gologin(stars: 1673)](https://github.com/dghubble/gologin) - chainable handlers for login with OAuth1 and OAuth2 authentication providers.
- [gorbac(stars: 1521)](https://github.com/mikespook/gorbac) - provides a lightweight role-based access control (RBAC) implementation in Golang.
- [gosession](http://github.com/Kwynto/gosession) - This is quick session for net/http in GoLang. This package is perhaps the best implementation of the session mechanism, or at least it tries to become one.
- [goth(stars: 4568)](https://github.com/markbates/goth) - provides a simple, clean, and idiomatic way to use OAuth and OAuth2. Handles multiple providers out of the box.
- [jeff(stars: 258)](https://github.com/abraithwaite/jeff) - Simple, flexible, secure, and idiomatic web session management with pluggable backends.
- [jwt(stars: 341)](https://github.com/pascaldekloe/jwt) - Lightweight JSON Web Token (JWT) library.
- [jwt(stars: 630)](https://github.com/cristalhq/jwt) - Safe, simple, and fast JSON Web Tokens for Go.
- [jwt-auth(stars: 231)](https://github.com/adam-hanna/jwt-auth) - JWT middleware for Golang http servers with many configuration options.
- [jwt-go(stars: 5819)](https://github.com/golang-jwt/jwt) - A full featured implementation of JSON Web Tokens (JWT). This library supports the parsing and verification as well as the generation and signing of JWTs.
- [keto(stars: 4391)](https://github.com/ory/keto) - Open Source (Go) implementation of "Zanzibar: Google's Consistent, Global Authorization System". Ships gRPC, REST APIs, newSQL, and an easy and granular permission language. Supports ACL, RBAC, and other access models.
- [loginsrv(stars: 1902)](https://github.com/tarent/loginsrv) - JWT login microservice with pluggable backends such as OAuth2 (Github), htpasswd, osiam.
- [oauth2(stars: 4992)](https://github.com/golang/oauth2) - Successor of goauth2. Generic OAuth 2.0 package that comes with JWT, Google APIs, Compute Engine, and App Engine support.
- [osin(stars: 1858)](https://github.com/openshift/osin) - Golang OAuth2 server library.
- [otpgen(stars: 133)](https://github.com/grijul/otpgen) - Library to generate TOTP/HOTP codes.
- [otpgo(stars: 63)](https://github.com/jltorresm/otpgo) - Time-Based One-Time Password (TOTP) and HMAC-Based One-Time Password (HOTP) library for Go.
- [paseto(stars: 776)](https://github.com/o1egl/paseto) - Golang implementation of Platform-Agnostic Security Tokens (PASETO).
- [permissions2(stars: 496)](https://github.com/xyproto/permissions2) - Library for keeping track of users, login states, and permissions. Uses secure cookies and bcrypt.
- [scope(stars: 33)](https://github.com/SonicRoshan/scope) - Easily Manage OAuth2 Scopes In Go.
- [scs(stars: 1770)](https://github.com/alexedwards/scs) - Session Manager for HTTP servers.
- [securecookie(stars: 74)](https://github.com/chmike/securecookie) - Efficient secure cookie encoding/decoding.
- [session(stars: 112)](https://github.com/icza/session) - Go session management for web servers (including support for Google App Engine - GAE).
- [sessions(stars: 75)](https://github.com/adam-hanna/sessions) - Dead simple, highly performant, highly customizable sessions service for go http servers.
- [sessionup(stars: 122)](https://github.com/swithek/sessionup) - Simple, yet effective HTTP session management and identification package.
- [sjwt(stars: 114)](https://github.com/brianvoe/sjwt) - Simple jwt generator and parser.

**[⬆ back to top](#contents)**

## Blockchain

_Tools for building blockchains._

- [cometbft(stars: 423)](https://github.com/cometbft/cometbft) - A distributed, Byzantine fault-tolerant, deterministic state machine replication engine. It is a fork of Tendermint Core and implements the Tendermint consensus algorithm.
- [cosmos-sdk(stars: 5621)](https://github.com/cosmos/cosmos-sdk) - A Framework for Building Public Blockchains in the Cosmos Ecosystem.
- [go-ethereum(stars: 44374)](https://github.com/ethereum/go-ethereum) - Official Go implementation of the Ethereum protocol.
- [gossamer(stars: 400)](https://github.com/ChainSafe/gossamer) - A Go implementation of the Polkadot Host.
- [kubo(stars: 15564)](https://github.com/ipfs/kubo) - A blockchain framework implemented in Go. It provides content-addressable storage which can be used for decentralized storage in DApps. It is based on the IPFS protocol.
- [solana-go(stars: 481)](https://github.com/gagliardetto/solana-go) - Go library to interface with Solana JSON RPC and WebSocket interfaces.
- [tendermint(stars: 5593)](https://github.com/tendermint/tendermint) - High-performance middleware for transforming a state machine written in any programming language into a Byzantine Fault Tolerant replicated state machine using the Tendermint consensus and blockchain protocols.

**[⬆ back to top](#contents)**

## Bot Building

_Libraries for building and working with bots._

- [bot(stars: 258)](https://github.com/go-telegram/bot) - Zero-dependencies Telegram Bot library with additional UI components
- [echotron(stars: 298)](https://github.com/NicoNex/echotron) - An elegant and concurrent library for Telegram Bots in Go.
- [ephemeral-roles(stars: 83)](https://github.com/ewohltman/ephemeral-roles) - A Discord bot for managing ephemeral roles based upon voice channel member presence.
- [go-chat-bot(stars: 806)](https://github.com/go-chat-bot/bot) - IRC, Slack & Telegram bot written in Go.
- [go-joe](https://joe-bot.net) - A general-purpose bot library inspired by Hubot but written in Go.
- [go-sarah(stars: 259)](https://github.com/oklahomer/go-sarah) - Framework to build a bot for desired chat services including LINE, Slack, Gitter, and more.
- [go-tgbot(stars: 117)](https://github.com/olebedev/go-tgbot) - Pure Golang Telegram Bot API wrapper, generated from swagger file, session-based router, and middleware.
- [go-twitch-irc(stars: 317)](https://github.com/gempir/go-twitch-irc) - Library to write bots for twitch.tv chat
- [Golang CryptoTrading Bot(stars: 966)](https://github.com/saniales/golang-crypto-trading-bot) - A golang implementation of a console-based trading bot for cryptocurrency exchanges.
- [govkbot(stars: 46)](https://github.com/nikepan/govkbot) - Simple Go [VK](https://vk.com) bot library.
- [hanu(stars: 150)](https://github.com/sbstjn/hanu) - Framework for writing Slack bots.
- [Kelp(stars: 1043)](https://github.com/stellar/kelp) - official trading and market-making bot for the [Stellar](https://www.stellar.org/) DEX. Works out-of-the-box, written in Golang, compatible with centralized exchanges and custom trading strategies.
- [larry(stars: 153)](https://github.com/ezeoleaf/larry) - Larry 🐦 is a really simple Twitter bot generator that tweets random repositories from Github built in Go.
- [margelet(stars: 81)](https://github.com/zhulik/margelet) - Framework for building Telegram bots.
- [micha(stars: 26)](https://github.com/onrik/micha) - Go Library for Telegram bot api.
- [olivia(stars: 3633)](https://github.com/olivia-ai/olivia) - A chatbot built with an artificial neural network.
- [slack-bot(stars: 158)](https://github.com/innogames/slack-bot) - Ready to use Slack Bot for lazy developers: Custom commands, Jenkins, Jira, Bitbucket, Github...
- [slacker(stars: 810)](https://github.com/shomali11/slacker) - Easy to use framework to create Slack bots.
- [slackscot(stars: 56)](https://github.com/alexandre-normand/slackscot) - Another framework for building Slack bots.
- [tbot(stars: 343)](https://github.com/yanzay/tbot) - Telegram bot server with API similar to net/http.
- [telebot(stars: 3426)](https://github.com/tucnak/telebot) - Telegram bot framework is written in Go.
- [telego(stars: 303)](https://github.com/mymmrac/telego) - Telegram Bot API library for Golang with full one-to-one API implementation.
- [telegram-bot-api(stars: 5184)](https://github.com/Syfaro/telegram-bot-api) - Simple and clean Telegram bot client.
- [teleterm(stars: 31)](https://github.com/alfiankan/teleterm) - Telegram Bot Exec Terminal Command.
- [Tenyks(stars: 176)](https://github.com/kyleterry/tenyks) - Service oriented IRC bot using Redis and JSON for messaging.
- [wayback(stars: 1537)](https://github.com/wabarc/wayback) - A bot for Telegram, Mastodon, Slack, and other messaging platforms archives webpages.

**[⬆ back to top](#contents)**

## Build Automation

_Libraries and tools help with build automation._

- [1build(stars: 219)](https://github.com/gopinath-langote/1build) - Command line tool to frictionlessly manage project-specific commands.
- [anko(stars: 35)](https://github.com/GuilhermeCaruso/anko) - Simple application watcher for multiple programming languages.
- [gaper(stars: 77)](https://github.com/maxcnunes/gaper) - Builds and restarts a Go project when it crashes or some watched file changes.
- [gilbert](https://go-gilbert.github.io) - Build system and task runner for Go projects.
- [goyek(stars: 408)](https://github.com/goyek/goyek) - Create build pipelines in Go.
- [mage(stars: 3787)](https://github.com/magefile/mage) - Mage is a make/rake-like build tool using Go.
- [mmake(stars: 1693)](https://github.com/tj/mmake) - Modern Make.
- [realize(stars: 4439)](https://github.com/tockins/realize) - Go build a system with file watchers and live to reload. Run, build and watch file changes with custom paths.
- [Task(stars: 9174)](https://github.com/go-task/task) - simple "Make" alternative.
- [taskctl(stars: 269)](https://github.com/taskctl/taskctl) - Concurrent task runner.
- [xc(stars: 928)](https://github.com/joerdav/xc) - Task runner with README.md defined tasks, executable markdown.

**[⬆ back to top](#contents)**

## Command Line

### Advanced Console UIs

_Libraries for building Console Applications and Console User Interfaces._

- [asciigraph(stars: 2380)](https://github.com/guptarohit/asciigraph) - Go package to make lightweight ASCII line graph ╭┈╯ in command line apps with no other dependencies.
- [aurora(stars: 1370)](https://github.com/logrusorgru/aurora) - ANSI terminal colors that support fmt.Printf/Sprintf.
- [box-cli-maker(stars: 434)](https://github.com/Delta456/box-cli-maker) - Make Highly Customized Boxes for your CLI.
- [bubble-table(stars: 279)](https://github.com/Evertras/bubble-table) - An interactive table component for bubbletea.
- [bubbles(stars: 4157)](https://github.com/charmbracelet/bubbles) - TUI components for bubbletea.
- [bubbletea(stars: 21625)](https://github.com/charmbracelet/bubbletea) - Go framework to build terminal apps, based on The Elm Architecture.
- [cfmt(stars: 100)](https://github.com/mingrammer/cfmt) - Contextual fmt inspired by bootstrap color classes.
- [cfmt(stars: 62)](https://github.com/i582/cfmt) - Simple and convenient formatted stylized output fully compatible with fmt library.
- [chalk(stars: 437)](https://github.com/ttacon/chalk) - Intuitive package for prettifying terminal/console output.
- [colourize(stars: 26)](https://github.com/TreyBastian/colourize) - Go library for ANSI colour text in terminals.
- [crab-config-files-templating(stars: 7)](https://github.com/alfiankan/crab-config-files-templating) - Dynamic configuration file templating tool for kubernetes manifest or general configuration files.
- [ctc(stars: 45)](https://github.com/wzshiming/ctc) - The non-invasive cross-platform terminal color library does not need to modify the Print method.
- [go-ataman(stars: 16)](https://github.com/workanator/go-ataman) - Go library for rendering ANSI colored text templates in terminals.
- [go-colorable(stars: 722)](https://github.com/mattn/go-colorable) - Colorable writer for windows.
- [go-colortext(stars: 214)](https://github.com/daviddengcn/go-colortext) - Go library for color output in terminals.
- [go-isatty(stars: 769)](https://github.com/mattn/go-isatty) - isatty for golang.
- [go-palette(stars: 14)](https://github.com/abusomani/go-palette) - Go library that provides elegant and convenient style definitions using ANSI colors. Fully compatible & wraps the [fmt library](https://pkg.go.dev/fmt) for nice terminal layouts.
- [go-prompt(stars: 5094)](https://github.com/c-bata/go-prompt) - Library for building a powerful interactive prompt, inspired by [python-prompt-toolkit](https://github.com/jonathanslenders/python-prompt-toolkit).
- [gocui(stars: 9505)](https://github.com/jroimartin/gocui) - Minimalist Go library aimed at creating Console User Interfaces.
- [gommon/color](https://github.com/labstack/gommon/tree/master/color) - Style terminal text.
- [gookit/color(stars: 1420)](https://github.com/gookit/color) - Terminal color rendering tool library, support 16 colors, 256 colors, RGB color rendering output, compatible with Windows.
- [lipgloss(stars: 6780)](https://github.com/charmbracelet/lipgloss) - Declaratively define styles for color, format and layout in the terminal.
- [marker(stars: 45)](https://github.com/cyucelen/marker) - Easiest way to match and mark strings for colorful terminal outputs.
- [mpb(stars: 2180)](https://github.com/vbauerster/mpb) - Multi progress bar for terminal applications.
- [progressbar(stars: 3643)](https://github.com/schollz/progressbar) - Basic thread-safe progress bar that works in every OS.
- [pterm(stars: 4373)](https://github.com/pterm/pterm) - A library to beautify console output on every platform with many combinable components.
- [simpletable(stars: 495)](https://github.com/alexeyco/simpletable) - Simple tables in a terminal with Go.
- [spinner(stars: 2206)](https://github.com/briandowns/spinner) - Go package to easily provide a terminal spinner with options.
- [tabby(stars: 341)](https://github.com/cheynewallace/tabby) - A tiny library for super simple Golang tables.
- [table(stars: 49)](https://github.com/tomlazar/table) - Small library for terminal color based tables .
- [tabular(stars: 76)](https://github.com/InVisionApp/tabular) - Print ASCII tables from command line utilities without the need to pass large sets of data to the API.
- [termbox-go(stars: 4600)](https://github.com/nsf/termbox-go) - Termbox is a library for creating cross-platform text-based interfaces.
- [termdash(stars: 2402)](https://github.com/mum4k/termdash) - Go terminal dashboard based on **termbox-go** and inspired by [termui](https://github.com/gizak/termui).
- [termenv(stars: 1564)](https://github.com/muesli/termenv) - Advanced ANSI style & color support for your terminal applications.
- [termui(stars: 12750)](https://github.com/gizak/termui) - Go terminal dashboard based on **termbox-go** and inspired by [blessed-contrib](https://github.com/yaronn/blessed-contrib).
- [uilive(stars: 1633)](https://github.com/gosuri/uilive) - Library for updating terminal output in real time.
- [uiprogress(stars: 2055)](https://github.com/gosuri/uiprogress) - Flexible library to render progress bars in terminal applications.
- [uitable(stars: 716)](https://github.com/gosuri/uitable) - Library to improve readability in terminal apps using tabular data.
- [yacspin(stars: 422)](https://github.com/theckman/yacspin) - Yet Another CLi Spinner package, for working with terminal spinners.

**[⬆ back to top](#contents)**

### Standard CLI

_Libraries for building standard or basic Command Line applications._

- [acmd(stars: 102)](https://github.com/cristalhq/acmd) - Simple, useful, and opinionated CLI package in Go.
- [argparse(stars: 574)](https://github.com/akamensky/argparse) - Command line argument parser inspired by Python's argparse module.
- [argv(stars: 39)](https://github.com/cosiner/argv) - Go library to split command line string as arguments array using the bash syntax.
- [carapace(stars: 146)](https://github.com/rsteube/carapace) - Command argument completion generator for spf13/cobra.
- [carapace-bin(stars: 492)](https://github.com/rsteube/carapace-bin) - Multi-shell multi-command argument completer.
- [carapace-spec(stars: 13)](https://github.com/rsteube/carapace-spec) - Define simple completions using a spec file.
- [cli(stars: 710)](https://github.com/mkideal/cli) - Feature-rich and easy to use command-line package based on golang struct tags.
- [cli(stars: 122)](https://github.com/teris-io/cli) - Simple and complete API for building command line interfaces in Go.
- [climax(stars: 214)](https://github.com/tucnak/climax) - Alternative CLI with "human face", in spirit of Go command.
- [clîr(stars: 161)](https://github.com/leaanthony/clir) - A Simple and Clear CLI library. Dependency free.
- [cmd(stars: 41)](https://github.com/posener/cmd) - Extends the standard `flag` package to support sub commands and more in idiomatic way.
- [cmdr(stars: 125)](https://github.com/hedzr/cmdr) - A POSIX/GNU style, getopt-like command-line UI Go library.
- [cobra(stars: 34632)](https://github.com/spf13/cobra) - Commander for modern Go CLI interactions.
- [command-chain(stars: 50)](https://github.com/rainu/go-command-chain) - A go library for configure and run command chains - such as pipelining in unix shells.
- [commandeer(stars: 171)](https://github.com/jaffee/commandeer) - Dev-friendly CLI apps: sets up flags, defaults, and usage based on struct fields and tags.
- [complete(stars: 903)](https://github.com/posener/complete) - Write bash completions in Go + Go command bash completion.
- [Dnote(stars: 2618)](https://github.com/dnote/dnote) - A simple command line notebook with multi-device sync.
- [elvish(stars: 5195)](https://github.com/elves/elvish) - An expressive programming language and a versatile interactive shell.
- [env(stars: 107)](https://github.com/codingconcepts/env) - Tag-based environment configuration for structs.
- [flag(stars: 129)](https://github.com/cosiner/flag) - Simple but powerful command line option parsing library for Go supporting subcommand.
- [flaggy(stars: 841)](https://github.com/integrii/flaggy) - A robust and idiomatic flags package with excellent subcommand support.
- [flagvar(stars: 43)](https://github.com/sgreben/flagvar) - A collection of flag argument types for Go's standard `flag` package.
- [go-andotp(stars: 25)](https://github.com/grijul/go-andotp) - A CLI program to encrypt/decrypt [andOTP](https://github.com/andOTP/andOTP) files. Can be used as a library as well.
- [go-arg(stars: 1789)](https://github.com/alexflint/go-arg) - Struct-based argument parsing in Go.
- [go-commander(stars: 35)](https://github.com/yitsushi/go-commander) - Go library to simplify CLI workflow.
- [go-flags(stars: 2471)](https://github.com/jessevdk/go-flags) - go command line option parser.
- [go-getoptions(stars: 53)](https://github.com/DavidGamba/go-getoptions) - Go option parser inspired by the flexibility of Perl’s GetOpt::Long.
- [gocmd(stars: 66)](https://github.com/devfacet/gocmd) - Go library for building command line applications.
- [hiboot cli](https://github.com/hidevopsio/hiboot/tree/master/pkg/app/cli) - cli application framework with auto configuration and dependency injection.
- [job(stars: 137)](https://github.com/liujianping/job) - JOB, make your short-term command as a long-term job.
- [kingpin(stars: 3398)](https://github.com/alecthomas/kingpin) - Command line and flag parser supporting sub commands (superseded by `kong`; see below).
- [liner(stars: 1016)](https://github.com/peterh/liner) - Go readline-like library for command-line interfaces.
- [mcli(stars: 24)](https://github.com/jxskiss/mcli) - A minimal but very powerful cli library for Go.
- [mitchellh/cli(stars: 1705)](https://github.com/mitchellh/cli) - Go library for implementing command-line interfaces.
- [mow.cli](https://github.com/jawher/mow.cli) - Go library for building CLI applications with sophisticated flag and argument parsing and validation.
- [ops(stars: 1146)](https://github.com/nanovms/ops) - Unikernel Builder/Orchestrator.
- [pflag(stars: 2247)](https://github.com/spf13/pflag) - Drop-in replacement for Go's flag package, implementing POSIX/GNU-style --flags.
- [readline](https://github.com/reeflective/readline) Shell library with modern and easy to use UI features.
- [sand(stars: 24)](https://github.com/Zaba505/sand) - Simple API for creating interpreters and so much more.
- [sflags(stars: 148)](https://github.com/octago/sflags) - Struct based flags generator for flag, urfave/cli, pflag, cobra, kingpin, and other libraries.
- [strumt(stars: 57)](https://github.com/antham/strumt) - Library to create prompt chain.
- [subcmd(stars: 6)](https://github.com/bobg/subcmd) - Another approach to parsing and running subcommands. Works alongside the standard `flag` package.
- [survey(stars: 3987)](https://github.com/go-survey/survey) - Build interactive and accessible prompts with full support for windows and posix terminals.
- [ts(stars: 20)](https://github.com/liujianping/ts) - Timestamp convert & compare tool.
- [ukautz/clif(stars: 125)](https://github.com/ukautz/clif) - Small command line interface framework.
- [urfave/cli(stars: 21143)](https://github.com/urfave/cli) - Simple, fast, and fun package for building command line apps in Go (formerly codegangsta/cli).
- [version(stars: 80)](https://github.com/mszostok/version) - Collects and displays CLI version information in multiple formats along with upgrade notice.
- [wlog(stars: 63)](https://github.com/dixonwille/wlog) - Simple logging interface that supports cross-platform color and concurrency.
- [wmenu(stars: 204)](https://github.com/dixonwille/wmenu) - Easy to use menu structure for cli applications that prompt users to make choices.

**[⬆ back to top](#contents)**

## Configuration

_Libraries for configuration parsing._

- [aconfig(stars: 498)](https://github.com/cristalhq/aconfig) - Simple, useful and opinionated config loader.
- [cleanenv(stars: 1288)](https://github.com/ilyakaznacheev/cleanenv) - Minimalistic configuration reader (from files, ENV, and wherever you want).
- [config(stars: 334)](https://github.com/JeremyLoy/config) - Cloud native application configuration. Bind ENV to structs in only two lines.
- [config(stars: 44)](https://github.com/num30/config) - configure you app using file, environment variables, or flags in two lines of code
- [config(stars: 268)](https://github.com/olebedev/config) - JSON or YAML configuration wrapper with environment variables and flags parsing.
- [configuration(stars: 84)](https://github.com/BoRuDar/configuration) - Library for initializing configuration structs from env variables, files, flags and 'default' tag.
- [configure(stars: 55)](https://github.com/paked/configure) - Provides configuration through multiple sources, including JSON, flags and environment variables.
- [configuro(stars: 89)](https://github.com/sherifabdlnaby/configuro) - opinionated configuration loading & validation framework from ENV and Files focused towards 12-Factor compliant applications.
- [confita(stars: 482)](https://github.com/heetch/confita) - Load configuration in cascade from multiple backends into a struct.
- [conflate(stars: 31)](https://github.com/the4thamigo-uk/conflate) - Library/tool to merge multiple JSON/YAML/TOML files from arbitrary URLs, validation against a JSON schema, and application of default values defined in the schema.
- [env(stars: 4084)](https://github.com/caarlos0/env) - Parse environment variables to Go structs (with defaults).
- [env(stars: 45)](https://github.com/junk1tm/env) - A lightweight package for loading environment variables into structs.
- [envcfg(stars: 102)](https://github.com/tomazk/envcfg) - Un-marshaling environment variables to Go structs.
- [envconf(stars: 11)](https://github.com/ian-kent/envconf) - Configuration from environment.
- [envconfig(stars: 235)](https://github.com/vrischmann/envconfig) - Read your configuration from environment variables.
- [envh(stars: 96)](https://github.com/antham/envh) - Helpers to manage environment variables.
- [fig(stars: 301)](https://github.com/kkyr/fig) - Tiny library for reading configuration from a file and from environment variables (with validation & defaults).
- [gcfg(stars: 165)](https://github.com/go-gcfg/gcfg) - read INI-style configuration files into Go structs; supports user-defined types and subsections.
- [genv(stars: 35)](https://github.com/sakirsensoy/genv) - Read environment variables easily with dotenv support.
- [go-aws-ssm(stars: 54)](https://github.com/PaddleHQ/go-aws-ssm) - Go package that fetches parameters from AWS System Manager - Parameter Store.
- [go-conf(stars: 10)](https://github.com/ThomasObenaus/go-conf) - Simple library for application configuration based on annotated structs. It supports reading the configuration from environment variables, config files and command line parameters.
- [go-ini(stars: 10)](https://github.com/subpop/go-ini) - A Go package that marshals and unmarshals INI-files.
- [go-ssm-config(stars: 18)](https://github.com/ianlopshire/go-ssm-config) - Go utility for loading configuration parameters from AWS SSM (Parameter Store).
- [go-up(stars: 43)](https://github.com/ufoscout/go-up) - A simple configuration library with recursive placeholders resolution and no magic.
- [goConfig(stars: 3)](https://github.com/crgimenes/goConfig) - Parses a struct as input and populates the fields of this struct with parameters from command line, environment variables and configuration file.
- [godotenv(stars: 7026)](https://github.com/joho/godotenv) - Go port of Ruby's dotenv library (Loads environment variables from `.env`).
- [gofigure(stars: 66)](https://github.com/ian-kent/gofigure) - Go application configuration made easy.
- [GoLobby/Config(stars: 352)](https://github.com/golobby/config) - GoLobby Config is a lightweight yet powerful configuration manager for the Go programming language.
- [gone/jconf](https://github.com/One-com/gone/tree/master/jconf) - Modular JSON configuration. Keep you config structs along with the code they configure and delegate parsing to submodules without sacrificing full config serialization.
- [gonfig(stars: 7)](https://github.com/milad-abbasi/gonfig) - Tag-based configuration parser which loads values from different providers into typesafe struct.
- [gookit/config(stars: 501)](https://github.com/gookit/config) - application config manage(load,get,set). support JSON, YAML, TOML, INI, HCL. multi file load, data override merge.
- [harvester(stars: 129)](https://github.com/beatlabs/harvester) - Harvester, a easy to use static and dynamic configuration package supporting seeding, env vars and Consul integration.
- [hjson(stars: 310)](https://github.com/hjson/hjson-go) - Human JSON, a configuration file format for humans. Relaxed syntax, fewer mistakes, more comments.
- [hocon(stars: 68)](https://github.com/gurkankaymak/hocon) - Configuration library for working with the HOCON(a human-friendly JSON superset) format, supports features like environment variables, referencing other values, comments and multiple files.
- [ingo(stars: 36)](https://github.com/schachmat/ingo) - Flags persisted in an ini-like config file.
- [ini(stars: 3364)](https://github.com/go-ini/ini) - Go package to read and write INI files.
- [ini(stars: 15)](https://github.com/wlevene/ini) - INI Parser & Write Library, Unmarshal to Struct,Marshal to Json,Write File,watch file.
- [joshbetz/config(stars: 216)](https://github.com/joshbetz/config) - Small configuration library for Go that parses environment variables, JSON files, and reloads automatically on SIGHUP.
- [kelseyhightower/envconfig(stars: 4766)](https://github.com/kelseyhightower/envconfig) - Go library for managing configuration data from environment variables.
- [koanf(stars: 2167)](https://github.com/knadh/koanf) - Light weight, extensible library for reading config in Go applications. Built in support for JSON, TOML, YAML, env, command line.
- [konfig(stars: 643)](https://github.com/lalamove/konfig) - Composable, observable and performant config handling for Go for the distributed processing era.
- [kong(stars: 1677)](https://github.com/alecthomas/kong) - Command-line parser with support for arbitrarily complex command-line structures and additional sources of configuration such as YAML, JSON, TOML, etc (successor to `kingpin`).
- [mini(stars: 34)](https://github.com/sasbury/mini) - Golang package for parsing ini-style configuration files.
- [nasermirzaei89/env(stars: 10)](https://github.com/nasermirzaei89/env) - Simple useful package for read environment variables.
- [nfigure(stars: 6)](https://github.com/muir/nfigure) - Per-library struct-tag based configuration from command lines (Posix & Go-style); environment, JSON, YAML
- [onion(stars: 114)](https://github.com/goraz/onion) - Layer based configuration for Go, Supports JSON, TOML, YAML, properties, etcd, env, and encryption using PGP.
- [piper(stars: 7)](https://github.com/Yiling-J/piper) - Viper wrapper with config inheritance and key generation.
- [store(stars: 274)](https://github.com/tucnak/store) - Lightweight configuration manager for Go.
- [swap(stars: 8)](https://github.com/oblq/swap) - Instantiate/configure structs recursively, based on build environment. (YAML, TOML, JSON and env).
- [typenv(stars: 9)](https://github.com/diegomarangoni/typenv) - Minimalistic, zero dependency, typed environment variables library.
- [uConfig(stars: 68)](https://github.com/omeid/uconfig) - Lightweight, zero-dependency, and extendable configuration management.
- [viper(stars: 24866)](https://github.com/spf13/viper) - Go configuration with fangs.
- [xdg(stars: 498)](https://github.com/adrg/xdg) - Go implementation of the [XDG Base Directory Specification](https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html) and [XDG user directories](https://wiki.archlinux.org/index.php/XDG_user_directories).
- [xdg(stars: 80)](https://github.com/OpenPeeDeeP/xdg) - Cross platform package that follows the [XDG Standard](https://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html).

**[⬆ back to top](#contents)**

## Continuous Integration

_Tools for help with continuous integration._

- [CDS(stars: 4351)](https://github.com/ovh/cds) - Enterprise-Grade CI/CD and DevOps Automation Open Source Platform.
- [drone(stars: 30904)](https://github.com/drone/drone) - Drone is a Continuous Integration platform built on Docker, written in Go.
- [duci(stars: 74)](https://github.com/duck8823/duci) - A simple ci server no needs domain specific languages.
- [go-fuzz-action(stars: 10)](https://github.com/jidicula/go-fuzz-action) - Use Go 1.18's built-in fuzz testing in GitHub Actions.
- [go-test-coverage(stars: 32)](https://github.com/vladopajic/go-test-coverage) - Tool and GitHub action which reports issues when test coverage is below set threshold.
- [gomason(stars: 60)](https://github.com/nikogura/gomason) - Test, Build, Sign, and Publish your go binaries from a clean workspace.
- [gotestfmt(stars: 448)](https://github.com/GoTestTools/gotestfmt) - go test output for humans.
- [goveralls(stars: 770)](https://github.com/mattn/goveralls) - Go integration for Coveralls.io continuous code coverage tracking system.
- [overalls(stars: 114)](https://github.com/go-playground/overalls) - Multi-Package go project coverprofile for tools like goveralls.
- [roveralls(stars: 19)](https://github.com/LawrenceWoodman/roveralls) - Recursive coverage testing tool.
- [woodpecker(stars: 3308)](https://github.com/woodpecker-ci/woodpecker) - Woodpecker is a community fork of the Drone CI system.

**[⬆ back to top](#contents)**

## CSS Preprocessors

_Libraries for preprocessing CSS files._

- [gcss(stars: 487)](https://github.com/yosssi/gcss) - Pure Go CSS Preprocessor.
- [go-libsass(stars: 199)](https://github.com/wellington/go-libsass) - Go wrapper to the 100% Sass compatible libsass project.

**[⬆ back to top](#contents)**

## Data Structures and Algorithms

### Bit-packing and Compression

- [bingo(stars: 29)](https://github.com/iancmcc/bingo) - Fast, zero-allocation, lexicographical-order-preserving packing of native types to bytes.
- [binpacker(stars: 216)](https://github.com/zhuangsirui/binpacker) - Binary packer and unpacker helps user build custom binary stream.
- [bit(stars: 156)](https://github.com/yourbasic/bit) - Golang set data structure with bonus bit-twiddling functions.
- [crunch(stars: 90)](https://github.com/superwhiskers/crunch) - Go package implementing buffers for handling various datatypes easily.
- [go-ef(stars: 31)](https://github.com/amallia/go-ef) - A Go implementation of the Elias-Fano encoding.
- [roaring(stars: 2227)](https://github.com/RoaringBitmap/roaring) - Go package implementing compressed bitsets.

### Bit Sets

- [bitmap(stars: 250)](https://github.com/kelindar/bitmap) - Dense, zero-allocation, SIMD-enabled bitmap/bitset in Go.
- [bitset(stars: 1213)](https://github.com/bits-and-blooms/bitset) - Go package implementing bitsets.

### Bloom and Cuckoo Filters

- [bloom(stars: 2181)](https://github.com/bits-and-blooms/bloom) - Go package implementing Bloom filters.
- [bloom(stars: 147)](https://github.com/zhenjl/bloom) - Bloom filters implemented in Go.
- [bloom(stars: 83)](https://github.com/yourbasic/bloom) - Golang Bloom filter implementation.
- [bloomfilter(stars: 16)](https://github.com/OldPanda/bloomfilter) - Yet another Bloomfilter implementation in Go, compatible with Java's Guava library.
- [boomfilters(stars: 1566)](https://github.com/tylertreat/BoomFilters) - Probabilistic data structures for processing continuous, unbounded streams.
- [cuckoo-filter(stars: 276)](https://github.com/linvon/cuckoo-filter) - Cuckoo filter: a comprehensive cuckoo filter, which is configurable and space optimized compared with other implements, and all features mentioned in original paper is available.
- [cuckoofilter(stars: 1038)](https://github.com/seiflotfy/cuckoofilter) - Cuckoo filter: a good alternative to a counting bloom filter implemented in Go.
- [ring(stars: 134)](https://github.com/TheTannerRyan/ring) - Go implementation of a high performance, thread safe bloom filter.

### Data Structure and Algorithm Collections

- [algorithms(stars: 751)](https://github.com/shady831213/algorithms) - Algorithms and data structures.CLRS study.
- [go-datastructures(stars: 7239)](https://github.com/Workiva/go-datastructures) - Collection of useful, performant, and thread-safe data structures.
- [gods(stars: 14799)](https://github.com/emirpasic/gods) - Go Data Structures. Containers, Sets, Lists, Stacks, Maps, BidiMaps, Trees, HashSet etc.
- [gostl(stars: 955)](https://github.com/liyue201/gostl) - Data structure and algorithm library for go, designed to provide functions similar to C++ STL.

### Iterators

- [goterator(stars: 16)](https://github.com/yaa110/goterator) - Iterator implementation to provide map and reduce functionalities.
- [iter(stars: 186)](https://github.com/disksing/iter) - Go implementation of C++ STL iterators and algorithms.

### Maps

See also [Database](#database) for more complex key-value stores, and [Trees](#trees) for
additional ordered map implementations.

- [cmap(stars: 79)](https://github.com/lrita/cmap) - a thread-safe concurrent map for go, support using `interface{}` as key and auto scale up shards.
- [dict(stars: 43)](https://github.com/srfrog/dict) - Python-like dictionaries (dict) for Go.
- [goradd/maps(stars: 30)](https://github.com/goradd/maps) - Go 1.18+ generic map interface for maps; safe maps; ordered maps; ordered, safe maps; etc.

### Miscellaneous Data Structures and Algorithms

- [concurrent-writer(stars: 52)](https://github.com/free/concurrent-writer) - Highly concurrent drop-in replacement for `bufio.Writer`.
- [conjungo(stars: 121)](https://github.com/InVisionApp/conjungo) - A small, powerful and flexible merge library.
- [count-min-log(stars: 66)](https://github.com/seiflotfy/count-min-log) - Go implementation Count-Min-Log sketch: Approximately counting with approximate counters (Like Count-Min sketch but using less memory).
- [fsm(stars: 54)](https://github.com/cocoonspace/fsm) - Finite-State Machine package.
- [genfuncs(stars: 46)](https://github.com/nwillc/genfuncs) - Go 1.18+ generics package inspired by Kotlin's Sequence and Map.
- [go-generics(stars: 40)](https://github.com/bobg/go-generics) - Generic slice, map, set, iterator, and goroutine utilities.
- [go-geoindex(stars: 353)](https://github.com/hailocab/go-geoindex) - In-memory geo index.
- [go-rampart(stars: 94)](https://github.com/francesconi/go-rampart) - Determine how intervals relate to each other.
- [go-rquad(stars: 129)](https://github.com/aurelien-rainone/go-rquad) - Region quadtrees with efficient point location and neighbour finding.
- [go-tuple(stars: 71)](https://github.com/barweiss/go-tuple) - Generic tuple implementation for Go 1.18+.
- [go18ds(stars: 39)](https://github.com/daichi-m/go18ds) - Go Data Structures using Go 1.18 generics.
- [gofal(stars: 18)](https://github.com/xxjwxc/gofal) - fractional api for Go.
- [gogu(stars: 89)](https://github.com/esimov/gogu) - A comprehensive, reusable and efficient concurrent-safe generics utility functions and data structures library.
- [gota(stars: 2821)](https://github.com/kniren/gota) - Implementation of dataframes, series, and data wrangling methods for Go.
- [hide(stars: 59)](https://github.com/emvi/hide) - ID type with marshalling to/from hash to prevent sending IDs to clients.
- [hilbert(stars: 274)](https://github.com/google/hilbert) - Go package for mapping values to and from space-filling curves, such as Hilbert and Peano curves.
- [hyperloglog(stars: 887)](https://github.com/axiomhq/hyperloglog) - HyperLogLog implementation with Sparse, LogLog-Beta bias correction and TailCut space reduction.
- [plinko(stars: 164)](https://github.com/shipt/plinko) - A finite state machine and workflow orchestrator that compiles for fast execution, easy debugging, auto-generated documentation. Includes advanced features such as side-effect hooks. 
- [quadtree(stars: 30)](https://github.com/s0rg/quadtree) - Generic, zero-alloc, 100%-test covered quadtree.
- [slices(stars: 16)](https://github.com/srfrog/slices) - Functions that operate on slices; like `package strings` but adapted to work with slices.
- [slices(stars: 15)](https://github.com/twharmon/slices) - Pure, generic functions for slices.

### Nullable Types

- [nan(stars: 80)](https://github.com/kak-tus/nan) - Zero allocation Nullable structures in one library with handy conversion functions, marshallers and unmarshallers.
- [null(stars: 32)](https://github.com/emvi/null) - Nullable Go types that can be marshalled/unmarshalled to/from JSON.
- [typ(stars: 44)](https://github.com/gurukami/typ) - Null Types, Safe primitive type conversion and fetching value from complex structures.

### Queues

- [deque(stars: 156)](https://github.com/edwingeng/deque) - A highly optimized double-ended queue.
- [deque(stars: 514)](https://github.com/gammazero/deque) - Fast ring-buffer deque (double-ended queue).
- [goconcurrentqueue(stars: 336)](https://github.com/enriquebris/goconcurrentqueue) - Concurrent FIFO queue.
- [memlog(stars: 110)](https://github.com/embano1/memlog) - An easy to use, lightweight, thread-safe and append-only in-memory data structure inspired by Apache Kafka.
- [queue(stars: 225)](https://github.com/adrianbrad/queue) - Multiple thread-safe, generic queue implementations for Go.

### Sets

- [dsu(stars: 13)](https://github.com/ihebu/dsu) - Disjoint Set data structure implementation in Go.
- [golang-set(stars: 3733)](https://github.com/deckarep/golang-set) - Thread-Safe and Non-Thread-Safe high-performance sets for Go.
- [goset(stars: 53)](https://github.com/zoumo/goset) - A useful Set collection implementation for Go.
- [set(stars: 26)](https://github.com/StudioSol/set) - Simple set data structure implementation in Go using LinkedHashMap.

### Text Analysis

- [bleve(stars: 9404)](https://github.com/blevesearch/bleve) - Modern text indexing library for go.
- [go-adaptive-radix-tree(stars: 322)](https://github.com/plar/go-adaptive-radix-tree) - Go implementation of Adaptive Radix Tree.
- [go-edlib(stars: 430)](https://github.com/hbollon/go-edlib) - Go string comparison and edit distance algorithms library (Levenshtein, LCS, Hamming, Damerau levenshtein, Jaro-Winkler, etc.) compatible with Unicode.
- [levenshtein(stars: 81)](https://github.com/agext/levenshtein) - Levenshtein distance and similarity metrics with customizable edit costs and Winkler-like bonus for common prefix.
- [levenshtein(stars: 315)](https://github.com/agnivade/levenshtein) - Implementation to calculate levenshtein distance in Go.
- [mspm(stars: 24)](https://github.com/BlackRabbitt/mspm) - Multi-String Pattern Matching Algorithm for information retrieval.
- [parsefields(stars: 8)](https://github.com/MonaxGT/parsefields) - Tools for parse JSON-like logs for collecting unique fields and events.
- [ptrie(stars: 35)](https://github.com/viant/ptrie) - An implementation of prefix tree.
- [trie(stars: 717)](https://github.com/derekparker/trie) - Trie implementation in Go.

### Trees

- [hashsplit](http://github.com/bobg/hashsplit) - Split byte streams into chunks, and arrange chunks into trees, with boundaries determined by content, not position.
- [merkle(stars: 12)](https://github.com/bobg/merkle) - Space-efficient computation of Merkle root hashes and inclusion proofs.
- [skiplist(stars: 264)](https://github.com/MauriceGit/skiplist) - Very fast Go Skiplist implementation.
- [skiplist(stars: 83)](https://github.com/gansidui/skiplist) - Skiplist implementation in Go.
- [treap(stars: 23)](https://github.com/perdata/treap) - Persistent, fast ordered map using tree heaps.
- [treemap(stars: 53)](https://github.com/igrmk/treemap) - Generic key-sorted map using a red-black tree under the hood.

### Pipes

- [ordered-concurrently(stars: 27)](https://github.com/tejzpr/ordered-concurrently) - Go module that processes work concurrently and returns output in a channel in the order of input.
- [parapipe(stars: 25)](https://github.com/nazar256/parapipe) - FIFO Pipeline which parallels execution on each stage while maintaining the order of messages and results.
- [pipeline(stars: 52)](https://github.com/hyfather/pipeline) - An implementation of pipelines with fan-in and fan-out.

**[⬆ back to top](#contents)**

## Database

### Caches

_Data stores with expiring records, in-memory distributed data stores, or in-memory subsets of file-based databases._

- [2q(stars: 33)](https://github.com/floatdrop/2q) - 2Q in-memory cache implementation.
- [bcache(stars: 141)](https://github.com/iwanbk/bcache) - Eventually consistent distributed in-memory cache Go library.
- [BigCache(stars: 6962)](https://github.com/allegro/bigcache) - Efficient key/value cache for gigabytes of data.
- [cache(stars: 163)](https://github.com/akyoto/cache) - In-memory key:value store with expiration time, 0 dependencies, <100 LoC, 100% coverage.
- [cache2go(stars: 2031)](https://github.com/muesli/cache2go) - In-memory key:value cache which supports automatic invalidation based on timeouts.
- [cachego(stars: 332)](https://github.com/faabiosr/cachego) - Golang Cache component for multiple drivers.
- [clusteredBigCache(stars: 44)](https://github.com/oaStuff/clusteredBigCache) - BigCache with clustering support and individual item expiration.
- [coherence-go-client(stars: 7)](https://github.com/oracle/coherence-go-client) - Full implementation of Oracle Coherence cache API for Go applications using gRPC as network transport.
- [couchcache(stars: 60)](https://github.com/codingsince1985/couchcache) - RESTful caching micro-service backed by Couchbase server.
- [fastcache(stars: 1921)](https://github.com/VictoriaMetrics/fastcache) - fast thread-safe inmemory cache for big number of entries. Minimizes GC overhead.
- [GCache(stars: 2478)](https://github.com/bluele/gcache) - Cache library with support for expirable Cache, LFU, LRU and ARC.
- [gdcache(stars: 11)](https://github.com/ulovecode/gdcache) - A pure non-intrusive cache library implemented by golang, you can use it to implement your own distributed cache.
- [go-cache(stars: 118)](https://github.com/viney-shih/go-cache) - A flexible multi-layer Go caching library to deal with in-memory and shared cache by adopting Cache-Aside pattern.
- [go-mcache(stars: 90)](https://github.com/OrlovEvgeny/go-mcache) - Fast in-memory key:value store/cache library. Pointer caches.
- [gocache(stars: 2086)](https://github.com/eko/gocache) - A complete Go cache library with multiple stores (memory, memcache, redis, ...), chainable, loadable, metrics cache and more.
- [gocache(stars: 4)](https://github.com/yuseferi/gocache) - A data race free Go ache library with high performance and auto pruge functionality
- [groupcache(stars: 12548)](https://github.com/golang/groupcache) - Groupcache is a caching and cache-filling library, intended as a replacement for memcached in many cases.
- [imcache(stars: 95)](https://github.com/erni27/imcache) - A generic in-memory cache Go library. It supports expiration, sliding expiration, max entries limit, eviction callbacks and sharding.
- [nscache(stars: 5)](https://github.com/no-src/nscache) - A Go caching framework that supports multiple data source drivers.
- [remember-go(stars: 137)](https://github.com/rocketlaunchr/remember-go) - A universal interface for caching slow database queries (backed by redis, memcached, ristretto, or in-memory).
- [theine(stars: 183)](https://github.com/Yiling-J/theine-go) - High performance, near optimal in-memory cache with proactive TTL expiration and generics.
- [timedmap(stars: 67)](https://github.com/zekroTJA/timedmap) - Map with expiring key-value pairs.
- [ttlcache(stars: 782)](https://github.com/jellydator/ttlcache) - An in-memory cache with item expiration and generics.
- [ttlcache(stars: 9)](https://github.com/cheshir/ttlcache) - In-memory key value storage with TTL for each record.

### Databases Implemented in Go

- [badger(stars: 13007)](https://github.com/dgraph-io/badger) - Fast key-value store in Go.
- [bbolt(stars: 7208)](https://github.com/etcd-io/bbolt) - An embedded key/value database for Go.
- [Bitcask](https://git.mills.io/prologic/bitcask) - Bitcask is an embeddable, persistent and fast key-value (KV) database written in pure Go with predictable read/write performance, low latency and high throughput thanks to the bitcask on-disk layout (LSM+WAL).
- [buntdb(stars: 4276)](https://github.com/tidwall/buntdb) - Fast, embeddable, in-memory key/value database for Go with custom indexing and spatial support.
- [clover(stars: 519)](https://github.com/ostafen/clover) - A lightweight document-oriented NoSQL database written in pure Golang.
- [cockroach(stars: 28346)](https://github.com/cockroachdb/cockroach) - Scalable, Geo-Replicated, Transactional Datastore.
- [Coffer(stars: 35)](https://github.com/claygod/coffer) - Simple ACID key-value database that supports transactions.
- [column(stars: 1324)](https://github.com/kelindar/column) - High-performance, columnar, embeddable in-memory store with bitmap indexing and transactions.
- [CovenantSQL(stars: 1462)](https://github.com/CovenantSQL/CovenantSQL) - CovenantSQL is a SQL database on blockchain.
- [Databunker(stars: 1144)](https://github.com/paranoidguy/databunker) - Personally identifiable information (PII) storage service built to comply with GDPR and CCPA.
- [dgraph(stars: 19819)](https://github.com/dgraph-io/dgraph) - Scalable, Distributed, Low Latency, High Throughput Graph Database.
- [diskv(stars: 1318)](https://github.com/peterbourgon/diskv) - Home-grown disk-backed key-value store.
- [dolt(stars: 16107)](https://github.com/dolthub/dolt) - Dolt – It's Git for Data.
- [dtf(stars: 216)](https://github.com/dtm-labs/dtf) - A distributed transaction manager. Support XA, TCC, SAGA, Reliable Messages.
- [eliasdb(stars: 984)](https://github.com/krotik/eliasdb) - Dependency-free, transactional graph database with REST API, phrase search and SQL-like query language.
- [godis(stars: 3138)](https://github.com/hdt3213/godis) - A Golang implemented high-performance Redis server and cluster.
- [goleveldb(stars: 5921)](https://github.com/syndtr/goleveldb) - Implementation of the [LevelDB](https://github.com/google/leveldb) key/value database in Go.
- [hare(stars: 84)](https://github.com/jameycribbs/hare) - A simple database management system that stores each table as a text file of line-delimited JSON.
- [immudb(stars: 8399)](https://github.com/codenotary/immudb) - immudb is a lightweight, high-speed immutable database for systems and applications written in Go.
- [influxdb(stars: 26979)](https://github.com/influxdb/influxdb) - Scalable datastore for metrics, events, and real-time analytics.
- [ledisdb(stars: 4029)](https://github.com/siddontang/ledisdb) - Ledisdb is a high performance NoSQL like Redis based on LevelDB.
- [levigo(stars: 417)](https://github.com/jmhodges/levigo) - Levigo is a Go wrapper for LevelDB.
- [libradb(stars: 128)](https://github.com/amit-davidson/LibraDB) - LibraDB is a simple database with less then 1000 lines of code for learning.
- [LinDB(stars: 2767)](https://github.com/lindb/lindb) - LinDB is a scalable, high performance, high availability distributed time series database.
- [lotusdb(stars: 1862)](https://github.com/flower-corp/lotusdb) - Fast k/v database compatible with lsm and b+tree.
- [Milvus(stars: 24824)](https://github.com/milvus-io/milvus) - Milvus is a vector database for embedding management, analytics and search.
- [moss(stars: 940)](https://github.com/couchbase/moss) - Moss is a simple LSM key-value storage engine written in 100% Go.
- [nutsdb(stars: 3206)](https://github.com/xujiajun/nutsdb) - Nutsdb is a simple, fast, embeddable, persistent key/value store written in pure Go. It supports fully serializable transactions and many data structures such as list, set, sorted set.
- [objectbox-go(stars: 935)](https://github.com/objectbox/objectbox-go) - High-performance embedded Object Database (NoSQL) with Go API.
- [piladb(stars: 200)](https://github.com/fern4lvarez/piladb) - Lightweight RESTful database engine based on stack data structures.
- [pogreb(stars: 1179)](https://github.com/akrylysov/pogreb) - Embedded key-value store for read-heavy workloads.
- [prometheus(stars: 51176)](https://github.com/prometheus/prometheus) - Monitoring system and time series database.
- [pudge(stars: 352)](https://github.com/recoilme/pudge) - Fast and simple key/value store written using Go's standard library.
- [regatta(stars: 58)](https://github.com/jamf/regatta) - Fast, simple, geo-distributed KV store built for cloud native era.
- [rosedb(stars: 4167)](https://github.com/roseduan/rosedb) - An embedded k-v database based on LSM+WAL, supports string, list, hash, set, zset.
- [rqlite(stars: 14369)](https://github.com/rqlite/rqlite) - The lightweight, distributed, relational database built on SQLite.
- [tempdb(stars: 17)](https://github.com/rafaeljesus/tempdb) - Key-value store for temporary items.
- [tidb(stars: 35456)](https://github.com/pingcap/tidb) - TiDB is a distributed SQL database. Inspired by the design of Google F1.
- [tiedot(stars: 2714)](https://github.com/HouzuoGuo/tiedot) - Your NoSQL database powered by Golang.
- [unitdb(stars: 113)](https://github.com/unit-io/unitdb) - Fast timeseries database for IoT, realtime messaging applications. Access unitdb with pubsub over tcp or websocket using github.com/unit-io/unitd application.
- [Vasto(stars: 252)](https://github.com/chrislusf/vasto) - A distributed high-performance key-value store. On Disk. Eventual consistent. HA. Able to grow or shrink without service interruption.
- [VictoriaMetrics(stars: 10034)](https://github.com/VictoriaMetrics/VictoriaMetrics) - fast, resource-effective and scalable open source time series database. May be used as long-term remote storage for Prometheus. Supports PromQL.

### Database Schema Migration

- [atlas(stars: 4207)](https://github.com/ariga/atlas) - A Database Toolkit. A CLI designed to help companies better work with their data.
- [avro(stars: 45)](https://github.com/khezen/avro) - Discover SQL schemas and convert them to AVRO schemas. Query SQL records into AVRO bytes.
- [bytebase(stars: 8696)](https://github.com/bytebase/bytebase) - Safe database schema change and version control for DevOps teams.
- [darwin(stars: 140)](https://github.com/GuiaBolso/darwin) - Database schema evolution library for Go.
- [dbmate(stars: 4070)](https://github.com/amacneil/dbmate) - A lightweight, framework-agnostic database migration tool.
- [go-fixtures(stars: 28)](https://github.com/RichardKnop/go-fixtures) - Django style fixtures for Golang's excellent built-in database/sql library.
- [go-pg-migrate(stars: 11)](https://github.com/lawzava/go-pg-migrate) - CLI-friendly package for go-pg migrations management.
- [go-pg-migrations(stars: 83)](https://github.com/robinjoseph08/go-pg-migrations) - A Go package to help write migrations with go-pg/pg.
- [goavro(stars: 928)](https://github.com/linkedin/goavro) - A Go package that encodes and decodes Avro data.
- [godfish(stars: 5)](https://github.com/rafaelespinoza/godfish) - Database migration manager, works with native query language. Support for cassandra, mysql, postgres, sqlite3.
- [goose(stars: 4896)](https://github.com/pressly/goose) - Database migration tool. You can manage your database's evolution by creating incremental SQL or Go scripts.
- [gorm-seeder(stars: 12)](https://github.com/Kachit/gorm-seeder) - Simple database seeder for Gorm ORM.
- [gormigrate(stars: 950)](https://github.com/go-gormigrate/gormigrate) - Database schema migration helper for Gorm ORM.
- [libschema(stars: 14)](https://github.com/muir/libschema) - Define your migrations separately in each library. Migrations for open source libraries. MySQL & PostgreSQL.
- [migrate(stars: 13089)](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
- [migrator(stars: 157)](https://github.com/lopezator/migrator) - Dead simple Go database migration library.
- [migrator(stars: 24)](https://github.com/larapulse/migrator) - MySQL database migrator designed to run migrations to your features and manage database schema update with intuitive go code.
- [schema(stars: 30)](https://github.com/adlio/schema) - Library to embed schema migrations for database/sql-compatible databases inside your Go binaries.
- [skeema(stars: 1193)](https://github.com/skeema/skeema) - Pure-SQL schema management system for MySQL, with support for sharding and external online schema change tools.
- [soda](https://github.com/gobuffalo/pop/tree/master/soda) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
- [sql-migrate(stars: 3010)](https://github.com/rubenv/sql-migrate) - Database migration tool. Allows embedding migrations into the application using go-bindata.
- [sqlize(stars: 74)](https://github.com/sunary/sqlize) - Database migration generator. Allows generate sql migration from model and existing sql by differ them.

### Database Tools

- [chproxy(stars: 1162)](https://github.com/Vertamedia/chproxy) - HTTP proxy for ClickHouse database.
- [clickhouse-bulk(stars: 448)](https://github.com/nikepan/clickhouse-bulk) - Collects small inserts and sends big requests to ClickHouse servers.
- [dbbench(stars: 91)](https://github.com/sj14/dbbench) - Database benchmarking tool with support for several databases and scripts.
- [dg(stars: 17)](https://github.com/codingconcepts/dg) - A fast data generator that produces CSV files from generated relational data.
- [dynago(stars: 10)](https://github.com/twharmon/dynago) - Simplify working with AWS DynamoDB.
- [go-mysql(stars: 4308)](https://github.com/siddontang/go-mysql) - Go toolset to handle MySQL protocol and replication.
- [go-mysql-elasticsearch(stars: 4036)](https://github.com/siddontang/go-mysql-elasticsearch) - Sync your MySQL data into Elasticsearch automatically.
- [hasql](https://golang.yandex/hasql) - Library for accessing multi-host SQL database installations.
- [kingshard(stars: 6297)](https://github.com/flike/kingshard) - kingshard is a high performance proxy for MySQL powered by Golang.
- [octillery(stars: 182)](https://github.com/knocknote/octillery) - Go package for sharding databases ( Supports every ORM or raw SQL ).
- [onedump(stars: 34)](https://github.com/liweiyi88/onedump) - Database backup from different drivers to different destinations with one command and configuration.
- [orchestrator(stars: 5341)](https://github.com/github/orchestrator) - MySQL replication topology manager & visualizer.
- [pg_timetable(stars: 979)](https://github.com/cybertec-postgresql/pg_timetable) - Advanced scheduling for PostgreSQL.
- [pgweb(stars: 8159)](https://github.com/sosedoff/pgweb) - Web-based PostgreSQL database browser.
- [prep(stars: 33)](https://github.com/hexdigest/prep) - Use prepared SQL statements without changing your code.
- [pREST(stars: 4003)](https://github.com/prest/prest) - Simplify and accelerate development, ⚡ instant, realtime, high-performance on any Postgres application, existing or new.
- [rdb(stars: 306)](https://github.com/HDT3213/rdb) - Redis RDB file parser for secondary development and memory analysis.
- [rwdb(stars: 18)](https://github.com/andizzle/rwdb) - rwdb provides read replica capability for multiple database servers setup.
- [vitess(stars: 17257)](https://github.com/youtube/vitess) - vitess provides servers and tools which facilitate scaling of MySQL databases for large scale web services.

### SQL Query Builders

_Libraries for building and using SQL._

- [bqb(stars: 117)](https://github.com/nullism/bqb) - Lightweight and easy to learn query builder.
- [buildsqlx(stars: 134)](https://github.com/arthurkushman/buildsqlx) - Go database query builder library for PostgreSQL.
- [builq(stars: 75)](https://github.com/cristalhq/builq) - Easily build SQL queries in Go.
- [dbq(stars: 386)](https://github.com/rocketlaunchr/dbq) - Zero boilerplate database operations for Go.
- [Dotsql(stars: 696)](https://github.com/gchaincl/dotsql) - Go library that helps you keep sql files in one place and use them with ease.
- [gendry(stars: 1552)](https://github.com/didi/gendry) - Non-invasive SQL builder and powerful data binder.
- [godbal(stars: 59)](https://github.com/xujiajun/godbal) - Database Abstraction Layer (dbal) for go. Support SQL builder and get result easily.
- [goqu(stars: 2100)](https://github.com/doug-martin/goqu) - Idiomatic SQL builder and query library.
- [gosql(stars: 29)](https://github.com/twharmon/gosql) - SQL Query builder with better null values support.
- [Hotcoal(stars: 12)](https://github.com/motrboat/hotcoal) - Secure your handcrafted SQL against injection.
- [igor(stars: 115)](https://github.com/galeone/igor) - Abstraction layer for PostgreSQL that supports advanced functionality and uses gorm-like syntax.
- [jet(stars: 1746)](https://github.com/go-jet/jet) - Framework for writing type-safe SQL queries in Go, with ability to easily convert database query result into desired arbitrary object structure.
- [ormlite(stars: 12)](https://github.com/pupizoid/ormlite) - Lightweight package containing some ORM-like features and helpers for sqlite databases.
- [ozzo-dbx(stars: 613)](https://github.com/go-ozzo/ozzo-dbx) - Powerful data retrieval methods as well as DB-agnostic query building capabilities.
- [qry(stars: 29)](https://github.com/HnH/qry) - Tool that generates constants from files with raw SQL queries.
- [sg(stars: 4)](https://github.com/go-the-way/sg) - A SQL Gen for generating standard SQLs(supports: CRUD) written in Go.
- [sq(stars: 195)](https://github.com/bokwoon95/go-structured-query) - Type-safe SQL builder and struct mapper for Go.
- [sqlc(stars: 9709)](https://github.com/kyleconroy/sqlc) - Generate type-safe code from SQL.
- [sqlf(stars: 136)](https://github.com/leporo/sqlf) - Fast SQL query builder.
- [sqlingo(stars: 353)](https://github.com/lqs/sqlingo) - A lightweight DSL to build SQL in Go.
- [sqrl(stars: 265)](https://github.com/elgris/sqrl) - SQL query builder, fork of Squirrel with improved performance.
- [Squalus](https://gitlab.com/qosenergy/squalus) - Thin layer over the Go SQL package that makes it easier to perform queries.
- [Squirrel(stars: 6271)](https://github.com/Masterminds/squirrel) - Go library that helps you build SQL queries.
- [xo(stars: 3499)](https://github.com/knq/xo) - Generate idiomatic Go code for databases based on existing schema definitions or custom queries supporting PostgreSQL, MySQL, SQLite, Oracle, and Microsoft SQL Server.

**[⬆ back to top](#contents)**

## Database Drivers

### Interfaces to Multiple Backends

- [cayley(stars: 14692)](https://github.com/google/cayley) - Graph database with support for multiple backends.
- [dsc(stars: 28)](https://github.com/viant/dsc) - Datastore connectivity for SQL, NoSQL, structured files.
- [go-transaction-manager(stars: 109)](https://github.com/avito-tech/go-transaction-manager) - Transaction manager with multiple adapters (sql, sqlx, gorm, mongo, ...) controls transaction boundaries.
- [gokv(stars: 609)](https://github.com/philippgille/gokv) - Simple key-value store abstraction and implementations for Go (Redis, Consul, etcd, bbolt, BadgerDB, LevelDB, Memcached, DynamoDB, S3, PostgreSQL, MongoDB, CockroachDB and many more).

### Relational Database Drivers

- [avatica(stars: 110)](https://github.com/apache/calcite-avatica-go) - Apache Avatica/Phoenix SQL driver for database/sql.
- [bgc(stars: 20)](https://github.com/viant/bgc) - Datastore Connectivity for BigQuery for go.
- [firebirdsql(stars: 210)](https://github.com/nakagami/firebirdsql) - Firebird RDBMS SQL driver for Go.
- [go-adodb(stars: 133)](https://github.com/mattn/go-adodb) - Microsoft ActiveX Object DataBase driver for go that uses database/sql.
- [go-mssqldb(stars: 1780)](https://github.com/denisenkom/go-mssqldb) - Microsoft MSSQL driver for Go.
- [go-oci8(stars: 619)](https://github.com/mattn/go-oci8) - Oracle driver for go that uses database/sql.
- [go-sql-driver/mysql(stars: 13928)](https://github.com/go-sql-driver/mysql) - MySQL driver for Go.
- [go-sqlite3(stars: 7119)](https://github.com/mattn/go-sqlite3) - SQLite3 driver for go that uses database/sql.
- [godror(stars: 475)](https://github.com/godror/godror) - Oracle driver for Go, using the ODPI-C driver.
- [gofreetds(stars: 109)](https://github.com/minus5/gofreetds) - Microsoft MSSQL driver. Go wrapper over [FreeTDS](https://www.freetds.org).
- [KSQL(stars: 253)](https://github.com/VinGarcia/ksql) - A Simple and Powerful Golang SQL Library
- [pgx(stars: 8725)](https://github.com/jackc/pgx) - PostgreSQL driver supporting features beyond those exposed by database/sql.
- [pig(stars: 15)](https://github.com/alexeyco/pig) - Simple [pgx](https://github.com/jackc/pgx) wrapper to execute and [scan](https://github.com/georgysavva/scany) query results easily.
- [pq(stars: 8504)](https://github.com/lib/pq) - Pure Go Postgres driver for database/sql.
- [Sqinn-Go(stars: 372)](https://github.com/cvilsmeier/sqinn-go) - SQLite with pure Go.
- [sqlhooks(stars: 624)](https://github.com/qustavo/sqlhooks) - Attach hooks to any database/sql driver.
- [surrealdb.go](https://github.com/surrealdb/surrealdb.go) - SurrealDB Driver for Go.
- [ydb-go-sdk(stars: 113)](https://github.com/ydb-platform/ydb-go-sdk) - native and database/sql driver YDB (Yandex Database)

### NoSQL Database Drivers

- [aerospike-client-go(stars: 417)](https://github.com/aerospike/aerospike-client-go) - Aerospike client in Go language.
- [arangolite(stars: 73)](https://github.com/solher/arangolite) - Lightweight golang driver for ArangoDB.
- [asc(stars: 9)](https://github.com/viant/asc) - Datastore Connectivity for Aerospike for go.
- [forestdb(stars: 34)](https://github.com/couchbase/goforestdb) - Go bindings for ForestDB.
- [go-couchbase(stars: 319)](https://github.com/couchbase/go-couchbase) - Couchbase client in Go.
- [go-pilosa(stars: 58)](https://github.com/pilosa/go-pilosa) - Go client library for Pilosa.
- [go-rejson(stars: 331)](https://github.com/nitishm/go-rejson) - Golang client for redislabs' ReJSON module using Redigo golang client. Store and manipulate structs as JSON objects in redis with ease.
- [gocb(stars: 353)](https://github.com/couchbase/gocb) - Official Couchbase Go SDK.
- [gocosmos(stars: 19)](https://github.com/btnguyen2k/gocosmos) - REST client and standard `database/sql` driver for Azure Cosmos DB.
- [gocql](https://gocql.github.io) - Go language driver for Apache Cassandra.
- [godis(stars: 109)](https://github.com/piaohao/godis) - redis client implement by golang, inspired by jedis.
- [godscache(stars: 11)](https://github.com/defcronyke/godscache) - A wrapper for the Google Cloud Platform Go Datastore package that adds caching using memcached.
- [gomemcache](https://github.com/bradfitz/gomemcache/) - memcache client library for the Go programming language.
- [gorethink(stars: 1644)](https://github.com/dancannon/gorethink) - Go language driver for RethinkDB.
- [goriak(stars: 30)](https://github.com/zegl/goriak) - Go language driver for Riak KV.
- [Kivik(stars: 286)](https://github.com/go-kivik/kivik) - Kivik provides a common Go and GopherJS client library for CouchDB, PouchDB, and similar databases.
- [mgm(stars: 708)](https://github.com/kamva/mgm) - MongoDB model-based ODM for Go (based on official MongoDB driver).
- [mgo(stars: 1962)](https://github.com/globalsign/mgo) - (unmaintained) MongoDB driver for the Go language that implements a rich and well tested selection of features under a very simple API following standard Go idioms.
- [mongo-go-driver(stars: 7760)](https://github.com/mongodb/mongo-go-driver) - Official MongoDB driver for the Go language.
- [neo4j(stars: 27)](https://github.com/cihangir/neo4j) - Neo4j Rest API Bindings for Golang.
- [Neo4j-GO(stars: 76)](https://github.com/davemeehan/Neo4j-GO) - Neo4j REST Client in golang.
- [neoism(stars: 389)](https://github.com/jmcvetta/neoism) - Neo4j client for Golang.
- [qmgo(stars: 1226)](https://github.com/qiniu/qmgo) - The MongoDB driver for Go. It‘s based on official MongoDB driver but easier to use like Mgo.
- [redeo(stars: 430)](https://github.com/bsm/redeo) - Redis-protocol compatible TCP servers/services.
- [redigo(stars: 9667)](https://github.com/gomodule/redigo) - Redigo is a Go client for the Redis database.
- [redis(stars: 18536)](https://github.com/redis/go-redis) - Redis client for Golang.
- [rueidis](http://github.com/rueian/rueidis) - Fast Redis RESP3 client with auto pipelining and server-assisted client side caching.
- [xredis(stars: 19)](https://github.com/shomali11/xredis) - Typesafe, customizable, clean & easy to use Redis client.

### Search and Analytic Databases

- [clickhouse-go](https://github.com/ClickHouse/clickhouse-go/) - ClickHouse SQL client for Go with a `database/sql` compatibility.
- [elastic(stars: 7263)](https://github.com/olivere/elastic) - Elasticsearch client for Go.
- [elasticsql(stars: 1120)](https://github.com/cch123/elasticsql) - Convert sql to elasticsearch dsl in Go.
- [elastigo(stars: 948)](https://github.com/mattbaird/elastigo) - Elasticsearch client library.
- [go-elasticsearch(stars: 5309)](https://github.com/elastic/go-elasticsearch) - Official Elasticsearch client for Go.
- [goes(stars: 30)](https://github.com/OwnLocal/goes) - Library to interact with Elasticsearch.
- [skizze(stars: 88)](https://github.com/seiflotfy/skizze) - probabilistic data-structures service and storage.

**[⬆ back to top](#contents)**

## Date and Time

_Libraries for working with dates and times._

- [approx(stars: 5)](https://github.com/goschtalt/approx) - A Duration extension supporting parsing/printing durations in days, weeks and years.
- [carbon(stars: 4172)](https://github.com/golang-module/carbon) - A simple, semantic and developer-friendly golang package for datetime.
- [carbon(stars: 767)](https://github.com/uniplaces/carbon) - Simple Time extension with a lot of util methods, ported from PHP Carbon library.
- [cronrange(stars: 18)](https://github.com/1set/cronrange) - Parses Cron-style time range expressions, checks if the given time is within any ranges.
- [date(stars: 118)](https://github.com/rickb777/date) - Augments Time for working with dates, date ranges, time spans, periods, and time-of-day.
- [dateparse(stars: 1945)](https://github.com/araddon/dateparse) - Parse date's without knowing format in advance.
- [durafmt(stars: 478)](https://github.com/hako/durafmt) - Time duration formatting library for Go.
- [feiertage(stars: 45)](https://github.com/wlbr/feiertage) - Set of functions to calculate public holidays in Germany, incl. specialization on the states of Germany (Bundesländer). Things like Easter, Pentecost, Thanksgiving...
- [go-anytime(stars: 13)](https://github.com/ijt/go-anytime) - Parse dates/times like "next dec 22nd at 3pm" and ranges like "from today until next thursday" without knowing the format in advance.
- [go-persian-calendar(stars: 177)](https://github.com/yaa110/go-persian-calendar) - The implementation of the Persian (Solar Hijri) Calendar in Go (golang).
- [go-str2duration(stars: 85)](https://github.com/xhit/go-str2duration) - Convert string to duration. Support time.Duration returned string and more.
- [go-sunrise(stars: 79)](https://github.com/nathan-osman/go-sunrise) - Calculate the sunrise and sunset times for a given location.
- [go-week(stars: 9)](https://github.com/stoewer/go-week) - An efficient package to work with ISO8601 week dates.
- [gostradamus(stars: 192)](https://github.com/bykof/gostradamus) - A Go package for working with dates.
- [iso8601(stars: 133)](https://github.com/relvacode/iso8601) - Efficiently parse ISO8601 date-times without regex.
- [kair(stars: 25)](https://github.com/GuilhermeCaruso/kair) - Date and Time - Golang Formatting Library.
- [now(stars: 4304)](https://github.com/jinzhu/now) - Now is a time toolkit for golang.
- [NullTime(stars: 14)](https://github.com/kirillDanshin/nulltime) - Nullable `time.Time`.
- [strftime(stars: 12)](https://github.com/awoodbeck/strftime) - C99-compatible strftime formatter.
- [timespan(stars: 82)](https://github.com/SaidinWoT/timespan) - For interacting with intervals of time, defined as a start time and a duration.
- [timeutil(stars: 194)](https://github.com/leekchan/timeutil) - Useful extensions (Timedelta, Strftime, ...) to the golang's time package.
- [tuesday(stars: 12)](https://github.com/osteele/tuesday) - Ruby-compatible Strftime function.

**[⬆ back to top](#contents)**

## Distributed Systems

_Packages that help with building Distributed Systems._

- [arpc(stars: 828)](https://github.com/lesismal/arpc) - More effective network communication, support two-way-calling, notify, broadcast.
- [celeriac](https://github.com/svcavallar/celeriac.v1) - Library for adding support for interacting and monitoring Celery workers, tasks and events in Go.
- [consistent(stars: 634)](https://github.com/buraksezer/consistent) - Consistent hashing with bounded loads.
- [consistenthash(stars: 22)](https://github.com/mbrostami/consistenthash) - Consistent hashing with configurable replicas.
- [dht(stars: 282)](https://github.com/anacrolix/dht) - BitTorrent Kademlia DHT implementation.
- [digota(stars: 487)](https://github.com/digota/digota) - grpc ecommerce microservice.
- [dot](https://github.com/dotchain/dot/) - distributed sync using operational transformation/OT.
- [doublejump(stars: 94)](https://github.com/edwingeng/doublejump) - A revamped Google's jump consistent hash.
- [dragonboat(stars: 4817)](https://github.com/lni/dragonboat) - A feature complete and high performance multi-group Raft library in Go.
- [Dragonfly(stars: 1790)](https://github.com/dragonflyoss/Dragonfly2) - Provide efficient, stable and secure file distribution and image acceleration based on p2p technology to be the best practice and standard solution in cloud native architectures.
- [drmaa(stars: 46)](https://github.com/dgruber/drmaa) - Job submission library for cluster schedulers based on the DRMAA standard.
- [dynamolock](https://cirello.io/dynamolock) - DynamoDB-backed distributed locking implementation.
- [dynatomic(stars: 16)](https://github.com/tylfin/dynatomic) - A library for using DynamoDB as an atomic counter.
- [emitter-io(stars: 3682)](https://github.com/emitter-io/emitter) - High performance, distributed, secure and low latency publish-subscribe platform built with MQTT, Websockets and love.
- [failured(stars: 11)](https://github.com/andy2046/failured) - adaptive accrual failure detector for distributed systems.
- [flowgraph(stars: 54)](https://github.com/vectaport/flowgraph) - flow-based programming package.
- [gleam(stars: 3286)](https://github.com/chrislusf/gleam) - Fast and scalable distributed map/reduce system written in pure Go and Luajit, combining Go's high concurrency with Luajit's high performance, runs standalone or distributed.
- [glow(stars: 3176)](https://github.com/chrislusf/glow) - Easy-to-Use scalable distributed big data processing, Map-Reduce, DAG execution, all in pure Go.
- [gmsec(stars: 23)](https://github.com/gmsec/micro) - A Go distributed systems development framework.
- [go-doudou(stars: 1261)](https://github.com/unionj-cloud/go-doudou) - A gossip protocol and OpenAPI 3.0 spec based decentralized microservice framework. Built-in go-doudou cli focusing on low-code and rapid dev can power up your productivity.
- [go-health(stars: 730)](https://github.com/InVisionApp/go-health) - Library for enabling asynchronous dependency health checks in your service.
- [go-jump(stars: 370)](https://github.com/dgryski/go-jump) - Port of Google's "Jump" Consistent Hash function.
- [go-kit(stars: 25782)](https://github.com/go-kit/kit) - Microservice toolkit with support for service discovery, load balancing, pluggable transports, request tracking, etc.
- [go-micro(stars: 21089)](https://github.com/micro/go-micro) - A distributed systems development framework.
- [go-mysql-lock(stars: 58)](https://github.com/sanketplus/go-mysql-lock) - MySQL based distributed lock.
- [go-pdu(stars: 46)](https://github.com/pdupub/go-pdu) - A decentralized identity-based social network.
- [go-sundheit(stars: 517)](https://github.com/AppsFlyer/go-sundheit) - A library built to provide support for defining async service health checks for golang services.
- [go-zero(stars: 26705)](https://github.com/tal-tech/go-zero) - A web and rpc framework. It's born to ensure the stability of the busy sites with resilient design. Builtin goctl greatly improves the development productivity.
- [gorpc(stars: 683)](https://github.com/valyala/gorpc) - Simple, fast and scalable RPC library for high load.
- [grpc-go(stars: 19306)](https://github.com/grpc/grpc-go) - The Go language implementation of gRPC. HTTP/2 based RPC.
- [hprose(stars: 1249)](https://github.com/hprose/hprose-golang) - Very newbility RPC Library, support 25+ languages now.
- [jsonrpc(stars: 184)](https://github.com/osamingo/jsonrpc) - The jsonrpc package helps implement of JSON-RPC 2.0.
- [jsonrpc(stars: 296)](https://github.com/ybbus/jsonrpc) - JSON-RPC 2.0 HTTP client implementation.
- [Kitex(stars: 6441)](https://github.com/cloudwego/kitex) - A high-performance and strong-extensibility Golang RPC framework that helps developers build microservices. If the performance and extensibility are the main concerns when you develop microservices, Kitex can be a good choice.
- [Kratos(stars: 21894)](https://github.com/go-kratos/kratos) - A modular-designed and easy-to-use microservices framework in Go.
- [liftbridge(stars: 2507)](https://github.com/liftbridge-io/liftbridge) - Lightweight, fault-tolerant message streams for NATS.
- [lura(stars: 5860)](https://github.com/luraproject/lura) - Ultra performant API Gateway framework with middlewares.
- [micro(stars: 11941)](https://github.com/micro/micro) - A distributed systems runtime for the cloud and beyond.
- [NATS(stars: 14086)](https://github.com/nats-io/gnatsd) - Lightweight, high performance messaging system for microservices, IoT, and cloud native systems.
- [outboxer(stars: 148)](https://github.com/italolelis/outboxer) - Outboxer is a go library that implements the outbox pattern.
- [pglock](https://cirello.io/pglock) - PostgreSQL-backed distributed locking implementation.
- [pjrpc](https://gitlab.com/pjrpc/pjrpc) - Golang JSON-RPC Server-Client with Protobuf spec.
- [raft(stars: 7601)](https://github.com/hashicorp/raft) - Golang implementation of the Raft consensus protocol, by HashiCorp.
- [raft(stars: 420)](https://github.com/etcd-io/raft) - Go implementation of the Raft consensus protocol, by CoreOS.
- [rain(stars: 896)](https://github.com/cenkalti/rain) - BitTorrent client and library.
- [redis-lock(stars: 1206)](https://github.com/bsm/redislock) - Simplified distributed locking implementation using Redis.
- [resgate](https://resgate.io/) - Realtime API Gateway for building REST, real time, and RPC APIs, where all clients are synchronized seamlessly.
- [ringpop-go(stars: 804)](https://github.com/uber/ringpop-go) - Scalable, fault-tolerant application-layer sharding for Go applications.
- [rpcx(stars: 7838)](https://github.com/smallnest/rpcx) - Distributed pluggable RPC service framework like alibaba Dubbo.
- [Semaphore(stars: 89)](https://github.com/jexia/semaphore) - A straightforward (micro) service orchestrator.
- [sleuth(stars: 368)](https://github.com/ursiform/sleuth) - Library for master-less p2p auto-discovery and RPC between HTTP services (using [ZeroMQ](https://github.com/zeromq/libzmq)).
- [Temporal(stars: 414)](https://github.com/temporalio/sdk-go) - Durable execution system for making code fault-tolerant and simple.
- [torrent(stars: 5030)](https://github.com/anacrolix/torrent) - BitTorrent client package.

**[⬆ back to top](#contents)**

## Dynamic DNS

_Tools for updating dynamic DNS records._

- [DDNS(stars: 239)](https://github.com/skibish/ddns) - Personal DDNS client with Digital Ocean Networking DNS as backend.
- [dyndns](https://gitlab.com/alcastle/dyndns) - Background Go process to regularly and automatically check your IP Address and make updates to (one or many) Dynamic DNS records for Google domains whenever your address changes.
- [GoDNS(stars: 1384)](https://github.com/timothyye/godns) - A dynamic DNS client tool, supports DNSPod & HE.net, written in Go.

**[⬆ back to top](#contents)**

## Email

_Libraries and tools that implement email creation and sending._

- [chasquid](https://blitiri.com.ar/p/chasquid) - SMTP server written in Go.
- [douceur(stars: 234)](https://github.com/aymerick/douceur) - CSS inliner for your HTML emails.
- [email(stars: 2492)](https://github.com/jordan-wright/email) - A robust and flexible email library for Go.
- [email-verifier(stars: 991)](https://github.com/AfterShip/email-verifier) - A Go library for email verification without sending any emails.
- [go-dkim(stars: 92)](https://github.com/toorop/go-dkim) - DKIM library, to sign & verify email.
- [go-email-normalizer(stars: 60)](https://github.com/dimuska139/go-email-normalizer) - Golang library for providing a canonical representation of email address.
- [go-email-validator(stars: 52)](https://github.com/go-email-validator/go-email-validator) - Modular email validator for syntax, disposable, smtp, etc... checking.
- [go-imap(stars: 1902)](https://github.com/emersion/go-imap) - IMAP library for clients and servers.
- [go-mail(stars: 374)](https://github.com/wneessen/go-mail) - A simple Go library for sending mails in Go.
- [go-message(stars: 342)](https://github.com/emersion/go-message) - Streaming library for the Internet Message Format and mail messages.
- [go-premailer(stars: 126)](https://github.com/vanng822/go-premailer) - Inline styling for HTML mail in Go.
- [go-simple-mail(stars: 553)](https://github.com/xhit/go-simple-mail) - Very simple package to send emails with SMTP Keep Alive and two timeouts: Connect and Send.
- [Hectane(stars: 218)](https://github.com/hectane/hectane) - Lightweight SMTP client providing an HTTP API.
- [hermes(stars: 2723)](https://github.com/matcornic/hermes) - Golang package that generates clean, responsive HTML e-mails.
- [Maddy(stars: 4459)](https://github.com/foxcpp/maddy) - All-in-one (SMTP, IMAP, DKIM, DMARC, MTA-STS, DANE) email server
- [mailchain(stars: 141)](https://github.com/mailchain/mailchain) - Send encrypted emails to blockchain addresses written in Go.
- [mailgun-go(stars: 673)](https://github.com/mailgun/mailgun-go) - Go library for sending mail with the Mailgun API.
- [MailHog(stars: 12900)](https://github.com/mailhog/MailHog) - Email and SMTP testing with web and API interface.
- [Mailpit(stars: 3711)](https://github.com/axllent/mailpit) - Email and SMTP testing tool for developers.
- [mailx(stars: 9)](https://github.com/valord577/mailx) - Mailx is a library that makes it easier to send email via SMTP. It is an enhancement of the golang standard library `net/smtp`.
- [SendGrid(stars: 942)](https://github.com/sendgrid/sendgrid-go) - SendGrid's Go library for sending email.
- [smtp(stars: 74)](https://github.com/mailhog/smtp) - SMTP server protocol state machine.
- [smtpmock(stars: 95)](https://github.com/mocktools/go-smtp-mock) - Lightweight configurable multithreaded fake SMTP server. Mimic any SMTP behaviour for your test environment.
- [truemail-go(stars: 68)](https://github.com/truemail-rb/truemail-go) - Configurable Golang email validator/verifier. Verify email via Regex, DNS, SMTP and even more.

**[⬆ back to top](#contents)**

## Embeddable Scripting Languages

_Embedding other languages inside your go code._

- [anko(stars: 1405)](https://github.com/mattn/anko) - Scriptable interpreter written in Go.
- [binder(stars: 68)](https://github.com/alexeyco/binder) - Go to Lua binding library, based on [gopher-lua](https://github.com/yuin/gopher-lua).
- [cel-go(stars: 1934)](https://github.com/google/cel-go) - Fast, portable, non-Turing complete expression evaluation with gradual typing.
- [ecal(stars: 37)](https://github.com/krotik/ecal) - A simple embeddable scripting language which supports concurrent event processing.
- [expr(stars: 4878)](https://github.com/antonmedv/expr) - Expression evaluation engine for Go: fast, non-Turing complete, dynamic typing, static typing.
- [gentee(stars: 123)](https://github.com/gentee/gentee) - Embeddable scripting programming language.
- [gisp(stars: 508)](https://github.com/jcla1/gisp) - Simple LISP in Go.
- [go-duktape(stars: 779)](https://github.com/olebedev/go-duktape) - Duktape JavaScript engine bindings for Go.
- [go-lua(stars: 2837)](https://github.com/Shopify/go-lua) - Port of the Lua 5.2 VM to pure Go.
- [go-php(stars: 914)](https://github.com/deuill/go-php) - PHP bindings for Go.
- [go-python(stars: 1514)](https://github.com/sbinet/go-python) - naive go bindings to the CPython C-API.
- [goja(stars: 4514)](https://github.com/dop251/goja) - ECMAScript 5.1(+) implementation in Go.
- [golua(stars: 630)](https://github.com/aarzilli/golua) - Go bindings for Lua C API.
- [gopher-lua(stars: 5799)](https://github.com/yuin/gopher-lua) - Lua 5.1 VM and compiler written in Go.
- [gval(stars: 681)](https://github.com/PaesslerAG/gval) - A highly customizable expression language written in Go.
- [metacall(stars: 1430)](https://github.com/metacall/core) - Cross-platform Polyglot Runtime which supports NodeJS, JavaScript, TypeScript, Python, Ruby, C#, WebAssembly, Java, Cobol and more.
- [ngaro(stars: 25)](https://github.com/db47h/ngaro) - Embeddable Ngaro VM implementation enabling scripting in Retro.
- [prolog(stars: 518)](https://github.com/ichiban/prolog) - Embeddable Prolog.
- [purl(stars: 39)](https://github.com/ian-kent/purl) - Perl 5.18.2 embedded in Go.
- [starlark-go(stars: 2138)](https://github.com/google/starlark-go) - Go implementation of Starlark: Python-like language with deterministic evaluation and hermetic execution.
- [tengo(stars: 3357)](https://github.com/d5/tengo) - Bytecode compiled script language for Go.

**[⬆ back to top](#contents)**

## Error Handling

_Libraries for handling errors._

- [emperror(stars: 317)](https://github.com/emperror/emperror) - Error handling tools and best practices for Go libraries and applications.
- [eris(stars: 1379)](https://github.com/rotisserie/eris) - A better way to handle, trace, and log errors in Go. Compatible with the standard error library and github.com/pkg/errors.
- [errlog(stars: 453)](https://github.com/snwfdhmp/errlog) - Hackable package that determines responsible source code for an error (and some other fast-debugging features). Pluggable to any logger in-place.
- [errors(stars: 191)](https://github.com/emperror/errors) - Drop-in replacement for the standard library errors package and github.com/pkg/errors. Provides various error handling primitives.
- [errors(stars: 6)](https://github.com/neuronlabs/errors) - Simple golang error handling with classification primitives.
- [errors(stars: 7)](https://github.com/PumpkinSeed/errors) - The most simple error wrapper with awesome performance and minimal memory overhead.
- [errors](https://gitlab.com/tozd/go/errors) - Providing errors with a stack trace and optional structured details. Compatible with github.com/pkg/errors API but does not use it internally.
- [errors(stars: 60)](https://github.com/bnkamalesh/errors) - Drop-in replacement for builtin Go errors. This is a minimal error handling package with custom error types, user friendly messages, Unwrap & Is. With very easy to use and straightforward helper functions.
- [errors(stars: 1877)](https://github.com/cockroachdb/errors) - Go error library with error portability over the network.
- [errorx(stars: 1045)](https://github.com/joomcode/errorx) - A feature rich error package with stack traces, composition of errors and more.
- [exception(stars: 26)](https://github.com/rbrahul/exception) - A simple utility package for exception handling with try-catch in Golang.
- [Falcon(stars: 10)](https://github.com/SonicRoshan/falcon) - A Simple Yet Highly Powerful Package For Error Handling.
- [Fault(stars: 145)](https://github.com/Southclaws/fault) - An ergonomic mechanism for wrapping errors in order to facilitate structured metadata and context for error values.
- [go-multierror(stars: 2096)](https://github.com/hashicorp/go-multierror) - Go (golang) package for representing a list of errors as a single error.
- [oops(stars: 130)](https://github.com/samber/oops) - Error handling with context, stack trace and source fragments.
- [tracerr(stars: 860)](https://github.com/ztrue/tracerr) - Golang errors with stack trace and source fragments.

**[⬆ back to top](#contents)**

## File Handling

_Libraries for handling files and file systems._

- [afero(stars: 5530)](https://github.com/spf13/afero) - FileSystem Abstraction System for Go.
- [afs(stars: 274)](https://github.com/viant/afs) - Abstract File Storage (mem, scp, zip, tar, cloud: s3, gs) for Go.
- [baraka(stars: 52)](https://github.com/xis/baraka) - A library to process http file uploads easily.
- [bigfile(stars: 246)](https://github.com/bigfile/bigfile) - A file transfer system, support to manage files with http api, rpc call and ftp client.
- [checksum(stars: 96)](https://github.com/codingsince1985/checksum) - Compute message digest, like MD5, SHA256, SHA1, CRC or BLAKE2s, for large files.
- [copy(stars: 634)](https://github.com/otiai10/copy) - Copy directory recursively.
- [flop(stars: 34)](https://github.com/homedepot/flop) - File operations library which aims to mirror feature parity with [GNU cp](https://www.gnu.org/software/coreutils/manual/html_node/cp-invocation.html).
- [gdu(stars: 2959)](https://github.com/dundee/gdu) - Disk usage analyzer with console interface.
- [go-csv-tag(stars: 111)](https://github.com/artonge/go-csv-tag) - Load csv file using tag.
- [go-decent-copy(stars: 21)](https://github.com/hugocarreira/go-decent-copy) - Copy files for humans.
- [go-exiftool(stars: 206)](https://github.com/barasher/go-exiftool) - Go bindings for ExifTool, the well-known library used to extract as much metadata as possible (EXIF, IPTC, ...) from files (pictures, PDF, office, ...).
- [go-gtfs(stars: 38)](https://github.com/artonge/go-gtfs) - Load gtfs files in go.
- [gofs(stars: 384)](https://github.com/no-src/gofs) - A cross-platform real-time file synchronization tool out of the box.
- [gut/yos(stars: 27)](https://github.com/1set/gut) - Simple and reliable package for file operations like copy/move/diff/list on files, directories and symbolic links.
- [higgs(stars: 18)](https://github.com/dastoori/higgs) - A tiny cross-platform Go library to hide/unhide files and directories.
- [iso9660(stars: 247)](https://github.com/kdomanski/iso9660) - A package for reading and creating ISO9660 disk images
- [notify(stars: 864)](https://github.com/rjeczalik/notify) - File system event notification library with simple API, similar to os/signal.
- [opc(stars: 74)](https://github.com/qmuntal/opc) - Load Open Packaging Conventions (OPC) files for Go.
- [parquet(stars: 87)](https://github.com/parsyl/parquet) - Read and write [parquet](https://parquet.apache.org) files.
- [pathtype(stars: 13)](https://github.com/jonchun/pathtype) - Treat paths as their own type instead of using strings.
- [pdfcpu(stars: 5767)](https://github.com/pdfcpu/pdfcpu) - PDF processor.
- [skywalker(stars: 95)](https://github.com/dixonwille/skywalker) - Package to allow one to concurrently go through a filesystem with ease.
- [stl](https://gitlab.com/russoj88/stl) - Modules to read and write STL (stereolithography) files. Concurrent algorithm for reading.
- [tarfs(stars: 59)](https://github.com/posener/tarfs) - Implementation of the [`FileSystem` interface](https://godoc.org/github.com/kr/fs#FileSystem) for tar files.
- [todotxt(stars: 20)](https://github.com/1set/todotxt) - Go library for Gina Trapani's [_todo.txt_](http://todotxt.org/) files, supports parsing and manipulating of task lists in the [_todo.txt_ format](https://github.com/todotxt/todo.txt).
- [vfs(stars: 263)](https://github.com/C2FO/vfs) - A pluggable, extensible, and opinionated set of filesystem functionality for Go across a number of filesystem types such as os, S3, and GCS.

**[⬆ back to top](#contents)**

## Financial

_Packages for accounting and finance._

- [accounting(stars: 842)](https://github.com/leekchan/accounting) - money and currency formatting for golang.
- [ach(stars: 413)](https://github.com/moov-io/ach) - A reader, writer, and validator for Automated Clearing House (ACH) files.
- [bbgo(stars: 1013)](https://github.com/c9s/bbgo) - A crypto trading bot framework written in Go. Including common crypto exchange API, standard indicators, back-testing and many built-in strategies.
- [currency(stars: 461)](https://github.com/bojanz/currency) - Handles currency amounts, provides currency information and formatting.
- [currency(stars: 56)](https://github.com/bnkamalesh/currency) - High performant & accurate currency computation package.
- [decimal(stars: 5573)](https://github.com/shopspring/decimal) - Arbitrary-precision fixed-point decimal numbers.
- [decimal(stars: 15)](https://github.com/govalues/decimal) - Immutable decimal numbers with panic-free arithmetic.
- [fpdecimal(stars: 25)](https://github.com/nikolaydubina/fpdecimal) - Fast and precise serialization and arithmetic for small fixed-point decimals
- [fpmoney(stars: 22)](https://github.com/nikolaydubina/fpmoney) - Fast and simple ISO4217 fixed-point decimal money.
- [go-finance(stars: 159)](https://github.com/alpeb/go-finance) - Library of financial functions for time value of money (annuities), cash flow, interest rate conversions, bonds and depreciation calculations.
- [go-finance(stars: 23)](https://github.com/pieterclaerhout/go-finance) - Module to fetch exchange rates, check VAT numbers via VIES and check IBAN bank account numbers.
- [go-finnhub(stars: 84)](https://github.com/m1/go-finnhub) - Client for stock market, forex and crypto data from finnhub.io. Access real-time financial market data from 60+ stock exchanges, 10 forex brokers, and 15+ crypto exchanges.
- [go-money(stars: 1435)](https://github.com/rhymond/go-money) - Implementation of Fowler's Money pattern.
- [go-nowpayments(stars: 2)](https://github.com/matm/go-nowpayments) - Library for the crypto NOWPayments API.
- [money(stars: 9)](https://github.com/govalues/money) - Immutable monetary amounts and exchange rates with panic-free arithmetic.
- [ofxgo(stars: 125)](https://github.com/aclindsa/ofxgo) - Query OFX servers and/or parse the responses (with example command-line client).
- [orderbook(stars: 376)](https://github.com/i25959341/orderbook) - Matching Engine for Limit Order Book in Golang.
- [payme(stars: 27)](https://github.com/jovandeginste/payme) - QR code generator (ASCII & PNG) for SEPA payments.
- [sleet(stars: 129)](https://github.com/BoltApp/sleet) - One unified interface for multiple Payment Service Providers (PsP) to process online payment.
- [techan(stars: 764)](https://github.com/sdcoffey/techan) - Technical analysis library with advanced market analysis and trading strategies.
- [ticker(stars: 4723)](https://github.com/achannarasappa/ticker) - Terminal stock watcher and stock position tracker.
- [transaction(stars: 126)](https://github.com/claygod/transaction) - Embedded transactional database of accounts, running in multithreaded mode.
- [vat(stars: 107)](https://github.com/dannyvankooten/vat) - VAT number validation & EU VAT rates.

**[⬆ back to top](#contents)**

## Forms

_Libraries for working with forms._

- [bind(stars: 30)](https://github.com/robfig/bind) - Bind form data to any Go values.
- [binding(stars: 794)](https://github.com/mholt/binding) - Binds form and JSON data from net/http Request to struct.
- [checker(stars: 5)](https://github.com/cinar/checker) - Checker helps validating user input through rules defined in struct tags or directly through functions.
- [conform(stars: 312)](https://github.com/leebenson/conform) - Keeps user input in check. Trims, sanitizes & scrubs data based on struct tags.
- [form(stars: 654)](https://github.com/go-playground/form) - Decodes url.Values into Go value(s) and Encodes Go value(s) into url.Values. Dual Array and Full map support.
- [formam(stars: 182)](https://github.com/monoculum/formam) - decode form's values into a struct.
- [forms(stars: 133)](https://github.com/albrow/forms) - Framework-agnostic library for parsing and validating form/JSON data which supports multipart forms and files.
- [gbind(stars: 8)](https://github.com/bdjimmy/gbind) - Bind data to any Go value. Can use built-in and custom expression binding capabilities; supports data validation
- [gorilla/csrf(stars: 960)](https://github.com/gorilla/csrf) - CSRF protection for Go web applications & services.
- [httpin(stars: 223)](https://github.com/ggicci/httpin) - Decode an HTTP request into a custom struct, including querystring, forms, HTTP headers, etc.
- [nosurf(stars: 1471)](https://github.com/justinas/nosurf) - CSRF protection middleware for Go.
- [qs(stars: 68)](https://github.com/sonh/qs) - Go module for encoding structs into URL query parameters.
- [queryparam(stars: 19)](https://github.com/tomwright/queryparam) - Decode `url.Values` into usable struct values of standard or custom types.

**[⬆ back to top](#contents)**

## Functional

_Packages to support functional programming in Go._

- [fp-go(stars: 270)](https://github.com/repeale/fp-go) - Collection of Functional Programming helpers powered by Golang 1.18+ generics.
- [fpGo(stars: 331)](https://github.com/TeaEntityLab/fpGo) - Monad, Functional Programming features for Golang.
- [fuego(stars: 141)](https://github.com/seborama/fuego) - Functional Experiment in Go.
- [go-functional(stars: 237)](https://github.com/BooleanCat/go-functional) - Functional programming in Go using generics
- [go-underscore(stars: 1292)](https://github.com/tobyhede/go-underscore) - Useful collection of helpfully functional Go collection utilities.
- [gofp(stars: 143)](https://github.com/rbrahul/gofp) - A lodash like powerful utility library for Golang.
- [mo(stars: 2039)](https://github.com/samber/mo) - Monads and popular FP abstractions, based on Go 1.18+ Generics (Option, Result, Either...).
- [underscore(stars: 94)](https://github.com/rjNemo/underscore) - Functional programming helpers for Go 1.18 and beyond.
- [valor(stars: 15)](https://github.com/phelmkamp/valor) - Generic option and result types that optionally contain a value.

**[⬆ back to top](#contents)**

## Game Development

_Awesome game development libraries._

- [Azul3D(stars: 596)](https://github.com/azul3d/engine) - 3D game engine written in Go.
- [Ebitengine(stars: 9274)](https://github.com/hajimehoshi/ebiten) - dead simple 2D game engine in Go.
- [engo(stars: 1685)](https://github.com/EngoEngine/engo) - Engo is an open-source 2D game engine written in Go. It follows the Entity-Component-System paradigm.
- [fantasyname(stars: 25)](https://github.com/s0rg/fantasyname) - Fantasy names generator.
- [g3n(stars: 2542)](https://github.com/g3n/engine) - Go 3D Game Engine.
- [go-astar(stars: 570)](https://github.com/beefsack/go-astar) - Go implementation of the A\* path finding algorithm.
- [go-sdl2(stars: 2067)](https://github.com/veandco/go-sdl2) - Go bindings for the [Simple DirectMedia Layer](https://www.libsdl.org/).
- [go3d(stars: 285)](https://github.com/ungerik/go3d) - Performance oriented 2D/3D math package for Go.
- [gonet(stars: 1239)](https://github.com/xtaci/gonet) - Game server skeleton implemented with golang.
- [goworld(stars: 2439)](https://github.com/xiaonanln/goworld) - Scalable game server engine, featuring space-entity framework and hot-swapping.
- [Harfang3D(stars: 434)](https://github.com/harfang3d/harfang3d) - 3D engine for the Go language, works on Windows and Linux ([Harfang on Go.dev](https://github.com/harfang3d/harfang-go)).
- [Leaf(stars: 5043)](https://github.com/name5566/leaf) - Lightweight game server framework.
- [nano(stars: 2571)](https://github.com/lonng/nano) - Lightweight, facility, high performance golang based game server framework.
- [Oak(stars: 1470)](https://github.com/oakmound/oak) - Pure Go game engine.
- [Pitaya(stars: 2050)](https://github.com/topfreegames/pitaya) - Scalable game server framework with clustering support and client libraries for iOS, Android, Unity and others through the C SDK.
- [Pixel(stars: 4364)](https://github.com/faiface/pixel) - Hand-crafted 2D game library in Go.
- [prototype(stars: 83)](https://github.com/gonutz/prototype) - Cross-platform (Windows/Linux/Mac) library for creating desktop games using a minimal API.
- [raylib-go(stars: 1168)](https://github.com/gen2brain/raylib-go) - Go bindings for [raylib](https://www.raylib.com/), a simple and easy-to-use library to learn videogames programming.
- [termloop(stars: 1375)](https://github.com/JoelOtter/termloop) - Terminal-based game engine for Go, built on top of Termbox.
- [tile(stars: 132)](https://github.com/kelindar/tile) - Data-oriented and cache-friendly 2D Grid library (TileMap), includes pathfinding, observers and import/export.

**[⬆ back to top](#contents)**

## Generators

_Tools that generate Go code._

- [convergen(stars: 19)](https://github.com/reedom/convergen) - Feature rich type-to-type copy code generator.
- [copygen(stars: 306)](https://github.com/switchupcb/copygen) - Generate type-to-type and type-based code without reflection.
- [generis(stars: 43)](https://github.com/senselogic/GENERIS) - Code generation tool providing generics, free-form macros, conditional compilation and HTML templating.
- [go-enum(stars: 594)](https://github.com/abice/go-enum) - Code generation for enums from code comments.
- [go-linq(stars: 3390)](https://github.com/ahmetalpbalkan/go-linq) - .NET LINQ-like query methods for Go.
- [goderive(stars: 1140)](https://github.com/awalterschulze/goderive) - Derives functions from input types.
- [gotype(stars: 57)](https://github.com/wzshiming/gotype) - Golang source code parsing, usage like reflect package.
- [goverter(stars: 386)](https://github.com/jmattheis/goverter) - Generate converters by defining an interface.
- [GoWrap(stars: 837)](https://github.com/hexdigest/gowrap) - Generate decorators for Go interfaces using simple templates.
- [interfaces(stars: 404)](https://github.com/rjeczalik/interfaces) - Command line tool for generating interface definitions.
- [jennifer(stars: 3098)](https://github.com/dave/jennifer) - Generate arbitrary Go code without templates.
- [oapi-codegen(stars: 4661)](https://github.com/deepmap/oapi-codegen) - This package contains a set of utilities for generating Go boilerplate code for services based on OpenAPI 3.0 API definitions.
- [typeregistry(stars: 23)](https://github.com/xiaoxin01/typeregistry) - A library to create type dynamically.

**[⬆ back to top](#contents)**

## Geographic

_Geographic tools and servers_

- [geoserver(stars: 85)](https://github.com/hishamkaram/geoserver) - geoserver Is a Go Package For Manipulating a GeoServer Instance via the GeoServer REST API.
- [gismanager(stars: 50)](https://github.com/hishamkaram/gismanager) - Publish Your GIS Data(Vector Data) to PostGIS and Geoserver.
- [godal(stars: 110)](https://github.com/airbusgeo/godal) - Go wrapper for GDAL.
- [H3(stars: 256)](https://github.com/uber/h3-go) - Go bindings for H3, a hierarchical hexagonal geospatial indexing system.
- [H3 GeoJSON(stars: 3)](https://github.com/mmadfox/go-geojson2h3) - Conversion utilities between H3 indexes and GeoJSON.
- [H3GeoDist(stars: 2)](https://github.com/mmadfox/go-h3geo-dist) - Distribution of Uber H3geo cells by virtual nodes.
- [mbtileserver(stars: 565)](https://github.com/consbio/mbtileserver) - A simple Go-based server for map tiles stored in mbtiles format.
- [osm(stars: 313)](https://github.com/paulmach/osm) - Library for reading, writing and working with OpenStreetMap data and APIs.
- [pbf(stars: 47)](https://github.com/maguro/pbf) - OpenStreetMap PBF golang encoder/decoder.
- [S2 geojson(stars: 26)](https://github.com/pantrif/s2-geojson) - Convert geojson to s2 cells & demonstrating some S2 geometry features on map.
- [S2 geometry(stars: 1606)](https://github.com/golang/geo) - S2 geometry library in Go.
- [simplefeatures(stars: 105)](https://github.com/peterstace/simplefeatures) - simplesfeatures is a 2D geometry library that provides Go types that model geometries, as well as algorithms that operate on them.
- [Tile38(stars: 8785)](https://github.com/tidwall/tile38) - Geolocation DB with spatial index and realtime geofencing.
- [Web-Mercator-Projection](https://github.com/jorelosorio/web-mercator-projection) A project to easily use and convert LonLat, Point and Tile to display info, markers, etc, in a map using the Web Mercator Projection.
- [WGS84(stars: 107)](https://github.com/wroge/wgs84) - Library for Coordinate Conversion and Transformation (ETRS89, OSGB36, NAD83, RGF93, Web Mercator, UTM).

**[⬆ back to top](#contents)**

## Go Compilers

_Tools for compiling Go to other languages._

- [c2go(stars: 295)](https://github.com/goplus/c2go) - Convert C code to Go code.
- [c4go(stars: 344)](https://github.com/Konstantin8105/c4go) - Transpile C code to Go code.
- [esp32(stars: 78)](https://github.com/andygeiss/esp32-transpiler) - Transpile Go into Arduino code.
- [f4go(stars: 40)](https://github.com/Konstantin8105/f4go) - Transpile FORTRAN 77 code to Go code.
- [gopherjs(stars: 12235)](https://github.com/gopherjs/gopherjs) - Compiler from Go to JavaScript.
- [tardisgo(stars: 428)](https://github.com/tardisgo/tardisgo) - Golang to Haxe to CPP/CSharp/Java/JavaScript transpiler.

**[⬆ back to top](#contents)**

## Goroutines

_Tools for managing and working with Goroutines._

- [ants(stars: 11530)](https://github.com/panjf2000/ants) - A high-performance and low-cost goroutine pool in Go.
- [artifex(stars: 178)](https://github.com/borderstech/artifex) - Simple in-memory job queue for Golang using worker-based dispatching.
- [async(stars: 162)](https://github.com/reugn/async) - An alternative sync library for Go (Future, Promise, Locks).
- [async(stars: 129)](https://github.com/studiosol/async) - A safe way to execute functions asynchronously, recovering them in case of panic.
- [async-job(stars: 8)](https://github.com/lab210-dev/async-job) - AsyncJob is an asynchronous queue job manager with light code, clear and speed.
- [breaker(stars: 16)](https://github.com/kamilsk/breaker) - Flexible mechanism to make execution flow interruptible.
- [channelify(stars: 27)](https://github.com/ddelizia/channelify) - Transform your function to return channels for easy and powerful parallel processing.
- [conc(stars: 7874)](https://github.com/sourcegraph/conc) - `conc` is your toolbelt for structured concurrency in go, making common tasks easier and safer.
- [concurrency-limiter(stars: 17)](https://github.com/vivek-ng/concurrency-limiter) - Concurrency limiter with support for timeouts , dynamic priority and context cancellation of goroutines.
- [conexec(stars: 14)](https://github.com/ITcathyh/conexec) - A concurrent toolkit to help execute funcs concurrently in an efficient and safe way. It supports specifying the overall timeout to avoid blocking and uses goroutine pool to improve efficiency.
- [cyclicbarrier(stars: 134)](https://github.com/marusama/cyclicbarrier) - CyclicBarrier for golang.
- [execpool(stars: 16)](https://github.com/hexdigest/execpool) - A pool built around exec.Cmd that spins up a given number of processes in advance and attaches stdin and stdout to them when needed. Very similar to FastCGI or Apache Prefork MPM but works for any command.
- [flowmatic(stars: 215)](https://github.com/carlmjohnson/flowmatic) - Structured concurrency made easy.
- [go-actor(stars: 90)](https://github.com/vladopajic/go-actor) - A tiny library for writing concurrent programs using actor model.
- [go-floc(stars: 264)](https://github.com/workanator/go-floc) - Orchestrate goroutines with ease.
- [go-flow(stars: 218)](https://github.com/kamildrazkiewicz/go-flow) - Control goroutines execution order.
- [go-tools/multithreading(stars: 13)](https://github.com/nikhilsaraf/go-tools) - Manage a pool of goroutines using this lightweight library with a simple API.
- [go-trylock(stars: 33)](https://github.com/subchen/go-trylock) - TryLock support on read-write lock for Golang.
- [go-waitgroup(stars: 44)](https://github.com/pieterclaerhout/go-waitgroup) - Like `sync.WaitGroup` with error handling and concurrency control.
- [go-workerpool(stars: 10)](https://github.com/zenthangplus/go-workerpool) - Inspired from Java Thread Pool, Go WorkerPool aims to control heavy Go Routines.
- [go-workers(stars: 161)](https://github.com/catmullet/go-workers) - Easily and safely run workers for large data processing pipelines.
- [goccm(stars: 66)](https://github.com/zenthangplus/goccm) - Go Concurrency Manager package limits the number of goroutines that allowed to run concurrently.
- [gohive(stars: 50)](https://github.com/loveleshsharma/gohive) - A highly performant and easy to use Goroutine pool for Go.
- [gollback(stars: 116)](https://github.com/vardius/gollback) - asynchronous simple function utilities, for managing execution of closures and callbacks.
- [gowl(stars: 59)](https://github.com/hamed-yousefi/gowl) - Gowl is a process management and process monitoring tool at once. An infinite worker pool gives you the ability to control the pool and processes and monitor their status.
- [goworker(stars: 2762)](https://github.com/benmanns/goworker) - goworker is a Go-based background worker.
- [gowp(stars: 487)](https://github.com/xxjwxc/gowp) - gowp is concurrency limiting goroutine pool.
- [gpool(stars: 88)](https://github.com/Sherifabdlnaby/gpool) - manages a resizeable pool of context-aware goroutines to bound concurrency.
- [grpool(stars: 735)](https://github.com/ivpusic/grpool) - Lightweight Goroutine pool.
- [hands(stars: 10)](https://github.com/duanckham/hands) - A process controller used to control the execution and return strategies of multiple goroutines.
- [Hunch(stars: 97)](https://github.com/AaronJan/Hunch) - Hunch provides functions like: `All`, `First`, `Retry`, `Waterfall` etc., that makes asynchronous flow control more intuitive.
- [kyoo(stars: 45)](https://github.com/dirkaholic/kyoo) - Provides an unlimited job queue and concurrent worker pools.
- [neilotoole/errgroup(stars: 156)](https://github.com/neilotoole/errgroup) - Drop-in alternative to `sync/errgroup`, limited to a pool of N worker goroutines.
- [nursery(stars: 62)](https://github.com/arunsworld/nursery) - Structured concurrency in Go.
- [oversight](https://pkg.go.dev/cirello.io/oversight) - Oversight is a complete implementation of the Erlang supervision trees.
- [parallel-fn(stars: 35)](https://github.com/rafaeljesus/parallel-fn) - Run functions in parallel.
- [pond(stars: 1080)](https://github.com/alitto/pond) - Minimalistic and High-performance goroutine worker pool written in Go.
- [pool(stars: 720)](https://github.com/go-playground/pool) - Limited consumer goroutine or unlimited goroutine pool for easier goroutine handling and cancellation.
- [queue(stars: 15)](https://github.com/AnikHasibul/queue) - Gives you a `sync.WaitGroup` like queue group accessibility. Helps you to throttle and limit goroutines, wait for the end of the all goroutines and much more.
- [routine(stars: 172)](https://github.com/timandy/routine) - `routine` is a `ThreadLocal` for go library. It encapsulates and provides some easy-to-use, non-competitive, high-performance `goroutine` context access interfaces, which can help you access coroutine context information more gracefully.
- [routine(stars: 61)](https://github.com/x-mod/routine) - go routine control with context, support: Main, Go, Pool and some useful Executors.
- [semaphore(stars: 95)](https://github.com/kamilsk/semaphore) - Semaphore pattern implementation with timeout of lock/unlock operations based on channel and context.
- [semaphore(stars: 165)](https://github.com/marusama/semaphore) - Fast resizable semaphore implementation based on CAS (faster than channel-based semaphore implementations).
- [stl(stars: 27)](https://github.com/ssgreg/stl) - Software transactional locks based on Software Transactional Memory (STM) concurrency control mechanism.
- [threadpool(stars: 96)](https://github.com/shettyh/threadpool) - Golang threadpool implementation.
- [tunny(stars: 3735)](https://github.com/Jeffail/tunny) - Goroutine pool for golang.
- [worker-pool(stars: 88)](https://github.com/vardius/worker-pool) - goworker is a Go simple async worker pool.
- [workerpool(stars: 1200)](https://github.com/gammazero/workerpool) - Goroutine pool that limits the concurrency of task execution, not the number of tasks queued.

**[⬆ back to top](#contents)**

## GUI

_Libraries for building GUI Applications._

_Toolkits_

- [app(stars: 7472)](https://github.com/murlokswarm/app) - Package to create apps with GO, HTML and CSS. Supports: MacOS, Windows in progress.
- [energy(stars: 196)](https://github.com/energye/energy) - Cross-platform based on LCL(Native System UI Control Library) and CEF(Chromium Embedded Framework) (Windows/ macOS / Linux)
- [fyne(stars: 22149)](https://github.com/fyne-io/fyne) - Cross platform native GUIs designed for Go based on Material Design. Supports: Linux, macOS, Windows, BSD, iOS and Android.
- [gio](https://gioui.org) - Gio is a library for writing cross-platform immediate mode GUI-s in Go. Gio supports all the major platforms: Linux, macOS, Windows, Android, iOS, FreeBSD, OpenBSD and WebAssembly.
- [go-astilectron(stars: 4880)](https://github.com/asticode/go-astilectron) - Build cross platform GUI apps with GO and HTML/JS/CSS (powered by Electron).
- [go-gtk](https://mattn.github.io/go-gtk/) - Go bindings for GTK.
- [go-sciter(stars: 2550)](https://github.com/sciter-sdk/go-sciter) - Go bindings for Sciter: the Embeddable HTML/CSS/script engine for modern desktop UI development. Cross platform.
- [Goey](https://bitbucket.org/rj/goey/src/master/) - Cross platform UI toolkit aggregator for Windows / Linux / Mac. GTK, Cocoa, Windows API
- [goradd/html5tag(stars: 6)](https://github.com/goradd/html5tag) - Library for outputting HTML5 tags.
- [gotk3(stars: 1978)](https://github.com/gotk3/gotk3) - Go bindings for GTK3.
- [gowd(stars: 417)](https://github.com/dtylman/gowd) - Rapid and simple desktop UI development with GO, HTML, CSS and NW.js. Cross platform.
- [qt(stars: 10074)](https://github.com/therecipe/qt) - Qt binding for Go (support for Windows / macOS / Linux / Android / iOS / Sailfish OS / Raspberry Pi).
- [ui(stars: 8312)](https://github.com/andlabs/ui) - Platform-native GUI library for Go. Cross platform.
- [unison(stars: 116)](https://github.com/richardwilkes/unison) - A unified graphical user experience toolkit for Go desktop applications. macOS, Windows, and Linux are supported.
- [Wails](https://wails.io) - Mac, Windows, Linux desktop apps with HTML UI using built-in OS HTML renderer.
- [walk(stars: 6635)](https://github.com/lxn/walk) - Windows application library kit for Go.
- [webview(stars: 11725)](https://github.com/zserge/webview) - Cross-platform webview window with simple two-way JavaScript bindings (Windows / macOS / Linux).

_Interaction_

- [AppIndicator Go(stars: 4)](https://github.com/gopherlibs/appindicator) - Go bindings for libappindicator3 C library.
- [gosx-notifier(stars: 581)](https://github.com/deckarep/gosx-notifier) - OSX Desktop Notifications library for Go.
- [mac-activity-tracker(stars: 26)](https://github.com/prashantgupta24/activity-tracker) - OSX library to notify about any (pluggable) activity on your machine.
- [mac-sleep-notifier(stars: 30)](https://github.com/prashantgupta24/mac-sleep-notifier) - OSX Sleep/Wake notifications in golang.
- [robotgo(stars: 8922)](https://github.com/go-vgo/robotgo) - Go Native cross-platform GUI system automation. Control the mouse, keyboard and other.
- [systray(stars: 2994)](https://github.com/getlantern/systray) - Cross platform Go library to place an icon and menu in the notification area.
- [trayhost(stars: 246)](https://github.com/shurcooL/trayhost) - Cross-platform Go library to place an icon in the host operating system's taskbar.
- [zenity(stars: 582)](https://github.com/ncruces/zenity) - Cross-platform Go library and CLI to create simple dialogs that interact graphically with the user.

**[⬆ back to top](#contents)**

## Hardware

_Libraries, tools, and tutorials for interacting with hardware._

- [arduino-cli(stars: 4077)](https://github.com/arduino/arduino-cli) - Official Arduino CLI and library. Can run standalone, or be incorporated into larger Go projects.
- [emgo(stars: 1039)](https://github.com/ziutek/emgo) - Go-like language for programming embedded systems (e.g. STM32 MCU).
- [ghw(stars: 1533)](https://github.com/jaypipes/ghw) - Golang hardware discovery/inspection library.
- [go-osc(stars: 179)](https://github.com/hypebeast/go-osc) - Open Sound Control (OSC) bindings for Go.
- [go-rpio(stars: 2091)](https://github.com/stianeikeland/go-rpio) - GPIO for Go, doesn't require cgo.
- [goroslib(stars: 286)](https://github.com/aler9/goroslib) - Robot Operating System (ROS) library for Go.
- [joystick(stars: 51)](https://github.com/0xcafed00d/joystick) - a polled API to read the state of an attached joystick.
- [sysinfo(stars: 469)](https://github.com/zcalusic/sysinfo) - A pure Go library providing Linux OS / kernel / hardware system information.

**[⬆ back to top](#contents)**

## Images

_Libraries for manipulating images._

- [bild(stars: 3874)](https://github.com/anthonynsimon/bild) - Collection of image processing algorithms in pure Go.
- [bimg(stars: 2448)](https://github.com/h2non/bimg) - Small package for fast and efficient image processing using libvips.
- [cameron(stars: 109)](https://github.com/aofei/cameron) - An avatar generator for Go.
- [canvas(stars: 1372)](https://github.com/tdewolff/canvas) - Vector graphics to PDF, SVG or rasterized image.
- [color-extractor(stars: 73)](https://github.com/marekm4/color-extractor) - Dominant color extractor with no external dependencies.
- [darkroom(stars: 224)](https://github.com/gojek/darkroom) - An image proxy with changeable storage backends and image processing engines with focus on speed and resiliency.
- [draft(stars: 572)](https://github.com/lucasepe/draft) - Generate High Level Microservice Architecture diagrams for GraphViz using simple YAML syntax.
- [geopattern(stars: 1259)](https://github.com/pravj/geopattern) - Create beautiful generative image patterns from a string.
- [gg(stars: 4176)](https://github.com/fogleman/gg) - 2D rendering in pure Go.
- [gift(stars: 1699)](https://github.com/disintegration/gift) - Package of image processing filters.
- [gltf(stars: 221)](https://github.com/qmuntal/gltf) - Efficient and robust glTF 2.0 reader, writer and validator.
- [go-cairo(stars: 136)](https://github.com/ungerik/go-cairo) - Go binding for the cairo graphics library.
- [go-gd(stars: 59)](https://github.com/bolknote/go-gd) - Go binding for GD library.
- [go-nude(stars: 395)](https://github.com/koyachi/go-nude) - Nudity detection with Go.
- [go-webcolors(stars: 27)](https://github.com/jyotiska/go-webcolors) - Port of webcolors library from Python to Go.
- [go-webp(stars: 173)](https://github.com/kolesa-team/go-webp) - Library for encode and decode webp pictures, using libwebp.
- [gocv(stars: 5945)](https://github.com/hybridgroup/gocv) - Go package for computer vision using OpenCV 3.3+.
- [goimagehash(stars: 668)](https://github.com/corona10/goimagehash) - Go Perceptual image hashing package.
- [goimghdr(stars: 40)](https://github.com/corona10/goimghdr) - The imghdr module determines the type of image contained in a file for Go.
- [govatar(stars: 573)](https://github.com/o1egl/govatar) - Library and CMD tool for generating funny avatars.
- [govips(stars: 1056)](https://github.com/davidbyttow/govips) - A lightning fast image processing and resizing library for Go.
- [gowitness(stars: 2486)](https://github.com/sensepost/gowitness) - Screenshoting webpages using go and headless chrome on command line.
- [gridder(stars: 69)](https://github.com/shomali11/gridder) - A Grid based 2D Graphics library.
- [image2ascii(stars: 815)](https://github.com/qeesung/image2ascii) - Convert image to ASCII.
- [imagick(stars: 1651)](https://github.com/gographics/imagick) - Go binding to ImageMagick's MagickWand C API.
- [imaginary(stars: 5200)](https://github.com/h2non/imaginary) - Fast and simple HTTP microservice for image resizing.
- [imaging(stars: 4949)](https://github.com/disintegration/imaging) - Simple Go image processing package.
- [img(stars: 151)](https://github.com/hawx/img) - Selection of image manipulation tools.
- [ln(stars: 3216)](https://github.com/fogleman/ln) - 3D line art rendering in Go.
- [mergi(stars: 214)](https://github.com/noelyahan/mergi) - Tool & Go library for image manipulation (Merge, Crop, Resize, Watermark, Animate).
- [mort(stars: 493)](https://github.com/aldor007/mort) - Storage and image processing server written in Go.
- [mpo(stars: 12)](https://github.com/donatj/mpo) - Decoder and conversion tool for MPO 3D Photos.
- [picfit(stars: 1940)](https://github.com/thoas/picfit) - An image resizing server written in Go.
- [pt(stars: 2059)](https://github.com/fogleman/pt) - Path tracing engine written in Go.
- [rez(stars: 209)](https://github.com/bamiaux/rez) - Image resizing in pure Go and SIMD.
- [scout(stars: 12)](https://github.com/jonoton/scout) - Scout is a standalone open source software solution for DIY video security.
- [smartcrop(stars: 1777)](https://github.com/muesli/smartcrop) - Finds good crops for arbitrary images and crop sizes.
- [steganography(stars: 277)](https://github.com/auyer/steganography) - Pure Go Library for LSB steganography.
- [stegify(stars: 1146)](https://github.com/DimitarPetrov/stegify) - Go tool for LSB steganography, capable of hiding any file within an image.
- [svgo(stars: 2066)](https://github.com/ajstarks/svgo) - Go Language Library for SVG generation.
- [tga(stars: 34)](https://github.com/ftrvxmtrx/tga) - Package tga is a TARGA image format decoder/encoder.
- [webp-server(stars: 66)](https://github.com/mehdipourfar/webp-server) - Simple and minimal image server capable of storing, resizing, converting and caching images.

**[⬆ back to top](#contents)**

## IoT (Internet of Things)

_Libraries for programming devices of the IoT._

- [connectordb(stars: 382)](https://github.com/connectordb/connectordb) - Open-Source Platform for Quantified Self & IoT.
- [devices(stars: 263)](https://github.com/goiot/devices) - Suite of libraries for IoT devices, experimental for x/exp/io.
- [ekuiper(stars: 1259)](https://github.com/lf-edge/ekuiper) - Lightweight data stream processing engine for IoT edge.
- [eywa(stars: 61)](https://github.com/xcodersun/eywa) - Project Eywa is essentially a connection manager that keeps track of connected devices.
- [flogo(stars: 2344)](https://github.com/tibcosoftware/flogo) - Project Flogo is an Open Source Framework for IoT Edge Apps & Integration.
- [gatt(stars: 1096)](https://github.com/paypal/gatt) - Gatt is a Go package for building Bluetooth Low Energy peripherals.
- [gobot](https://github.com/hybridgroup/gobot/) - Gobot is a framework for robotics, physical computing, and the Internet of Things.
- [huego(stars: 243)](https://github.com/amimof/huego) - An extensive Philips Hue client library for Go.
- [iot](https://github.com/vaelen/iot/) - IoT is a simple framework for implementing a Google IoT Core device.
- [mainflux(stars: 3)](https://github.com/Mainflux/mainflux) - Industrial IoT Messaging and Device Management Server.
- [periph](https://periph.io/) - Peripherals I/O to interface with low-level board facilities.
- [sensorbee(stars: 227)](https://github.com/sensorbee/sensorbee) - Lightweight stream processing engine for IoT.

**[⬆ back to top](#contents)**

## Job Scheduler

_Libraries for scheduling jobs._

- [cdule(stars: 43)](https://github.com/deepaksinghvi/cdule) - Job scheduler library with database support
- [cheek(stars: 110)](https://github.com/datarootsio/cheek) - A simple crontab like scheduler that aims to offer a KISS approach to job scheduling.
- [clockwerk(stars: 139)](https://github.com/onatm/clockwerk) - Go package to schedule periodic jobs using a simple, fluent syntax.
- [cronticker(stars: 11)](https://github.com/krayzpipes/cronticker) - A ticker implementation to support cron schedules.
- [Dagu(stars: 997)](https://github.com/dagu-go/dagu) - No-code workflow executor. it executes DAGs defined in a simple YAML format.
- [go-cron(stars: 221)](https://github.com/rk/go-cron) - Simple Cron library for go that can execute closures or functions at varying intervals, from once a second to once a year on a specific date and time. Primarily for web applications and long running daemons.
- [go-quartz(stars: 1564)](https://github.com/reugn/go-quartz) - Simple, zero-dependency scheduling library for Go.
- [gocron(stars: 4476)](https://github.com/go-co-op/gocron) - Easy and fluent Go job scheduling. This is an actively maintained fork of [jasonlvhit/gocron](https://github.com/jasonlvhit/gocron).
- [goflow(stars: 221)](https://github.com/fieldryand/goflow) - A simple but powerful DAG scheduler and dashboard.
- [gron(stars: 998)](https://github.com/roylee0704/gron) - Define time-based tasks using a simple Go API and Gron’s scheduler will run them accordingly.
- [gronx(stars: 322)](https://github.com/adhocore/gronx) - Cron expression parser, task runner and daemon consuming crontab like task list.
- [JobRunner(stars: 994)](https://github.com/bamzi/jobrunner) - Smart and featureful cron job scheduler with job queuing and live monitoring built in.
- [jobs(stars: 494)](https://github.com/albrow/jobs) - Persistent and flexible background jobs library.
- [leprechaun(stars: 100)](https://github.com/kilgaloon/leprechaun) - Job scheduler that supports webhooks, crons and classic scheduling.
- [sched(stars: 27)](https://github.com/romshark/sched) - A job scheduler with the ability to fast-forward time.
- [scheduler(stars: 435)](https://github.com/carlescere/scheduler) - Cronjobs scheduling made easy.
- [tasks(stars: 217)](https://github.com/madflojo/tasks) - An easy to use in-process scheduler for recurring tasks in Go.

**[⬆ back to top](#contents)**

## JSON

_Libraries for working with JSON._

- [ajson(stars: 198)](https://github.com/spyzhov/ajson) - Abstract JSON for golang with JSONPath support.
- [ask(stars: 31)](https://github.com/simonnilsson/ask) - Easy access to nested values in maps and slices. Works in combination with encoding/json and other packages that "Unmarshal" arbitrary data into Go data-types.
- [dynjson(stars: 16)](https://github.com/cocoonspace/dynjson) - Client-customizable JSON formats for dynamic APIs.
- [ej(stars: 10)](https://github.com/lucassscaravelli/ej) - Write and read JSON from different sources succinctly.
- [epoch(stars: 13)](https://github.com/vtopc/epoch) - Contains primitives for marshaling/unmarshalling Unix timestamp/epoch to/from build-in time.Time type in JSON.
- [fastjson(stars: 2070)](https://github.com/valyala/fastjson) - Fast JSON parser and validator for Go. No custom structs, no code generation, no reflection.
- [gjo(stars: 122)](https://github.com/skanehira/gjo) - Small utility to create JSON objects.
- [GJSON(stars: 13203)](https://github.com/tidwall/gjson) - Get a JSON value with one line of code.
- [go-jsonerror(stars: 14)](https://github.com/ddymko/go-jsonerror) - Go-JsonError is meant to allow us to easily create json response errors that follow the JsonApi spec.
- [go-respond(stars: 52)](https://github.com/nicklaw5/go-respond) - Go package for handling common HTTP JSON responses.
- [gojmapr(stars: 21)](https://github.com/limiu82214/gojmapr) - Get simple struct from complex json by json path.
- [gojq(stars: 191)](https://github.com/elgs/gojq) - JSON query in Golang.
- [gojson(stars: 2631)](https://github.com/ChimeraCoder/gojson) - Automatically generate Go (golang) struct definitions from example JSON.
- [htmljson(stars: 5)](https://github.com/nikolaydubina/htmljson) - Rich rendering of JSON as HTML in Go.
- [JayDiff(stars: 103)](https://github.com/yazgazan/jaydiff) - JSON diff utility written in Go.
- [jettison(stars: 166)](https://github.com/wI2L/jettison) - Fast and flexible JSON encoder for Go.
- [jscan(stars: 74)](https://github.com/romshark/jscan) - High performance zero-allocation JSON iterator.
- [JSON-to-Go](https://mholt.github.io/json-to-go/) - Convert JSON to Go struct.
- [JSON-to-Proto](https://json-to-proto.github.io/) - Convert JSON to Protobuf online.
- [json2go(stars: 124)](https://github.com/m-zajac/json2go) - Advanced JSON to Go struct conversion. Provides package that can parse multiple JSON documents and create struct to fit them all.
- [jsonapi-errors(stars: 14)](https://github.com/AmuzaTkts/jsonapi-errors) - Go bindings based on the JSON API errors reference.
- [jsoncolor(stars: 38)](https://github.com/neilotoole/jsoncolor) - Drop-in replacement for `encoding/json` that outputs colorized JSON.
- [jsondiff(stars: 366)](https://github.com/wI2L/jsondiff) - JSON diff library for Go based on RFC6902 (JSON Patch).
- [jsonf(stars: 65)](https://github.com/miolini/jsonf) - Console tool for highlighted formatting and struct query fetching JSON.
- [jsongo(stars: 109)](https://github.com/ricardolonga/jsongo) - Fluent API to make it easier to create Json objects.
- [jsonhal(stars: 14)](https://github.com/RichardKnop/jsonhal) - Simple Go package to make custom structs marshal into HAL compatible JSON responses.
- [jsonhandlers(stars: 2)](https://github.com/abusomani/jsonhandlers) - JSON library to expose simple handlers that lets you easily read and write json from various sources.
- [jsonic(stars: 11)](https://github.com/sinhashubham95/jsonic) - Utilities to handle and query JSON without defining structs in a type safe manner.
- [jsonvalue](https://github.com/Andrew-M-C/go.jsonvalue) - A fast and convenient library for unstructured JSON data, replacing `encoding/json`.
- [jzon(stars: 12)](https://github.com/zerosnake0/jzon) - JSON library with standard compatible API/behavior.
- [kazaam(stars: 270)](https://github.com/Qntfy/kazaam) - API for arbitrary transformation of JSON documents.
- [mapslice-json(stars: 16)](https://github.com/mickep76/mapslice-json) - Go MapSlice for ordered marshal/ unmarshal of maps in JSON.
- [marshmallow(stars: 331)](https://github.com/PerimeterX/marshmallow) - Performant JSON unmarshalling for flexible use cases.
- [mp(stars: 47)](https://github.com/sanbornm/mp) - Simple cli email parser. It currently takes stdin and outputs JSON.
- [OjG(stars: 752)](https://github.com/ohler55/ojg) - Optimized JSON for Go is a high performance parser with a variety of additional JSON tools including JSONPath.
- [omg.jsonparser](https://github.com/dedalqq/omg.jsonparser) - Simple JSON parser with validation by condition via golang struct fields tags.
- [ujson(stars: 72)](https://github.com/olvrng/ujson) - Fast and minimal JSON parser and transformer that works on unstructured JSON.
- [vjson(stars: 38)](https://github.com/miladibra10/vjson) - Go package for validating JSON objects with declaring a JSON schema with fluent API.

**[⬆ back to top](#contents)**

## Logging

_Libraries for generating and working with log files._

- [distillog(stars: 32)](https://github.com/amoghe/distillog) - distilled levelled logging (think of it as stdlib + log levels).
- [glg(stars: 186)](https://github.com/kpango/glg) - glg is simple and fast leveled logging library for Go.
- [glo(stars: 15)](https://github.com/lajosbencz/glo) - PHP Monolog inspired logging facility with identical severity levels.
- [glog(stars: 3482)](https://github.com/golang/glog) - Leveled execution logs for Go.
- [go-cronowriter(stars: 55)](https://github.com/utahta/go-cronowriter) - Simple writer that rotate log files automatically based on current date and time, like cronolog.
- [go-log(stars: 10)](https://github.com/pieterclaerhout/go-log) - A logging library with stack traces, object dumping and optional timestamps.
- [go-log(stars: 14)](https://github.com/subchen/go-log) - Simple and configurable Logging in Go, with level, formatters and writers.
- [go-log(stars: 32)](https://github.com/siddontang/go-log) - Log lib supports level and multi handlers.
- [go-log(stars: 42)](https://github.com/ian-kent/go-log) - Log4j implementation in Go.
- [go-logger(stars: 285)](https://github.com/apsdehal/go-logger) - Simple logger of Go Programs, with level handlers.
- [gologger(stars: 40)](https://github.com/sadlil/gologger) - Simple easy to use log lib for go, logs in Colored Console, Simple Console, File or Elasticsearch.
- [gomol(stars: 18)](https://github.com/aphistic/gomol) - Multiple-output, structured logging for Go with extensible logging outputs.
- [gone/log](https://github.com/One-com/gone/tree/master/log) - Fast, extendable, full-featured, std-lib source compatible log library.
- [httpretty(stars: 381)](https://github.com/henvic/httpretty) - Pretty-prints your regular HTTP requests on your terminal for debugging (similar to http.DumpRequest).
- [journald(stars: 38)](https://github.com/ssgreg/journald) - Go implementation of systemd Journal's native API for logging.
- [kemba(stars: 11)](https://github.com/clok/kemba) - A tiny debug logging tool inspired by [debug](https://github.com/visionmedia/debug), great for CLI tools and applications.
- [log(stars: 10)](https://github.com/aerogo/log) - An O(1) logging system that allows you to connect one log to multiple writers (e.g. stdout, a file and a TCP connection).
- [log(stars: 1341)](https://github.com/apex/log) - Structured logging package for Go.
- [log(stars: 288)](https://github.com/go-playground/log) - Simple, configurable and scalable Structured Logging for Go.
- [log(stars: 25)](https://github.com/teris-io/log) - Structured log interface for Go cleanly separates logging facade from its implementation.
- [log(stars: 14)](https://github.com/heartwilltell/log) - Simple leveled logging wrapper around standard log package.
- [log(stars: 3)](https://github.com/no-src/log) - A simple logging framework out of the box.
- [log-voyage(stars: 93)](https://github.com/firstrow/logvoyage) - Full-featured logging saas written in golang.
- [log15(stars: 1097)](https://github.com/inconshreveable/log15) - Simple, powerful logging for Go.
- [logdump(stars: 11)](https://github.com/ewwwwwqm/logdump) - Package for multi-level logging.
- [logex(stars: 42)](https://github.com/chzyer/logex) - Golang log lib, supports tracking and level, wrap by standard log lib.
- [logger(stars: 157)](https://github.com/azer/logger) - Minimalistic logging library for Go.
- [logmatic(stars: 16)](https://github.com/borderstech/logmatic) - Colorized logger for Golang with dynamic log level configuration.
- [logo(stars: 11)](https://github.com/mbndr/logo) - Golang logger to different configurable writers.
- [logrus(stars: 23610)](https://github.com/Sirupsen/logrus) - Structured logger for Go.
- [logrusiowriter(stars: 16)](https://github.com/cabify/logrusiowriter) - `io.Writer` implementation using [logrus](https://github.com/sirupsen/logrus) logger.
- [logrusly(stars: 28)](https://github.com/sebest/logrusly) - [logrus](https://github.com/sirupsen/logrus) plug-in to send errors to a [Loggly](https://www.loggly.com/).
- [logur(stars: 198)](https://github.com/logur/logur) - An opinionated logger interface and collection of logging best practices with adapters and integrations for well-known libraries ([logrus](https://github.com/sirupsen/logrus), [go-kit log](https://github.com/go-kit/kit/tree/master/log), [zap](https://github.com/uber-go/zap), [zerolog](https://github.com/rs/zerolog), etc).
- [logutils(stars: 358)](https://github.com/hashicorp/logutils) - Utilities for slightly better logging in Go (Golang) extending the standard logger.
- [logxi(stars: 352)](https://github.com/mgutz/logxi) - 12-factor app logger that is fast and makes you happy.
- [lumberjack(stars: 4417)](https://github.com/natefinch/lumberjack) - Simple rolling logger, implements io.WriteCloser.
- [mlog(stars: 32)](https://github.com/jbrodriguez/mlog) - Simple logging module for go, with 5 levels, an optional rotating logfile feature and stdout/stderr output.
- [noodlog(stars: 44)](https://github.com/gyozatech/noodlog) - Parametrized JSON logging library which lets you obfuscate sensitive data and marshal any kind of content. No more printed pointers instead of values, nor escape chars for the JSON strings.
- [onelog(stars: 414)](https://github.com/francoispqt/onelog) - Onelog is a dead simple but very efficient JSON logger. It is the fastest JSON logger out there in all scenarios. Also, it is one of the logger with the lowest allocation.
- [ozzo-log(stars: 122)](https://github.com/go-ozzo/ozzo-log) - High performance logging supporting log severity, categorization, and filtering. Can send filtered log messages to various targets (e.g. console, network, mail).
- [phuslu/log(stars: 528)](https://github.com/phuslu/log) - High performance structured logging.
- [pp(stars: 1711)](https://github.com/k0kubun/pp) - Colored pretty printer for Go language.
- [rollingwriter(stars: 288)](https://github.com/arthurkiller/rollingWriter) - RollingWriter is an auto-rotate `io.Writer` implementation with multi policies to provide log file rotation.
- [seelog(stars: 1640)](https://github.com/cihub/seelog) - Logging functionality with flexible dispatching, filtering, and formatting.
- [slf4g(stars: 3)](https://github.com/echocat/slf4g) - Simple Logging Facade for Golang: Simple structured logging; but powerful, extendable and customizable, with huge amount of learnings from decades of past logging frameworks.
- [slog(stars: 311)](https://github.com/gookit/slog) - Lightweight, configurable, extensible logger for Go.
- [slog-formatter(stars: 59)](https://github.com/samber/slog-formatter) - Common formatters for slog and helpers to build your own.
- [slog-multi(stars: 192)](https://github.com/samber/slog-multi) - Chain of slog.Handler (pipeline, fanout...).
- [spew(stars: 5802)](https://github.com/davecgh/go-spew) - Implements a deep pretty printer for Go data structures to aid in debugging.
- [sqldb-logger(stars: 330)](https://github.com/simukti/sqldb-logger) - A logger for Go SQL database driver without modify existing \*sql.DB stdlib usage.
- [stdlog(stars: 46)](https://github.com/alexcesaro/log) - Stdlog is an object-oriented library providing leveled logging. It is very useful for cron jobs.
- [structy/log(stars: 5)](https://github.com/structy/log) - A simple to use log system, minimalist but with features for debugging and differentiation of messages.
- [tail(stars: 2639)](https://github.com/hpcloud/tail) - Go package striving to emulate the features of the BSD tail program.
- [tint(stars: 424)](https://github.com/lmittmann/tint) - A slog.Handler that writes tinted logs.
- [xlog(stars: 8)](https://github.com/xfxdev/xlog) - Plugin architecture and flexible log system for Go, with level ctrl, multiple log target and custom log format.
- [xlog(stars: 137)](https://github.com/rs/xlog) - Structured logger for `net/context` aware HTTP handlers with flexible dispatching.
- [xylog(stars: 15)](https://github.com/xybor-x/xylog) - Leveled and structured logging, dynamic fields, high performance, zone management, simple configuration, and readable syntax.
- [yell(stars: 1)](https://github.com/jfcg/yell) - Yet another minimalistic logging library.
- [zap(stars: 20244)](https://github.com/uber-go/zap) - Fast, structured, leveled logging in Go.
- [zax(stars: 12)](https://github.com/yuseferi/zax) - Integrate Context with Zap logger, which leads to more flexibility in Go logging.
- [zerolog(stars: 9198)](https://github.com/rs/zerolog) - Zero-allocation JSON logger.
- [zkits-logger(stars: 25)](https://github.com/edoger/zkits-logger) - A powerful zero-dependency JSON logger.
- [zl(stars: 4)](https://github.com/nkmr-jp/zl) - High Developer Experience, zap based logger. It offers rich functionality but is easy to configure.

**[⬆ back to top](#contents)**

## Machine Learning

_Libraries for Machine Learning._

- [bayesian(stars: 778)](https://github.com/jbrukh/bayesian) - Naive Bayesian Classification for Golang.
- [CloudForest(stars: 734)](https://github.com/ryanbressler/CloudForest) - Fast, flexible, multi-threaded ensembles of decision trees for machine learning in pure Go.
- [ddt(stars: 33)](https://github.com/sgrodriguez/ddt) - Dynamic decision tree, create trees defining customizable rules.
- [eaopt(stars: 864)](https://github.com/MaxHalford/eaopt) - An evolutionary optimization library.
- [evoli(stars: 28)](https://github.com/khezen/evoli) - Genetic Algorithm and Particle Swarm Optimization library.
- [fonet(stars: 79)](https://github.com/Fontinalis/fonet) - A Deep Neural Network library written in Go.
- [go-cluster(stars: 40)](https://github.com/e-XpertSolutions/go-cluster) - Go implementation of the k-modes and k-prototypes clustering algorithms.
- [go-deep(stars: 489)](https://github.com/patrikeh/go-deep) - A feature-rich neural network library in Go.
- [go-fann(stars: 115)](https://github.com/white-pony/go-fann) - Go bindings for Fast Artificial Neural Networks(FANN) library.
- [go-featureprocessing(stars: 105)](https://github.com/nikolaydubina/go-featureprocessing) - Fast and convenient feature processing for low latency machine learning in Go.
- [go-galib(stars: 195)](https://github.com/thoj/go-galib) - Genetic Algorithms library written in Go / golang.
- [go-pr(stars: 62)](https://github.com/daviddengcn/go-pr) - Pattern recognition package in Go lang.
- [gobrain(stars: 548)](https://github.com/goml/gobrain) - Neural Networks written in go.
- [godist(stars: 36)](https://github.com/e-dard/godist) - Various probability distributions, and associated methods.
- [goga(stars: 205)](https://github.com/tomcraven/goga) - Genetic algorithm library for Go.
- [GoLearn(stars: 9080)](https://github.com/sjwhitworth/golearn) - General Machine Learning library for Go.
- [golinear(stars: 44)](https://github.com/danieldk/golinear) - liblinear bindings for Go.
- [GoMind(stars: 70)](https://github.com/surenderthakran/gomind) - A simplistic Neural Network Library in Go.
- [goml(stars: 1521)](https://github.com/cdipaolo/goml) - On-line Machine Learning in Go.
- [gonet(stars: 80)](https://github.com/dathoangnd/gonet) - Neural Network for Go.
- [Goptuna(stars: 244)](https://github.com/c-bata/goptuna) - Bayesian optimization framework for black-box functions written in Go. Everything will be optimized.
- [goRecommend(stars: 196)](https://github.com/timkaye11/goRecommend) - Recommendation Algorithms library written in Go.
- [gorgonia(stars: 5205)](https://github.com/gorgonia/gorgonia) - graph-based computational library like Theano for Go that provides primitives for building various machine learning and neural network algorithms.
- [gorse(stars: 7799)](https://github.com/zhenghaoz/gorse) - An offline recommender system backend based on collaborative filtering written in Go.
- [goscore(stars: 93)](https://github.com/asafschers/goscore) - Go Scoring API for PMML.
- [gosseract(stars: 2348)](https://github.com/otiai10/gosseract) - Go package for OCR (Optical Character Recognition), by using Tesseract C++ library.
- [libsvm(stars: 73)](https://github.com/datastream/libsvm) - libsvm golang version derived work based on LIBSVM 3.14.
- [m2cgen(stars: 2670)](https://github.com/BayesWitnesses/m2cgen) - A CLI tool to transpile trained classic ML models into a native Go code with zero dependencies, written in Python with Go language support.
- [neat(stars: 69)](https://github.com/jinyeom/neat) - Plug-and-play, parallel Go framework for NeuroEvolution of Augmenting Topologies (NEAT).
- [neural-go(stars: 67)](https://github.com/schuyler/neural-go) - Multilayer perceptron network implemented in Go, with training via backpropagation.
- [ocrserver(stars: 627)](https://github.com/otiai10/ocrserver) - A simple OCR API server, seriously easy to be deployed by Docker and Heroku.
- [onnx-go(stars: 575)](https://github.com/owulveryck/onnx-go) - Go Interface to Open Neural Network Exchange (ONNX).
- [probab(stars: 19)](https://github.com/ThePaw/probab) - Probability distribution functions. Bayesian inference. Written in pure Go.
- [randomforest(stars: 38)](https://github.com/malaschitz/randomForest) - Easy to use Random Forest library for Go.
- [regommend(stars: 305)](https://github.com/muesli/regommend) - Recommendation & collaborative filtering engine.
- [shield(stars: 155)](https://github.com/eaigner/shield) - Bayesian text classifier with flexible tokenizers and storage backends for Go.
- [tfgo(stars: 2329)](https://github.com/galeone/tfgo) - Easy to use Tensorflow bindings: simplifies the usage of the official Tensorflow Go bindings. Define computational graphs in Go, load and execute models trained in Python.
- [Varis(stars: 53)](https://github.com/Xamber/Varis) - Golang Neural Network.

**[⬆ back to top](#contents)**

## Messaging

_Libraries that implement messaging systems._

- [ami(stars: 27)](https://github.com/kak-tus/ami) - Go client to reliable queues based on Redis Cluster Streams.
- [amqp(stars: 1132)](https://github.com/rabbitmq/amqp091-go) - Go RabbitMQ Client Library.
- [APNs2(stars: 2893)](https://github.com/sideshow/apns2) - HTTP/2 Apple Push Notification provider for Go — Send push notifications to iOS, tvOS, Safari and OSX apps.
- [Asynq(stars: 7900)](https://github.com/hibiken/asynq) - A simple, reliable, and efficient distributed task queue for Go built on top of Redis.
- [Beaver(stars: 1489)](https://github.com/Clivern/Beaver) - A real time messaging server to build a scalable in-app notifications, multiplayer games, chat apps in web and mobile apps.
- [Benthos(stars: 6938)](https://github.com/Jeffail/benthos) - A message streaming bridge between a range of protocols.
- [Bus(stars: 320)](https://github.com/mustafaturan/bus) - Minimalist message bus implementation for internal communication.
- [Centrifugo(stars: 7581)](https://github.com/centrifugal/centrifugo) - Real-time messaging (Websockets or SockJS) server in Go.
- [Chanify(stars: 1188)](https://github.com/chanify/chanify) - A push notification server send message to your iOS devices.
- [Commander(stars: 65)](https://github.com/jeroenrinzema/commander) - A high-level event driven consumer/producer supporting various "dialects" such as Apache Kafka.
- [Confluent Kafka Golang Client(stars: 4266)](https://github.com/confluentinc/confluent-kafka-go) - confluent-kafka-go is Confluent's Golang client for Apache Kafka and the Confluent Platform.
- [dbus(stars: 906)](https://github.com/godbus/dbus) - Native Go bindings for D-Bus.
- [drone-line(stars: 79)](https://github.com/appleboy/drone-line) - Sending [Line](https://at.line.me/en) notifications using a binary, docker or Drone CI.
- [emitter(stars: 489)](https://github.com/olebedev/emitter) - Emits events using Go way, with wildcard, predicates, cancellation possibilities and many other good wins.
- [event(stars: 54)](https://github.com/agoalofalife/event) - Implementation of the pattern observer.
- [EventBus(stars: 1574)](https://github.com/asaskevich/EventBus) - The lightweight event bus with async compatibility.
- [gaurun-client(stars: 11)](https://github.com/osamingo/gaurun-client) - Gaurun Client written in Go.
- [Glue(stars: 410)](https://github.com/desertbit/glue) - Robust Go and Javascript Socket Library (Alternative to Socket.io).
- [go-eventbus(stars: 1)](https://github.com/stanipetrosyan/go-eventbus) - Simple Event Bus package for Go.
- [go-mq(stars: 88)](https://github.com/cheshir/go-mq) - RabbitMQ client with declarative configuration.
- [go-notify(stars: 67)](https://github.com/TheCreeper/go-notify) - Native implementation of the freedesktop notification spec.
- [go-nsq(stars: 2495)](https://github.com/nsqio/go-nsq) - the official Go package for NSQ.
- [go-res(stars: 62)](https://github.com/jirenius/go-res) - Package for building REST/real-time services where clients are synchronized seamlessly, using NATS and Resgate.
- [go-socket.io](https://github.com/googollee/go-socket.io) - socket.io library for golang, a realtime application framework.
- [go-vitotrol(stars: 23)](https://github.com/maxatome/go-vitotrol) - Client library to Viessmann Vitotrol web service.
- [Gollum(stars: 935)](https://github.com/trivago/gollum) - A n:m multiplexer that gathers messages from different sources and broadcasts them to a set of destinations.
- [golongpoll(stars: 641)](https://github.com/jcuga/golongpoll) - HTTP longpoll server library that makes web pub-sub simple.
- [gopush-cluster(stars: 2079)](https://github.com/Terry-Mao/gopush-cluster) - gopush-cluster is a go push server cluster.
- [gorush(stars: 7418)](https://github.com/appleboy/gorush) - Push notification server using [APNs2](https://github.com/sideshow/apns2) and google [GCM](https://github.com/google/go-gcm).
- [gosd(stars: 23)](https://github.com/alexsniffin/gosd) - A library for scheduling when to dispatch a message to a channel.
- [guble(stars: 155)](https://github.com/smancke/guble) - Messaging server using push notifications (Google Firebase Cloud Messaging, Apple Push Notification services, SMS) as well as websockets, a REST API, featuring distributed operation and message-persistence.
- [hare(stars: 51)](https://github.com/leozz37/hare) - A user friendly library for sending messages and listening to TCP sockets.
- [hub(stars: 139)](https://github.com/leandro-lugaresi/hub) - A Message/Event Hub for Go applications, using publish/subscribe pattern with support for alias like rabbitMQ exchanges.
- [jazz(stars: 18)](https://github.com/socifi/jazz) - A simple RabbitMQ abstraction layer for queue administration and publishing and consuming of messages.
- [machinery(stars: 7130)](https://github.com/RichardKnop/machinery) - Asynchronous task queue/job queue based on distributed message passing.
- [mangos(stars: 617)](https://github.com/nanomsg/mangos) - Pure go implementation of the Nanomsg ("Scalability Protocols") with transport interoperability.
- [melody(stars: 3383)](https://github.com/olahol/melody) - Minimalist framework for dealing with websocket sessions, includes broadcasting and automatic ping/pong handling.
- [Mercure(stars: 3606)](https://github.com/dunglas/mercure) - Server and library to dispatch server-sent updates using the Mercure protocol (built on top of Server-Sent Events).
- [messagebus(stars: 260)](https://github.com/vardius/message-bus) - messagebus is a Go simple async message bus, perfect for using as event bus when doing event sourcing, CQRS, DDD.
- [NATS Go Client(stars: 4930)](https://github.com/nats-io/nats) - Lightweight and high performance publish-subscribe and distributed queueing messaging system - this is the Go library.
- [nsq-event-bus(stars: 77)](https://github.com/rafaeljesus/nsq-event-bus) - A tiny wrapper around NSQ topic and channel.
- [oplog(stars: 113)](https://github.com/dailymotion/oplog) - Generic oplog/replication system for REST APIs.
- [pubsub(stars: 393)](https://github.com/tuxychandru/pubsub) - Simple pubsub package for go.
- [Quamina(stars: 353)](https://github.com/timbray/quamina) - Fast pattern-matching for filtering messages and events.
- [rabbus(stars: 97)](https://github.com/rafaeljesus/rabbus) - A tiny wrapper over amqp exchanges and queues.
- [rabtap(stars: 247)](https://github.com/jandelgado/rabtap) - RabbitMQ swiss army knife cli app.
- [RapidMQ(stars: 67)](https://github.com/sybrexsys/RapidMQ) - RapidMQ is a lightweight and reliable library for managing of the local messages queue.
- [Ratus(stars: 91)](https://github.com/hyperonym/ratus) - Ratus is a RESTful asynchronous task queue server.
- [redisqueue(stars: 115)](https://github.com/robinjoseph08/redisqueue) - redisqueue provides a producer and consumer of a queue that uses Redis streams.
- [rmqconn(stars: 21)](https://github.com/sbabiv/rmqconn) - RabbitMQ Reconnection. Wrapper over amqp.Connection and amqp.Dial. Allowing to do a reconnection when the connection is broken before forcing the call to the Close () method to be closed.
- [sarama(stars: 10655)](https://github.com/Shopify/sarama) - Go library for Apache Kafka.
- [Uniqush-Push(stars: 1519)](https://github.com/uniqush/uniqush-push) - Redis backed unified push service for server-side notifications to mobile devices.
- [Watermill(stars: 6421)](https://github.com/ThreeDotsLabs/watermill) - Working efficiently with message streams. Building event driven applications, enabling event sourcing, RPC over messages, sagas. Can use conventional pub/sub implementations like Kafka or RabbitMQ, but also HTTP or MySQL binlog.
- [zmq4(stars: 1104)](https://github.com/pebbe/zmq4) - Go interface to ZeroMQ version 4. Also available for [version 3](https://github.com/pebbe/zmq3) and [version 2](https://github.com/pebbe/zmq2).

**[⬆ back to top](#contents)**

## Microsoft Office

- [unioffice(stars: 4043)](https://github.com/unidoc/unioffice) - Pure go library for creating and processing Office Word (.docx), Excel (.xlsx) and Powerpoint (.pptx) documents.

### Microsoft Excel

_Libraries for working with Microsoft Excel._

- [excelize(stars: 16443)](https://github.com/xuri/excelize) - Golang library for reading and writing Microsoft Excel&trade; (XLSX) files.
- [exl(stars: 23)](https://github.com/go-the-way/exl) - Excel binding to struct written in Go.(Only supports Go1.18+)
- [go-excel(stars: 181)](https://github.com/szyhf/go-excel) - A simple and light reader to read a relate-db-like excel as a table.
- [goxlsxwriter(stars: 21)](https://github.com/fterrag/goxlsxwriter) - Golang bindings for libxlsxwriter for writing XLSX (Microsoft Excel) files.
- [xlsx(stars: 5620)](https://github.com/tealeg/xlsx) - Library to simplify reading the XML format used by recent version of Microsoft Excel in Go programs.
- [xlsx(stars: 170)](https://github.com/plandem/xlsx) - Fast and safe way to read/update your existing Microsoft Excel files in Go programs.

**[⬆ back to top](#contents)**

## Miscellaneous

### Dependency Injection

_Libraries for working with dependency injection._

- [alice(stars: 51)](https://github.com/magic003/alice) - Additive dependency injection container for Golang.
- [boot-go](http://github.com/boot-go/boot) - Component-based development with dependency injection using reflections for Go developers.
- [di(stars: 213)](https://github.com/goava/di) - A dependency injection container for go programming language.
- [dig(stars: 3535)](https://github.com/uber-go/dig) - A reflection based dependency injection toolkit for Go.
- [dingo(stars: 171)](https://github.com/i-love-flamingo/dingo) - A dependency injection toolkit for Go, based on Guice.
- [do(stars: 1416)](https://github.com/samber/do) - A dependency injection framework based on Generics.
- [fx(stars: 4780)](https://github.com/uber-go/fx) - A dependency injection based application framework for Go (built on top of dig).
- [gocontainer(stars: 19)](https://github.com/vardius/gocontainer) - Simple Dependency Injection Container.
- [goioc/di(stars: 313)](https://github.com/goioc/di) - Spring-inspired Dependency Injection Container.
- [GoLobby/Container(stars: 515)](https://github.com/golobby/container) - GoLobby Container is a lightweight yet powerful IoC dependency injection container for the Go programming language.
- [gontainer(stars: 17)](https://github.com/NVIDIA/gontainer) - A dependency injection service container for Go projects.
- [gontainer/gontainer(stars: 12)](https://github.com/gontainer/gontainer) - A YAML-based Dependency Injection container for GO. It supports dependencies' scopes, and auto-detection of circular dependencies. Gontainer is concurrent-safe.
- [google/wire(stars: 11714)](https://github.com/google/wire) - Automated Initialization in Go.
- [HnH/di(stars: 7)](https://github.com/HnH/di) - DI container library that is focused on clean API and flexibility.
- [kinit(stars: 9)](https://github.com/go-kata/kinit) - Customizable dependency injection container with the global mode, cascade initialization and panic-safe finalization.
- [linker(stars: 35)](https://github.com/logrange/linker) - A reflection based dependency injection and inversion of control library with components lifecycle support.
- [nject(stars: 28)](https://github.com/muir/nject) - A type safe, reflective framework for libraries, tests, http endpoints, and service startup.
- [wire(stars: 38)](https://github.com/Fs02/wire) - Strict Runtime Dependency Injection for Golang.

**[⬆ back to top](#contents)**

### Project Layout

_**Unofficial** set of patterns for structuring projects._

- [ardanlabs/service(stars: 3137)](https://github.com/ardanlabs/service) - A [starter kit](https://github.com/ardanlabs/service/wiki) for building production grade scalable web service applications.
- [cookiecutter-golang(stars: 649)](https://github.com/lacion/cookiecutter-golang) - A Go application boilerplate template for quick starting projects following production best practices.
- [go-module(stars: 19)](https://github.com/octomation/go-module) - Template for a typical module written on Go.
- [go-sample(stars: 125)](https://github.com/zitryss/go-sample) - A sample layout for Go application projects with the real code.
- [go-starter(stars: 404)](https://github.com/allaboutapps/go-starter) - An opinionated production-ready RESTful JSON backend template, highly integrated with VSCode DevContainers.
- [go-todo-backend(stars: 274)](https://github.com/Fs02/go-todo-backend) - Go Todo Backend example using modular project layout for product microservice.
- [gobase(stars: 52)](https://github.com/wajox/gobase) - A simple skeleton for golang application with basic setup for real golang application.
- [golang-standards/project-layout(stars: 43666)](https://github.com/golang-standards/project-layout) - Set of common historical and emerging project layout patterns in the Go ecosystem. Note: despite the org-name they do not represent official golang standards, see [this issue](https://github.com/golang-standards/project-layout/issues/117) for more information. Nonetheless, some may find the layout useful.
- [golang-templates/seed(stars: 412)](https://github.com/golang-templates/seed) - Go application GitHub repository template.
- [insidieux/inizio(stars: 17)](https://github.com/insidieux/inizio) - Golang project layout generator with plugins.
- [modern-go-application(stars: 1721)](https://github.com/sagikazarmark/modern-go-application) - Go application boilerplate and example applying modern practices.
- [nunu(stars: 859)](https://github.com/go-nunu/nunu) - Nunu is a scaffolding tool for building Go applications. 
- [pagoda(stars: 1109)](https://github.com/mikestefanello/pagoda) - Rapid, easy full-stack web development starter kit built in Go.
- [scaffold(stars: 146)](https://github.com/catchplay/scaffold) - Scaffold generates a starter Go project layout. Lets you focus on business logic implemented.
- [wangyoucao577/go-project-layout(stars: 23)](https://github.com/wangyoucao577/go-project-layout) - Set of practices and discussions on how to structure Go project layout.

**[⬆ back to top](#contents)**

### Strings

_Libraries for working with strings._

- [bexp](https://github.com/happy-sdk/happy-go/tree/main/strings/bexp) - Go implementation of Brace Expansion mechanism to generate arbitrary strings.
- [caps(stars: 50)](https://github.com/chanced/caps) - A case conversion library.
- [go-formatter](https://gitlab.com/tymonx/go-formatter) - Implements **replacement fields** surrounded by curly braces `{}` format strings.
- [gobeam/Stringy(stars: 194)](https://github.com/gobeam/Stringy) - String manipulation library to convert string to camel case, snake case, kebab case / slugify etc.
- [strutil(stars: 198)](https://github.com/ozgio/strutil) - String utilities.
- [sttr(stars: 763)](https://github.com/abhimanyu003/sttr) - cross-platform, cli app to perform various operations on string.
- [xstrings(stars: 1264)](https://github.com/huandu/xstrings) - Collection of useful string functions ported from other languages.

**[⬆ back to top](#contents)**

### Uncategorized

_These libraries were placed here because none of the other categories seemed to fit._

- [anagent(stars: 16)](https://github.com/mudler/anagent) - Minimalistic, pluggable Golang evloop/timer handler with dependency-injection.
- [antch(stars: 252)](https://github.com/antchfx/antch) - A fast, powerful and extensible web crawling & scraping framework.
- [archiver(stars: 4144)](https://github.com/mholt/archiver) - Library and command for making and extracting .zip and .tar.gz archives.
- [autoflags(stars: 39)](https://github.com/artyom/autoflags) - Go package to automatically define command line flags from struct fields.
- [avgRating(stars: 15)](https://github.com/kirillDanshin/avgRating) - Calculate average score and rating based on Wilson Score Equation.
- [banner(stars: 443)](https://github.com/dimiro1/banner) - Add beautiful banners into your Go applications.
- [base64Captcha(stars: 1930)](https://github.com/mojocn/base64Captcha) - Base64captch supports digit, number, alphabet, arithmetic, audio and digit-alphabet captcha.
- [basexx(stars: 4)](https://github.com/bobg/basexx) - Convert to, from, and between digit strings in various number bases.
- [battery(stars: 230)](https://github.com/distatus/battery) - Cross-platform, normalized battery information library.
- [bitio(stars: 225)](https://github.com/icza/bitio) - Highly optimized bit-level Reader and Writer for Go.
- [browscap_go(stars: 46)](https://github.com/digitalcrab/browscap_go) - GoLang Library for [Browser Capabilities Project](https://browscap.org/).
- [captcha(stars: 138)](https://github.com/steambap/captcha) - Package captcha provides an easy to use, unopinionated API for captcha generation.
- [common(stars: 4)](https://github.com/kubeservice-stack/common) - A library for server framework.
- [conv(stars: 383)](https://github.com/cstockton/go-conv) - Package conv provides fast and intuitive conversions across Go types.
- [datacounter(stars: 47)](https://github.com/miolini/datacounter) - Go counters for readers/writer/http.ResponseWriter.
- [faker(stars: 10)](https://github.com/neotoolkit/faker) - Fake data generator.
- [faker(stars: 86)](https://github.com/pioz/faker) - Random fake data and struct generator for Go.
- [ffmt(stars: 303)](https://github.com/go-ffmt/ffmt) - Beautify data display for Humans.
- [gatus(stars: 4289)](https://github.com/TwinProduction/gatus) - Automated service health dashboard.
- [go-commandbus(stars: 11)](https://github.com/lana/go-commandbus) - A slight and pluggable command-bus for Go.
- [go-commons-pool(stars: 1188)](https://github.com/jolestar/go-commons-pool) - Generic object pool for Golang.
- [go-openapi](https://github.com/go-openapi) - Collection of packages to parse and utilize open-api schemas.
- [go-resiliency(stars: 1932)](https://github.com/eapache/go-resiliency) - Resiliency patterns for golang.
- [go-unarr(stars: 260)](https://github.com/gen2brain/go-unarr) - Decompression library for RAR, TAR, ZIP and 7z archives.
- [gofakeit(stars: 3883)](https://github.com/brianvoe/gofakeit) - Random data generator written in go.
- [gommit(stars: 106)](https://github.com/antham/gommit) - Analyze git commit messages to ensure they follow defined patterns.
- [gopsutil(stars: 9716)](https://github.com/shirou/gopsutil) - Cross-platform library for retrieving process and system utilization(CPU, Memory, Disks, etc).
- [gosh(stars: 35)](https://github.com/osamingo/gosh) - Provide Go Statistics Handler, Struct, Measure Method.
- [gosms(stars: 1438)](https://github.com/haxpax/gosms) - Your own local SMS gateway in Go that can be used to send SMS.
- [gotoprom(stars: 108)](https://github.com/cabify/gotoprom) - Type-safe metrics builder wrapper library for the official Prometheus client.
- [gountries(stars: 400)](https://github.com/pariz/gountries) - Package that exposes country and subdivision data.
- [gtree(stars: 174)](https://github.com/ddddddO/gtree) - Provide CLI, Package and Web for tree output and directories creation from Markdown or programmatically.
- [health(stars: 724)](https://github.com/alexliesenfeld/health) - A simple and flexible health check library for Go.
- [health(stars: 447)](https://github.com/dimiro1/health) - Easy to use, extensible health check library.
- [healthcheck(stars: 259)](https://github.com/etherlabsio/healthcheck) - An opinionated and concurrent health-check HTTP handler for RESTful services.
- [hostutils(stars: 12)](https://github.com/Wing924/hostutils) - A golang library for packing and unpacking FQDNs list.
- [indigo(stars: 105)](https://github.com/osamingo/indigo) - Distributed unique ID generator of using Sonyflake and encoded by Base58.
- [lk(stars: 331)](https://github.com/hyperboloide/lk) - A simple licensing library for golang.
- [llvm(stars: 1113)](https://github.com/llir/llvm) - Library for interacting with LLVM IR in pure Go.
- [metrics(stars: 27)](https://github.com/pascaldekloe/metrics) - Library for metrics instrumentation and Prometheus exposition.
- [morse(stars: 79)](https://github.com/alwindoss/morse) - Library to convert to and from morse code.
- [numa(stars: 19)](https://github.com/lrita/numa) - NUMA is a utility library, which is written in go. It help us to write some NUMA-AWARED code.
- [openapi(stars: 10)](https://github.com/neotoolkit/openapi) - OpenAPI 3.x parser.
- [pdfgen(stars: 66)](https://github.com/hyperboloide/pdfgen) - HTTP service to generate PDF from Json requests.
- [persian(stars: 78)](https://github.com/mavihq/persian) - Some utilities for Persian language in go.
- [sandid(stars: 41)](https://github.com/aofei/sandid) - Every grain of sand on earth has its own ID.
- [shellwords(stars: 21)](https://github.com/Wing924/shellwords) - A Golang library to manipulate strings according to the word parsing rules of the UNIX Bourne shell.
- [shortid(stars: 903)](https://github.com/teris-io/shortid) - Distributed generation of super short, unique, non-sequential, URL friendly IDs.
- [shoutrrr(stars: 774)](https://github.com/containrrr/shoutrrr) - Notification library providing easy access to various messaging services like slack, mattermost, gotify and smtp among others.
- [sitemap-format(stars: 4)](https://github.com/mingard/sitemap-format) - A simple sitemap generator, with a little syntactic sugar.
- [stateless(stars: 726)](https://github.com/qmuntal/stateless) - A fluent library for creating state machines.
- [stats(stars: 168)](https://github.com/go-playground/stats) - Monitors Go MemStats + System stats such as Memory, Swap and CPU and sends via UDP anywhere you want for logging etc...
- [turtle(stars: 153)](https://github.com/hackebrot/turtle) - Emojis for Go.
- [url-shortener(stars: 44)](https://github.com/pantrif/url-shortener) - A modern, powerful, and robust URL shortener microservice with mysql support.
- [VarHandler](https://github.com/azr/generators/tree/master/varhandler) - Generate boilerplate http input and output handling.
- [varint(stars: 11)](https://github.com/chmike/varint) - A faster varying length integer encoder/decoder than the one provided in the standard library.
- [xdg(stars: 36)](https://github.com/rkoesters/xdg) - FreeDesktop.org (xdg) Specs implemented in Go.
- [xkg(stars: 56)](https://github.com/go-xkg/xkg) - X Keyboard Grabber.
- [xz(stars: 437)](https://github.com/ulikunitz/xz) - Pure golang package for reading and writing xz-compressed files.

**[⬆ back to top](#contents)**

## Natural Language Processing

_Libraries for working with human languages._

See also [Text Processing](#text-processing) and [Text Analysis](#text-analysis).

### Language Detection

- [detectlanguage(stars: 23)](https://github.com/detectlanguage/detectlanguage-go) - Language Detection API Go Client. Supports batch requests, short phrase or single word language detection.
- [getlang(stars: 161)](https://github.com/rylans/getlang) - Fast natural language detection package.
- [guesslanguage(stars: 57)](https://github.com/endeveit/guesslanguage) - Functions to determine the natural language of a unicode text.
- [whatlanggo(stars: 606)](https://github.com/abadojack/whatlanggo) - Natural language detection package for Go. Supports 84 languages and 24 scripts (writing systems e.g. Latin, Cyrillic, etc).

### Morphological Analyzers

- [go-stem(stars: 76)](https://github.com/agonopol/go-stem) - Implementation of the porter stemming algorithm.
- [go2vec(stars: 50)](https://github.com/danieldk/go2vec) - Reader and utility functions for word2vec embeddings.
- [golibstemmer(stars: 20)](https://github.com/rjohnsondev/golibstemmer) - Go bindings for the snowball libstemmer library including porter 2.
- [gosentiwordnet(stars: 11)](https://github.com/dinopuguh/gosentiwordnet) - Sentiment analyzer using sentiwordnet lexicon in Go.
- [govader(stars: 38)](https://github.com/jonreiter/govader) - Go implementation of [VADER Sentiment Analysis](https://github.com/cjhutto/vaderSentiment).
- [govader-backend(stars: 6)](https://github.com/PIMPfiction/govader_backend) - Microservice implementation of [GoVader](https://github.com/jonreiter/govader).
- [kagome(stars: 766)](https://github.com/ikawaha/kagome) - JP morphological analyzer written in pure Go.
- [libtextcat(stars: 13)](https://github.com/goodsign/libtextcat) - Cgo binding for libtextcat C library. Guaranteed compatibility with version 2.2.
- [nlp(stars: 386)](https://github.com/Shixzie/nlp) - Extract values from strings and fill your structs with nlp.
- [nlp(stars: 424)](https://github.com/james-bowman/nlp) - Go Natural Language Processing library supporting LSA (Latent Semantic Analysis).
- [paicehusk(stars: 29)](https://github.com/rookii/paicehusk) - Golang implementation of the Paice/Husk Stemming Algorithm.
- [porter(stars: 12)](https://github.com/a2800276/porter) - This is a fairly straightforward port of Martin Porter's C implementation of the Porter stemming algorithm.
- [porter2(stars: 46)](https://github.com/zhenjl/porter2) - Really fast Porter 2 stemmer.
- [RAKE.go](https://github.com/afjoseph/RAKE.Go) - Go port of the Rapid Automatic Keyword Extraction Algorithm (RAKE).
- [snowball(stars: 36)](https://github.com/goodsign/snowball) - Snowball stemmer port (cgo wrapper) for Go. Provides word stem extraction functionality [Snowball native](http://snowball.tartarus.org/).
- [spaGO(stars: 1682)](https://github.com/nlpodyssey/spago) - Self-contained Machine Learning and Natural Language Processing library in Go.
- [spelling-corrector(stars: 2)](https://github.com/jorelosorio/spellingcorrector) - A spelling corrector for the Spanish language or create your own.

### Slugifiers

- [go-slugify(stars: 88)](https://github.com/mozillazg/go-slugify) - Make pretty slug with multiple languages support.
- [slug(stars: 1040)](https://github.com/gosimple/slug) - URL-friendly slugify with multiple languages support.
- [Slugify(stars: 32)](https://github.com/avelino/slugify) - Go slugify application that handles string.

### Tokenizers

- [gojieba(stars: 2269)](https://github.com/yanyiwu/gojieba) - This is a Go implementation of [jieba](https://github.com/fxsjy/jieba) which a Chinese word splitting algorithm.
- [gotokenizer(stars: 18)](https://github.com/xujiajun/gotokenizer) - A tokenizer based on the dictionary and Bigram language models for Golang. (Now only support chinese segmentation)
- [gse(stars: 2394)](https://github.com/go-ego/gse) - Go efficient text segmentation; support english, chinese, japanese and other.
- [MMSEGO(stars: 62)](https://github.com/awsong/MMSEGO) - This is a GO implementation of [MMSEG](http://technology.chtsai.org/mmseg/) which a Chinese word splitting algorithm.
- [prose(stars: 3025)](https://github.com/jdkato/prose) - Library for text processing that supports tokenization, part-of-speech tagging, named-entity extraction, and more. English only.
- [segment(stars: 83)](https://github.com/blevesearch/segment) - Go library for performing Unicode Text Segmentation as described in [Unicode Standard Annex #29](https://www.unicode.org/reports/tr29/)
- [sentences(stars: 398)](https://github.com/neurosnap/sentences) - Sentence tokenizer: converts text into a list of sentences.
- [shamoji(stars: 13)](https://github.com/osamingo/shamoji) - The shamoji is word filtering package written in Go.
- [stemmer(stars: 50)](https://github.com/dchest/stemmer) - Stemmer packages for Go programming language. Includes English and German stemmers.
- [textcat(stars: 71)](https://github.com/pebbe/textcat) - Go package for n-gram based text categorization, with support for utf-8 and raw text.

### Translation

- [go-i18n](https://github.com/nicksnyder/go-i18n/) - Package and an accompanying tool to work with localized text.
- [go-localize(stars: 55)](https://github.com/m1/go-localize) - Simple and easy to use i18n (Internationalization and localization) engine - used for translating locale strings.
- [go-mystem(stars: 30)](https://github.com/dveselov/mystem) - CGo bindings to Yandex.Mystem - russian morphology analyzer.
- [go-pinyin(stars: 1516)](https://github.com/mozillazg/go-pinyin) - CN Hanzi to Hanyu Pinyin converter.
- [gotext(stars: 409)](https://github.com/leonelquinteros/gotext) - GNU gettext utilities for Go.
- [icu(stars: 21)](https://github.com/goodsign/icu) - Cgo binding for icu4c C library detection and conversion functions. Guaranteed compatibility with version 50.1.
- [iuliia-go(stars: 42)](https://github.com/mehanizm/iuliia-go) - Transliterate Cyrillic → Latin in every possible way.
- [spreak(stars: 29)](https://github.com/vorlif/spreak) - Flexible translation and humanization library for Go, based on the concepts behind gettext.
- [t(stars: 17)](https://github.com/youthlin/t) - Another i18n pkg for golang, which follows GNU gettext style and supports .po/.mo files: `t.T (gettext)`, `t.N (ngettext)`, etc. And it contains a cmd tool [xtemplate](https://github.com/youthlin/t/blob/main/cmd/xtemplate), which can extract messages as a pot file from text/html template.

### Transliteration

- [enca(stars: 16)](https://github.com/endeveit/enca) - Minimal cgo bindings for [libenca](https://cihar.com/software/enca/), which detects character encodings.
- [go-unidecode(stars: 117)](https://github.com/mozillazg/go-unidecode) - ASCII transliterations of Unicode text.
- [gounidecode(stars: 79)](https://github.com/fiam/gounidecode) - Unicode transliterator (also known as unidecode) for Go.
- [transliterator(stars: 38)](https://github.com/alexsergivan/transliterator) - Provides one-way string transliteration with supporting of language-specific transliteration rules.

**[⬆ back to top](#contents)**

## Networking

_Libraries for working with various layers of the network._

- [arp(stars: 349)](https://github.com/mdlayher/arp) - Package arp implements the ARP protocol, as described in RFC 826.
- [buffstreams(stars: 254)](https://github.com/stabbycutyou/buffstreams) - Streaming protocolbuffer data over TCP made easy.
- [canopus(stars: 153)](https://github.com/zubairhamed/canopus) - CoAP Client/Server implementation (RFC 7252).
- [cidranger(stars: 866)](https://github.com/yl2chen/cidranger) - Fast IP to CIDR lookup for Go.
- [dhcp6(stars: 76)](https://github.com/mdlayher/dhcp6) - Package dhcp6 implements a DHCPv6 server, as described in RFC 3315.
- [dns(stars: 7588)](https://github.com/miekg/dns) - Go library for working with DNS.
- [dnsmonster(stars: 287)](https://github.com/mosajjal/dnsmonster) - Passive DNS Capture/Monitoring Framework.
- [easytcp(stars: 781)](https://github.com/DarthPestilane/easytcp) - A light-weight TCP framework written in Go (Golang), built with message router. EasyTCP helps you build a TCP server easily fast and less painful.
- [ether(stars: 79)](https://github.com/songgao/ether) - Cross-platform Go package for sending and receiving ethernet frames.
- [ethernet(stars: 267)](https://github.com/mdlayher/ethernet) - Package ethernet implements marshaling and unmarshalling of IEEE 802.3 Ethernet II frames and IEEE 802.1Q VLAN tags.
- [event(stars: 153)](https://github.com/cheng-zhongliang/event) - Simple I/O event notification library written in Golang.
- [fasthttp(stars: 20532)](https://github.com/valyala/fasthttp) - Package fasthttp is a fast HTTP implementation for Go, up to 10 times faster than net/http.
- [fortio(stars: 3082)](https://github.com/fortio/fortio) - Load testing library and command line tool, advanced echo server and web UI. Allows to specify a set query-per-second load and record latency histograms and other useful stats and graph them. Tcp, Http, gRPC.
- [ftp(stars: 1196)](https://github.com/jlaffaye/ftp) - Package ftp implements a FTP client as described in [RFC 959](https://tools.ietf.org/html/rfc959).
- [ftpserverlib(stars: 380)](https://github.com/fclairamb/ftpserverlib) - Fully featured FTP server library.
- [fullproxy(stars: 68)](https://github.com/shoriwe/fullproxy) - A fully featured scriptable and daemon configurable proxy and pivoting toolkit with SOCKS5, HTTP, raw ports and reverse proxy protocols.
- [gaio(stars: 529)](https://github.com/xtaci/gaio) - High performance async-io networking for Golang in proactor mode.
- [gev(stars: 1687)](https://github.com/Allenxuxu/gev) - gev is a lightweight, fast non-blocking TCP network library based on Reactor mode.
- [gldap(stars: 103)](https://github.com/jimlambrt/gldap) - gldap provides an ldap server implementation and you provide handlers for its ldap operations.
- [gmqtt(stars: 939)](https://github.com/DrmagicE/gmqtt) - Gmqtt is a flexible, high-performance MQTT broker library that fully implements the MQTT protocol V3.1.1.
- [gnet(stars: 8281)](https://github.com/panjf2000/gnet) - `gnet` is a high-performance, lightweight, non-blocking, event-driven networking framework written in pure Go.
- [gnet(stars: 12)](https://github.com/fish-tennis/gnet) - `gnet` is a high-performance networking framework,especially for game servers.
- [gNxI(stars: 247)](https://github.com/google/gnxi) - A collection of tools for Network Management that use the gNMI and gNOI protocols.
- [go-getter(stars: 1544)](https://github.com/hashicorp/go-getter) - Go library for downloading files or directories from various sources using a URL.
- [go-powerdns(stars: 76)](https://github.com/joeig/go-powerdns) - PowerDNS API bindings for Golang.
- [go-sse(stars: 6)](https://github.com/lampctl/go-sse) - Go client and server implementation of HTML server-sent events.
- [go-stun(stars: 614)](https://github.com/ccding/go-stun) - Go implementation of the STUN client (RFC 3489 and RFC 5389).
- [gobgp(stars: 3374)](https://github.com/osrg/gobgp) - BGP implemented in the Go Programming Language.
- [golibwireshark(stars: 29)](https://github.com/sunwxg/golibwireshark) - Package golibwireshark use libwireshark library to decode pcap file and analyse dissection data.
- [gopacket(stars: 5954)](https://github.com/google/gopacket) - Go library for packet processing with libpcap bindings.
- [gopcap(stars: 480)](https://github.com/akrennmair/gopcap) - Go wrapper for libpcap.
- [goshark(stars: 18)](https://github.com/sunwxg/goshark) - Package goshark use tshark to decode IP packet and create data struct to analyse packet.
- [gosnmp(stars: 1030)](https://github.com/soniah/gosnmp) - Native Go library for performing SNMP actions.
- [gotcp(stars: 508)](https://github.com/gansidui/gotcp) - Go package for quickly writing tcp applications.
- [grab(stars: 1338)](https://github.com/cavaliercoder/grab) - Go package for managing file downloads.
- [graval(stars: 28)](https://github.com/koofr/graval) - Experimental FTP server framework.
- [gws(stars: 846)](https://github.com/lxzan/gws) - High-Performance WebSocket Server & Client With AsyncIO Supporting .
- [HTTPLab(stars: 3940)](https://github.com/gchaincl/httplab) - HTTPLabs let you inspect HTTP requests and forge responses.
- [httpproxy(stars: 21)](https://github.com/wzshiming/httpproxy) - HTTP proxy handler and dialer.
- [iplib(stars: 126)](https://github.com/c-robinson/iplib) - Library for working with IP addresses (net.IP, net.IPNet), inspired by python [ipaddress](https://docs.python.org/3/library/ipaddress.html) and ruby [ipaddr](https://ruby-doc.org/stdlib-2.5.1/libdoc/ipaddr/rdoc/IPAddr.html)
- [jazigo(stars: 203)](https://github.com/udhos/jazigo) - Jazigo is a tool written in Go for retrieving configuration for multiple network devices.
- [kcp-go(stars: 3830)](https://github.com/xtaci/kcp-go) - KCP - Fast and Reliable ARQ Protocol.
- [kcptun(stars: 13551)](https://github.com/xtaci/kcptun) - Extremely simple & fast udp tunnel based on KCP protocol.
- [lhttp(stars: 691)](https://github.com/fanux/lhttp) - Powerful websocket framework, build your IM server more easily.
- [linkio(stars: 52)](https://github.com/ian-kent/linkio) - Network link speed simulation for Reader/Writer interfaces.
- [llb(stars: 14)](https://github.com/kirillDanshin/llb) - It's a very simple but quick backend for proxy servers. Can be useful for fast redirection to predefined domain with zero memory allocation and fast response.
- [mdns(stars: 1067)](https://github.com/hashicorp/mdns) - Simple mDNS (Multicast DNS) client/server library in Golang.
- [mqttPaho](https://eclipse.org/paho/clients/golang/) - The Paho Go Client provides an MQTT client library for connection to MQTT brokers via TCP, TLS or WebSockets.
- [nbio(stars: 1857)](https://github.com/lesismal/nbio) - Pure Go 1000k+ connections solution, support tls/http1.x/websocket and basically compatible with net/http, with high-performance and low memory cost, non-blocking, event-driven, easy-to-use.
- [netpoll(stars: 3766)](https://github.com/cloudwego/netpoll) - A high-performance non-blocking I/O networking framework, which focused on RPC scenarios, developed by ByteDance.
- [NFF-Go(stars: 1336)](https://github.com/intel-go/nff-go) - Framework for rapid development of performant network functions for cloud and bare-metal (former YANFF).
- [packet(stars: 76)](https://github.com/aerogo/packet) - Send packets over TCP and UDP. It can buffer messages and hot-swap connections if needed.
- [panoptes-stream(stars: 40)](https://github.com/yahoo/panoptes-stream) - A cloud native distributed streaming network telemetry (gNMI, Juniper JTI and Cisco MDT).
- [peerdiscovery(stars: 615)](https://github.com/schollz/peerdiscovery) - Pure Go library for cross-platform local peer discovery using UDP multicast.
- [portproxy(stars: 53)](https://github.com/aybabtme/portproxy) - Simple TCP proxy which adds CORS support to API's which don't support it.
- [publicip(stars: 27)](https://github.com/polera/publicip) - Package publicip returns your public facing IPv4 address (internet egress).
- [quic-go(stars: 9058)](https://github.com/lucas-clemente/quic-go) - An implementation of the QUIC protocol in pure Go.
- [raw(stars: 422)](https://github.com/mdlayher/raw) - Package raw enables reading and writing data at the device driver level for a network interface.
- [sftp(stars: 1425)](https://github.com/pkg/sftp) - Package sftp implements the SSH File Transfer Protocol as described in <https://filezilla-project.org/specs/draft-ietf-secsh-filexfer-02.txt>.
- [ssh(stars: 3216)](https://github.com/gliderlabs/ssh) - Higher-level API for building SSH servers (wraps crypto/ssh).
- [sslb(stars: 145)](https://github.com/eduardonunesp/sslb) - It's a Super Simples Load Balancer, just a little project to achieve some kind of performance.
- [stun(stars: 487)](https://github.com/go-rtc/stun) - Go implementation of RFC 5389 STUN protocol.
- [tcp_server(stars: 424)](https://github.com/firstrow/tcp_server) - Go library for building tcp servers faster.
- [tcpack(stars: 176)](https://github.com/lim-yoona/tcpack) - tcpack is an application protocol based on TCP to Pack and Unpack bytes stream in go program.
- [tspool(stars: 14)](https://github.com/two/tspool) - A TCP Library use worker pool to improve performance and protect your server.
- [utp(stars: 170)](https://github.com/anacrolix/utp) - Go uTP micro transport protocol implementation.
- [vssh(stars: 923)](https://github.com/yahoo/vssh) - Go library for building network and server automation over SSH protocol.
- [water(stars: 1791)](https://github.com/songgao/water) - Simple TUN/TAP library.
- [webrtc(stars: 12121)](https://github.com/pions/webrtc) - A pure Go implementation of the WebRTC API.
- [winrm(stars: 410)](https://github.com/masterzen/winrm) - Go WinRM client to remotely execute commands on Windows machines.
- [xtcp(stars: 148)](https://github.com/xfxdev/xtcp) - TCP Server Framework with simultaneous full duplex communication, graceful shutdown, and custom protocol.

**[⬆ back to top](#contents)**

### HTTP Clients

_Libraries for making HTTP requests._

- [fast-shot(stars: 19)](https://github.com/opus-domini/fast-shot) - Hit your API targets with rapid-fire precision using Go's fastest and simple HTTP Client.
- [gentleman(stars: 1031)](https://github.com/h2non/gentleman) - Full-featured plugin-driven HTTP client library.
- [go-cleanhttp(stars: 352)](https://github.com/hashicorp/go-cleanhttp) - Get easily stdlib HTTP client, which does not share any state with other clients.
- [go-http-client(stars: 64)](https://github.com/bozd4g/go-http-client) - Make http calls simply and easily.
- [go-otelroundtripper(stars: 74)](https://github.com/NdoleStudio/go-otelroundtripper) - Go http.RoundTripper that emits open telemetry metrics for HTTP requests.
- [go-req(stars: 21)](https://github.com/wenerme/go-req) - Declarative golang HTTP client.
- [go-retryablehttp(stars: 1748)](https://github.com/hashicorp/go-retryablehttp) - Retryable HTTP client in Go.
- [go-zoox/fetch(stars: 61)](https://github.com/go-zoox/fetch) - A Powerful, Lightweight, Easy Http Client, inspired by Web Fetch API.
- [grequests(stars: 2008)](https://github.com/levigross/grequests) - A Go "clone" of the great and famous Requests library.
- [heimdall(stars: 2525)](https://github.com/gojektech/heimdall) - An enhanced http client with retry and hystrix capabilities.
- [httpretry(stars: 40)](https://github.com/ybbus/httpretry) - Enriches the default go HTTP client with retry functionality.
- [pester(stars: 624)](https://github.com/sethgrid/pester) - Go HTTP client calls with retries, backoff, and concurrency.
- [req(stars: 3832)](https://github.com/imroc/req) - Simple Go HTTP client with Black Magic (Less code and More efficiency).
- [request(stars: 268)](https://github.com/monaco-io/request) - HTTP client for golang. If you have experience about axios or requests, you will love it. No 3rd dependency.
- [requests(stars: 1301)](https://github.com/carlmjohnson/requests) - HTTP requests for Gophers. Uses context.Context and doesn't hide the underlying net/http.Client, making it compatible with standard Go APIs. Also includes testing tools.
- [resty(stars: 8829)](https://github.com/go-resty/resty) - Simple HTTP and REST client for Go inspired by Ruby rest-client.
- [rq(stars: 51)](https://github.com/ddo/rq) - A nicer interface for golang stdlib HTTP client.
- [sling(stars: 1614)](https://github.com/dghubble/sling) - Sling is a Go HTTP client library for creating and sending API requests.

**[⬆ back to top](#contents)**

## OpenGL

_Libraries for using OpenGL in Go._

- [gl(stars: 1003)](https://github.com/go-gl/gl) - Go bindings for OpenGL (generated via glow).
- [glfw(stars: 1467)](https://github.com/go-gl/glfw) - Go bindings for GLFW 3.
- [go-glmatrix(stars: 7)](https://github.com/technohippy/go-glmatrix) - Go port of [glMatrix](https://glmatrix.net/) library.
- [goxjs/gl(stars: 169)](https://github.com/goxjs/gl) - Go cross-platform OpenGL bindings (OS X, Linux, Windows, browsers, iOS, Android).
- [goxjs/glfw(stars: 79)](https://github.com/goxjs/glfw) - Go cross-platform glfw library for creating an OpenGL context and receiving events.
- [mathgl(stars: 517)](https://github.com/go-gl/mathgl) - Pure Go math package specialized for 3D math, with inspiration from GLM.

**[⬆ back to top](#contents)**

## ORM

_Libraries that implement Object-Relational Mapping or datamapping techniques._

- [bun(stars: 2668)](https://github.com/uptrace/bun) - SQL-first Golang ORM. Successor of go-pg.
- [cacheme(stars: 24)](https://github.com/Yiling-J/cacheme-go) - Schema based, typed Redis caching/memoize framework for Go.
- [ent(stars: 14432)](https://github.com/facebook/ent) - An entity framework for Go. Simple, yet powerful ORM for modeling and querying data.
- [go-dbw(stars: 10)](https://github.com/hashicorp/go-dbw) - A simple package that encapsulates database operations.
- [go-firestorm(stars: 47)](https://github.com/jschoedt/go-firestorm) - A simple ORM for Google/Firebase Cloud Firestore.
- [go-sql(stars: 176)](https://github.com/rushteam/gosql) - A easy ORM for mysql.
- [go-sqlbuilder(stars: 1148)](https://github.com/huandu/go-sqlbuilder) - A flexible and powerful SQL string builder library plus a zero-config ORM.
- [go-store(stars: 112)](https://github.com/gosuri/go-store) - Simple and fast Redis backed key-value store library for Go.
- [golobby/orm(stars: 150)](https://github.com/golobby/orm) - Simple, fast, type-safe, generic orm for developer happiness.
- [GORM(stars: 34443)](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang, aims to be developer friendly.
- [gormt(stars: 2286)](https://github.com/xxjwxc/gormt) - Mysql database to golang gorm struct.
- [gorp(stars: 3700)](https://github.com/go-gorp/gorp) - Go Relational Persistence, ORM-ish library for Go.
- [grimoire(stars: 160)](https://github.com/Fs02/grimoire) - Grimoire is a database access layer and validation for golang. (Support: MySQL, PostgreSQL and SQLite3).
- [lore(stars: 14)](https://github.com/abrahambotros/lore) - Simple and lightweight pseudo-ORM/pseudo-struct-mapping environment for Go.
- [marlow(stars: 14)](https://github.com/marlow/marlow) - Generated ORM from project structs for compile time safety assurances.
- [pop/soda(stars: 1390)](https://github.com/gobuffalo/pop) - Database migration, creation, ORM, etc... for MySQL, PostgreSQL, and SQLite.
- [Prisma(stars: 1783)](https://github.com/prisma/prisma-client-go) - Prisma Client Go, Typesafe database access for Go.
- [reform(stars: 1429)](https://github.com/go-reform/reform) - Better ORM for Go, based on non-empty interfaces and code generation.
- [rel(stars: 698)](https://github.com/go-rel/rel) - Modern Database Access Layer for Golang - Testable, Extendable and Crafted Into a Clean and Elegant API.
- [SQLBoiler(stars: 6190)](https://github.com/volatiletech/sqlboiler) - ORM generator. Generate a featureful and blazing-fast ORM tailored to your database schema.
- [upper.io/db(stars: 3424)](https://github.com/upper/db) - Single interface for interacting with different data sources through the use of adapters that wrap mature database drivers.
- [XORM](https://gitea.com/xorm/xorm) - Simple and powerful ORM for Go. (Support: MySQL, MyMysql, PostgreSQL, Tidb, SQLite3, MsSql and Oracle).
- [Zoom(stars: 305)](https://github.com/albrow/zoom) - Blazing-fast datastore and querying engine built on Redis.

**[⬆ back to top](#contents)**

## Package Management

_Official tooling for dependency and package management_

- [go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) - Modules are the unit of source code interchange and versioning. The go command has direct support for working with modules, including recording and resolving dependencies on other modules.

_Official experimental tooling for package management_

- [dep(stars: 12913)](https://github.com/golang/dep) - Go dependency tool.
- [vgo](https://go.googlesource.com/vgo/) - Versioned Go.

_Unofficial libraries for package and dependency management._

- [glide(stars: 8167)](https://github.com/Masterminds/glide) - Manage your golang vendor and vendored packages with ease. Inspired by tools like Maven, Bundler, and Pip.
- [godep(stars: 5565)](https://github.com/tools/godep) - dependency tool for go, godep helps build packages reproducibly by fixing their dependencies.
- [gom(stars: 1390)](https://github.com/mattn/gom) - Go Manager - bundle for go.
- [goop(stars: 780)](https://github.com/nitrous-io/goop) - Simple dependency manager for Go (golang), inspired by Bundler.
- [gop(stars: 50)](https://github.com/lunny/gop) - Build and manage your Go applications out of GOPATH.
- [gopm(stars: 2473)](https://github.com/gpmgo/gopm) - Go Package Manager.
- [govendor(stars: 4941)](https://github.com/kardianos/govendor) - Go Package Manager. Go vendor tool that works with the standard vendor file.
- [gpm(stars: 1194)](https://github.com/pote/gpm) - Barebones dependency manager for Go.
- [gup(stars: 239)](https://github.com/nao1215/gup) - Update binaries installed by "go install".
- [johnny-deps(stars: 214)](https://github.com/VividCortex/johnny-deps) - Minimal dependency version using Git.
- [modgv(stars: 475)](https://github.com/lucasepe/modgv) - Converts 'go mod graph' output into Graphviz's DOT language.
- [mvn-golang(stars: 158)](https://github.com/raydac/mvn-golang) - plugin that provides way for auto-loading of Golang SDK, dependency management and start build environment in Maven project infrastructure.
- [nut(stars: 234)](https://github.com/jingweno/nut) - Vendor Go dependencies.
- [VenGO(stars: 123)](https://github.com/DamnWidget/VenGO) - create and manage exportable isolated go virtual environments.

**[⬆ back to top](#contents)**

## Performance

- [go-instrument(stars: 113)](https://github.com/nikolaydubina/go-instrument) - Automatically add spans to all methods and functions.
- [jaeger(stars: 18834)](https://github.com/jaegertracing/jaeger) - A distributed tracing system.
- [pixie(stars: 5045)](https://github.com/pixie-labs/pixie) - No instrumentation tracing for Golang applications via eBPF.
- [profile(stars: 1924)](https://github.com/pkg/profile) - Simple profiling support package for Go.
- [statsviz(stars: 3026)](https://github.com/arl/statsviz) - Live visualization of your Go application runtime statistics.
- [tracer(stars: 82)](https://github.com/kamilsk/tracer) - Simple, lightweight tracing.

**[⬆ back to top](#contents)**

## Query Language

- [api-fu(stars: 52)](https://github.com/ccbrown/api-fu) - Comprehensive GraphQL implementation.
- [dasel(stars: 4676)](https://github.com/tomwright/dasel) - Query and update data structures using selectors from the command line. Comparable to jq/yq but supports JSON, YAML, TOML and XML with zero runtime dependencies.
- [gojsonq(stars: 2126)](https://github.com/thedevsaddam/gojsonq) - A simple Go package to Query over JSON Data.
- [goven(stars: 57)](https://github.com/SeldonIO/goven) - A drop-in query language for any database schema.
- [gqlgen(stars: 9376)](https://github.com/99designs/gqlgen) - go generate based graphql server library.
- [grapher(stars: 3)](https://github.com/reaganiwadha/grapher) - A GraphQL field builder utilizing Go generics with extra utilities and features.
- [graphql(stars: 57)](https://github.com/tmc/graphql) - graphql parser + utilities.
- [graphql(stars: 4554)](https://github.com/neelance/graphql-go) - GraphQL server with a focus on ease of use.
- [graphql-go(stars: 9556)](https://github.com/graphql-go/graphql) - Implementation of GraphQL for Go.
- [gws(stars: 7)](https://github.com/Zaba505/gws) - Apollos' "GraphQL over Websocket" client and server implementation.
- [jsonpath(stars: 20)](https://github.com/AsaiYusuke/jsonpath) - A query library for retrieving part of JSON based on JSONPath syntax.
- [jsonql(stars: 271)](https://github.com/elgs/jsonql) - JSON query expression library in Golang.
- [jsonslice(stars: 79)](https://github.com/bhmj/jsonslice) - Jsonpath queries with advanced filters.
- [mql(stars: 24)](https://github.com/hashicorp/mql) - Model Query Language (mql) is a query language for your database models.
- [rql(stars: 318)](https://github.com/a8m/rql) - Resource Query Language for REST API.
- [rqp(stars: 61)](https://github.com/timsolov/rest-query-parser) - Query Parser for REST API. Filtering, validations, both `AND`, `OR` operations are supported directly in the query.
- [straf(stars: 36)](https://github.com/SonicRoshan/straf) - Easily Convert Golang structs to GraphQL objects.

**[⬆ back to top](#contents)**

## Resource Embedding

- [debme(stars: 27)](https://github.com/leaanthony/debme) - Create an `embed.FS` from an existing `embed.FS` subdirectory.
- [esc(stars: 637)](https://github.com/mjibson/esc) - Embeds files into Go programs and provides http.FileSystem interfaces to them.
- [fileb0x(stars: 630)](https://github.com/UnnoTed/fileb0x) - Simple tool to embed files in go with focus on "customization" and ease to use.
- [go-resources(stars: 176)](https://github.com/omeid/go-resources) - Unfancy resources embedding with Go.
- [go.rice](https://github.com/GeertJohan/go.rice) - go.rice is a Go package that makes working with resources such as HTML, JS, CSS, images, and templates very easy.
- [mule(stars: 13)](https://github.com/wlbr/mule) - Embed external resources like images, movies ... into Go source code to create single file binaries using `go generate`. Focused on simplicity.
- [packr(stars: 3409)](https://github.com/gobuffalo/packr) - The simple and easy way to embed static files into Go binaries.
- [rebed(stars: 28)](https://github.com/soypat/rebed) - Recreate folder structures and files from Go 1.16's `embed.FS` type
- [statics(stars: 67)](https://github.com/go-playground/statics) - Embeds static resources into go files for single binary compilation + works with http.FileSystem + symlinks.
- [statik(stars: 3693)](https://github.com/rakyll/statik) - Embeds static files into a Go executable.
- [templify(stars: 30)](https://github.com/wlbr/templify) - Embed external template files into Go code to create single file binaries.
- [vfsgen(stars: 980)](https://github.com/shurcooL/vfsgen) - Generates a vfsdata.go file that statically implements the given virtual filesystem.

**[⬆ back to top](#contents)**

## Science and Data Analysis

_Libraries for scientific computing and data analyzing._

- [assocentity(stars: 14)](https://github.com/ndabAP/assocentity) - Package assocentity returns the average distance from words to a given entity.
- [bradleyterry(stars: 9)](https://github.com/seanhagen/bradleyterry) - Provides a Bradley-Terry Model for pairwise comparisons.
- [calendarheatmap(stars: 381)](https://github.com/nikolaydubina/calendarheatmap) - Calendar heatmap in plain Go inspired by Github contribution activity.
- [chart(stars: 761)](https://github.com/vdobler/chart) - Simple Chart Plotting library for Go. Supports many graphs types.
- [dataframe-go(stars: 1083)](https://github.com/rocketlaunchr/dataframe-go) - Dataframes for machine-learning and statistics (similar to pandas).
- [decimal(stars: 38)](https://github.com/db47h/decimal) - Package decimal implements arbitrary-precision decimal floating-point arithmetic.
- [evaler(stars: 52)](https://github.com/soniah/evaler) - Simple floating point arithmetic expression evaluator.
- [ewma(stars: 419)](https://github.com/VividCortex/ewma) - Exponentially-weighted moving averages.
- [geom(stars: 55)](https://github.com/skelterjohn/geom) - 2D geometry for golang.
- [go-dsp(stars: 833)](https://github.com/mjibson/go-dsp) - Digital Signal Processing for Go.
- [go-estimate(stars: 105)](https://github.com/milosgajdos/go-estimate) - State estimation and filtering algorithms in Go.
- [go-gt(stars: 10)](https://github.com/ThePaw/go-gt) - Graph theory algorithms written in "Go" language.
- [godesim(stars: 22)](https://github.com/soypat/godesim) - Extended/multivariable ODE solver framework for event-based simulations with simple API.
- [goent(stars: 34)](https://github.com/kzahedi/goent) - GO Implementation of Entropy Measures.
- [gograph(stars: 40)](https://github.com/hmdsefi/gograph) - A golang generic graph library that provides mathematical graph-theory and algorithms.
- [gohistogram(stars: 172)](https://github.com/VividCortex/gohistogram) - Approximate histograms for data streams.
- [gonum(stars: 7029)](https://github.com/gonum/gonum) - Gonum is a set of numeric libraries for the Go programming language. It contains libraries for matrices, statistics, optimization, and more.
- [gonum/plot(stars: 2568)](https://github.com/gonum/plot) - gonum/plot provides an API for building and drawing plots in Go.
- [goraph(stars: 726)](https://github.com/gyuho/goraph) - Pure Go graph theory library(data structure, algorithm visualization).
- [gosl(stars: 1789)](https://github.com/cpmech/gosl) - Go scientific library for linear algebra, FFT, geometry, NURBS, numerical methods, probabilities, optimisation, differential equations, and more.
- [GoStats(stars: 20)](https://github.com/OGFris/GoStats) - GoStats is an Open Source GoLang library for math statistics mostly used in Machine Learning domains, it covers most of the Statistical measures functions.
- [graph(stars: 665)](https://github.com/yourbasic/graph) - Library of basic graph algorithms.
- [jsonl-graph(stars: 69)](https://github.com/nikolaydubina/jsonl-graph) - Tool to manipulate JSONL graphs with graphviz support.
- [ode(stars: 22)](https://github.com/ChristopherRabotin/ode) - Ordinary differential equation (ODE) solver which supports extended states and channel-based iteration stop conditions.
- [orb(stars: 772)](https://github.com/paulmach/orb) - 2D geometry types with clipping, GeoJSON and Mapbox Vector Tile support.
- [pagerank(stars: 78)](https://github.com/alixaxel/pagerank) - Weighted PageRank algorithm implemented in Go.
- [piecewiselinear(stars: 26)](https://github.com/sgreben/piecewiselinear) - Tiny linear interpolation library.
- [PiHex(stars: 20)](https://github.com/claygod/PiHex) - Implementation of the "Bailey-Borwein-Plouffe" algorithm for the hexadecimal number Pi.
- [rootfinding(stars: 11)](https://github.com/khezen/rootfinding) - root-finding algorithms library for finding roots of quadratic functions.
- [sparse(stars: 148)](https://github.com/james-bowman/sparse) - Go Sparse matrix formats for linear algebra supporting scientific and machine learning applications, compatible with gonum matrix libraries.
- [stats(stars: 2842)](https://github.com/montanaflynn/stats) - Statistics package with common functions missing from the Golang standard library.
- [streamtools(stars: 1314)](https://github.com/nytlabs/streamtools) - general purpose, graphical tool for dealing with streams of data.
- [TextRank(stars: 191)](https://github.com/DavidBelicza/TextRank) - TextRank implementation in Golang with extendable features (summarization, weighting, phrase extraction) and multithreading (goroutine) support.
- [triangolatte(stars: 36)](https://github.com/tchayen/triangolatte) - 2D triangulation library. Allows translating lines and polygons (both based on points) to the language of GPUs.

**[⬆ back to top](#contents)**

## Security

_Libraries that are used to help make your application more secure._

- [acmetool(stars: 2014)](https://github.com/hlandau/acme) - ACME (Let's Encrypt) client tool with automatic renewal.
- [acra(stars: 1244)](https://github.com/cossacklabs/acra) - Network encryption proxy to protect database-based applications from data leaks: strong selective encryption, SQL injections prevention, intrusion detection system.
- [age(stars: 14564)](https://github.com/FiloSottile/age) - A simple, modern and secure encryption tool (and Go library) with small explicit keys, no config options, and UNIX-style composability.
- [argon2-hashing(stars: 20)](https://github.com/andskur/argon2-hashing) - light wrapper around Go's argon2 package that closely mirrors with Go's standard library Bcrypt and simple-scrypt package.
- [argon2pw(stars: 90)](https://github.com/raja/argon2pw) - Argon2 password hash generation with constant-time password comparison.
- [autocert](https://godoc.org/golang.org/x/crypto/acme/autocert) - Auto provision Let's Encrypt certificates and start a TLS server.
- [BadActor(stars: 318)](https://github.com/jaredfolkins/badactor) - In-memory, application-driven jailer built in the spirit of fail2ban.
- [beelzebub(stars: 484)](https://github.com/mariocandela/beelzebub) - A secure low code honeypot framework, leveraging AI for System Virtualization.
- [booster(stars: 417)](https://github.com/anatol/booster) - Fast initramfs generator with full-disk encryption support.
- [Cameradar(stars: 3744)](https://github.com/Ullaakut/cameradar) - Tool and library to remotely hack RTSP streams from surveillance cameras.
- [certificates(stars: 36)](https://github.com/mvmaasakkers/certificates) - An opinionated tool for generating tls certificates.
- [CertMagic(stars: 4688)](https://github.com/caddyserver/certmagic) - Mature, robust, and powerful ACME client integration for fully-managed TLS certificate issuance and renewal.
- [Coraza(stars: 1567)](https://github.com/corazawaf/coraza) - Enterprise-ready, modsecurity and OWASP CRS compatible WAF library.
- [dongle(stars: 815)](https://github.com/golang-module/dongle) - A simple, semantic and developer-friendly golang package for encoding&decoding and encryption&decryption.
- [firewalld-rest(stars: 334)](https://github.com/prashantgupta24/firewalld-rest) - A rest application to dynamically update firewalld rules on a linux server.
- [go-generate-password(stars: 50)](https://github.com/m1/go-generate-password) - Password generator that can be used on the cli or as a library.
- [go-htpasswd(stars: 33)](https://github.com/tg123/go-htpasswd) - Apache htpasswd Parser for Go.
- [go-password-validator(stars: 451)](https://github.com/lane-c-wagner/go-password-validator) - Password validator based on raw cryptographic entropy values.
- [go-yara(stars: 329)](https://github.com/hillu/go-yara) - Go Bindings for [YARA](https://github.com/plusvic/yara), the "pattern matching swiss knife for malware researchers (and everyone else)".
- [goArgonPass(stars: 16)](https://github.com/dwin/goArgonPass) - Argon2 password hash and verification designed to be compatible with existing Python and PHP implementations.
- [goSecretBoxPassword(stars: 58)](https://github.com/dwin/goSecretBoxPassword) - A probably paranoid package for securely hashing and encrypting passwords.
- [Interpol(stars: 2)](https://github.com/avahidi/interpol) - Rule-based data generator for fuzzing and penetration testing.
- [lego(stars: 6631)](https://github.com/go-acme/lego) - Pure Go ACME client library and CLI tool (for use with Let's Encrypt).
- [memguard(stars: 2440)](https://github.com/awnumar/memguard) - A pure Go library for handling sensitive values in memory.
- [nacl(stars: 536)](https://github.com/kevinburke/nacl) - Go implementation of the NaCL set of API's.
- [optimus-go(stars: 350)](https://github.com/pjebs/optimus-go) - ID hashing and Obfuscation using Knuth's Algorithm.
- [passlib(stars: 287)](https://github.com/hlandau/passlib) - Futureproof password hashing library.
- [secret(stars: 24)](https://github.com/rsjethani/secret) - Prevent your secrets from leaking into logs, std\* etc.
- [secure(stars: 2176)](https://github.com/unrolled/secure) - HTTP middleware for Go that facilitates some quick security wins.
- [secureio(stars: 30)](https://github.com/xaionaro-go/secureio) - An keyexchanging+authenticating+encrypting wrapper and multiplexer for `io.ReadWriteCloser` based on XChaCha20-poly1305, ECDH and ED25519.
- [simple-scrypt(stars: 190)](https://github.com/elithrar/simple-scrypt) - Scrypt package with a simple, obvious API and automatic cost calibration built-in.
- [ssh-vault(stars: 412)](https://github.com/ssh-vault/ssh-vault) - encrypt/decrypt using ssh keys.
- [sslmgr(stars: 21)](https://github.com/adrianosela/sslmgr) - SSL certificates made easy with a high level wrapper around acme/autocert.
- [teler-waf(stars: 209)](https://github.com/kitabisa/teler-waf) - teler-waf is a Go HTTP middleware that provide teler IDS functionality to protect against web-based attacks and improve the security of Go-based web applications. It is highly configurable and easy to integrate into existing Go applications.
- [themis(stars: 1780)](https://github.com/cossacklabs/themis) - high-level cryptographic library for solving typical data security tasks (secure data storage, secure messaging, zero-knowledge proof authentication), available for 14 languages, best fit for multi-platform apps.

**[⬆ back to top](#contents)**

## Serialization

_Libraries and tools for binary serialization._

- [asn1(stars: 53)](https://github.com/PromonLogicalis/asn1) - Asn.1 BER and DER encoding library for golang.
- [bambam(stars: 67)](https://github.com/glycerine/bambam) - generator for Cap'n Proto schemas from go.
- [bel(stars: 35)](https://github.com/32leaves/bel) - Generate TypeScript interfaces from Go structs/interfaces. Useful for JSON RPC.
- [binstruct(stars: 80)](https://github.com/ghostiam/binstruct) - Golang binary decoder for mapping data into the structure.
- [cbor(stars: 631)](https://github.com/fxamacker/cbor) - Small, safe, and easy CBOR encoding and decoding library.
- [colfer(stars: 725)](https://github.com/pascaldekloe/colfer) - Code generation for the Colfer binary format.
- [csvutil(stars: 854)](https://github.com/jszwec/csvutil) - High Performance, idiomatic CSV record encoding and decoding to native Go structures.
- [elastic(stars: 24)](https://github.com/epiclabs-io/elastic) - Convert slices, maps or any other unknown value across different types at run-time, no matter what.
- [fixedwidth(stars: 9)](https://github.com/huydang284/fixedwidth) - Fixed-width text formatting (UTF-8 supported).
- [fwencoder(stars: 26)](https://github.com/o1egl/fwencoder) - Fixed width file parser (encoding and decoding library) for Go.
- [go-capnproto(stars: 287)](https://github.com/glycerine/go-capnproto) - Cap'n Proto library and parser for go.
- [go-codec(stars: 1785)](https://github.com/ugorji/go) - High Performance, feature-Rich, idiomatic encode, decode and rpc library for msgpack, cbor and json, with runtime-based OR code-generation support.
- [go-lctree(stars: 4)](https://github.com/sbourlon/go-lctree) - Provides a CLI and primitives to serialize and deserialize [LeetCode binary trees](https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation).
- [gogoprotobuf(stars: 5615)](https://github.com/gogo/protobuf) - Protocol Buffers for Go with Gadgets.
- [goprotobuf(stars: 9404)](https://github.com/golang/protobuf) - Go support, in the form of a library and protocol compiler plugin, for Google's protocol buffers.
- [gotiny(stars: 15)](https://github.com/raszia/gotiny) - Efficient Go serialization library, gotiny is almost as fast as serialization libraries that generate code.
- [jsoniter(stars: 12823)](https://github.com/json-iterator/go) - High-performance 100% compatible drop-in replacement of "encoding/json".
- [mapstructure(stars: 7363)](https://github.com/mitchellh/mapstructure) - Go library for decoding generic map values into native Go structures.
- [php_session_decoder(stars: 160)](https://github.com/yvasiyarov/php_session_decoder) - GoLang library for working with PHP session format and PHP Serialize/Unserialize functions.
- [pletter(stars: 19)](https://github.com/vimeda/pletter) - A standard way to wrap a proto message for message brokers.
- [structomap(stars: 143)](https://github.com/tuvistavie/structomap) - Library to easily and dynamically generate maps from static structures.
- [unitpacking(stars: 6)](https://github.com/recolude/unitpacking) - Library to pack unit vectors into as fewest bytes as possible.

**[⬆ back to top](#contents)**

## Server Applications

- [algernon(stars: 2516)](https://github.com/xyproto/algernon) - HTTP/2 web server with built-in support for Lua, Markdown, GCSS and Amber.
- [Caddy(stars: 51335)](https://github.com/caddyserver/caddy) - Caddy is an alternative, HTTP/2 web server that's easy to configure and use.
- [consul](https://www.consul.io/) - Consul is a tool for service discovery, monitoring and configuration.
- [cortex-tenant(stars: 82)](https://github.com/blind-oracle/cortex-tenant) - Prometheus remote write proxy that adds add Cortex tenant ID header based on metric labels.
- [devd(stars: 3376)](https://github.com/cortesi/devd) - Local webserver for developers.
- [discovery(stars: 1766)](https://github.com/Bilibili/discovery) - A registry for resilient mid-tier load balancing and failover.
- [dudeldu(stars: 144)](https://github.com/krotik/dudeldu) - A simple SHOUTcast server.
- [dummy(stars: 177)](https://github.com/neotoolkit/dummy) - Run mock server based off an API contract with one command.
- [Easegress(stars: 5599)](https://github.com/megaease/easegress) - A cloud native high availability/performance traffic orchestration system with observability and extensibility.
- [etcd(stars: 45450)](https://github.com/etcd-io/etcd) - Highly-available key value store for shared configuration and service discovery.
- [Euterpe(stars: 504)](https://github.com/ironsmile/euterpe) - Self-hosted music streaming server with built-in web UI and REST API.
- [Fider(stars: 2478)](https://github.com/getfider/fider) - Fider is an open platform to collect and organize customer feedback.
- [Flagr(stars: 2302)](https://github.com/checkr/flagr) - Flagr is an open-source feature flagging and A/B testing service.
- [flipt(stars: 3079)](https://github.com/markphelps/flipt) - A self contained feature flag solution written in Go and Vue.js
- [go-feature-flag(stars: 925)](https://github.com/thomaspoignant/go-feature-flag) - A feature flag solution, with only a YAML file in the backend (S3, GitHub, HTTP, local file ...), no server to install, just add a file in a central system and refer to it.
- [go-proxy-cache(stars: 108)](https://github.com/fabiocicerchia/go-proxy-cache) - Simple Reverse Proxy with Caching, written in Go, using Redis.
- [jackal(stars: 1440)](https://github.com/ortuman/jackal) - An XMPP server written in Go.
- [lets-proxy2(stars: 83)](https://github.com/rekby/lets-proxy2) - Reverse proxy for handle https with issue certificates in fly from lets-encrypt.
- [minio(stars: 42293)](https://github.com/minio/minio) - Minio is a distributed object storage server.
- [Moxy(stars: 12)](https://github.com/sinhashubham95/moxy) - Moxy is a simple mocker and proxy application server, you can create mock endpoints as well as proxy requests in case no mock exists for the endpoint.
- [nginx-prometheus(stars: 38)](https://github.com/blind-oracle/nginx-prometheus) - Nginx log parser and exporter to Prometheus.
- [nsq](https://nsq.io/) - A realtime distributed messaging platform.
- [protoxy(stars: 33)](https://github.com/camgraff/protoxy) - A proxy server that converts JSON request bodies to Protocol Buffers.
- [psql-streamer(stars: 53)](https://github.com/blind-oracle/psql-streamer) - Stream database events from PostgreSQL to Kafka.
- [riemann-relay(stars: 2)](https://github.com/blind-oracle/riemann-relay) - Relay to load-balance Riemann events and/or convert them to Carbon.
- [RoadRunner(stars: 7528)](https://github.com/spiral/roadrunner) - High-performance PHP application server, load-balancer and process manager.
- [SFTPGo(stars: 7339)](https://github.com/drakkan/sftpgo) - Fully featured and highly configurable SFTP server with optional FTP/S and WebDAV support. It can serve local filesystem and Cloud Storage backends such as S3 and Google Cloud Storage.
- [simple-jwt-provider(stars: 37)](https://github.com/leberKleber/simple-jwt-provider) - Simple and lightweight provider which exhibits JWTs, supports login, password-reset (via mail) and user management.
- [Trickster(stars: 1928)](https://github.com/tricksterproxy/trickster) - HTTP reverse proxy cache and time series accelerator.
- [Wish(stars: 2444)](https://github.com/charmbracelet/wish) - Make SSH apps, just like that!

**[⬆ back to top](#contents)**

## Stream Processing

_Libraries and tools for stream processing and reactive programming._

- [go-streams(stars: 1634)](https://github.com/reugn/go-streams) - Go stream processing library.
- [goio(stars: 75)](https://github.com/primetalk/goio) - An implementation of IO, Stream, Fiber for Golang, inspired by awesome Scala libraries cats and fs2.
- [machine(stars: 140)](https://github.com/whitaker-io/machine) - Go library for writing and generating stream workers with built in metrics and traceability.
- [stream(stars: 84)](https://github.com/youthlin/stream) - Go Stream, like Java 8 Stream: Filter/Map/FlatMap/Peek/Sorted/ForEach/Reduce...

**[⬆ back to top](#contents)**

## Template Engines

_Libraries and tools for templating and lexing._

- [ego(stars: 567)](https://github.com/benbjohnson/ego) - Lightweight templating language that lets you write templates in Go. Templates are translated into Go and compiled.
- [extemplate(stars: 56)](https://github.com/dannyvankooten/extemplate) - Tiny wrapper around html/template to allow for easy file-based template inheritance.
- [fasttemplate(stars: 785)](https://github.com/valyala/fasttemplate) - Simple and fast template engine. Substitutes template placeholders up to 10x faster than [text/template](https://golang.org/pkg/text/template/).
- [gospin(stars: 52)](https://github.com/m1/gospin) - Article spinning and spintax/spinning syntax engine, useful for A/B, testing pieces of text/articles and creating more natural conversations.
- [got(stars: 22)](https://github.com/goradd/got) - A Go code generator inspired by Hero and Fasttemplate. Has include files, custom tag definitions, injected Go code, language translation, and more.
- [goview(stars: 383)](https://github.com/foolin/goview) - Goview is a lightweight, minimalist and idiomatic template library based on golang html/template for building Go web application.
- [jet(stars: 1107)](https://github.com/CloudyKit/jet) - Jet template engine.
- [liquid(stars: 246)](https://github.com/osteele/liquid) - Go implementation of Shopify Liquid templates.
- [maroto(stars: 1318)](https://github.com/johnfercher/maroto) - A maroto way to create PDFs. Maroto is inspired in Bootstrap and uses gofpdf. Fast and simple.
- [pongo2(stars: 2717)](https://github.com/flosch/pongo2) - Django-like template-engine for Go.
- [quicktemplate(stars: 2939)](https://github.com/valyala/quicktemplate) - Fast, powerful, yet easy to use template engine. Converts templates into Go code and then compiles it.
- [raymond(stars: 569)](https://github.com/aymerick/raymond) - Complete handlebars implementation in Go.
- [Razor(stars: 837)](https://github.com/sipin/gorazor) - Razor view engine for Golang.
- [Soy(stars: 173)](https://github.com/robfig/soy) - Closure templates (aka Soy templates) for Go, following the [official spec](https://developers.google.com/closure/templates/).
- [sprig(stars: 3854)](https://github.com/Masterminds/sprig) - Useful template functions for Go templates.
- [tbd(stars: 25)](https://github.com/lucasepe/tbd) - A really simple way to create text templates with placeholders - exposes extra builtin Git repo metadata.
- [templ(stars: 4344)](https://github.com/a-h/templ) - A HTML templating language that has great developer tooling.

**[⬆ back to top](#contents)**

## Testing

_Libraries for testing codebases and generating test data._

- Testing Frameworks

  - [apitest](https://apitest.dev) - Simple and extensible behavioural testing library for REST based services or HTTP handlers that supports mocking external http calls and rendering of sequence diagrams.
  - [assert(stars: 62)](https://github.com/go-playground/assert) - Basic Assertion Library used along side native go testing, with building blocks for custom assertions.
  - [badio(stars: 11)](https://github.com/cavaliercoder/badio) - Extensions to Go's `testing/iotest` package.
  - [baloo(stars: 770)](https://github.com/h2non/baloo) - Expressive and versatile end-to-end HTTP API testing made easy.
  - [be(stars: 84)](https://github.com/carlmjohnson/be) - The minimalist generic test assertion library.
  - [biff(stars: 13)](https://github.com/fulldump/biff) - Bifurcation testing framework, BDD compatible.
  - [charlatan(stars: 199)](https://github.com/percolate/charlatan) - Tool to generate fake interface implementations for tests.
  - [commander(stars: 217)](https://github.com/SimonBaeumer/commander) - Tool for testing cli applications on windows, linux and osx.
  - [covergates(stars: 59)](https://github.com/covergates/covergates) - Self-hosted code coverage report review and management service.
  - [cupaloy(stars: 282)](https://github.com/bradleyjkemp/cupaloy) - Simple snapshot testing addon for your test framework.
  - [dbcleaner(stars: 157)](https://github.com/khaiql/dbcleaner) - Clean database for testing purpose, inspired by `database_cleaner` in Ruby.
  - [dsunit(stars: 43)](https://github.com/viant/dsunit) - Datastore testing for SQL, NoSQL, structured files.
  - [embedded-postgres(stars: 687)](https://github.com/fergusstrange/embedded-postgres) - Run a real Postgres database locally on Linux, OSX or Windows as part of another Go application or test.
  - [endly(stars: 247)](https://github.com/viant/endly) - Declarative end to end functional testing.
  - [fixenv(stars: 26)](https://github.com/rekby/fixenv) - Fixture manage engine, inspired by pytest fixtures.
  - [fluentassert(stars: 36)](https://github.com/fluentassert/verify) - Extensible, type-safe, fluent assertion Go library.
  - [flute(stars: 18)](https://github.com/suzuki-shunsuke/flute) - HTTP client testing framework.
  - [frisby(stars: 275)](https://github.com/verdverm/frisby) - REST API testing framework.
  - [gherkingen(stars: 66)](https://github.com/hedhyw/gherkingen) - BDD boilerplate generator and framework.
  - [ginkgo](https://onsi.github.io/ginkgo/) - BDD Testing Framework for Go.
  - [gnomock(stars: 1274)](https://github.com/orlangure/gnomock) - integration testing with real dependencies (database, cache, even Kubernetes or AWS) running in Docker, without mocks.
  - [go-carpet(stars: 240)](https://github.com/msoap/go-carpet) - Tool for viewing test coverage in terminal.
  - [go-cmp(stars: 3841)](https://github.com/google/go-cmp) - Package for comparing Go values in tests.
  - [go-hit(stars: 242)](https://github.com/Eun/go-hit) - Hit is an http integration test framework written in golang.
  - [go-mutesting(stars: 614)](https://github.com/zimmski/go-mutesting) - Mutation testing for Go source code.
  - [go-mysql-test-container(stars: 2)](https://github.com/arikama/go-mysql-test-container) - Golang MySQL testcontainer to help with MySQL integration testing.
  - [go-snaps](http://github.com/gkampitakis/go-snaps) - Jest-like snapshot testing in Golang.
  - [go-testdeep(stars: 404)](https://github.com/maxatome/go-testdeep) - Extremely flexible golang deep comparison, extends the go testing package.
  - [go-testpredicate(stars: 5)](https://github.com/maargenton/go-testpredicate) - Test predicate style assertions library with extensive diagnostics output.
  - [go-vcr(stars: 1064)](https://github.com/dnaeon/go-vcr) - Record and replay your HTTP interactions for fast, deterministic and accurate tests.
  - [goblin(stars: 888)](https://github.com/franela/goblin) - Mocha like testing framework of Go.
  - [goc(stars: 738)](https://github.com/qiniu/goc) - Goc is a comprehensive coverage testing system for The Go Programming Language.
  - [gocheck](https://labix.org/gocheck) - More advanced testing framework alternative to gotest.
  - [GoConvey](https://github.com/smartystreets/goconvey/) - BDD-style framework with web UI and live reload.
  - [gocrest(stars: 97)](https://github.com/corbym/gocrest) - Composable hamcrest-like matchers for Go assertions.
  - [godog(stars: 2115)](https://github.com/cucumber/godog) - Cucumber BDD framework for Go.
  - [gofight(stars: 437)](https://github.com/appleboy/gofight) - API Handler Testing for Golang Router framework.
  - [gogiven(stars: 14)](https://github.com/corbym/gogiven) - YATSPEC-like BDD testing framework for Go.
  - [gomatch(stars: 45)](https://github.com/jfilipczyk/gomatch) - library created for testing JSON against patterns.
  - [gomega](https://onsi.github.io/gomega/) - Rspec like matcher/assertion library.
  - [Gont(stars: 63)](https://github.com/stv0g/gont) - Go network testing toolkit for testing building complex network topologies using Linux namespaces.
  - [GoSpec(stars: 114)](https://github.com/orfjackal/gospec) - BDD-style testing framework for the Go programming language.
  - [gospecify(stars: 53)](https://github.com/stesla/gospecify) - This provides a BDD syntax for testing your Go code. It should be familiar to anybody who has used libraries such as rspec.
  - [gosuite(stars: 12)](https://github.com/pavlo/gosuite) - Brings lightweight test suites with setup/teardown facilities to `testing` by leveraging Go1.7's Subtests.
  - [got(stars: 264)](https://github.com/ysmood/got) - An enjoyable golang test framework.
  - [gotest.tools](https://github.com/gotestyourself/gotest.tools) - A collection of packages to augment the go testing package and support common patterns.
  - [Hamcrest(stars: 29)](https://github.com/rdrdr/hamcrest) - fluent framework for declarative Matcher objects that, when applied to input values, produce self-describing results.
  - [httpexpect(stars: 2387)](https://github.com/gavv/httpexpect) - Concise, declarative, and easy to use end-to-end HTTP and REST API testing.
  - [is(stars: 1657)](https://github.com/matryer/is) - Professional lightweight testing mini-framework for Go.
  - [jsonassert(stars: 115)](https://github.com/kinbiko/jsonassert) - Package for verifying that your JSON payloads are serialized correctly.
  - [omg.testingtools](https://github.com/dedalqq/omg.testingtools) - The simple library for change a values of private fields for testing.
  - [restit(stars: 55)](https://github.com/yookoala/restit) - Go micro framework to help writing RESTful API integration test.
  - [schema(stars: 20)](https://github.com/jgroeneveld/schema) - Quick and easy expression matching for JSON schemas used in requests and responses.
  - [stop-and-go(stars: 10)](https://github.com/elgohr/stop-and-go) - Testing helper for concurrency.
  - [testcase(stars: 114)](https://github.com/adamluzsi/testcase) - Idiomatic testing framework for Behavior Driven Development.
  - [testcontainers-go(stars: 2654)](https://github.com/testcontainers/testcontainers-go) - A Go package that makes it simple to create and clean up container-based dependencies for automated integration/smoke tests. The clean, easy-to-use API enables developers to programmatically define containers that should be run as part of a test and clean up those resources when the test is done.
  - [testfixtures(stars: 999)](https://github.com/go-testfixtures/testfixtures) - A helper for Rails' like test fixtures to test database applications.
  - [Testify(stars: 21208)](https://github.com/stretchr/testify) - Sacred extension to the standard go testing package.
  - [testmd](https://godoc.org/github.com/tvastar/test/cmd/testmd) - Convert markdown snippets into testable go code.
  - [testsql(stars: 17)](https://github.com/zhulongcheng/testsql) - Generate test data from SQL files before testing and clear it after finished.
  - [testza(stars: 414)](https://github.com/MarvinJWendt/testza) - Full-featured test framework with nice colorized output.
  - [trial(stars: 6)](https://github.com/jgroeneveld/trial) - Quick and easy extendable assertions without introducing much boilerplate.
  - [Tt(stars: 6)](https://github.com/vcaesar/tt) - Simple and colorful test tools.
  - [wstest(stars: 101)](https://github.com/posener/wstest) - Websocket client for unit-testing a websocket http.Handler.

- Mock

  - [counterfeiter(stars: 893)](https://github.com/maxbrunsfeld/counterfeiter) - Tool for generating self-contained mock objects.
  - [genmock](https://gitlab.com/so_literate/genmock) - Go mocking system with code generator for building calls of the interface methods.
  - [go-localstack(stars: 73)](https://github.com/elgohr/go-localstack) - Tool for using localstack in AWS testing.
  - [go-sqlmock(stars: 5626)](https://github.com/DATA-DOG/go-sqlmock) - Mock SQL driver for testing database interactions.
  - [go-txdb(stars: 589)](https://github.com/DATA-DOG/go-txdb) - Single transaction based database driver mainly for testing purposes.
  - [gock(stars: 1982)](https://github.com/h2non/gock) - Versatile HTTP mocking made easy.
  - [gomock(stars: 9154)](https://github.com/golang/mock) - Mocking framework for the Go programming language.
  - [govcr(stars: 161)](https://github.com/seborama/govcr) - HTTP mock for Golang: record and replay HTTP interactions for offline testing.
  - [hoverfly(stars: 2283)](https://github.com/SpectoLabs/hoverfly) - HTTP(S) proxy for recording and simulating REST/SOAP APIs with extensible middleware and easy-to-use CLI.
  - [httpmock(stars: 1810)](https://github.com/jarcoal/httpmock) - Easy mocking of HTTP responses from external resources.
  - [minimock(stars: 523)](https://github.com/gojuno/minimock) - Mock generator for Go interfaces.
  - [mockery(stars: 5311)](https://github.com/vektra/mockery) - Tool to generate Go interfaces.
  - [mockhttp(stars: 23)](https://github.com/tv42/mockhttp) - Mock object for Go http.ResponseWriter.
  - [mockit(stars: 16)](https://github.com/pasdam/mockit) - Allows functions and method easy mocking, without defining new types; it's similar to Mockito for Java.
  - [monkey(stars: 2)](https://github.com/awterman/monkey) - One line to mock functions/methods/variables in place without dependency injection or code generation.
  - [mooncake(stars: 18)](https://github.com/GuilhermeCaruso/mooncake) - A simple way to generate mocks for multiple purposes
  - [timex(stars: 70)](https://github.com/cabify/timex) - A test-friendly replacement for the native `time` package.

- Fuzzing and delta-debugging/reducing/shrinking.

  - [go-fuzz(stars: 4681)](https://github.com/dvyukov/go-fuzz) - Randomized testing system.
  - [gofuzz(stars: 1476)](https://github.com/google/gofuzz) - Library for populating go objects with random values.
  - [Tavor(stars: 241)](https://github.com/zimmski/tavor) - Generic fuzzing and delta-debugging framework.

- Selenium and browser control tools.

  - [cdp(stars: 702)](https://github.com/mafredri/cdp) - Type-safe bindings for the Chrome Debugging Protocol that can be used with browsers or other debug targets that implement it.
  - [chromedp(stars: 9876)](https://github.com/knq/chromedp) - a way to drive/test Chrome, Safari, Edge, Android Webviews, and other browsers supporting the Chrome Debugging Protocol.
  - [ggr(stars: 303)](https://github.com/aerokube/ggr) - a lightweight server that routes and proxies Selenium WebDriver requests to multiple Selenium hubs.
  - [playwright-go(stars: 1582)](https://github.com/mxschmitt/playwright-go) - browser automation library to control Chromium, Firefox and WebKit with a single API.
  - [rod(stars: 4438)](https://github.com/go-rod/rod) - A Devtools driver to make web automation and scraping easy.
  - [selenoid(stars: 2432)](https://github.com/aerokube/selenoid) - alternative Selenium hub server that launches browsers within containers.

- Fail injection
  - [failpoint(stars: 794)](https://github.com/pingcap/failpoint) - An implementation of [failpoints](https://www.freebsd.org/cgi/man.cgi?query=fail) for Golang.

**[⬆ back to top](#contents)**

## Text Processing

_Libraries for parsing and manipulating texts._

See also [Natural Language Processing](#natural-language-processing) and [Text Analysis](#text-analysis).

### Formatters

- [address(stars: 69)](https://github.com/bojanz/address) - Handles address representation, validation and formatting.
- [align(stars: 83)](https://github.com/Guitarbum722/align) - A general purpose application that aligns text.
- [bytes](https://github.com/labstack/gommon/tree/master/bytes) - Formats and parses numeric byte values (10K, 2M, 3G, etc.).
- [go-fixedwidth(stars: 79)](https://github.com/ianlopshire/go-fixedwidth) - Fixed-width text formatting (encoder/decoder with reflection).
- [go-humanize(stars: 3882)](https://github.com/dustin/go-humanize) - Formatters for time, numbers, and memory size to human readable format.
- [gotabulate(stars: 305)](https://github.com/bndr/gotabulate) - Easily pretty-print your tabular data with Go.
- [textwrap(stars: 5)](https://github.com/isbm/textwrap) - Wraps text at end of lines. Implementation of `textwrap` module from Python.

### Markup Languages

- [bafi(stars: 89)](https://github.com/mmalcek/bafi) - Universal JSON, BSON, YAML, XML translator to ANY format using templates.
- [bbConvert(stars: 7)](https://github.com/CalebQ42/bbConvert) - Converts bbCode to HTML that allows you to add support for custom bbCode tags.
- [blackfriday(stars: 5281)](https://github.com/russross/blackfriday) - Markdown processor in Go.
- [github_flavored_markdown](https://godoc.org/github.com/shurcooL/github_flavored_markdown) - GitHub Flavored Markdown renderer (using blackfriday) with fenced code block highlighting, clickable header anchor links.
- [go-output-format(stars: 13)](https://github.com/drewstinnett/go-output-format) - Output go structures into multiple formats (YAML/JSON/etc) in your command line app.
- [go-toml(stars: 1535)](https://github.com/pelletier/go-toml) - Go library for the TOML format with query support and handy cli tools.
- [goldmark(stars: 3064)](https://github.com/yuin/goldmark) - A Markdown parser written in Go. Easy to extend, standard (CommonMark) compliant, well structured.
- [goq(stars: 247)](https://github.com/andrewstuart/goq) - Declarative unmarshalling of HTML using struct tags with jQuery syntax (uses GoQuery).
- [html-to-markdown(stars: 641)](https://github.com/JohannesKaufmann/html-to-markdown) - Convert HTML to Markdown. Even works with entire websites and can be extended through rules.
- [htmlquery(stars: 657)](https://github.com/antchfx/htmlquery) - An XPath query package for HTML, lets you extract data or evaluate from HTML documents by an XPath expression.
- [htmlyaml(stars: 3)](https://github.com/nikolaydubina/htmlyaml) -  Rich rendering of YAML as HTML in Go
- [mxj(stars: 591)](https://github.com/clbanning/mxj) - Encode / decode XML as JSON or map[string]interface{}; extract values with dot-notation paths and wildcards. Replaces x2j and j2x packages.
- [toml(stars: 4349)](https://github.com/BurntSushi/toml) - TOML configuration format (encoder/decoder with reflection).

### Parsers/Encoders/Decoders

- [allot(stars: 58)](https://github.com/sbstjn/allot) - Placeholder and wildcard text parsing for CLI tools and bots.
- [codetree(stars: 23)](https://github.com/aerogo/codetree) - Parses indented code (python, pixy, scarlet, etc.) and returns a tree structure.
- [commonregex(stars: 874)](https://github.com/mingrammer/commonregex) - A collection of common regular expressions for Go.
- [did(stars: 80)](https://github.com/ockam-network/did) - DID (Decentralized Identifiers) Parser and Stringer in Go.
- [doi(stars: 8)](https://github.com/hscells/doi) - Document object identifier (doi) parser in Go.
- [editorconfig-core-go(stars: 125)](https://github.com/editorconfig/editorconfig-core-go) - Editorconfig file parser and manipulator for Go.
- [encdec(stars: 8)](https://github.com/mickep76/encdec) - Package provides a generic interface to encoders and decoders.
- [go-fasttld(stars: 28)](https://github.com/elliotwutingfeng/go-fasttld) - High performance effective top level domains (eTLD) extraction module.
- [go-nmea(stars: 202)](https://github.com/adrianmo/go-nmea) - NMEA parser library for the Go language.
- [go-vcard(stars: 95)](https://github.com/emersion/go-vcard) - Parse and format vCard.
- [gofeed(stars: 2317)](https://github.com/mmcdole/gofeed) - Parse RSS and Atom feeds in Go.
- [gographviz(stars: 533)](https://github.com/awalterschulze/gographviz) - Parses the Graphviz DOT language.
- [gonameparts(stars: 39)](https://github.com/polera/gonameparts) - Parses human names into individual name parts.
- [ltsv(stars: 8)](https://github.com/Wing924/ltsv) - High performance [LTSV (Labeled Tab Separated Value)](http://ltsv.org/) reader for Go.
- [normalize(stars: 37)](https://github.com/avito-tech/normalize) - Sanitize, normalize and compare fuzzy text.
- [omniparser(stars: 603)](https://github.com/jf-tech/omniparser) - A versatile ETL library that parses text input (CSV/txt/JSON/XML/EDI/X12/EDIFACT/etc) in streaming fashion and transforms data into JSON output using data-driven schema.
- [parseargs-go(stars: 10)](https://github.com/nproc/parseargs-go) - string argument parser that understands quotes and backslashes.
- [parth(stars: 46)](https://github.com/codemodus/parth) - URL path segmentation parsing.
- [sdp(stars: 112)](https://github.com/gortc/sdp) - SDP: Session Description Protocol [[RFC 4566](https://tools.ietf.org/html/rfc4566)].
- [sh(stars: 6408)](https://github.com/mvdan/sh) - Shell parser and formatter.
- [tokenizer(stars: 70)](https://github.com/bzick/tokenizer) - Parse any string, slice or infinite buffer to any tokens.
- [when(stars: 1303)](https://github.com/olebedev/when) - Natural EN and RU language date/time parser with pluggable rules.
- [xj2go(stars: 34)](https://github.com/stackerzzq/xj2go) - Convert xml or json to go struct.

### Regular Expressions

- [genex(stars: 76)](https://github.com/alixaxel/genex) - Count and expand Regular Expressions into all matching Strings.
- [go-wildcard(stars: 59)](https://github.com/IGLOU-EU/go-wildcard) - Simple and lightweight wildcard pattern matching.
- [goregen(stars: 89)](https://github.com/zach-klippenstein/goregen) - Library for generating random strings from regular expressions.
- [regroup(stars: 137)](https://github.com/oriser/regroup) - Match regex expression named groups into go struct using struct tags and automatic parsing.
- [rex(stars: 172)](https://github.com/hedhyw/rex) - Regular expressions builder.

### Sanitation

- [bluemonday(stars: 2869)](https://github.com/microcosm-cc/bluemonday) - HTML Sanitizer.
- [gofuckyourself(stars: 59)](https://github.com/JoshuaDoes/gofuckyourself) - A sanitization-based swear filter for Go.

### Scrapers

- [colly(stars: 21398)](https://github.com/asciimoo/colly) - Fast and Elegant Scraping Framework for Gophers.
- [dataflowkit(stars: 626)](https://github.com/slotix/dataflowkit) - Web scraping Framework to turn websites into structured data.
- [go-recipe(stars: 25)](https://github.com/kkyr/go-recipe) - A package for scraping recipes from websites.
- [GoQuery(stars: 13232)](https://github.com/PuerkitoBio/goquery) - GoQuery brings a syntax and a set of features similar to jQuery to the Go language.
- [gospider(stars: 202)](https://github.com/zhshch2002/gospider) - A simple golang spider/scraping framework,build a spider in 3 lines. migrated from [goribot](https://github.com/zhshch2002/goribot)
- [pagser(stars: 93)](https://github.com/foolin/pagser) - Pagser is a simple, extensible, configurable parse and deserialize html page to struct based on goquery and struct tags for golang crawler.
- [Tagify(stars: 34)](https://github.com/zoomio/tagify) - Produces a set of tags from given source.
- [walker(stars: 7)](https://github.com/cyucelen/walker) - Seamlessly fetch paginated data from any source. Simple and high performance API scraping included.
- [xurls(stars: 1102)](https://github.com/mvdan/xurls) - Extract urls from text.

### RSS

- [podcast(stars: 126)](https://github.com/eduncan911/podcast) - iTunes Compliant and RSS 2.0 Podcast Generator in Golang

### Utility/Miscellaneous

- [go-runewidth(stars: 565)](https://github.com/mattn/go-runewidth) - Functions to get fixed width of the character or string.
- [go-zero-width(stars: 111)](https://github.com/trubitsyn/go-zero-width) - Zero-width character detection and removal for Go.
- [kace(stars: 20)](https://github.com/codemodus/kace) - Common case conversions covering common initialisms.
- [petrovich(stars: 44)](https://github.com/striker2000/petrovich) - Petrovich is the library which inflects Russian names to given grammatical case.
- [radix(stars: 186)](https://github.com/yourbasic/radix) - Fast string sorting algorithm.
- [TySug(stars: 17)](https://github.com/Dynom/TySug) - Alternative suggestions with respect to keyboard layouts.

**[⬆ back to top](#contents)**

## Third-party APIs

_Libraries for accessing third party APIs._

- [airtable(stars: 68)](https://github.com/mehanizm/airtable) - Go client library for the [Airtable API](https://airtable.com/api).
- [amazon-product-advertising-api(stars: 57)](https://github.com/ngs/go-amazon-product-advertising-api) - Go Client Library for [Amazon Product Advertising API](https://affiliate-program.amazon.com/gp/advertising/api/detail/main.html).
- [anaconda(stars: 1142)](https://github.com/ChimeraCoder/anaconda) - Go client library for the Twitter 1.1 API.
- [appstore-sdk-go(stars: 5)](https://github.com/Kachit/appstore-sdk-go) - Unofficial Golang SDK for AppStore Connect API.
- [aws-sdk-go(stars: 2242)](https://github.com/aws/aws-sdk-go-v2) - The official AWS SDK for the Go programming language.
- [bqwriter(stars: 13)](https://github.com/OTA-Insight/bqwriter) - High Level Go Library to write data into [Google BigQuery](https://cloud.google.com/bigquery) at a high throughout.
- [brewerydb(stars: 19)](https://github.com/naegelejd/brewerydb) - Go library for accessing the BreweryDB API.
- [cachet(stars: 90)](https://github.com/andygrunwald/cachet) - Go client library for [Cachet (open source status page system)](https://cachethq.io/).
- [circleci(stars: 66)](https://github.com/jszwedko/go-circleci) - Go client library for interacting with CircleCI's API.
- [clarifai(stars: 55)](https://github.com/samuelcouch/clarifai) - Go client library for interfacing with the Clarifai API.
- [codeship-go(stars: 18)](https://github.com/codeship/codeship-go) - Go client library for interacting with Codeship's API v2.
- [coinpaprika-go(stars: 23)](https://github.com/coinpaprika/coinpaprika-api-go-client) - Go client library for interacting with Coinpaprika's API.
- [device-check-go(stars: 23)](https://github.com/rinchsan/device-check-go) - Go client library for interacting with [iOS DeviceCheck API](https://developer.apple.com/documentation/devicecheck) v1.
- [discordgo(stars: 4450)](https://github.com/bwmarrin/discordgo) - Go bindings for the Discord Chat API.
- [dusupay-sdk-go(stars: 3)](https://github.com/Kachit/dusupay-sdk-go) - Unofficial Dusupay payment gateway API Client for Go
- [ethrpc(stars: 260)](https://github.com/onrik/ethrpc) - Go bindings for Ethereum JSON RPC API.
- [facebook(stars: 1235)](https://github.com/huandu/facebook) - Go Library that supports the Facebook Graph API.
- [fasapay-sdk-go(stars: 2)](https://github.com/Kachit/fasapay-sdk-go) - Unofficial Fasapay payment gateway XML API Client for Golang.
- [fcm(stars: 51)](https://github.com/maddevsio/fcm) - Go library for Firebase Cloud Messaging.
- [gads(stars: 48)](https://github.com/emiddleton/gads) - Google Adwords Unofficial API.
- [gami(stars: 32)](https://github.com/bit4bit/gami) - Go library for Asterisk Manager Interface.
- [gcm(stars: 31)](https://github.com/Aorioli/gcm) - Go library for Google Cloud Messaging.
- [geo-golang(stars: 491)](https://github.com/codingsince1985/geo-golang) - Go Library to access [Google Maps](https://developers.google.com/maps/documentation/geocoding/intro), [MapQuest](https://developer.mapquest.com/documentation/geocoding-api/), [Nominatim](https://developer.mapquest.com/documentation/open/nominatim-search), [OpenCage](https://opencagedata.com/api), [Bing](https://msdn.microsoft.com/en-us/library/ff701715.aspx), [Mapbox](https://www.mapbox.com/developers/api/geocoding/), and [OpenStreetMap](https://wiki.openstreetmap.org/wiki/Nominatim) geocoding / reverse geocoding APIs.
- [github(stars: 9857)](https://github.com/google/go-github) - Go library for accessing the GitHub REST API v3.
- [githubql(stars: 1057)](https://github.com/shurcooL/githubql) - Go library for accessing the GitHub GraphQL API v4.
- [go-atlassian(stars: 96)](https://github.com/ctreminiom/go-atlassian) - Go library for accessing the [Atlassian Cloud](https://www.atlassian.com/enterprise/cloud) services (Jira, Jira Service Management, Jira Agile, Confluence, Admin Cloud)
- [go-aws-news(stars: 18)](https://github.com/circa10a/go-aws-news) - Go application and library to fetch what's new from AWS.
- [go-chronos(stars: 8)](https://github.com/axelspringer/go-chronos) - Go library for interacting with the [Chronos](https://mesos.github.io/chronos/) Job Scheduler
- [go-hacknews(stars: 17)](https://github.com/PaulRosset/go-hacknews) - Tiny Go client for HackerNews API.
- [go-here(stars: 13)](https://github.com/abdullahselek/go-here) - Go client library around the HERE location based APIs.
- [go-hibp(stars: 8)](https://github.com/wneessen/go-hibp) - Simple Go binding to the "Have I Been Pwned" APIs.
- [go-imgur(stars: 25)](https://github.com/koffeinsource/go-imgur) - Go client library for [imgur](https://imgur.com)
- [go-jira(stars: 1386)](https://github.com/andygrunwald/go-jira) - Go client library for [Atlassian JIRA](https://www.atlassian.com/software/jira)
- [go-lark(stars: 172)](https://github.com/go-lark/lark) - An easy-to-use unofficial SDK for [Feishu](https://open.feishu.cn/) and [Lark](https://open.larksuite.com/) Open Platform.
- [go-marathon(stars: 200)](https://github.com/gambol99/go-marathon) - Go library for interacting with Mesosphere's Marathon PAAS.
- [go-myanimelist(stars: 37)](https://github.com/nstratos/go-myanimelist) - Go client library for accessing the [MyAnimeList API](https://myanimelist.net/apiconfig/references/api/v2).
- [go-openai(stars: 7399)](https://github.com/sashabaranov/go-openai) - OpenAI ChatGPT, DALL·E, Whisper API library for Go.
- [go-openproject(stars: 13)](https://github.com/manuelbcd/go-openproject) - Go client library for interacting with [OpenProject](https://docs.openproject.org/api/) API.
- [go-postman-collection(stars: 71)](https://github.com/rbretecher/go-postman-collection) - Go module to work with [Postman Collections](https://learning.getpostman.com/docs/postman/collections/creating-collections/) (compatible with Insomnia).
- [go-redoc(stars: 53)](https://github.com/mvrilo/go-redoc) - Embedded OpenAPI/Swagger documentation ui for Go using [ReDoc](https://redocly.com/).
- [go-restcountries(stars: 4)](https://github.com/chriscross0/go-restcountries) - Go library for the [REST Countries API](https://countrylayer.com/).
- [go-sophos(stars: 12)](https://github.com/esurdam/go-sophos) - Go client library for the [Sophos UTM REST API](https://www.sophos.com/en-us/medialibrary/PDFs/documentation/UTMonAWS/Sophos-UTM-RESTful-API.pdf?la=en) with zero dependencies.
- [go-sptrans(stars: 8)](https://github.com/sergioaugrod/go-sptrans) - Go client library for the SPTrans Olho Vivo API.
- [go-swagger-ui(stars: 11)](https://github.com/esurdam/go-swagger-ui) - Go library containing precompiled [Swagger UI](https://swagger.io/tools/swagger-ui/) for serving swagger json.
- [go-telegraph](https://gitlab.com/toby3d/telegraph) - Telegraph publishing platform API client.
- [go-trending(stars: 138)](https://github.com/andygrunwald/go-trending) - Go library for accessing [trending repositories](https://github.com/trending) and [developers](https://github.com/trending/developers) at Github.
- [go-twitter(stars: 1588)](https://github.com/dghubble/go-twitter) - Go client library for the Twitter v1.1 APIs.
- [go-unsplash(stars: 72)](https://github.com/hbagdi/go-unsplash) - Go client library for the [Unsplash.com](https://unsplash.com) API.
- [go-xkcd(stars: 51)](https://github.com/nishanths/go-xkcd) - Go client for the xkcd API.
- [go-yapla](https://git.iglou.eu/Production/go-yapla) - Go client library for the Yapla v2.0 API.
- [goagi(stars: 9)](https://github.com/staskobzar/goagi) - Go library to build Asterisk PBX agi/fastagi applications.
- [goami2(stars: 14)](https://github.com/staskobzar/goami2) - AMI v2 library for Asterisk PBX.
- [GoFreeDB(stars: 29)](https://github.com/FreeLeh/GoFreeDB) - Golang library providing common and simple database abstractions on top of Google Sheets.
- [gogtrends(stars: 75)](https://github.com/groovili/gogtrends) - Google Trends Unofficial API.
- [golang-tmdb(stars: 88)](https://github.com/cyruzin/golang-tmdb) - Golang wrapper for The Movie Database API v3.
- [golyrics(stars: 40)](https://github.com/mamal72/golyrics) - Golyrics is a Go library to fetch music lyrics data from the Wikia website.
- [gomalshare(stars: 12)](https://github.com/MonaxGT/gomalshare) - Go library MalShare API [malshare.com](https://www.malshare.com/)
- [GoMusicBrainz(stars: 52)](https://github.com/michiwend/gomusicbrainz) - Go MusicBrainz WS2 client library.
- [google(stars: 3664)](https://github.com/google/google-api-go-client) - Auto-generated Google APIs for Go.
- [google-analytics(stars: 14)](https://github.com/chonthu/go-google-analytics) - Simple wrapper for easy google analytics reporting.
- [google-cloud(stars: 3479)](https://github.com/GoogleCloudPlatform/gcloud-golang) - Google Cloud APIs Go Client Library.
- [google-email-audit-api(stars: 8)](https://github.com/ngs/go-google-email-audit-api) - Go client library for [Google G Suite Email Audit API](https://developers.google.com/admin-sdk/email-audit/).
- [google-play-scraper(stars: 77)](https://github.com/n0madic/google-play-scraper) - Get data from Google Play Store.
- [gopaapi5(stars: 15)](https://github.com/utekaravinash/gopaapi5) - Go Client Library for [Amazon Product Advertising API 5.0](https://webservices.amazon.com/paapi5/documentation/).
- [gosip(stars: 124)](https://github.com/koltyakov/gosip) - Client library for SharePoint.
- [gostorm(stars: 129)](https://github.com/jsgilmore/gostorm) - GoStorm is a Go library that implements the communications protocol required to write Storm spouts and Bolts in Go that communicate with the Storm shells.
- [hipchat(stars: 105)](https://github.com/andybons/hipchat) - This project implements a golang client library for the Hipchat API.
- [hipchat (xmpp)(stars: 110)](https://github.com/daneharrigan/hipchat) - A golang package to communicate with HipChat over XMPP.
- [igdb(stars: 79)](https://github.com/Henry-Sarabia/igdb) - Go client for the [Internet Game Database API](https://api.igdb.com/).
- [jokeapi-go(stars: 22)](https://github.com/icelain/jokeapi) - Go client for [JokeAPI](https://sv443.net/jokeapi/v2/).
- [lark(stars: 363)](https://github.com/chyroc/lark) - [Feishu](https://open.feishu.cn/)/[Lark](https://open.larksuite.com/) Open API Go SDK, Support ALL Open API and Event Callback.
- [lastpass-go(stars: 33)](https://github.com/ansd/lastpass-go) - Go client library for the [LastPass](https://www.lastpass.com/) API.
- [libgoffi(stars: 9)](https://github.com/clevabit/libgoffi) - Library adapter toolbox for native [libffi](https://sourceware.org/libffi/) integration
- [Medium(stars: 139)](https://github.com/Medium/medium-sdk-go) - Golang SDK for Medium's OAuth2 API.
- [megos(stars: 54)](https://github.com/andygrunwald/megos) - Client library for accessing an [Apache Mesos](https://mesos.apache.org/) cluster.
- [minio-go(stars: 2138)](https://github.com/minio/minio-go) - Minio Go Library for Amazon S3 compatible cloud storage.
- [mixpanel(stars: 60)](https://github.com/dukex/mixpanel) - Mixpanel is a library for tracking events and sending Mixpanel profile updates to Mixpanel from your go applications.
- [newsapi-go(stars: 6)](https://github.com/jellydator/newsapi-go) - Go client for [NewsAPI](https://newsapi.org/).
- [openaigo(stars: 280)](https://github.com/otiai10/openaigo) - OpenAI GPT3/GPT3.5 ChatGPT API client library for Go.
- [patreon-go(stars: 38)](https://github.com/mxpv/patreon-go) - Go library for Patreon API.
- [paypal(stars: 621)](https://github.com/logpacker/PayPal-Go-SDK) - Wrapper for PayPal payment API.
- [playlyfe(stars: 2)](https://github.com/playlyfe/playlyfe-go-sdk) - The Playlyfe Rest API Go SDK.
- [pushover(stars: 139)](https://github.com/gregdel/pushover) - Go wrapper for the Pushover API.
- [rawg-sdk-go(stars: 9)](https://github.com/dimuska139/rawg-sdk-go) - Go library for the [RAWG Video Games Database](https://rawg.io/) API
- [rrdaclient(stars: 9)](https://github.com/Omie/rrdaclient) - Go Library to access statdns.com API, which is in turn RRDA API. DNS Queries over HTTP.
- [shopify(stars: 25)](https://github.com/rapito/go-shopify) - Go Library to make CRUD request to the Shopify API.
- [simples3(stars: 119)](https://github.com/rhnvrm/simples3) - Simple no frills AWS S3 Library using REST with V4 Signing written in Go.
- [slack(stars: 4467)](https://github.com/slack-go/slack) - Slack API in Go.
- [smite(stars: 11)](https://github.com/sergiotapia/smitego) - Go package to wraps access to the Smite game API.
- [spotify(stars: 48)](https://github.com/rapito/go-spotify) - Go Library to access Spotify WEB API.
- [steam(stars: 31)](https://github.com/sostronk/go-steam) - Go Library to interact with Steam game servers.
- [stripe(stars: 1922)](https://github.com/stripe/stripe-go) - Go client for the Stripe API.
- [swag(stars: 43)](https://github.com/zc2638/swag) - No comments, simple go wrapper to create swagger 2.0 compatible APIs. Support most routing frameworks, such as built-in, gin, chi, mux, echo, httprouter, fasthttp and more.
- [textbelt(stars: 19)](https://github.com/dietsche/textbelt) - Go client for the textbelt.com txt messaging API.
- [translate(stars: 33)](https://github.com/poorny/translate) - Go online translation package.
- [Trello(stars: 215)](https://github.com/adlio/trello) - Go wrapper for the Trello API.
- [TripAdvisor(stars: 2)](https://github.com/mrbenosborne/tripadvisor-golang) - Go wrapper for the TripAdvisor API.
- [tumblr(stars: 9)](https://github.com/mattcunningham/gumblr) - Go wrapper for the Tumblr v2 API.
- [twitter-scraper(stars: 810)](https://github.com/n0madic/twitter-scraper) - Scrape the Twitter Frontend API without authentication and limits.
- [uptimerobot(stars: 57)](https://github.com/bitfield/uptimerobot) - Go wrapper and command-line client for the Uptime Robot v2 API.
- [vl-go(stars: 2)](https://github.com/verifid/vl-go) - Go client library around the VerifID identity verification layer API.
- [webhooks(stars: 894)](https://github.com/go-playground/webhooks) - Webhook receiver for GitHub and Bitbucket.
- [wit-go(stars: 148)](https://github.com/wit-ai/wit-go) - Go client for wit.ai HTTP API.
- [ynab](https://github.com/brunomvsouza/ynab.go) - Go wrapper for the YNAB API.
- [zooz(stars: 7)](https://github.com/gojuno/go-zooz) - Go client for the Zooz API.

**[⬆ back to top](#contents)**

## Utilities

_General utilities and tools to make your life easier._

- [air(stars: 13308)](https://github.com/cosmtrek/air) - Air - Live reload for Go apps.
- [apm(stars: 164)](https://github.com/topfreegames/apm) - Process manager for Golang applications with an HTTP API.
- [backscanner(stars: 55)](https://github.com/icza/backscanner) - A scanner similar to bufio.Scanner, but it reads and returns lines in reverse order, starting at a given position and going backward.
- [beyond(stars: 56)](https://github.com/wesovilabs/beyond) - The Go tool that will drive you to the AOP world!
- [blank(stars: 11)](https://github.com/Henry-Sarabia/blank) - Verify or remove blanks and whitespace from strings.
- [bleep(stars: 10)](https://github.com/sinhashubham95/bleep) - Perform any number of actions on any set of OS signals in Go.
- [boilr(stars: 1671)](https://github.com/tmrts/boilr) - Blazingly fast CLI tool for creating projects from boilerplate templates.
- [changie(stars: 552)](https://github.com/miniscruff/changie) - Automated changelog tool for preparing releases with lots of customization options.
- [chyle(stars: 154)](https://github.com/antham/chyle) - Changelog generator using a git repository with multiple configuration possibilities.
- [circuit(stars: 720)](https://github.com/cep21/circuit) - An efficient and feature complete Hystrix like Go implementation of the circuit breaker pattern.
- [circuitbreaker(stars: 1094)](https://github.com/rubyist/circuitbreaker) - Circuit Breakers in Go.
- [clipboard(stars: 489)](https://github.com/golang-design/clipboard) - 📋 cross-platform clipboard package in Go.
- [clockwork(stars: 563)](https://github.com/jonboulle/clockwork) - A simple fake clock for golang.
- [cmd(stars: 140)](https://github.com/SimonBaeumer/cmd) - Library for executing shell commands on osx, windows and linux.
- [command(stars: 14)](https://github.com/txgruppi/command) - Command pattern for Go with thread safe serial and parallel dispatcher.
- [contextplus(stars: 17)](https://github.com/contextplus/contextplus) - Package contextplus provide more easy to use functions for contexts.
- [copy(stars: 38)](https://github.com/gotidy/copy) - Package for fast copying structs of different types.
- [copy-pasta(stars: 50)](https://github.com/jutkko/copy-pasta) - Universal multi-workstation clipboard that uses S3 like backend for the storage.
- [countries(stars: 319)](https://github.com/biter777/countries) - Full implementation of ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standards.
- [countries(stars: 85)](https://github.com/pioz/countries) - All you need when you are working with countries in Go.
- [create-go-app(stars: 2261)](https://github.com/create-go-app/cli) - A powerful CLI for create a new production-ready project with backend (Golang), frontend (JavaScript, TypeScript) & deploy automation (Ansible, Docker) by running one command.
- [cryptgo(stars: 142)](https://github.com/Gituser143/cryptgo) - Crytpgo is a TUI based application written purely in Go to monitor and observe cryptocurrency prices in real time!
- [ctop(stars: 14869)](https://github.com/bcicen/ctop) - [Top-like](https://ctop.sh) interface (e.g. htop) for container metrics.
- [ctxutil(stars: 25)](https://github.com/posener/ctxutil) - A collection of utility functions for contexts.
- [cvt(stars: 44)](https://github.com/shockerli/cvt) - Easy and safe convert any value to another type.
- [dbt(stars: 58)](https://github.com/nikogura/dbt) - A framework for running self-updating signed binaries from a central, trusted repository.
- [Death(stars: 191)](https://github.com/vrecan/death) - Managing go application shutdown with signals.
- [Deepcopier(stars: 431)](https://github.com/ulule/deepcopier) - Simple struct copying for Go.
- [delve(stars: 612)](https://github.com/derekparker/delve) - Go debugger.
- [dive(stars: 39960)](https://github.com/wagoodman/dive) - A tool for exploring each layer in a Docker image.
- [dlog(stars: 17)](https://github.com/kirillDanshin/dlog) - Compile-time controlled logger to make your release smaller without removing debug calls.
- [EaseProbe(stars: 1964)](https://github.com/megaease/easeprobe) - A simple, standalone, and lightWeight tool that can do health/status checking daemon, support HTTP/TCP/SSH/Shell/Client/... probes, and Slack/Discord/Telegram/SMS... notification.
- [equalizer(stars: 43)](https://github.com/reugn/equalizer) - Quota manager and rate limiter collection for Go.
- [ergo(stars: 594)](https://github.com/cristianoliveira/ergo) - The management of multiple local services running over different ports made easy.
- [evaluator(stars: 37)](https://github.com/nullne/evaluator) - Evaluate an expression dynamically based on s-expression. It's simple and easy to extend.
- [filetype(stars: 1939)](https://github.com/h2non/filetype) - Small package to infer the file type checking the magic numbers signature.
- [filler(stars: 17)](https://github.com/yaronsumel/filler) - small utility to fill structs using "fill" tag.
- [filter(stars: 136)](https://github.com/gookit/filter) - provide filtering, sanitizing, and conversion of Go data.
- [fzf(stars: 56410)](https://github.com/junegunn/fzf) - Command-line fuzzy finder written in Go.
- [generate(stars: 30)](https://github.com/go-playground/generate) - runs go generate recursively on a specified path or environment variable and can filter by regex.
- [ghokin(stars: 40)](https://github.com/antham/ghokin) - Parallelized formatter with no external dependencies for gherkin (cucumber, behat...).
- [git-time-metric(stars: 956)](https://github.com/git-time-metric/gtm) - Simple, seamless, lightweight time tracking for Git.
- [gitbatch(stars: 1517)](https://github.com/isacikgoz/gitbatch) - manage your git repositories in one place.
- [go-actuator(stars: 14)](https://github.com/sinhashubham95/go-actuator) - Production ready features for Go based web frameworks.
- [go-astitodo(stars: 61)](https://github.com/asticode/go-astitodo) - Parse TODOs in your GO code.
- [go-bind-plugin(stars: 184)](https://github.com/wendigo/go-bind-plugin) - go:generate tool for wrapping symbols exported by golang plugins (1.8 only).
- [go-bsdiff(stars: 149)](https://github.com/gabstv/go-bsdiff) - Pure Go bsdiff and bspatch libraries and CLI tools.
- [go-clip(stars: 12)](https://github.com/prashantgupta24/go-clip) - A minimalistic clipboard manager for Mac.
- [go-convert(stars: 22)](https://github.com/Eun/go-convert) - Package go-convert enables you to convert a value into another type.
- [go-countries(stars: 13)](https://github.com/mikekonan/go-countries) - Lightweight lookup over ISO-3166 codes.
- [go-dry(stars: 488)](https://github.com/ungerik/go-dry) - DRY (don't repeat yourself) package for Go.
- [go-funk(stars: 4500)](https://github.com/thoas/go-funk) - Modern Go utility library which provides helpers (map, find, contains, filter, chunk, reverse, ...).
- [go-health(stars: 94)](https://github.com/Talento90/go-health) - Health package simplifies the way you add health check to your services.
- [go-httpheader(stars: 46)](https://github.com/mozillazg/go-httpheader) - Go library for encoding structs into Header fields.
- [go-lock(stars: 104)](https://github.com/viney-shih/go-lock) - go-lock is a lock library implementing read-write mutex and read-write trylock without starvation.
- [go-pattern-match(stars: 83)](https://github.com/PhakornKiong/go-pattern-match) - A Pattern matching library inspired by ts-pattern.
- [go-pkg(stars: 6)](https://github.com/chenquan/go-pkg) - A go toolkit.
- [go-problemdetails(stars: 17)](https://github.com/mvmaasakkers/go-problemdetails) - Go package for working with Problem Details.
- [go-qr(stars: 9)](https://github.com/piglig/go-qr) - A native, high-quality and minimalistic QR code generator.
- [go-rate(stars: 393)](https://github.com/beefsack/go-rate) - Timed rate limiter for Go.
- [go-sitemap-generator(stars: 210)](https://github.com/ikeikeikeike/go-sitemap-generator) - XML Sitemap generator written in Go.
- [go-trigger(stars: 237)](https://github.com/sadlil/go-trigger) - Go-lang global event triggerer, Register Events with an id and trigger the event from anywhere from your project.
- [go-type(stars: 17)](https://github.com/mikekonan/go-types) - Library providing Go types for store/validation and transfer of ISO-4217, ISO-3166, and other types.
- [goback(stars: 48)](https://github.com/carlescere/goback) - Go simple exponential backoff package.
- [goctx(stars: 7)](https://github.com/zerosnake0/goctx) - Get your context value with high performance.
- [godaemon(stars: 492)](https://github.com/VividCortex/godaemon) - Utility to write daemons.
- [godropbox(stars: 4162)](https://github.com/dropbox/godropbox) - Common libraries for writing Go services/applications from Dropbox.
- [gofn(stars: 28)](https://github.com/tiendc/gofn) - High performance utility functions written using Generics for Go 1.18+.
- [gohper(stars: 256)](https://github.com/cosiner/gohper) - Various tools/modules help for development.
- [golarm(stars: 51)](https://github.com/msempere/golarm) - Fire alarms with system events.
- [golog(stars: 60)](https://github.com/mlimaloureiro/golog) - Easy and lightweight CLI tool to time track your tasks.
- [gopencils(stars: 450)](https://github.com/bndr/gopencils) - Small and simple package to easily consume REST APIs.
- [goplaceholder(stars: 28)](https://github.com/michiwend/goplaceholder) - a small golang lib to generate placeholder images.
- [goreadability(stars: 69)](https://github.com/philipjkim/goreadability) - Webpage summary extractor using Facebook Open Graph and arc90's readability.
- [goreleaser(stars: 12429)](https://github.com/goreleaser/goreleaser) - Deliver Go binaries as fast and easily as possible.
- [goreporter(stars: 3090)](https://github.com/wgliang/goreporter) - Golang tool that does static analysis, unit testing, code review and generate code quality report.
- [goseaweedfs(stars: 111)](https://github.com/linxGnu/goseaweedfs) - SeaweedFS client library with almost full features.
- [gostrutils(stars: 43)](https://github.com/ik5/gostrutils) - Collections of string manipulation and conversion functions.
- [gotenv(stars: 276)](https://github.com/subosito/gotenv) - Load environment variables from `.env` or any `io.Reader` in Go.
- [goval(stars: 136)](https://github.com/maja42/goval) - Evaluate arbitrary expressions in Go.
- [gpath(stars: 40)](https://github.com/tenntenn/gpath) - Library to simplify access struct fields with Go's expression in reflection.
- [graterm(stars: 27)](https://github.com/skovtunenko/graterm) - Provides primitives to perform ordered (sequential/concurrent) GRAceful TERMination (aka shutdown) in Go application.
- [grofer(stars: 314)](https://github.com/pesos/grofer) - A system and resource monitoring tool written in Golang!
- [gubrak(stars: 469)](https://github.com/novalagung/gubrak) - Golang utility library with syntactic sugar. It's like lodash, but for golang.
- [handy(stars: 76)](https://github.com/miguelpragier/handy) - Many utilities and helpers like string handlers/formatters and validators.
- [hostctl(stars: 971)](https://github.com/guumaster/hostctl) - A CLI tool to manage /etc/hosts with easy commands.
- [htcat(stars: 552)](https://github.com/htcat/htcat) - Parallel and Pipelined HTTP GET Utility.
- [hub(stars: 22603)](https://github.com/github/hub) - wrap git commands with additional functionality to interact with github from the terminal.
- [hystrix-go(stars: 4088)](https://github.com/afex/hystrix-go) - Implements Hystrix patterns of programmer-defined fallbacks aka circuit breaker.
- [immortal(stars: 781)](https://github.com/immortal/immortal) - \*nix cross-platform (OS agnostic) supervisor.
- [intrinsic(stars: 46)](https://github.com/mengzhuo/intrinsic) - Use x86 SIMD without writing any assembly code.
- [jsend(stars: 21)](https://github.com/clevergo/jsend) - JSend's implementation written in Go.
- [jump(stars: 1711)](https://github.com/gsamokovarov/jump) - Jump helps you navigate faster by learning your habits.
- [just(stars: 18)](https://github.com/kazhuravlev/just) - Just a collection of useful functions for working with generic data structures.
- [koazee(stars: 519)](https://github.com/wesovilabs/koazee) - Library inspired in Lazy evaluation and functional programming that takes the hassle out of working with arrays.
- [lancet(stars: 3328)](https://github.com/duke-git/lancet) - A comprehensive, efficient, and reusable util function library of go.
- [lets-go(stars: 6)](https://github.com/aplescia-chwy/lets-go) - Go module that provides common utilities for Cloud Native REST API development. Also contains AWS Specific utilities.
- [limiters(stars: 259)](https://github.com/mennanov/limiters) - Rate limiters for distributed applications in Golang with configurable back-ends and distributed locks.
- [lo(stars: 14294)](https://github.com/samber/lo) - A Lodash like Go library based on Go 1.18+ Generics (map, filter, contains, find...)
- [loncha(stars: 7)](https://github.com/kazu/loncha) - A high-performance slice Utilities.
- [lrserver(stars: 124)](https://github.com/jaschaephraim/lrserver) - LiveReload server for Go.
- [mani(stars: 404)](https://github.com/alajmo/mani) - CLI tool to help you manage multiple repositories.
- [mc(stars: 2609)](https://github.com/minio/mc) - Minio Client provides minimal tools to work with Amazon S3 compatible cloud storage and filesystems.
- [mergo(stars: 2624)](https://github.com/imdario/mergo) - Helper to merge structs and maps in Golang. Useful for configuration default values, avoiding messy if-statements.
- [mimemagic(stars: 96)](https://github.com/zRedShift/mimemagic) - Pure Go ultra performant MIME sniffing library/utility.
- [mimesniffer(stars: 32)](https://github.com/aofei/mimesniffer) - A MIME type sniffer for Go.
- [mimetype(stars: 1308)](https://github.com/gabriel-vasile/mimetype) - Package for MIME type detection based on magic numbers.
- [minify(stars: 3494)](https://github.com/tdewolff/minify) - Fast minifiers for HTML, CSS, JS, XML, JSON and SVG file formats.
- [minquery(stars: 61)](https://github.com/icza/minquery) - MongoDB / mgo.v2 query that supports efficient pagination (cursors to continue listing documents where we left off).
- [moldova(stars: 167)](https://github.com/StabbyCutyou/moldova) - Utility for generating random data based on an input template.
- [mole(stars: 1668)](https://github.com/davrodpin/mole) - cli app to easily create ssh tunnels.
- [mongo-go-pagination(stars: 126)](https://github.com/gobeam/mongo-go-pagination) - Mongodb Pagination for official mongodb/mongo-go-driver package which supports both normal queries and Aggregation pipelines.
- [mssqlx(stars: 101)](https://github.com/linxGnu/mssqlx) - Database client library, proxy for any master slave, master master structures. Lightweight and auto balancing in mind.
- [multitick(stars: 69)](https://github.com/VividCortex/multitick) - Multiplexor for aligned tickers.
- [myhttp(stars: 34)](https://github.com/inancgumus/myhttp) - Simple API to make HTTP GET requests with timeout support.
- [netbug(stars: 73)](https://github.com/e-dard/netbug) - Easy remote profiling of your services.
- [nfdump(stars: 8)](https://github.com/chrispassas/nfdump) - Read nfdump netflow files.
- [nostromo(stars: 138)](https://github.com/pokanop/nostromo) - CLI for building powerful aliases.
- [objwalker(stars: 2)](https://github.com/rekby/objwalker) - Walk by go objects with reflection.
- [okrun(stars: 16)](https://github.com/xta/okrun) - go run error steamroller.
- [olaf(stars: 6)](https://github.com/btnguyen2k/olaf) - Twitter Snowflake implemented in Go.
- [onecache(stars: 133)](https://github.com/adelowo/onecache) - Caching library with support for multiple backend stores (Redis, Memcached, filesystem etc).
- [panicparse(stars: 3465)](https://github.com/maruel/panicparse) - Groups similar goroutines and colorizes stack dump.
- [pattern-match(stars: 233)](https://github.com/alexpantyukhin/go-pattern-match) - Pattern matching library.
- [peco(stars: 7491)](https://github.com/peco/peco) - Simplistic interactive filtering tool.
- [pgo(stars: 84)](https://github.com/arthurkushman/pgo) - Convenient functions for PHP community.
- [pm(stars: 79)](https://github.com/VividCortex/pm) - Process (i.e. goroutine) manager with an HTTP API.
- [pointer(stars: 40)](https://github.com/xorcare/pointer) - Package pointer contains helper routines for simplifying the creation of optional fields of basic type.
- [ptr(stars: 23)](https://github.com/gotidy/ptr) - Package that provide functions for simplified creation of pointers from constants of basic types.
- [rclient(stars: 35)](https://github.com/zpatrick/rclient) - Readable, flexible, simple-to-use client for REST APIs.
- [reflectutils(stars: 7)](https://github.com/muir/reflectutils) - Helpers for working with reflection: struct tag parsing; recursive walking; fill value from string.
- [remote-touchpad(stars: 449)](https://github.com/Unrud/remote-touchpad) - Control mouse and keyboard from a smartphone.
- [repeat(stars: 84)](https://github.com/ssgreg/repeat) - Go implementation of different backoff strategies useful for retrying operations and heartbeating.
- [request(stars: 426)](https://github.com/mozillazg/request) - Go HTTP Requests for Humans™.
- [rerun(stars: 167)](https://github.com/ivpusic/rerun) - Recompiling and rerunning go apps when source changes.
- [rest-go(stars: 16)](https://github.com/edermanoel94/rest-go) - A package that provide many helpful methods for working with rest api.
- [retry(stars: 334)](https://github.com/kamilsk/retry) - The most advanced functional mechanism to perform actions repetitively until successful.
- [retry(stars: 10)](https://github.com/percolate/retry) - A simple but highly configurable retry package for Go.
- [retry(stars: 63)](https://github.com/thedevsaddam/retry) - Simple and easy retry mechanism package for Go.
- [retry(stars: 12)](https://github.com/shafreeck/retry) - A pretty simple library to ensure your work to be done.
- [retry-go(stars: 2077)](https://github.com/avast/retry-go) - Simple library for retry mechanism.
- [retry-go(stars: 47)](https://github.com/rafaeljesus/retry-go) - Retrying made simple and easy for golang.
- [robustly(stars: 157)](https://github.com/VividCortex/robustly) - Runs functions resiliently, catching and restarting panics.
- [rospo(stars: 255)](https://github.com/ferama/rospo) - Simple and reliable ssh tunnels with embedded ssh server in Golang.
- [scan(stars: 478)](https://github.com/blockloop/scan) - Scan golang `sql.Rows` directly to structs, slices, or primitive types.
- [scan(stars: 54)](https://github.com/wroge/scan) - Scan sql rows into any type powered by generics.
- [scany(stars: 1092)](https://github.com/georgysavva/scany) - Library for scanning data from a database into Go structs and more.
- [serve(stars: 319)](https://github.com/syntaqx/serve) - A static http server anywhere you need.
- [set(stars: 46)](https://github.com/nofeaturesonlybugs/set) - Performant and flexible struct mapping and loose type conversion.
- [shutdown(stars: 55)](https://github.com/ztrue/shutdown) - App shutdown hooks for `os.Signal` handling.
- [silk(stars: 13)](https://github.com/chrispassas/silk) - Read silk netflow files.
- [slice(stars: 51)](https://github.com/psampaz/slice) - Type-safe functions for common Go slice operations.
- [sliceconv(stars: 8)](https://github.com/Henry-Sarabia/sliceconv) - Slice conversion between primitive types.
- [slicer(stars: 44)](https://github.com/leaanthony/slicer) - Makes working with slices easier.
- [sorty(stars: 124)](https://github.com/jfcg/sorty) - Fast Concurrent / Parallel Sorting.
- [sqlx(stars: 14781)](https://github.com/jmoiron/sqlx) - provides a set of extensions on top of the excellent built-in database/sql package.
- [sshman(stars: 45)](https://github.com/shoobyban/sshman) - SSH Manager for authorized_keys files on multiple remote servers.
- [statiks(stars: 10)](https://github.com/janiltonmaciel/statiks) - Fast, zero-configuration, static HTTP filer server.
- [Storm(stars: 2017)](https://github.com/asdine/storm) - Simple and powerful toolkit for BoltDB.
- [structs(stars: 24)](https://github.com/PumpkinSeed/structs) - Implement simple functions to manipulate structs.
- [throttle(stars: 34)](https://github.com/yudppp/throttle) - Throttle is an object that will perform exactly one action per duration.
- [tik(stars: 5)](https://github.com/andy2046/tik) - Simple and easy timing wheel package for Go.
- [tome(stars: 35)](https://github.com/cyruzin/tome) - Tome was designed to paginate simple RESTful APIs.
- [toolbox(stars: 194)](https://github.com/viant/toolbox) - Slice, map, multimap, struct, function, data conversion utilities. Service router, macro evaluator, tokenizer.
- [ugo(stars: 27)](https://github.com/alxrm/ugo) - ugo is slice toolbox with concise syntax for Go.
- [UNIS(stars: 70)](https://github.com/esemplastic/unis) - Common Architecture™ for String Utilities in Go.
- [upterm(stars: 702)](https://github.com/owenthereal/upterm) - A tool for developers to share terminal/tmux sessions securely over the web. It’s perfect for remote pair programming, accessing computers behind NATs/firewalls, remote debugging, and more.
- [usql(stars: 8378)](https://github.com/knq/usql) - usql is a universal command-line interface for SQL databases.
- [util(stars: 281)](https://github.com/shomali11/util) - Collection of useful utility functions. (strings, concurrency, manipulations, ...).
- [watchhttp(stars: 27)](https://github.com/nikolaydubina/watchhttp) - Run command periodically and expose latest STDOUT or its rich delta as HTTP endpoint.
- [wifiqr(stars: 246)](https://github.com/reugn/wifiqr) - Wi-Fi QR Code Generator.
- [wuzz(stars: 10409)](https://github.com/asciimoo/wuzz) - Interactive cli tool for HTTP inspection.
- [xferspdy(stars: 99)](https://github.com/monmohan/xferspdy) - Xferspdy provides binary diff and patch library in golang.
- [yogo(stars: 35)](https://github.com/antham/yogo) - Check yopmail mails from command line.

**[⬆ back to top](#contents)**

## UUID

_Libraries for working with UUIDs._

- [goid(stars: 39)](https://github.com/jakehl/goid) - Generate and Parse RFC4122 compliant V4 UUIDs.
- [gouid(stars: 22)](https://github.com/twharmon/gouid) - Generate cryptographically secure random string IDs with just one allocation.
- [nanoid(stars: 58)](https://github.com/aidarkhanov/nanoid) - A tiny and efficient Go unique string ID generator.
- [sno(stars: 89)](https://github.com/muyo/sno) - Compact, sortable and fast unique IDs with embedded metadata.
- [ulid(stars: 3919)](https://github.com/oklog/ulid) - Go implementation of ULID (Universally Unique Lexicographically Sortable Identifier).
- [uniq](https://gitlab.com/skilstak/code/go/uniq) - No hassle safe, fast unique identifiers with commands.
- [uuid(stars: 18)](https://github.com/agext/uuid) - Generate, encode, and decode UUIDs v1 with fast or cryptographic-quality random node identifier.
- [uuid(stars: 1462)](https://github.com/gofrs/uuid) - Implementation of Universally Unique Identifier (UUID). Supports both creation and parsing of UUIDs. Actively maintained fork of satori uuid.
- [uuid(stars: 4781)](https://github.com/google/uuid) - Go package for UUIDs based on RFC 4122 and DCE 1.1: Authentication and Security Services.
- [wuid(stars: 505)](https://github.com/edwingeng/wuid) - An extremely fast globally unique number generator.
- [xid(stars: 3611)](https://github.com/rs/xid) - Xid is a globally unique id generator library, ready to be safely used directly in your server code.

**[⬆ back to top](#contents)**

## Validation

_Libraries for validation._

- [checkdigit(stars: 106)](https://github.com/osamingo/checkdigit) - Provide check digit algorithms (Luhn, Verhoeff, Damm) and calculators (ISBN, EAN, JAN, UPC, etc.).
- [gody(stars: 65)](https://github.com/guiferpa/gody) - :balloon: A lightweight struct validator for Go.
- [govalid(stars: 36)](https://github.com/twharmon/govalid) - Fast, tag-based validation for structs.
- [govalidator(stars: 5866)](https://github.com/asaskevich/govalidator) - Validators and sanitizers for strings, numerics, slices and structs.
- [govalidator(stars: 1255)](https://github.com/thedevsaddam/govalidator) - Validate Golang request data with simple rules. Highly inspired by Laravel's request validation.
- [jio(stars: 91)](https://github.com/faceair/jio) - jio is a json schema validator similar to [joi](https://github.com/hapijs/joi).
- [ozzo-validation(stars: 3448)](https://github.com/go-ozzo/ozzo-validation) - Supports validation of various data types (structs, strings, maps, slices, etc.) with configurable and extensible validation rules specified in usual code constructs instead of struct tags.
- [validate(stars: 970)](https://github.com/gookit/validate) - Go package for data validation and filtering. support validate Map, Struct, Request(Form, JSON, url.Values, Uploaded Files) data and more features.
- [validate(stars: 91)](https://github.com/gobuffalo/validate) - This package provides a framework for writing validations for Go applications.
- [validator(stars: 14727)](https://github.com/go-playground/validator) - Go Struct and Field validation, including Cross Field, Cross Struct, Map, Slice and Array diving.
- [Validator(stars: 7)](https://github.com/go-the-way/validator) - A lightweight model validator written in Go.Contains VFs:Min, Max, MinLength, MaxLength, Length, Enum, Regex.
- [valix](https://github.com/marrow16/valix) Go package for validating requests

**[⬆ back to top](#contents)**

## Version Control

_Libraries for version control._

- [cli](https://gitlab.com/gitlab-org/cli) - An open-source GitLab command line tool bringing GitLab's cool features to your command line.
- [froggit-go(stars: 41)](https://github.com/jfrog/froggit-go) - Froggit-Go is a Go library, allowing to perform actions on VCS providers.
- [gh(stars: 82)](https://github.com/rjeczalik/gh) - Scriptable server and net/http middleware for GitHub Webhooks.
- [git2go(stars: 1881)](https://github.com/libgit2/git2go) - Go bindings for libgit2.
- [githooks(stars: 83)](https://github.com/gabyx/githooks) - Per-repo and shared Git hooks with version control and auto update.
- [go-git(stars: 5233)](https://github.com/go-git/go-git) - highly extensible Git implementation in pure Go.
- [go-vcs(stars: 77)](https://github.com/sourcegraph/go-vcs) - manipulate and inspect VCS repositories in Go.
- [hercules(stars: 1972)](https://github.com/src-d/hercules) - gaining advanced insights from Git repository history.
- [hgo(stars: 17)](https://github.com/beyang/hgo) - Hgo is a collection of Go packages providing read-access to local Mercurial repositories.

**[⬆ back to top](#contents)**

## Video

_Libraries for manipulating video._

- [gmf(stars: 870)](https://github.com/3d0c/gmf) - Go bindings for FFmpeg av\* libraries.
- [go-astisub(stars: 524)](https://github.com/asticode/go-astisub) - Manipulate subtitles in GO (.srt, .stl, .ttml, .webvtt, .ssa/.ass, teletext, .smi, etc.).
- [go-astits(stars: 511)](https://github.com/asticode/go-astits) - Parse and demux MPEG Transport Streams (.ts) natively in GO.
- [go-m3u8(stars: 24)](https://github.com/etherlabsio/go-m3u8) - Parser and generator library for Apple m3u8 playlists. Actively maintained version of quangngotan95/go-m3u8 with improvements and latest HLS playlist parsing compatibility.
- [go-mpd(stars: 19)](https://github.com/unki2aut/go-mpd) - Parser and generator library for MPEG-DASH manifest files.
- [goav(stars: 2052)](https://github.com/giorgisio/goav) - Comprehensive Go bindings for FFmpeg.
- [gortsplib(stars: 535)](https://github.com/aler9/gortsplib) - Pure Go RTSP server and client library.
- [gst(stars: 168)](https://github.com/ziutek/gst) - Go bindings for GStreamer.
- [libgosubs(stars: 24)](https://github.com/wargarblgarbl/libgosubs) - Subtitle format support for go. Supports .srt, .ttml, and .ass.
- [libvlc-go(stars: 388)](https://github.com/adrg/libvlc-go) - Go bindings for libvlc 2.X/3.X/4.X (used by the VLC media player).
- [m3u8(stars: 1118)](https://github.com/grafov/m3u8) - Parser and generator library of M3U8 playlists for Apple HLS.
- [v4l(stars: 77)](https://github.com/korandiz/v4l) - Video capture library for Linux, written in Go.

**[⬆ back to top](#contents)**

## Web Frameworks

_Full stack web frameworks._

- [aah](https://aahframework.org) - Scalable, performant, rapid development Web framework for Go.
- [Aero(stars: 550)](https://github.com/aerogo/aero) - High-performance web framework for Go, reaches top scores in Lighthouse.
- [Air(stars: 437)](https://github.com/aofei/air) - An ideally refined web framework for Go.
- [anoweb(stars: 5)](https://github.com/go-the-way/anoweb) - The lightweight and powerful web framework using the new way for Go.Another go the way.
- [appy(stars: 131)](https://github.com/appist/appy) - An opinionated productive web framework that helps scaling business easier.
- [Atreugo(stars: 1151)](https://github.com/savsgio/atreugo) - High performance and extensible micro web framework with zero memory allocations in hot paths.
- [Banjo(stars: 21)](https://github.com/nsheremet/banjo) - Very simple and fast web framework for Go.
- [Beego(stars: 30499)](https://github.com/beego/beego) - beego is an open-source, high-performance web framework for the Go programming language.
- [Buffalo](https://gobuffalo.io) - Bringing the productivity of Rails to Go!
- [Confetti Framework](https://confetti-framework.github.io/docs/) - Confetti is a Go web application framework with an expressive, elegant syntax. Confetti combines the elegance of Laravel and the simplicity of Go.
- [Don(stars: 33)](https://github.com/abemedia/go-don) - A highly performant and simple to use API framework.
- [Echo(stars: 27444)](https://github.com/labstack/echo) - High performance, minimalist Go web framework.
- [Fiber(stars: 29730)](https://github.com/gofiber/fiber) - An Express.js inspired web framework build on Fasthttp.
- [Fireball(stars: 58)](https://github.com/zpatrick/fireball) - More "natural" feeling web framework.
- [Flamingo(stars: 417)](https://github.com/i-love-flamingo/flamingo) - Framework for pluggable web projects. Including a concept for modules and offering features for DI, Configareas, i18n, template engines, graphql, observability, security, events, routing & reverse routing etc.
- [Flamingo Commerce(stars: 463)](https://github.com/i-love-flamingo/flamingo-commerce) - Providing e-commerce features using clean architecture like DDD and ports and adapters, that you can use to build flexible e-commerce applications.
- [Gearbox(stars: 730)](https://github.com/abahmed/gearbox) - A web framework written in Go with a focus on high performance and memory optimization.
- [Gin(stars: 73320)](https://github.com/gin-gonic/gin) - Gin is a web framework written in Go! It features a martini-like API with much better performance, up to 40 times faster. If you need performance and good productivity.
- [Ginrpc(stars: 286)](https://github.com/xxjwxc/ginrpc) - Gin parameter automatic binding tool,gin rpc tools.
- [Gizmo(stars: 3748)](https://github.com/NYTimes/gizmo) - Microservice toolkit used by the New York Times.
- [go-json-rest(stars: 3511)](https://github.com/ant0ine/go-json-rest) - Quick and easy way to setup a RESTful JSON API.
- [go-rest(stars: 127)](https://github.com/ungerik/go-rest) - Small and evil REST framework for Go.
- [Goa(stars: 5367)](https://github.com/goadesign/goa) - Goa provides a holistic approach for developing remote APIs and microservices in Go.
- [goa(stars: 49)](https://github.com/goa-go/goa) - goa is just like koajs for golang, it is a flexible, light, high-performance and extensible web framework based on middleware.
- [GoFr(stars: 659)](https://github.com/gofr-dev/gofr) - Gofr is an opinionated microservice development framework.
- [GoFrame(stars: 10388)](https://github.com/gogf/gf) - GoFrame is a modular, full-featured and production-ready application development framework of golang.
- [golamb(stars: 6)](https://github.com/twharmon/golamb) - Golamb makes it easier to write API endpoints for use with AWS Lambda and API Gateway.
- [Golax(stars: 75)](https://github.com/fulldump/golax) - A non Sinatra fast HTTP framework with support for Google custom methods, deep interceptors, recursion and more.
- [Golf(stars: 267)](https://github.com/dinever/golf) - Golf is a fast, simple and lightweight micro-web framework for Go. It comes with powerful features and has no dependencies other than the Go Standard Library.
- [Gondola(stars: 311)](https://github.com/rainycape/gondola) - The web framework for writing faster sites, faster.
- [gongular(stars: 500)](https://github.com/mustafaakin/gongular) - Fast Go web framework with input mapping/validation and (DI) Dependency Injection.
- [GoTuna(stars: 46)](https://github.com/gotuna/gotuna) - Minimalistic web framework for Go with mux router, middlewares, sessions, templates, embedded views and static files.
- [goweb(stars: 36)](https://github.com/twharmon/goweb) - Web framework with routing, websockets, logging, middleware, static file server (optional gzip), and automatic TLS.
- [Goyave(stars: 1397)](https://github.com/go-goyave/goyave) - Feature-complete REST API framework aimed at clean code and fast development, with powerful built-in functionalities.
- [Hertz(stars: 4435)](https://github.com/cloudwego/hertz) - A high-performance and strong-extensibility Go HTTP framework that helps developers build microservices.
- [hiboot(stars: 180)](https://github.com/hidevopsio/hiboot) - hiboot is a high performance web application framework with auto configuration and dependency injection support.
- [Huma](https://github.com/danielgtaylor/huma/) - Framework for modern REST/GraphQL APIs with built-in OpenAPI 3, generated documentation, and a CLI.
- [Macaron(stars: 3445)](https://github.com/go-macaron/macaron) - Macaron is a high productive and modular design web framework in Go.
- [mango(stars: 372)](https://github.com/paulbellamy/mango) - Mango is a modular web-application framework for Go, inspired by Rack, and PEP333.
- [Microservice(stars: 114)](https://github.com/claygod/microservice) - The framework for the creation of microservices, written in Golang.
- [neo(stars: 418)](https://github.com/ivpusic/neo) - Neo is minimal and fast Go Web Framework with extremely simple API.
- [patron(stars: 124)](https://github.com/beatlabs/patron) - Patron is a microservice framework following best cloud practices with a focus on productivity.
- [Pnutmux](https://gitlab.com/fruitygo/pnutmux) - Pnutmux is a powerful Go web framework that uses regex for matching and handling HTTP requests. It offers features such as CORS handling, structured logging, URL parameters extraction, middlewares, and concurrency limiting.
- [Pulse(stars: 37)](https://github.com/gopulse/pulse) - Pulse is an HTTP web framework written in Go (Golang)
- [Resoursea(stars: 34)](https://github.com/resoursea/api) - REST framework for quickly writing resource based services.
- [REST Layer(stars: 1237)](https://github.com/rs/rest-layer) - Framework to build REST/GraphQL API on top of databases with mostly configuration over code.
- [Revel(stars: 13036)](https://github.com/revel/revel) - High-productivity web framework for the Go language.
- [rex(stars: 33)](https://github.com/goanywhere/rex) - Rex is a library for modular development built upon gorilla/mux, fully compatible with `net/http`.
- [rk-boot(stars: 466)](https://github.com/rookie-ninja/rk-boot) - A bootstrapper library for building enterprise go microservice with Gin and gRPC quickly and easily.
- [rux(stars: 94)](https://github.com/gookit/rux) - Simple and fast web framework for build golang HTTP applications.
- [tango(stars: 835)](https://github.com/lunny/tango) - Micro & pluggable web framework for Go.
- [tigertonic(stars: 996)](https://github.com/rcrowley/go-tigertonic) - Go framework for building JSON web services inspired by Dropwizard.
- [uAdmin(stars: 308)](https://github.com/uadmin/uadmin) - Fully featured web framework for Golang, inspired by Django.
- [utron(stars: 2221)](https://github.com/gernest/utron) - Lightweight MVC framework for Go(Golang).
- [vox(stars: 81)](https://github.com/aisk/vox) - A golang web framework for humans, inspired by Koa heavily.
- [WebGo(stars: 292)](https://github.com/bnkamalesh/webgo) - A micro-framework to build web apps; with handler chaining, middleware and context injection. With standard library compliant HTTP handlers(i.e. http.HandlerFunc).
- [YARF(stars: 67)](https://github.com/yarf-framework/yarf) - Fast micro-framework designed to build REST APIs and web services in a fast and simple way.

**[⬆ back to top](#contents)**

### Middlewares

#### Actual middlewares

- [client-timing(stars: 24)](https://github.com/posener/client-timing) - An HTTP client for Server-Timing header.
- [CORS(stars: 2476)](https://github.com/rs/cors) - Easily add CORS capabilities to your API.
- [echo-middleware(stars: 13)](https://github.com/faabiosr/echo-middleware) - Middleware for Echo framework with logging and metrics.
- [formjson(stars: 38)](https://github.com/rs/formjson) - Transparently handle JSON input as a standard form POST.
- [go-fault(stars: 483)](https://github.com/github/go-fault) - Fault injection middleware for Go.
- [go-server-timing(stars: 854)](https://github.com/mitchellh/go-server-timing) - Add/parse Server-Timing header.
- [Limiter(stars: 1884)](https://github.com/ulule/limiter) - Dead simple rate limit middleware for Go.
- [ln-paywall(stars: 137)](https://github.com/philippgille/ln-paywall) - Go middleware for monetizing APIs on a per-request basis with the Lightning Network (Bitcoin).
- [mid(stars: 7)](https://github.com/bobg/mid) - Miscellaneous HTTP middleware features: idiomatic error return from handlers; receive/respond with JSON data; request tracing; and more.
- [rk-gin(stars: 42)](https://github.com/rookie-ninja/rk-gin) - Middleware for Gin framework with logging, metrics, auth, tracing etc.
- [rk-grpc(stars: 68)](https://github.com/rookie-ninja/rk-grpc) - Middleware for gRPC with logging, metrics, auth, tracing etc.
- [Tollbooth(stars: 2538)](https://github.com/didip/tollbooth) - Rate limit HTTP request handler.
- [XFF(stars: 98)](https://github.com/sebest/xff) - Handle `X-Forwarded-For` header and friends.

#### Libraries for creating HTTP middlewares

- [alice(stars: 2904)](https://github.com/justinas/alice) - Painless middleware chaining for Go.
- [catena(stars: 9)](https://github.com/codemodus/catena) - http.Handler wrapper catenation (same API as "chain").
- [chain(stars: 63)](https://github.com/codemodus/chain) - Handler wrapper chaining with scoped data (net/context-based "middleware").
- [gores(stars: 102)](https://github.com/alioygur/gores) - Go package that handles HTML, JSON, XML and etc. responses. Useful for RESTful APIs.
- [interpose(stars: 293)](https://github.com/carbocation/interpose) - Minimalist net/http middleware for golang.
- [mediary(stars: 89)](https://github.com/HereMobilityDevelopers/mediary) - add interceptors to `http.Client` to allow dumping/shaping/tracing/... of requests/responses.
- [muxchain(stars: 208)](https://github.com/stephens2424/muxchain) - Lightweight middleware for net/http.
- [negroni(stars: 7391)](https://github.com/urfave/negroni) - Idiomatic HTTP middleware for Golang.
- [render(stars: 1870)](https://github.com/unrolled/render) - Go package for easily rendering JSON, XML, and HTML template responses.
- [renderer(stars: 258)](https://github.com/thedevsaddam/renderer) - Simple, lightweight and faster response (JSON, JSONP, XML, YAML, HTML, File) rendering package for Go.
- [rye(stars: 102)](https://github.com/InVisionApp/rye) - Tiny Go middleware library (with canned Middlewares) that supports JWT, CORS, Statsd, and Go 1.7 context.
- [stats(stars: 593)](https://github.com/thoas/stats) - Go middleware that stores various information about your web application.

**[⬆ back to top](#contents)**

### Routers

- [alien(stars: 127)](https://github.com/gernest/alien) - Lightweight and fast http router from outer space.
- [bellt(stars: 54)](https://github.com/GuilhermeCaruso/bellt) - A simple Go HTTP router.
- [Bone(stars: 1295)](https://github.com/go-zoo/bone) - Lightning Fast HTTP Multiplexer.
- [Bxog(stars: 104)](https://github.com/claygod/Bxog) - Simple and fast HTTP router for Go. It works with routes of varying difficulty, length and nesting. And he knows how to create a URL from the received parameters.
- [chi(stars: 16199)](https://github.com/go-chi/chi) - Small, fast and expressive HTTP router built on net/context.
- [fasthttprouter(stars: 874)](https://github.com/buaazp/fasthttprouter) - High performance router forked from `httprouter`. The first router fit for `fasthttp`.
- [FastRouter(stars: 23)](https://github.com/razonyang/fastrouter) - a fast, flexible HTTP router written in Go.
- [goblin(stars: 72)](https://github.com/bmf-san/goblin) - A golang http router based on trie tree.
- [gocraft/web(stars: 1500)](https://github.com/gocraft/web) - Mux and middleware package in Go.
- [Goji(stars: 947)](https://github.com/goji/goji) - Goji is a minimalistic and flexible HTTP request multiplexer with support for `net/context`.
- [GoLobby/Router(stars: 21)](https://github.com/golobby/router) - GoLobby Router is a lightweight yet powerful HTTP router for the Go programming language.
- [goroute(stars: 9)](https://github.com/goroute/route) - Simple yet powerful HTTP request multiplexer.
- [GoRouter(stars: 150)](https://github.com/vardius/gorouter) - GoRouter is a Server/API micro framework, HTTP request router, multiplexer, mux that provides request router with middleware supporting `net/context`.
- [gowww/router(stars: 186)](https://github.com/gowww/router) - Lightning fast HTTP router fully compatible with the net/http.Handler interface.
- [httprouter(stars: 16021)](https://github.com/julienschmidt/httprouter) - High performance router. Use this and the standard http handlers to form a very high performance web framework.
- [httptreemux(stars: 610)](https://github.com/dimfeld/httptreemux) - High-speed, flexible tree-based HTTP router for Go. Inspiration from httprouter.
- [lars(stars: 389)](https://github.com/go-playground/lars) - Is a lightweight, fast and extensible zero allocation HTTP router for Go used to create customizable frameworks.
- [mux(stars: 19692)](https://github.com/gorilla/mux) - Powerful URL router and dispatcher for golang.
- [nchi(stars: 13)](https://github.com/muir/nchi) - chi-like router built on httprouter with dependency injection based middleware wrappers
- [ngamux(stars: 62)](https://github.com/ngamux/ngamux) - Simple HTTP router for Go.
- [ozzo-routing(stars: 446)](https://github.com/go-ozzo/ozzo-routing) - An extremely fast Go (golang) HTTP router that supports regular expression route matching. Comes with full support for building RESTful APIs.
- [pure(stars: 145)](https://github.com/go-playground/pure) - Is a lightweight HTTP router that sticks to the std "net/http" implementation.
- [Siesta(stars: 350)](https://github.com/VividCortex/siesta) - Composable framework to write middleware and handlers.
- [vestigo(stars: 267)](https://github.com/husobee/vestigo) - Performant, stand-alone, HTTP compliant URL Router for go web applications.
- [violetear(stars: 107)](https://github.com/nbari/violetear) - Go HTTP router.
- [xmux(stars: 98)](https://github.com/rs/xmux) - High performance muxer based on `httprouter` with `net/context` support.
- [xujiajun/gorouter(stars: 531)](https://github.com/xujiajun/gorouter) - A simple and fast HTTP router for Go.

**[⬆ back to top](#contents)**

## WebAssembly

- [dom(stars: 482)](https://github.com/dennwc/dom) - DOM library.
- [go-canvas(stars: 217)](https://github.com/markfarnan/go-canvas) - Library to use HTML5 Canvas, with all drawing within go code.
- [tinygo(stars: 13873)](https://github.com/tinygo-org/tinygo) - Go compiler for small places. Microcontrollers, WebAssembly, and command-line tools. Based on LLVM.
- [vert(stars: 94)](https://github.com/norunners/vert) - Interop between Go and JS values.
- [wasmbrowsertest(stars: 162)](https://github.com/agnivade/wasmbrowsertest) - Run Go WASM tests in your browser.
- [webapi(stars: 149)](https://github.com/gowebapi/webapi) - Bindings for DOM and HTML generated from WebIDL.

**[⬆ back to top](#contents)**

## Windows

- [d3d9(stars: 152)](https://github.com/gonutz/d3d9) - Go bindings for Direct3D9.
- [go-ole(stars: 1062)](https://github.com/go-ole/go-ole) - Win32 OLE implementation for golang.
- [gosddl(stars: 10)](https://github.com/MonaxGT/gosddl) - Converter from SDDL-string to user-friendly JSON. SDDL consist of four part: Owner, Primary Group, DACL, SACL.

**[⬆ back to top](#contents)**

## XML

_Libraries and tools for manipulating XML._

- [XML-Comp(stars: 21)](https://github.com/xml-comp/xml-comp) - Simple command line XML comparer that generates diffs of folders, files and tags.
- [xml2map(stars: 59)](https://github.com/sbabiv/xml2map) - XML to MAP converter written Golang.
- [xmlwriter(stars: 25)](https://github.com/shabbyrobe/xmlwriter) - Procedural XML generation API based on libxml2's xmlwriter module.
- [xpath(stars: 620)](https://github.com/antchfx/xpath) - XPath package for Go.
- [xquery(stars: 157)](https://github.com/antchfx/xquery) - XQuery lets you extract data from HTML/XML documents using XPath expression.
- [zek(stars: 710)](https://github.com/miku/zek) - Generate a Go struct from XML.

## Zero Trust

_Libraries and tools to implement Zero Trust architectures._

- [Cosign(stars: 3800)](https://github.com/sigstore/cosign) - Container Signing, Verification and Storage in an OCI registry.
- [in-toto(stars: 112)](https://github.com/in-toto/in-toto-golang) - Go implementation of the in-toto (provides a framework to protect the integrity of the software supply chain) python reference implementation.
- [Spiffe-Vault(stars: 66)](https://github.com/philips-labs/spiffe-vault) - Utilizes Spiffe JWT authentication with Hashicorp Vault for secretless authentication.
- [Spire(stars: 1592)](https://github.com/spiffe/spire) - SPIRE (the SPIFFE Runtime Environment) is a toolchain of APIs for establishing trust between software systems across a wide variety of hosting platforms.

## Code Analysis

_Source code analysis tools, also known as Static Application Security Testing (SAST) Tools._

- [apicompat(stars: 177)](https://github.com/bradleyfalzon/apicompat) - Checks recent changes to a Go project for backwards incompatible changes.
- [asty(stars: 66)](https://github.com/asty-org/asty) - Converts golang AST to JSON and JSON to AST.
- [blanket](https://gitlab.com/verygoodsoftwarenotvirus/blanket) - blanket is a tool that helps you catch functions which don't have direct unit tests in your Go packages.
- [ChainJacking(stars: 53)](https://github.com/Checkmarx/chainjacking) - Find which of your Go lang direct GitHub dependencies is susceptible to ChainJacking attack.
- [Chronos(stars: 413)](https://github.com/amit-davidson/Chronos) - Detects race conditions statically
- [dupl(stars: 324)](https://github.com/mibk/dupl) - Tool for code clone detection.
- [errcheck(stars: 2174)](https://github.com/kisielk/errcheck) - Errcheck is a program for checking for unchecked errors in Go programs.
- [gcvis(stars: 1093)](https://github.com/davecheney/gcvis) - Visualise Go program GC trace data in real time.
- [go-checkstyle(stars: 128)](https://github.com/qiniu/checkstyle) - checkstyle is a style check tool like java checkstyle. This tool inspired by java checkstyle, golint. The style referred to some points in Go Code Review Comments.
- [go-cleanarch(stars: 795)](https://github.com/roblaszczak/go-cleanarch) - go-cleanarch was created to validate Clean Architecture rules, like a The Dependency Rule and interaction between packages in your Go projects.
- [go-critic(stars: 1712)](https://github.com/go-critic/go-critic) - source code linter that brings checks that are currently not implemented in other linters.
- [go-mod-outdated(stars: 641)](https://github.com/psampaz/go-mod-outdated) - An easy way to find outdated dependencies of your Go projects.
- [go-outdated(stars: 45)](https://github.com/firstrow/go-outdated) - Console application that displays outdated packages.
- [goast-viewer(stars: 740)](https://github.com/yuroyoro/goast-viewer) - Web based Golang AST visualizer.
- [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) - Tool to fix (add, remove) your Go imports automatically.
- [golang-ifood-sdk(stars: 11)](https://github.com/arxdsilva/golang-ifood-sdk) - iFood API SDK.
- [golangci-lint](https://github.com/golangci/golangci-lint) – A fast Go linters runner. It runs linters in parallel, uses caching, supports `yaml` config, has integrations with all major IDE and has dozens of linters included.
- [golines(stars: 744)](https://github.com/segmentio/golines) - Formatter that automatically shortens long lines in Go code.
- [GoPlantUML(stars: 1609)](https://github.com/jfeliu007/goplantuml) - Library and CLI that generates text plantump class diagram containing information about structures and interfaces with the relationship among them.
- [goreturns](https://sourcegraph.com/github.com/sqs/goreturns) - Adds zero-value return statements to match the func return types.
- [gostatus(stars: 242)](https://github.com/shurcooL/gostatus) - Command line tool, shows the status of repositories that contain Go packages.
- [lint(stars: 66)](https://github.com/surullabs/lint) - Run linters as part of go test.
- [php-parser(stars: 935)](https://github.com/z7zmey/php-parser) - A Parser for PHP written in Go.
- [revive](https://github.com/mgechev/revive) – ~6x faster, stricter, configurable, extensible, and beautiful drop-in replacement for `golint`.
- [staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) - staticcheck is `go vet` on steroids, applying a ton of static analysis checks you might be used to from tools like ReSharper for C#.
- [testifylint](https://github.com/Antonboom/testifylint) – A linter that checks usage of [github.com/stretchr/testify](https://github.com/stretchr/testify).
- [tickgit(stars: 316)](https://github.com/augmentable-dev/tickgit) - CLI and go package for surfacing code comment TODOs (in any language) and applying a `git blame`to identify the author.
- [todocheck(stars: 406)](https://github.com/preslavmihaylov/todocheck) - Static code analyser which links TODO comments in code with issues in your issue tracker.
- [unconvert(stars: 368)](https://github.com/mdempsky/unconvert) - Remove unnecessary type conversions from Go source.
- [usestdlibvars(stars: 41)](https://github.com/sashamelentyev/usestdlibvars) - A linter that detect the possibility to use variables/constants from the Go standard library.
- [vacuum(stars: 336)](https://github.com/daveshanley/vacuum) - An ultra-super-fast, lightweight OpenAPI linter and quality checking tool.
- [validate(stars: 61)](https://github.com/mccoyst/validate) - Automatically validates struct fields with tags.

**[⬆ back to top](#contents)**

## Editor Plugins

_Plugin for text editors and IDEs._

- [coc-go language server extension for Vim/Neovim(stars: 554)](https://github.com/josa42/coc-go) - This plugin adds [gopls](https://github.com/golang/tools/blob/master/gopls/README.md) features to Vim/Neovim.
- [Go Doc(stars: 7)](https://github.com/msyrus/vscode-go-doc) - A Visual Studio Code extension for showing definition in output and generating go doc.
- [Go plugin for JetBrains IDEs](https://plugins.jetbrains.com/plugin/9568-go) - Go plugin for JetBrains IDEs.
- [go-language-server(stars: 32)](https://github.com/theia-ide/go-language-server) - A wrapper to turn the VSCode go extension into a language server supporting the language-server-protocol.
- [go-mode](https://github.com/dominikh/go-mode.el) - Go mode for GNU/Emacs.
- [go-plus(stars: 1512)](https://github.com/joefitzgerald/go-plus) - Go (Golang) Package For Atom That Adds Autocomplete, Formatting, Syntax Checking, Linting and Vetting.
- [gocode(stars: 5010)](https://github.com/nsf/gocode) - Autocompletion daemon for the Go programming language.
- [goimports-reviser(stars: 480)](https://github.com/incu6us/goimports-reviser) - Formatting tool for imports.
- [goprofiling](https://marketplace.visualstudio.com/items?itemName=MaxMedia.go-prof) - This extension adds benchmark profiling support for the Go language to VS Code.
- [GoSublime(stars: 3421)](https://github.com/DisposaBoy/GoSublime) - Golang plugin collection for the text editor SublimeText 3 providing code completion and other IDE-like features.
- [gounit-vim(stars: 22)](https://github.com/hexdigest/gounit-vim) - Vim plugin for generating Go tests based on the function's or method's signature.
- [theia-go-extension(stars: 15)](https://github.com/theia-ide/theia-go-extension) - Go language support for the Theia IDE.
- [vim-compiler-go(stars: 87)](https://github.com/rjohnsondev/vim-compiler-go) - Vim plugin to highlight syntax errors on save.
- [vim-go(stars: 15704)](https://github.com/fatih/vim-go) - Go development plugin for Vim.
- [vscode-go(stars: 3632)](https://github.com/golang/vscode-go) - Extension for Visual Studio Code (VS Code) which provides support for the Go language.
- [Watch(stars: 199)](https://github.com/eaburns/Watch) - Runs a command in an acme win on file changes.

**[⬆ back to top](#contents)**

## Go Generate Tools

- [envdoc(stars: 13)](https://github.com/g4s8/envdoc) -  generate documentation for environment variables from Go source files.
- [generic(stars: 48)](https://github.com/usk81/generic) - flexible data type for Go.
- [genny(stars: 1689)](https://github.com/cheekybits/genny) - Elegant generics for Go.
- [gocontracts(stars: 104)](https://github.com/Parquery/gocontracts) - brings design-by-contract to Go by synchronizing the code with the documentation.
- [godal(stars: 17)](https://github.com/mafulong/godal) - Generate orm models corresponding to golang by specifying sql ddl file, which can be used by gorm.
- [gonerics(stars: 114)](https://github.com/bouk/gonerics) - Idiomatic Generics in Go.
- [gotests(stars: 4770)](https://github.com/cweill/gotests) - Generate Go tests from your source code.
- [gounit(stars: 75)](https://github.com/hexdigest/gounit) - Generate Go tests using your own templates.
- [hasgo(stars: 133)](https://github.com/DylanMeeus/hasgo) - Generate Haskell inspired functions for your slices.
- [options-gen(stars: 47)](https://github.com/kazhuravlev/options-gen) - Functional options described by Dave Cheney's post "Functional options for friendly APIs".
- [re2dfa](https://gitlab.com/opennota/re2dfa) - Transform regular expressions into finite state machines and output Go source code.
- [sqlgen(stars: 71)](https://github.com/anqiansong/sqlgen) - Generate gorm, xorm, sqlx, bun, sql code from SQL file or DSN.
- [TOML-to-Go](https://xuri.me/toml-to-go) - Translates TOML into a Go type in the browser instantly.
- [xgen(stars: 278)](https://github.com/xuri/xgen) - XSD (XML Schema Definition) parser and Go/C/Java/Rust/TypeScript code generator.

**[⬆ back to top](#contents)**

## Go Tools

- [colorgo(stars: 113)](https://github.com/songgao/colorgo) - Wrapper around `go` command for colorized `go build` output.
- [depth(stars: 863)](https://github.com/KyleBanks/depth) - Visualize dependency trees of any package by analyzing imports.
- [docs(stars: 30)](https://github.com/go-oas/docs) - Automatically generate RESTful API documentation for GO projects - aligned with Open API Specification standard.
- [go-callvis(stars: 5553)](https://github.com/TrueFurby/go-callvis) - Visualize call graph of your Go program using dot format.
- [go-james(stars: 61)](https://github.com/pieterclaerhout/go-james) - Go project skeleton creator, builds and tests your projects without the manual setup.
- [go-pkg-complete(stars: 43)](https://github.com/skelterjohn/go-pkg-complete) - Bash completion for go and wgo.
- [go-swagger(stars: 9062)](https://github.com/go-swagger/go-swagger) - Swagger 2.0 implementation for go. Swagger is a simple yet powerful representation of your RESTful API.
- [godbg(stars: 199)](https://github.com/tylerwince/godbg) - Implementation of Rusts `dbg!` macro for quick and easy debugging during development.
- [gomodrun](https://github.com/dustinblackman/gomodrun/) - Go tool that executes and caches binaries included in go.mod files.
- [gotemplate.io](https://gotemplate.io/) - Online tool to preview `text/template` templates live.
- [gotestdox(stars: 73)](https://github.com/bitfield/gotestdox) - Show Go test results as readable sentences.
- [gothanks(stars: 118)](https://github.com/psampaz/gothanks) - GoThanks automatically stars your go.mod github dependencies, sending this way some love to their maintainers.
- [igo(stars: 62)](https://github.com/rocketlaunchr/igo) - An igo to go transpiler (new language features for Go language!)
- [modver(stars: 12)](https://github.com/bobg/modver) - Compare two versions of a Go module to check the version-number change required (major, minor, or patchlevel), according to [semver](https://semver.org/) rules.
- [OctoLinker(stars: 5212)](https://github.com/OctoLinker/browser-extension) - Navigate through go files efficiently with the OctoLinker browser extension for GitHub.
- [richgo(stars: 822)](https://github.com/kyoh86/richgo) - Enrich `go test` outputs with text decorations.
- [roumon(stars: 158)](https://github.com/becheran/roumon) - Monitor current state of all active goroutines via a command line interface.
- [rts(stars: 248)](https://github.com/galeone/rts) - RTS: response to struct. Generates Go structs from server responses.
- [textra(stars: 5)](https://github.com/ravsii/textra) - Extract Go struct field names, types and tags for filtering and exporting.
- [typex(stars: 185)](https://github.com/dtgorski/typex) - Examine Go types and their transitive dependencies, alternatively export results as TypeScript value objects (or types) declaration.

**[⬆ back to top](#contents)**

## Software Packages

_Software written in Go._

**[⬆ back to top](#contents)**

### DevOps Tools

- [abbreviate(stars: 207)](https://github.com/dnnrly/abbreviate) - abbreviate is a tool turning long strings in to shorter ones with configurable separators, for example to embed branch names in to deployment stack IDs.
- [aptly(stars: 10)](https://github.com/smira/aptly) - aptly is a Debian repository management tool.
- [aurora(stars: 588)](https://github.com/xuri/aurora) - Cross-platform web-based Beanstalkd queue server console.
- [awsenv(stars: 34)](https://github.com/soniah/awsenv) - Small binary that loads Amazon (AWS) environment variables for a profile.
- [Balerter(stars: 295)](https://github.com/balerter/balerter) - A self-hosted script-based alerting manager.
- [Blast(stars: 214)](https://github.com/dave/blast) - A simple tool for API load testing and batch jobs.
- [bombardier(stars: 4972)](https://github.com/codesenberg/bombardier) - Fast cross-platform HTTP benchmarking tool.
- [bosun(stars: 3364)](https://github.com/bosun-monitor/bosun) - Time Series Alerting Framework.
- [cassowary(stars: 699)](https://github.com/rogerwelin/cassowary) - Modern cross-platform HTTP load-testing tool written in Go.
- [Ddosify(stars: 8011)](https://github.com/ddosify/ddosify) - High-performance load testing tool, written in Golang.
- [decompose(stars: 59)](https://github.com/s0rg/decompose) - tool to generate and process Docker containers connections graphs.
- [DepCharge(stars: 23)](https://github.com/centerorbit/depcharge) - Helps orchestrating the execution of commands across the many dependencies in larger projects.
- [Docker](https://www.docker.com/) - Open platform for distributed applications for developers and sysadmins.
- [docker-go-mingw(stars: 40)](https://github.com/x1unix/docker-go-mingw) - Docker image for building Go binaries for Windows with MinGW toolchain.
- [Dockerfile-Generator(stars: 160)](https://github.com/ozankasikci/dockerfile-generator) - A go library and an executable that produces valid Dockerfiles using various input channels.
- [dogo(stars: 258)](https://github.com/liudng/dogo) - Monitoring changes in the source file and automatically compile and run (restart).
- [drone-jenkins(stars: 38)](https://github.com/appleboy/drone-jenkins) - Trigger downstream Jenkins jobs using a binary, docker or Drone CI.
- [drone-scp(stars: 131)](https://github.com/appleboy/drone-scp) - Copy files and artifacts via SSH using a binary, docker or Drone CI.
- [Dropship(stars: 64)](https://github.com/chrismckenzie/dropship) - Tool for deploying code via cdn.
- [easyssh-proxy(stars: 299)](https://github.com/appleboy/easyssh-proxy) - Golang package for easy remote execution through SSH and SCP downloading via `ProxyCommand`.
- [fac(stars: 1822)](https://github.com/mkchoi212/fac) - Command-line user interface to fix git merge conflicts.
- [Flannel(stars: 8327)](https://github.com/flannel-io/flannel) - Flannel is a network fabric for containers, designed for Kubernetes.
- [Fleet device management(stars: 1766)](https://github.com/fleetdm/fleet) - Lightweight, programmable telemetry for servers and workstations.
- [gaia(stars: 5137)](https://github.com/gaia-pipeline/gaia) - Build powerful pipelines in any programming language.
- [ghorg(stars: 1390)](https://github.com/gabrie30/ghorg) - Quickly clone an entire org/users repositories into one directory - Supports GitHub, GitLab, Gitea, and Bitbucket.
- [Gitea(stars: 39884)](https://github.com/go-gitea/gitea) - Fork of Gogs, entirely community driven.
- [gitea-github-migrator](https://git.jonasfranz.software/JonasFranzDEV/gitea-github-migrator) - Migrate all your GitHub repositories, issues, milestones and labels to your Gitea instance.
- [go-furnace(stars: 97)](https://github.com/go-furnace/go-furnace) - Hosting solution written in Go. Deploy your Application with ease on AWS, GCP or DigitalOcean.
- [go-rocket-update(stars: 89)](https://github.com/mouuff/go-rocket-update) - A simple way to make self updating Go applications - Supports Github and Gitlab.
- [go-selfupdate(stars: 1353)](https://github.com/sanbornm/go-selfupdate) - Enable your Go applications to self update.
- [gobrew(stars: 194)](https://github.com/cryptojuice/gobrew) - gobrew lets you easily switch between multiple versions of go.
- [gobrew(stars: 311)](https://github.com/kevincobain2000/gobrew) - Go version manager. Super simple tool to install and manage Go versions. Install go without root. Gobrew doesn't require shell rehash.
- [godbg(stars: 226)](https://github.com/sirnewton01/godbg) - Web-based gdb front-end application.
- [Gogs](https://gogs.io/) - A Self Hosted Git Service in the Go Programming Language.
- [gonative(stars: 337)](https://github.com/inconshreveable/gonative) - Tool which creates a build of Go that can cross compile to all platforms while still using the Cgo-enabled versions of the stdlib packages.
- [govvv(stars: 534)](https://github.com/ahmetalpbalkan/govvv) - “go build” wrapper to easily add version information into Go binaries.
- [gox(stars: 4582)](https://github.com/mitchellh/gox) - Dead simple, no frills Go cross compile tool.
- [goxc(stars: 1682)](https://github.com/laher/goxc) - build tool for Go, with a focus on cross-compiling and packaging.
- [grapes(stars: 166)](https://github.com/yaronsumel/grapes) - Lightweight tool designed to distribute commands over ssh with ease.
- [GVM(stars: 9169)](https://github.com/moovweb/gvm) - GVM provides an interface to manage Go versions.
- [Hey(stars: 16740)](https://github.com/rakyll/hey) - Hey is a tiny program that sends some load to a web application.
- [httpref(stars: 35)](https://github.com/dnnrly/httpref) - httpref is a handy CLI reference for HTTP methods, status codes, headers, and TCP and UDP ports.
- [jcli(stars: 374)](https://github.com/jenkins-zh/jenkins-cli) - Jenkins CLI allows you manage your Jenkins as an easy way.
- [k3d(stars: 4841)](https://github.com/k3d-io/k3d) - Little helper to run CNCF's k3s in Docker.
- [k3s(stars: 25367)](https://github.com/k3s-io/k3s) - Lightweight Kubernetes.
- [k6(stars: 22324)](https://github.com/grafana/k6) - A modern load testing tool, using Go and JavaScript.
- [kala(stars: 2059)](https://github.com/ajvb/kala) - Simplistic, modern, and performant job scheduler.
- [kcli(stars: 202)](https://github.com/cswank/kcli) - Command line tool for inspecting kafka topics/partitions/messages.
- [kind(stars: 12345)](https://github.com/kubernetes-sigs/kind) - Kubernetes IN Docker - local clusters for testing Kubernetes.
- [ko(stars: 6844)](https://github.com/google/ko) - Command line tool for building and deploying Go applications on Kubernetes
- [kool(stars: 657)](https://github.com/kool-dev/kool) - Command line tool for managing Docker environments as an easy way.
- [kubernetes(stars: 104106)](https://github.com/kubernetes/kubernetes) - Container Cluster Manager from Google.
- [kubeshark(stars: 10178)](https://github.com/kubeshark/kubeshark) - API traffic analyzer for Kubernetes, inspired by Wireshark, purposely built for Kubernetes.
- [KubeVela(stars: 5827)](https://github.com/kubevela/kubevela) - Cloud native application delivery.
- [kwatch(stars: 867)](https://github.com/abahmed/kwatch) - Monitor & detect crashes in your Kubernetes(K8s) cluster instantly.
- [lstags(stars: 315)](https://github.com/ivanilves/lstags) - Tool and API to sync Docker images across different registries.
- [lwc(stars: 32)](https://github.com/timdp/lwc) - A live-updating version of the UNIX wc command.
- [manssh(stars: 291)](https://github.com/xwjdsh/manssh) - manssh is a command line tool for managing your ssh alias config easily.
- [Mantil(stars: 105)](https://github.com/mantil-io/mantil) - Go specific framework for building serverless applications on AWS that enables you to focus on pure Go code while Mantil takes care of the infrastructure.
- [minikube(stars: 27758)](https://github.com/kubernetes/minikube) - Run Kubernetes locally.
- [Moby(stars: 67103)](https://github.com/moby/moby) - Collaborative project for the container ecosystem to assemble container-based systems.
- [Mora(stars: 313)](https://github.com/emicklei/mora) - REST server for accessing MongoDB documents and meta data.
- [ostent(stars: 177)](https://github.com/ostrost/ostent) - collects and displays system metrics and optionally relays to Graphite and/or InfluxDB.
- [Packer(stars: 14734)](https://github.com/mitchellh/packer) - Packer is a tool for creating identical machine images for multiple platforms from a single source configuration.
- [Pewpew(stars: 395)](https://github.com/bengadbois/pewpew) - Flexible HTTP command line stress tester.
- [PipeCD(stars: 910)](https://github.com/pipe-cd/pipecd) - A GitOps-style continuous delivery platform that provides consistent deployment and operations experience for any applications.
- [Pomerium(stars: 3745)](https://github.com/pomerium/pomerium) - Pomerium is an identity-aware access proxy.
- [Rodent(stars: 33)](https://github.com/alouche/rodent) - Rodent helps you manage Go versions, projects and track dependencies.
- [s3-proxy(stars: 237)](https://github.com/oxyno-zeta/s3-proxy) - S3 Proxy with GET, PUT and DELETE methods and authentication (OpenID Connect and Basic Auth).
- [s3gof3r(stars: 1144)](https://github.com/rlmcpherson/s3gof3r) - Small utility/library optimized for high speed transfer of large objects into and out of Amazon S3.
- [s5cmd(stars: 2113)](https://github.com/peak/s5cmd) - Blazing fast S3 and local filesystem execution tool.
- [Scaleway-cli(stars: 862)](https://github.com/scaleway/scaleway-cli) - Manage BareMetal Servers from Command Line (as easily as with Docker).
- [script(stars: 4760)](https://github.com/bitfield/script) - Making it easy to write shell-like scripts in Go for DevOps and system administration tasks.
- [sg(stars: 8)](https://github.com/ChristopherRabotin/sg) - Benchmarks a set of HTTP endpoints (like ab), with possibility to use the response code and data between each call for specific server stress based on its previous response.
- [skm(stars: 915)](https://github.com/TimothyYe/skm) - SKM is a simple and powerful SSH Keys Manager, it helps you to manage your multiple SSH keys easily!
- [StatusOK(stars: 1599)](https://github.com/sanathp/statusok) - Monitor your Website and REST APIs.Get Notified through Slack, E-mail when your server is down or response time is more than expected.
- [terraform-provider-openapi(stars: 266)](https://github.com/dikhan/terraform-provider-openapi) - Terraform provider plugin that dynamically configures itself at runtime based on an OpenAPI document (formerly known as swagger file) containing the definitions of the APIs exposed.
- [tf-profile(stars: 135)](https://github.com/datarootsio/tf-profile) - Profiler for Terraform runs. Generate global stats, resource-level stats or visualizations.
- [traefik(stars: 45985)](https://github.com/containous/traefik) - Reverse proxy and load balancer with support for multiple backends.
- [trubka(stars: 330)](https://github.com/xitonix/trubka) - A CLI tool to manage and troubleshoot Apache Kafka clusters with the ability of generically publishing/consuming protocol buffer and plain text events to/from Kafka.
- [uTask(stars: 1039)](https://github.com/ovh/utask) - Automation engine that models and executes business processes declared in yaml.
- [Vegeta(stars: 22295)](https://github.com/tsenart/vegeta) - HTTP load testing tool and library. It's over 9000!
- [wait-for(stars: 18)](https://github.com/dnnrly/wait-for) - Wait for something to happen (from the command line) before continuing. Easy orchestration of Docker services and other things.
- [webhook(stars: 9531)](https://github.com/adnanh/webhook) - Tool which allows user to create HTTP endpoints (hooks) that execute commands on the server.
- [Wide](https://wide.b3log.org/login) - Web-based IDE for Teams using Golang.
- [winrm-cli(stars: 160)](https://github.com/masterzen/winrm-cli) - Cli tool to remotely execute commands on Windows machines.

**[⬆ back to top](#contents)**

### Other Software

- [Better Go Playground](https://goplay.tools) - Go playground with syntax highlight, code completion and other features.
- [blocky(stars: 3228)](https://github.com/0xERR0R/blocky) - Fast and lightweight DNS proxy as ad-blocker for local network with many features.
- [borg(stars: 1593)](https://github.com/crufter/borg) - Terminal based search engine for bash snippets.
- [boxed(stars: 80)](https://github.com/tejo/boxed) - Dropbox based blog engine.
- [Cherry(stars: 295)](https://github.com/rafael-santiago/cherry) - Tiny webchat server in Go.
- [Circuit(stars: 1972)](https://github.com/gocircuit/circuit) - Circuit is a programmable platform-as-a-service (PaaS) and/or Infrastructure-as-a-Service (IaaS), for management, discovery, synchronization and orchestration of services and hosts comprising cloud applications.
- [Comcast(stars: 10136)](https://github.com/tylertreat/Comcast) - Simulate bad network connections.
- [confd(stars: 8222)](https://github.com/kelseyhightower/confd) - Manage local application configuration files using templates and data from etcd or consul.
- [crawley(stars: 211)](https://github.com/s0rg/crawley) - Web scraper/crawler for cli.
- [croc(stars: 25323)](https://github.com/schollz/croc) - Easily and securely send files or folders from one computer to another.
- [Documize(stars: 2003)](https://github.com/documize/community) - Modern wiki software that integrates data from SaaS tools.
- [dp(stars: 82)](https://github.com/scryinfo/dp) - Through SDK for data exchange with blockchain, developers can get easy access to DAPP development.
- [drive(stars: 6590)](https://github.com/odeke-em/drive) - Google Drive client for the commandline.
- [Duplicacy(stars: 4845)](https://github.com/gilbertchen/duplicacy) - A cross-platform network and cloud backup tool based on the idea of lock-free deduplication.
- [fjira(stars: 59)](https://github.com/mk-5/fjira) - A fuzzy-search based terminal UI application for Attlasian Jira
- [Gebug(stars: 626)](https://github.com/moshebe/gebug) - A tool that makes debugging of Dockerized Go applications super easy by enabling Debugger and Hot-Reload features, seamlessly.
- [gfile(stars: 718)](https://github.com/Antonito/gfile) - Securely transfer files between two computers, without any third party, over WebRTC.
- [Go Package Store(stars: 898)](https://github.com/shurcooL/Go-Package-Store) - App that displays updates for the Go packages in your GOPATH.
- [go-peerflix(stars: 461)](https://github.com/Sioro-Neoku/go-peerflix) - Video streaming torrent client.
- [goblin](https://goblin.reaper.im) - Golang binaries in a curl, built by goblins.
- [GoBoy(stars: 2566)](https://github.com/Humpheh/goboy) - Nintendo Game Boy Color emulator written in Go.
- [gocc(stars: 582)](https://github.com/goccmack/gocc) - Gocc is a compiler kit for Go written in Go.
- [GoDocTooltip(stars: 13)](https://github.com/diankong/GoDocTooltip) - Chrome extension for Go Doc sites, which shows function description as tooltip at function list.
- [Gokapi(stars: 850)](https://github.com/Forceu/gokapi) - Lightweight server to share files, which expire after a set amount of downloads or days. Similar to Firefox Send, but without public upload.
- [GoLand](https://jetbrains.com/go) - Full featured cross-platform Go IDE.
- [GoNB(stars: 337)](https://github.com/janpfeifer/gonb) - Interactive Go programming with Jupyter Notebooks (also works in VSCode, Binder and Google's Colab).
- [Gor(stars: 18017)](https://github.com/buger/gor) - Http traffic replication tool, for replaying traffic from production to stage/dev environments in real-time.
- [Guora(stars: 660)](https://github.com/meloalright/guora) - A self-hosted Quora like web application written in Go.
- [hoofli(stars: 8)](https://github.com/dnnrly/hoofli) - Generate PlantUML diagrams from Chrome or Firefox network inspections.
- [hotswap(stars: 271)](https://github.com/edwingeng/hotswap) - A complete solution to reload your go code without restarting your server, interrupting or blocking any ongoing procedure.
- [hugo](https://gohugo.io/) - Fast and Modern Static Website Engine.
- [ide(stars: 357)](https://github.com/thestrukture/ide) - Browser accessible IDE. Designed for Go with Go.
- [ipe(stars: 366)](https://github.com/dimiro1/ipe) - Open source Pusher server implementation compatible with Pusher client libraries written in GO.
- [joincap(stars: 194)](https://github.com/assafmo/joincap) - Command-line utility for merging multiple pcap files together.
- [Juju](https://jujucharms.com/) - Cloud-agnostic service deployment and orchestration - supports EC2, Azure, Openstack, MAAS and more.
- [Leaps(stars: 738)](https://github.com/jeffail/leaps) - Pair programming service using Operational Transforms.
- [lgo(stars: 2335)](https://github.com/yunabe/lgo) - Interactive Go programming with Jupyter. It supports code completion, code inspection and 100% Go compatibility.
- [limetext](https://limetext.github.io) - Lime Text is a powerful and elegant text editor primarily developed in Go that aims to be a Free and open-source software successor to Sublime Text.
- [LiteIDE(stars: 7381)](https://github.com/visualfc/liteide) - LiteIDE is a simple, open source, cross-platform Go IDE.
- [mockingjay(stars: 547)](https://github.com/quii/mockingjay-server) - Fake HTTP servers and consumer driven contracts from one configuration file. You can also make the server randomly misbehave to help do more realistic performance tests.
- [myLG(stars: 2677)](https://github.com/mehrdadrad/mylg) - Command Line Network Diagnostic tool written in Go.
- [naclpipe(stars: 23)](https://github.com/unix4fun/naclpipe) - Simple NaCL EC25519 based crypto pipe tool written in Go.
- [Neo-cowsay(stars: 272)](https://github.com/Code-Hex/Neo-cowsay) - 🐮 cowsay is reborn. for a New Era.
- [nes(stars: 5330)](https://github.com/fogleman/nes) - Nintendo Entertainment System (NES) emulator written in Go.
- [Orbit(stars: 177)](https://github.com/gulien/orbit) - A simple tool for running commands and generating files from templates.
- [peg(stars: 946)](https://github.com/pointlander/peg) - Peg, Parsing Expression Grammar, is an implementation of a Packrat parser generator.
- [Plik(stars: 1336)](https://github.com/root-gg/plik) - Plik is a temporary file upload system (Wetransfer like) in Go.
- [portal(stars: 1044)](https://github.com/SpatiumPortae/portal) - Portal is a quick and easy command-line file transfer utility from any computer to another.
- [protoncheck(stars: 4)](https://github.com/servusdei2018/protoncheck) - ProtonMail module for waybar/polybar/yabar/i3blocks.
- [restic(stars: 21987)](https://github.com/restic/restic) - De-duplicating backup program.
- [sake(stars: 632)](https://github.com/alajmo/sake) - sake is a command runner for local and remote hosts.
- [scc(stars: 5343)](https://github.com/boyter/scc) - Sloc Cloc and Code, a very fast accurate code counter with complexity calculations and COCOMO estimates.
- [Seaweed File System(stars: 18926)](https://github.com/chrislusf/seaweedfs) - Fast, Simple and Scalable Distributed File System with O(1) disk seek.
- [shell2http(stars: 1221)](https://github.com/msoap/shell2http) - Executing shell commands via http server (for prototyping or remote control).
- [snap(stars: 1799)](https://github.com/intelsdi-x/snap) - Powerful telemetry framework.
- [Snitch(stars: 17)](https://github.com/lucasgomide/snitch) - Simple way to notify your team and many tools when someone has deployed any application via Tsuru.
- [Stack Up(stars: 2453)](https://github.com/pressly/sup) - Stack Up, a super simple deployment tool - just Unix - think of it like 'make' for a network of servers.
- [stew(stars: 122)](https://github.com/marwanhawari/stew) - An independent package manager for compiled binaries.
- [syncthing](https://syncthing.net/) - Open, decentralized file synchronization tool and protocol.
- [tcpdog(stars: 247)](https://github.com/mehrdadrad/tcpdog) - eBPF based TCP observability.
- [tcpprobe(stars: 348)](https://github.com/mehrdadrad/tcpprobe) - TCP tool for network performance and path monitoring, including socket statistics.
- [term-quiz(stars: 26)](https://github.com/crazcalm/term-quiz) - Quizzes for your terminal.
- [toxiproxy(stars: 9983)](https://github.com/shopify/toxiproxy) - Proxy to simulate network and system conditions for automated tests.
- [tsuru](https://tsuru.io/) - Extensible and open source Platform as a Service software.
- [vaku(stars: 149)](https://github.com/lingrino/vaku) - CLI & API for folder-based functions in Vault like copy, move, and search.
- [vFlow(stars: 1030)](https://github.com/VerizonDigital/vflow) - High-performance, scalable and reliable IPFIX, sFlow and Netflow collector.
- [wellington(stars: 305)](https://github.com/wellington/wellington) - Sass project management tool, extends the language with sprite functions (like Compass).
- [woke(stars: 409)](https://github.com/get-woke/woke) - Detect non-inclusive language in your source code.
- [yai(stars: 448)](https://github.com/ekkinox/yai) - AI powered terminal assistant.
- [zs](https://git.mills.io/prologic/zs) - an extremely minimal static site generator.

**[⬆ back to top](#contents)**

# Resources

_Where to discover new Go libraries._

**[⬆ back to top](#contents)**

## Benchmarks

- [autobench(stars: 98)](https://github.com/davecheney/autobench) - Framework to compare the performance between different Go versions.
- [go-benchmark-app(stars: 27)](https://github.com/mrLSD/go-benchmark-app) - Powerful HTTP-benchmark tool mixed with Аb, Wrk, Siege tools. Gathering statistics and various parameters for benchmarks and comparison results.
- [go-benchmarks(stars: 145)](https://github.com/tylertreat/go-benchmarks) - Few miscellaneous Go microbenchmarks. Compare some language features to alternative approaches.
- [go-http-routing-benchmark(stars: 1627)](https://github.com/julienschmidt/go-http-routing-benchmark) - Go HTTP request router benchmark and comparison.
- [go-json-benchmark(stars: 9)](https://github.com/zerosnake0/go-json-benchmark) - Go JSON benchmark.
- [go-ml-benchmarks(stars: 29)](https://github.com/nikolaydubina/go-ml-benchmarks) - benchmarks for machine learning inference in Go.
- [go-web-framework-benchmark(stars: 1965)](https://github.com/smallnest/go-web-framework-benchmark) - Go web framework benchmark.
- [go_serialization_benchmarks(stars: 1503)](https://github.com/alecthomas/go_serialization_benchmarks) - Benchmarks of Go serialization methods.
- [gocostmodel(stars: 61)](https://github.com/PuerkitoBio/gocostmodel) - Benchmarks of common basic operations for the Go language.
- [golang-sql-benchmark(stars: 65)](https://github.com/tyler-smith/golang-sql-benchmark) - Collection of benchmarks for popular Go database/SQL utilities.
- [gospeed(stars: 113)](https://github.com/feyeleanor/GoSpeed) - Go micro-benchmarks for calculating the speed of language constructs.
- [kvbench(stars: 26)](https://github.com/jimrobinson/kvbench) - Key/Value database benchmark.
- [skynet(stars: 1031)](https://github.com/atemerev/skynet) - Skynet 1M threads microbenchmark.
- [speedtest-resize(stars: 237)](https://github.com/fawick/speedtest-resize) - Compare various Image resize algorithms for the Go language.

**[⬆ back to top](#contents)**

## Conferences

- [GoCon](https://gocon.connpass.com/) - Tokyo, Japan.
- [GoDays](https://www.godays.io/) - Berlin, Germany.
- [GoLab](https://golab.io/) - Florence, Italy.
- [GopherChina](https://gopherchina.org) - Shanghai, China.
- [GopherCon](https://www.gophercon.com/) - Denver, USA.
- [GopherCon Australia](https://gophercon.com.au/) - Sydney, Australia.
- [GopherCon Brazil](https://gopherconbr.org) - Florianópolis, Brazil.
- [GopherCon Europe](https://gophercon.eu/) - Berlin, Germany.
- [GopherCon India](https://gopherconindia.org/) - Pune, India.
- [GopherCon Israel](https://www.gophercon.org.il/) - Tel Aviv, Israel.
- [GopherCon Russia](https://www.gophercon-russia.ru) - Moscow, Russia.
- [GopherCon Singapore](https://gophercon.sg) - Mapletree Business City, Singapore.
- [GopherCon UK](https://www.gophercon.co.uk/) - London, UK.
- [GopherCon Vietnam](https://gophercon.vn/) - Ho Chi Minh City, Vietnam.
- [GoWest Conference](https://www.gowestconf.com/) - Lehi, USA.

**[⬆ back to top](#contents)**

## E-Books

### E-books for purchase

- [100 Go Mistakes: How to Avoid Them](https://www.manning.com/books/100-go-mistakes-how-to-avoid-them)
- [Black Hat Go](https://nostarch.com/blackhatgo) - Go programming for hackers and pentesters.
- [Build an Orchestrator in Go](https://www.manning.com/books/build-an-orchestrator-in-go)
- [Continuous Delivery in Go](https://www.manning.com/books/continuous-delivery-in-go) - This practical guide to continuous delivery shows you how to rapidly establish an automated pipeline that will improve your testing, code quality, and final product.
- [Creative DIY Microcontroller Project With TinyGo and WebAssembly](https://www.packtpub.com/product/creative-diy-microcontroller-projects-with-tinygo-and-webassembly/9781800560208) - An introduction into the TinyGo compiler with projects involving Arduino and WebAssembly.
- [Effective Go: Elegant, efficient, and testable code](https://www.manning.com/books/effective-go) - Unlock Go’s unique perspective on program design, and start writing simple, maintainable, and testable Go code.
- [For the Love of Go](https://bitfieldconsulting.com/books/love) - An introductory book for Go beginners.
- [Go Faster](https://leanpub.com/gofaster) - This book seeks to shorten your learning curve and help you become a proficient Go programmer, faster.
- [Know Go: Generics](https://bitfieldconsulting.com/books/generics) - A guide to understanding and using generics in Go.
- [Lets-Go](https://lets-go.alexedwards.net) - A step-by-step guide to creating fast, secure and maintanable web applications with Go.
- [Lets-Go-Further](https://lets-go-further.alexedwards.net) - Advanced patterns for building APIs and web applications in Go.
- [The Power of Go: Tests](https://bitfieldconsulting.com/books/tests) - A guide to testing in Go.
- [The Power of Go: Tools](https://bitfieldconsulting.com/books/tools) - A guide to writing command-line tools in Go.
- [Writing A Compiler In Go](https://compilerbook.com)
- [Writing An Interpreter In Go](https://interpreterbook.com) - Book that introduces dozens of techniques for writing idiomatic, expressive, and efficient Go code that avoids common pitfalls.

### Free e-books

- [A Go Developer's Notebook](https://leanpub.com/GoNotebook/read)
- [An Introduction to Programming in Go](http://www.golang-book.com/)
- [Build Web Application with Golang](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/)
- [Building Web Apps With Go](https://codegangsta.gitbooks.io/building-web-apps-with-go/content/)
- [Go 101](https://go101.org) - A book focusing on Go syntax/semantics and all kinds of details.
- [Go AST Book (Chinese)(stars: 5190)](https://github.com/chai2010/go-ast-book) - A book focusing on Go `go/*` packages.
- [Go Succinctly(stars: 23)](https://github.com/thedevsir/gosuccinctly) - in Persian.
- [Go with the domain](https://threedots.tech/go-with-the-domain/) - A book showing how to apply DDD, Clean Architecture, and CQRS by practical refactoring.
- [GoBooks(stars: 15354)](https://github.com/dariubs/GoBooks) - A curated list of Go books.
- [How To Code in Go eBook](https://www.digitalocean.com/community/books/how-to-code-in-go-ebook) - A 600 page introduction to Go aimed at first time developers.
- [Learning Go](https://www.miek.nl/downloads/Go/Learning-Go-latest.pdf)
- [Network Programming With Go](https://jan.newmarch.name/golang/)
- [Practical Go Lessons](https://www.practical-go-lessons.com/)
- [Spaceship Go A Journey to the Standard Library](https://blasrodri.github.io/spaceship-go-gh-pages/)
- [The Go Programming Language](https://www.gopl.io/)
- [The Golang Standard Library by Example (Chinese)](https://github.com/polaris1119/The-Golang-Standard-Library-by-Example)
- [Web Application with Go the Anti-Textbook](https://github.com/thewhitetulip/web-dev-golang-anti-textbook/)

**[⬆ back to top](#contents)**

## Gophers

- [Free Gophers Pack(stars: 3216)](https://github.com/MariaLetta/free-gophers-pack) - Gopher graphics pack by Maria Letta with illustrations and emotional characters in vector and raster.
- [Go-gopher-Vector(stars: 68)](https://github.com/keygx/Go-gopher-Vector) - Go gopher Vector Data [.ai, .svg].
- [gopher-logos(stars: 117)](https://github.com/GolangUA/gopher-logos) - adorable gopher logos.
- [gopher-stickers](https://github.com/tenntenn/gopher-stickers)
- [gophericons](https://github.com/shalakhin/gophericons)
- [gopherize.me](https://github.com/matryer/gopherize.me) - Gopherize yourself.
- [gophers(stars: 2877)](https://github.com/ashleymcnamara/gophers) - Gopher artworks by Ashley McNamara.
- [gophers(stars: 3288)](https://github.com/egonelbre/gophers) - Free gophers.
- [gophers(stars: 58)](https://github.com/rogeralsing/gophers) - random gopher graphics.
- [gophers(stars: 130)](https://github.com/sillecelik/go-gopher) - Gopher amigurumi toy pattern.
- [gophers(stars: 28)](https://github.com/scraly/gophers) - Gophers by Aurélie Vache.

**[⬆ back to top](#contents)**

## Meetups

- [Basel Go Meetup](https://www.meetup.com/Basel-Go-Meetup/)
- [Belfast Gophers](https://www.meetup.com/Belfast-Gophers/)
- [Belgrade Golang Meetup](https://www.meetup.com/golang-serbia/)
- [Berlin Golang](https://www.meetup.com/golang-users-berlin/)
- [Brisbane Gophers](https://www.meetup.com/Brisbane-Golang-Meetup/)
- [Bärner Go Meetup - Berne, Switzerland](https://www.meetup.com/berner-go-meetup/)
- [Canberra Gophers](https://www.meetup.com/Canberra-Gophers/)
- [Go Language NYC](https://www.meetup.com/golanguagenewyork/)
- [Go London User Group](https://www.meetup.com/Go-London-User-Group/)
- [Go Remote Meetup](https://www.meetup.com/Go-Remote-Meetup/)
- [Go Toronto](https://www.meetup.com/go-toronto/)
- [Go User Group Atlanta](https://www.meetup.com/Go-Users-Group-Atlanta/)
- [GoBandung](https://www.meetup.com/GoBandung/)
- [GoBridge, San Francisco, CA](https://www.meetup.com/gobridge/)
- [GoCracow - Krakow, Poland](https://www.meetup.com/GoCracow/)
- [GoJakarta](https://www.meetup.com/GoJakarta/)
- [Golang Amsterdam](https://www.meetup.com/golang-amsterdam/)
- [Golang Argentina](https://www.meetup.com/Golang-Argentina/)
- [Golang Athens](https://www.meetup.com/Athens-Gophers/)
- [Golang Baltimore, MD](https://www.meetup.com/BaltimoreGolang/)
- [Golang Bangalore](https://www.meetup.com/Golang-Bangalore/)
- [Golang Belo Horizonte - Brazil](https://www.meetup.com/go-belo-horizonte/)
- [Golang Boston](https://www.meetup.com/bostongo/)
- [Golang Bulgaria](https://www.meetup.com/Golang-Bulgaria/)
- [Golang Cardiff, UK](https://www.meetup.com/Cardiff-Go-Meetup/)
- [Golang Copenhagen](https://www.meetup.com/Go-Cph/)
- [Golang Curitiba - Brazil](https://www.meetup.com/GolangCWB/)
- [Golang DC, Arlington, VA](https://www.meetup.com/Golang-DC/)
- [Golang Dorset, UK](https://www.meetup.com/golang-dorset/)
- [Golang Estonia](https://www.meetup.com/Golang-Estonia/)
- [Golang Gurgaon, India](https://www.meetup.com/Gurgaon-Go-Meetup/)
- [Golang Hamburg - Germany](https://www.meetup.com/Go-User-Group-Hamburg/)
- [Golang Israel](https://www.meetup.com/Go-Israel/)
- [Golang Kathmandu](https://www.meetup.com/Golang-Kathmandu/)
- [Golang Korea](https://www.meetup.com/GDG-Golang-Korea/)
- [Golang Lima - Peru](https://www.meetup.com/Golang-Peru/)
- [Golang Lyon](https://www.meetup.com/Golang-Lyon/)
- [Golang Marseille](https://www.meetup.com/fr-FR/Golang-Marseille/)
- [Golang Melbourne](https://www.meetup.com/golang-mel/)
- [Golang Mountain View](https://www.meetup.com/Golang-Mountain-View/)
- [Golang North East](https://www.meetup.com/en-AU/Golang-North-East/)
- [Golang Paris](https://www.meetup.com/Golang-Paris/)
- [Golang Poland](https://www.meetup.com/Golang-Poland/)
- [Golang Pune](https://www.meetup.com/Golang-Pune/)
- [Golang Rotterdam](https://www.meetup.com/golang-rotterdam/)
- [Golang Singapore](https://www.meetup.com/golangsg/)
- [Golang Stockholm](https://www.meetup.com/Go-Stockholm/)
- [Golang Sydney, AU](https://www.meetup.com/golang-syd/)
- [Golang São Paulo - Brazil](https://www.meetup.com/golangbr/)
- [Golang Taipei](https://www.meetup.com/golang-taipei-meetup/)
- [Golang Thessaloniki](https://www.meetup.com/thessaloniki-golang-meetup/)
- [Golang Turkey](https://kommunity.com/goturkiye)
- [Golang Vancouver, BC](https://www.meetup.com/golangvan/)
- [Golang Vienna, Austria](https://www.meetup.com/viennago/)
- [Golang Москва](https://www.meetup.com/Golang-Moscow/)
- [GoSF - San Francisco, CA](https://www.meetup.com/golangsf)
- [Istanbul Golang](https://www.meetup.com/Istanbul-Golang/)
- [Seattle Go Programmers](https://www.meetup.com/golang/)
- [Ukrainian Golang User Groups](https://www.meetup.com/uagolang/)
- [Utah Go User Group](https://www.meetup.com/utahgophers/)
- [Women Who Go - San Francisco, CA](https://www.meetup.com/Women-Who-Go/)
- [Zürich Gophers - Zurich, Switzerland](https://www.meetup.com/zurich-gophers/)

_Add the group of your city/country here (send **PR**)_

**[⬆ back to top](#contents)**

## Style Guides

- [bahlo/go-styleguide](https://github.com/bahlo/go-styleguide)
- [CockroachDB](https://github.com/cockroachdb/cockroach/blob/master/docs/style.md)
- [GitLab](https://docs.gitlab.com/ee/development/go_guide/)
- [Google](https://google.github.io/styleguide/go/)
- [Hyperledger](https://github.com/hyperledger/fabric/blob/release-1.4/docs/source/style-guides/go-style.rst)
- [Magnetico](https://github.com/boramalper/magnetico/wiki/magnetico-Design-Specification)
- [Sourcegraph](https://docs.sourcegraph.com/dev/background-information/languages/go)
- [Thanos](https://thanos.io/tip/contributing/coding-style-guide.md/)
- [Trybe](https://github.com/betrybe/playbook-go/blob/main/README_EN.md)
- [Uber](https://github.com/uber-go/guide/blob/master/style.md)

**[⬆ back to top](#contents)**

## Social Media

### Twitter

- [@GoDiscussions](https://twitter.com/GoDiscussions)
- [@golang](https://twitter.com/golang)
- [@golang_news](https://twitter.com/golang_news)
- [@golangch](https://twitter.com/golangch)
- [@golangflow](https://twitter.com/golangflow)
- [@golangweekly](https://twitter.com/golangweekly)

**[⬆ back to top](#contents)**

### Reddit

- [r/golang](https://www.reddit.com/r/golang/)

**[⬆ back to top](#contents)**

## Websites

- [Awesome Go @LibHunt](https://go.libhunt.com) - Your go-to Go Toolbox.
- [Awesome Golang Workshops(stars: 488)](https://github.com/amit-davidson/awesome-golang-workshops) - A curated list of awesome golang workshops.
- [Awesome Remote Job(stars: 27205)](https://github.com/lukasz-madon/awesome-remote-job) - Curated list of awesome remote jobs. A lot of them are looking for Go hackers.
- [awesome-awesomeness(stars: 30922)](https://github.com/bayandin/awesome-awesomeness) - List of other amazingly awesome lists.
- [awesome-go-extra(stars: 22)](https://github.com/xwjdsh/awesome-go-extra) - Parse awesome-go README file and generate a new README file with repo info.
- [Code with Mukesh](https://codewithmukesh.com/blog/category/golang) - Software Engineer and Blogs @ codewithmukesh.com.
- [Coding Mystery](https://codingmystery.com) - Solve exciting escape-room-inspired programming challenges using Go.
- [CodinGame](https://www.codingame.com/) - Learn Go by solving interactive tasks using small games as practical examples.
- [Explore Go Libraries & Projects](https://kandi.openweaver.com/explore/go) - Discover & find a curated list of popular & new Go libraries, top authors, trending project kits, discussions, tutorials & learning resources on kandi.
- [Go Blog](https://blog.golang.org) - The official Go blog.
- [Go Code Club](https://www.youtube.com/watch?v=nvoIPQYdx9g&list=PLEcwzBXTPUE_YQR7R0BRtHBYJ0LN3Y0i3) - A group of Gophers read and discuss a different Go project every week.
- [Go Community on Hashnode](https://hashnode.com/n/go) - Community of Gophers on Hashnode.
- [Go Forum](https://forum.golangbridge.org) - Forum to discuss Go.
- [Go Projects](https://github.com/golang/go/wiki/Projects) - List of projects on the Go community wiki.
- [Go Proverbs](https://go-proverbs.github.io/) - Go Proverbs by Rob Pike.
- [Go Report Card](https://goreportcard.com) - A report card for your Go package.
- [go.dev](https://go.dev/) - A hub for Go developers.
- [gocryforhelp(stars: 41)](https://github.com/ninedraft/gocryforhelp) - Collection of Go projects that needs help. Good place to start your open-source way in Go.
- [godoc.org](https://godoc.org/) - Documentation for open source Go packages.
- [Golang Developer Jobs](https://golangjob.xyz) - Developer Jobs exclusively for Golang related Roles.
- [Golang Flow](https://golangflow.io) - Post Updates, News, Packages and more.
- [Golang News](https://golangnews.com) - Links and news about Go programming.
- [Golang Resources](https://golangresources.com) - A curation of the best articles, exercises, talks and videos to learn Go.
- [Golang Weekly](https://discu.eu/weekly/golang/) - Each monday projects, tutorials and articles about Go.
- [golang-graphics(stars: 139)](https://github.com/mholt/golang-graphics) - Collection of Go images, graphics, and art.
- [golang-nuts](https://groups.google.com/forum/#!forum/golang-nuts) - Go mailing list.
- [Google Plus Community](https://plus.google.com/communities/114112804251407510571) - The Google+ community for #golang enthusiasts.
- [Gopher Community Chat](https://invite.slack.golangbridge.org) - Join Our New Slack Community For Gophers ([Understand how it came](https://blog.gopheracademy.com/gophers-slack-community/)).
- [Gophercises](https://gophercises.com/) - Free coding exercises for budding gophers.
- [gowalker.org](https://gowalker.org) - Go Project API documentation.
- [json2go](https://m-zajac.github.io/json2go) - Advanced JSON to Go struct conversion - online tool.
- [justforfunc](https://www.youtube.com/c/justforfunc) - Youtube channel dedicated to Go programming language tips and tricks, hosted by Francesc Campoy [@francesc](https://twitter.com/francesc).
- [Learn Go Programming](https://blog.learngoprogramming.com) - Learn Go concepts with illustrations.
- [Made with Golang](https://madewithgolang.com/?ref=awesome-go)
- [r/Golang](https://www.reddit.com/r/golang) - News about Go.
- [studygolang](https://studygolang.com) - The community of studygolang in China.
- [Trending Go repositories on GitHub today](https://github.com/trending?l=go) - Good place to find new Go libraries.
- [TutorialEdge - Golang](https://tutorialedge.net/course/golang/)

**[⬆ back to top](#contents)**

### Tutorials

- [50 Shades of Go](https://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/) - Traps, Gotchas, and Common Mistakes for New Golang Devs.
- [A Comprehensive Guide to Structured Logging in Go](https://betterstack.com/community/guides/logging/logging-in-go/) - Delve deep into the world of structured logging in Go with a specific focus on recently accepted slog proposal which aims to bring high performance structured logging with levels to the standard library.
- [A Guide to Golang E-Commerce](https://snipcart.com/blog/golang-ecommerce-ponzu-cms-demo?utm_term=golang-ecommerce-ponzu-cms-demo) - Building a Golang site for e-commerce (demo included).
- [A Tour of Go](https://tour.golang.org/) - Interactive tour of Go.
- [Build a Database in 1000 lines of code]( https://link.medium.com/O9YQlx89Htb) - Build a NoSQL Database From Zero in 1000 Lines of Code.
- [Build web application with Golang(stars: 42582)](https://github.com/astaxie/build-web-application-with-golang) - Golang ebook intro how to build a web app with golang.
- [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql) - We’ll write an API with the help of the powerful Gorilla Mux.
- [Building Go Web Applications and Microservices Using Gin](https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin) - Get familiar with Gin and find out how it can help you reduce boilerplate code and build a request handling pipeline.
- [Caching Slow Database Queries](https://medium.com/@rocketlaunchr.cloud/caching-slow-database-queries-1085d308a0c9) - How to cache slow database queries.
- [Canceling MySQL](https://medium.com/@rocketlaunchr.cloud/canceling-mysql-in-go-827ed8f83b30) - How to cancel MySQL queries.
- [CodeCrafters Golang Track](https://app.codecrafters.io/tracks/go) - Achieve mastery in advanced Go by building your own Redis, Docker, Git, and SQLite. Featuring goroutines, systems programming, file I/O, and more.
- [Debugged.it Go patterns(stars: 11)](https://github.com/haveyoudebuggedit/go-patterns) - Advanced Go patterns with ready-to-run examples.
- [Design Patterns in Go(stars: 108)](https://github.com/shubhamzanwar/design-patterns) - Collection of programming design patterns implemented in Go.
- [Ethereum Development with Go(stars: 1656)](https://github.com/miguelmota/ethereum-development-with-go-book) - A little e-book on Ethereum Development with Go.
- [Games With Go](https://www.youtube.com/watch?v=9D4yH7e_ea8&list=PLDZujg-VgQlZUy1iCqBbe5faZLMkA3g2x) - A video series teaching programming and game development.
- [Go By Example](https://gobyexample.com/) - Hands-on introduction to Go using annotated example programs.
- [Go Cheat Sheet(stars: 7962)](https://github.com/a8m/go-lang-cheat-sheet) - Go's reference card.
- [Go database/sql tutorial](http://go-database-sql.org/) - Introduction to database/sql.
- [Go in 7 days(stars: 133)](https://github.com/harrytran103/7_days_of_go) - Learn everything about Go in 7 days (from a Nodejs developer).
- [Go Language Tutorial](https://www.javatpoint.com/go-tutorial) - Learn Go language Tutorial.
- [Go Tutorial](https://www.tutorialspoint.com/go/index.htm) - Learn Go programming.
- [Go WebAssembly Tutorial - Building a Simple Calculator](https://tutorialedge.net/golang/go-webassembly-tutorial/)
- [go-clean-template(stars: 5698)](https://github.com/evrone/go-clean-template) - Clean Architecture template for Golang services.
- [go-patterns(stars: 23332)](https://github.com/tmrts/go-patterns) - Curated list of Go design patterns, recipes and idioms.
- [goapp(stars: 695)](https://github.com/bnkamalesh/goapp) - An opinionated guideline to structure & develop a Go web application/service.
- [Golang for Node.js Developers(stars: 4423)](https://github.com/miguelmota/golang-for-nodejs-developers) - Examples of Golang compared to Node.js for learning.
- [Golang Tutorial Guide](https://www.freecodecamp.org/news/golang-tutorial-list-free-courses-learn-go-programming-language/) - A List of Free Courses to Learn the Go Programming Language.
- [Golangbot](https://golangbot.com/learn-golang-series/) - Tutorials to get started with programming in Go.
- [GopherCoding](https://gophercoding.com/) - Collection of code snippets and tutorials to help tackle every day issues.
- [GopherSnippets](https://gophersnippets.com/) - Code snippets with tests and testable examples for the Go programming language.
- [Gosamples](https://gosamples.dev/) - Collection of code snippets that let you solve everyday code problems.
- [Hackr.io](https://hackr.io/tutorials/learn-golang) - Learn Go from the best online golang tutorials submitted & voted by the golang programming community.
- [Hex Monscape(stars: 48)](https://github.com/Haraj-backend/hex-monscape) - Getting started guidelines in writing maintainable code using Hexagonal Architecture.
- [How to Benchmark: dbq vs sqlx vs GORM](https://medium.com/@rocketlaunchr.cloud/how-to-benchmark-dbq-vs-sqlx-vs-gorm-e814caacecb5) - Learn how to benchmark in Go. As a case-study, we will benchmark dbq, sqlx and GORM.
- [How To Deploy a Go Web Application with Docker](https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker) - Learn how to use Docker for Go development and how to build production Docker images.
- [How to Use Godog for Behavior-driven Development in Go](https://semaphoreci.com/community/tutorials/how-to-use-godog-for-behavior-driven-development-in-go) - Get started with Godog — a Behavior-driven development framework for building and testing Go applications.
- [Learn Go with 1000+ Exercises(stars: 17984)](https://github.com/inancgumus/learngo) - Learn Go with thousands of examples, exercises, and quizzes.
- [Learn Go with TDD(stars: 20647)](https://github.com/quii/learn-go-with-tests) - Learn Go with test-driven development.
- [Learning Go by examples](https://dev.to/aurelievache/learning-go-by-examples-introduction-448n) - Series of articles in order to learn Golang language by concrete applications as example.
- [Microservices with Go](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_) - Dive deep into building microservices using Go, including gRPC.
- [package main](https://www.youtube.com/packagemain) - YouTube channel about Programming in Go.
- [Programming with Google Go](https://www.coursera.org/specializations/google-golang) - Coursera Specialization to learn about Go from scratch.
- [Saving a Third of Our Memory by Re-ordering Go Struct Fields](https://qvault.io/golang/struct-field-ordering-memory/) - How inefficient field ordering in Go structs.
- [Scaling Go Applications](https://betterstack.com/community/guides/scaling-go/) - Everything about building, deploying and scaling Go applications in production.
- [The world’s easiest introduction to WebAssembly with Golang](https://medium.com/@martinolsansky/webassembly-with-golang-is-fun-b243c0e34f02)
- [W3basic Go Tutorials](https://www.w3basic.com/golang/) - W3Basic provides an in-depth tutorial and well-organized content to learn Golang programming.
- [Working with Go(stars: 1158)](https://github.com/mkaz/working-with-go) - Intro to go for experienced programmers.
- [Your basic Go](https://yourbasic.org/golang) - Huge collection of tutorials and how to's.

**[⬆ back to top](#contents)**

### Guided Learning

- [The Go Developer Roadmap](https://roadmap.sh/golang) - A visual roadmap that new Go developers can follow through to help them learn Go.
- [The Go Learning Path](https://tutorialedge.net/paths/golang/) - A guided learning path containing a mix of free and premium resources.

**[⬆ back to top](#contents)**
