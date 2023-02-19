# in-memory-notepad

A small project to try out Golang (following the In-Memory Notepad project from Hyperskill).
This program is supposed to work like a simple notepad app in which the user can keep various plain-text notes.
It first asks for the maximum number of notes the user intends to keep. Afterwards, it takes the following commands:

* **create _some text_**: Creates a new note with the specified text.
* **list**: Lists all notes that have been created so far together with their respective position.
* **delete _position_**: Deletes the note at the specified position (numbering starts from 1).
* **update _position_ _some new text_**: Changes the text of the note at this position (numbering starts from 1) to the newly specified text.
* **clear**: Deletes all notes. 
* **exit**: Terminates the program.
