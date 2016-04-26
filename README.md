# srbuf
Simple Ring Buffer in Go

Small, lightweight, efficient, and no-nonsense. Geared towards single byte read and write actions.
Buffer size is static and declared when buffer is instantiated.
Data overflow will simply overwrite oldest bytes. It is up to user application to read from buffer often enough to avoid overflow.
