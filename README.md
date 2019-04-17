streamstatus
==============================
[![Build Status](https://travis-ci.com/andresterba/streamstatus.svg?branch=master)](https://travis-ci.com/andresterba/streamstatus) [![Go Report Card](https://goreportcard.com/badge/github.com/andresterba/streamstatus)](https://goreportcard.com/report/github.com/andresterba/streamstatus)

Streamstatus is a small cli tool to check the status of your favorite twitch streamers.

## Installation

- If you are running [archlinux](https://www.archlinux.org/) use the [AUR-package](https://aur.archlinux.org/packages/streamstatus/)

- If not clone this repo and build with `go build -o streamstatus`

## Usage

Add `.streamstatus` to your home directory. Then add streamer and category.

Example:
```
~ cat .streamstatus
demuslim sc2
basetradetv sc2 
TaKeTV sc2 
Lamboking sc2 
Adam13531 code
```

## Run

Simply run `streamstatus show` to see all streamers.

```
~ streamstatus show
Your streams are currently :
[ 0] demuslim         sc2    online
[ 1] basetradetv      sc2    offline
[ 2] TaKeTV           sc2    online
[ 3] Lamboking        sc2    offline
[ 4] Adam13531        code   online
```

Run `streamstatus` with a category filter eg. `streamstatus filter code`to see all streamers in this category.

```
~ streamstatus filter code
Your streams are currently :
[ 0] Adam13531        code   online
```

