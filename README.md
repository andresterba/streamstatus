streamstatus
==============================

Streamstatus is a small cli tool to check the status of your favorite twitch streamers.

## Build

Build with `go build -o streamstatus`

## Usage

Add a file `.streamstatus` in your home directory. Add streamer and add category.

Example:
```
~ cat .streamstatus
demuslim sc2
TaKeTV sc2
sentdex python
jakenbakelive travel
```

## Output
```
~ streamstatus
Your streams are currently :
demuslim           sc2        online 
TaKeTV             sc2        offline 
sentdex            code       offline 
jakenbakelive      travel     offline 
```

