# Hitbox Input Viewer

This is a (currently) very simple web app for showing your hitbox's inputs on stream. I hadn't been able to find one that didn't require Windows, and as a Linux user that primarily uses Proton + Steam to game now, wanted something that could work no matter what OS I boot into. This was made solely with Svelte, JS, and CSS, so it's also really lightweight.

Currently, this only supports Brook UFB as an input source (also I've only tested by forcing Xbox 360 mode, given I'm a Linux user). Also, this only works with default mappings. PRs to detect and set other input devices, such as Cthulhu/PS360+/anything else that can be SOCD cleaned for hitbox usage, would be very welcome!

## Usage

1. Navigate to https://hitbox-input-viewer.vercel.app or copy that link.
2. Add browser source to OBS or Xsplit with the URL.
3. Set browser source height to 720px and width to 1280px (This is for safety sake as I did not create a responsive design for this app -- it's supposed to be viewed as an OBS/Xsplit browser source after all!)
4. Press a button on your hitbox to display the input (will have to be done after every disconnect and on every launch of OBS/Xsplit)

## Roadmap

- Add ability to change colours
- Add ability to change mapping
- Add ability to have a custom URL parameter set your custom colours/mapping

## Donations

Not really necessary; I did this for myself as I wanted this available on my stream. If you use this app and like it, you can always [donate to me via streamlabs](https://streamlabs.com/mssngprsn/tip)
