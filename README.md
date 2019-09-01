# Nintaco Go API - Screenshot Example

### About

[The Nintaco NES/Famicom emulator](https://nintaco.com/) provides [a Go API](https://github.com/meatfighter/nintaco-go-api) that enables externally running programs to control the emulator at a very granular level. This example captures gameplay screenshots.

### Launch

1. Start Nintaco and launch a game.
2. Open the Start Program Server window via Tools | Start Program Server...
3. Press Start Server.
4. From the command-line, launch this Go program.
5. Press the Select gamepad button at any time to capture a screenshot. By default, Select is mapped to the apostrophe key.

### Getting Pixels

This example demonstrates how to use `API.GetPixels`, which captures the emulator's screen into a provided array. When `API.GetPixels` is invoked within a `FrameListener` receiver, the pixels values will be from the frame that is about to be displayed. If it is inovked prior to the listener callback, some of the values could be from a prior frame.

`API.GetPixels` obtains 9-bit extended palette indices. The lower 6 bits represent one of the 64 colors and the upper 3 bits describe if and how that color is emphasized. Usually there is no emphasis (the upper 3 bits will all be 0). This example uses a table to convert the extended palette indices into RGBA's.

A bot controlling Nintaco could potentially call `API.GetPixels` on every frame and analyze the image returned to determine what actions to carry out. However, it's usually far more practical and efficient to obtain the game state by reading CPU memory directly. But that technique requires researching how the game objects are represented in memory.  