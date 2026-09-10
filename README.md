# femboysay

its amazing, isn't it. now hyprland users can **EXPRESS** themselves.

```
╭─────────────────╮
│ femboys say hi! │
╰─────────────────╯

         \
          \
           \
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⠶⣶⣤⣔⣶⡶⣦⣤⣠⣶⡄ ...
```

*This idea is by Noel. IYKYK. CodeVerse Discord btw.*

But the programming part? **me**.

## Install

Requires Go 1.23+.

```sh
go build -o femboysay .
# optional: install to $GOPATH/bin
go install gitlab.com/lalaggi/femboysay@latest
```

## Usage

```sh
femboysay hello world
echo hello | femboysay
femboysay -w 60 hello world
femboysay -style double hello world
femboysay -h
femboysay -v
```

Options:

- `-w <n>` — maximum width of the speech bubble (default: 40)
- `-style <name>` — bubble style: `rounded` (default), `sharp`, `double`, `heavy`, `classic` (old cowsay look)
- `-h` — show help
- `-v` — print version

Reads the message from arguments, or from stdin when piped.

## License

[WTFPL](LICENSE) — do what the fuck you want to.
