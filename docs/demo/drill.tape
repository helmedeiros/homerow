Output docs/demo/drill.gif

Set Shell "bash"
Set FontSize 16
Set Width 900
Set Height 540
Set Padding 20
Set Theme "Catppuccin Mocha"
Set TypingSpeed 150ms

Hide
Type "make build >/dev/null && export PS1= HOMEROW_DATA=$(mktemp -d)/progress.json && clear && ./bin/homerow"
Enter
Wait+Screen /enter start/
Down
Enter
Wait+Screen /NEW KEY/
Type "f"
Sleep 300ms
Type "j"
Wait+Screen /backspace fix/
Show
Sleep 1.5s
Type "f f f ff ff fff j j j jj "
Type "jk"
Sleep 1.5s
Backspace
Sleep 500ms
Type "j jjj fj jf ffjj jjff fjfj jfjf"
Sleep 2.5s
Hide
Escape
Sleep 500ms
Type "q"
Sleep 500ms
