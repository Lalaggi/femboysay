package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

const (
	defaultWidth = 40
	defaultStyle = "rounded"
	version      = "1.1.0"
	femboyArt    = `
         \
          \
           \
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⠶⣶⣤⣔⣶⡶⣦⣤⣠⣶⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣯⡏⠉⣴⣿⢿⢿⣿⢷⣶⣍⡻⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢠⡶⢤⣴⣿⠟⠀⢁⡞⠛⡟⢣⠐⢳⡀⠈⢻⣯⡺⢦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢸⣇⣾⣟⠎⠀⢠⡞⠀⡀⢀⠆⠀⠀⣧⠀⡀⠹⡳⣦⣿⡦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⡸⣻⣟⡟⠀⠀⢸⠁⢀⢧⢸⡾⡀⠶⣿⡿⣳⣅⣼⢻⠙⢿⡶⢤⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠞⠉⡿⢡⢡⠀⠀⣿⠀⣿⣯⣭⣹⢝⡊⢸⠚⢻⠝⢺⣗⢆⣸⡷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⢠⢳⣮⣣⣤⣄⡟⣇⡿⣻⡿⣿⠉⠻⣜⠐⠿⣛⡻⠿⠿⠵⡳⣿⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢀⡿⢻⢡⡿⣼⡃⢡⠈⣴⠈⠒⠋⠀⠀⠐⠀⠀⠀⠌⡟⡦⣌⡉⠻⡉⢣⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢸⠃⢹⣾⠀⢹⣷⡜⣷⣿⣧⠀⠀⢴⠞⠛⣿⠀⢀⢧⡜⡟⠀⠉⠓⢄⡀⠈⠓⠤⣀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠈⠀⢸⢹⡆⢸⣿⡛⡷⣝⣽⠳⢦⣌⡓⠀⣢⠔⡡⢻⡋⠀⠀⠀⠀⠀⠑⢦⡀⠀⠀⠑⠢⣄⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠈⣼⢇⡼⠿⣿⣿⡿⣿⡄⠀⠀⢹⢉⡰⠛⠁⠈⠁⠀⠀⠀⠀⠀⠀⠀⠙⢶⡀⠀⠀⠀⠙⠢⡀⠀
⠀⠀⠀⠀⠀⢏⣞⠝⠀⠀⠸⣿⣿⡺⣮⣤⢀⣘⠓⠶⣶⡔⢶⣖⢤⡤⠤⠤⠠⣖⢲⣦⡤⠽⢦⡀⠀⠀⠀⠘⣆
⠀⠀⠀⢀⢴⣾⡾⠀⠀⠀⢀⠚⣿⣷⡈⠻⣆⡈⠁⢁⣈⣇⠘⣿⣄⠀⠀⠀⠀⢸⠈⡿⡇⠀⠀⠉⠀⠀⠀⠀⢸
⠀⠀⠀⡎⣫⠟⡇⠀⠀⠀⠘⠀⠈⠻⣿⣶⣾⣽⣲⣗⠋⢻⣦⣿⠀⠀⠀⠀⠀⢠⣤⣆⡧⠤⠤⠄⠤⠐⠒⠂⠁
⠀⠀⢀⢾⢯⡎⠀⠀⠀⠀⠀⣴⠁⠀⠀⠉⠙⢟⠻⠿⣿⣼⡟⣬⠂⣧⠴⠐⠂⠚⠋⠛⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⣸⡼⢣⣿⡖⠒⠂⠀⠒⡳⠀⠀⠀⠀⠀⠈⠑⠚⠓⣟⣉⣳⠀⢱⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⣿⡧⢺⢻⣿⣿⣓⣲⣾⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⡇⡇⠫⣆⠸⠄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⡻⡅⠊⣶⡇⠀⠀⠀⢸⡆⠀⠀⠀⠀⠀⠀⠀⠀⢀⡇⠀⠀⠸⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⢻⡿⡾⠈⡇⠀⠀⠀⢸⢇⠀⠀⠀⠀⠀⠀⠀⠀⠘⣴⡀⠀⠀⢸⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠰⣝⠿⣐⡏⠀⠀⠀⣼⡸⠀⠀⠀⠀⠀⠀⠀⠀⠀⢣⠱⡄⠀⡟⣷⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠙⣿⣿⣯⠀⠀⠀⡟⡇⠀⠀⠀⠀⠀⣀⣀⡀⠤⠬⢆⡿⣄⣃⡈⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠹⣿⣿⠀⠀⠀⢱⣷⠈⠉⠉⠉⠉⠀⠀⠀⠀⠀⠘⠁⢸⠧⠞⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠈⠯⣆⠀⢀⡴⠁⠀⠀⠀⠀⠀⠀⠀⠀⢀⠆⠀⢀⡎⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠈⣳⠟⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⡿⠀⠀⢸⡃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢠⡾⢇⠀⠀⠉⠓⠲⠤⠤⠤⠤⠤⠤⠤⠲⢲⣻⡁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⣰⠏⠀⡸⠓⢵⣀⡀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣨⣼⣇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⢀⣾⠛⢀⡞⠁⠀⠀⡞⠉⠉⠉⠁⡖⠀⠈⠹⡍⠉⠈⠙⣞⣆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⢠⡴⡟⠉⠀⡾⠀⠀⠀⡸⠀⠀⠀⠀⢠⠇⠀⠀⠀⡇⠀⠀⠀⠘⡜⢦⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢸⣿⡽⠀⠀⡰⠁⠀⠀⢰⠃⠀⠀⠀⠀⣼⠄⠀⠀⠀⢱⡄⠀⠀⠀⠹⣄⢳⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⢨⣿⣄⠀⣴⠃⠀⠀⢀⡏⠀⠀⠀⠀⠀⣯⠀⠀⠀⠀⢸⣳⠀⠀⠀⠀⠙⣦⢝⣶⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠻⢿⣿⣥⣀⡀⠀⡼⠀⠀⠀⠀⠀⠀⣇⠀⠀⠀⠀⠀⡟⢇⠀⢀⠀⣀⣌⣿⡃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠸⠟⠙⢻⣿⣶⣿⣿⣷⣶⣤⣀⣰⡿⣿⣶⣶⣦⣤⣼⣾⠿⠛⠛⠋⠉⠉⢳⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⢧⠉⠙⠁⠈⠉⠛⠛⠚⠛⢀⡇⠉⢯⠙⠉⠁⠀⠀⠀⠀⠀⠀⠀⢣⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠸⡄⠀⠀⠀⠀⠀⠀⠀⠀⠰⠂⠀⠀⢧⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢧⠀⠀⠀⠀⠀⠀⠀⠀⠀`
)

// Message holds the wrapped lines of a speech bubble.
type Message struct {
	lines    []string
	maxWidth int
}

func printArt() {
	fmt.Println(femboyArt)
}

func newMessage() *Message {
	return &Message{
		lines:    make([]string, 0),
		maxWidth: 0,
	}
}

func displayWidth(s string) int {
	return utf8.RuneCountInString(s)
}

// addLine appends a line and tracks the widest line (in runes).
func (m *Message) addLine(line string) {
	m.lines = append(m.lines, line)
	if w := displayWidth(line); w > m.maxWidth {
		m.maxWidth = w
	}
}

// splitWord breaks a single over-long word into width-sized chunks.
func splitWord(word string, width int) []string {
	if width <= 0 || displayWidth(word) <= width {
		return []string{word}
	}
	var chunks []string
	var current strings.Builder
	count := 0
	for _, r := range word {
		if count >= width {
			chunks = append(chunks, current.String())
			current.Reset()
			count = 0
		}
		current.WriteRune(r)
		count++
	}
	if current.Len() > 0 {
		chunks = append(chunks, current.String())
	}
	return chunks
}

// wrapText word-wraps text into the message, preserving explicit newlines.
func (m *Message) wrapText(text string, width int) {
	for _, para := range strings.Split(text, "\n") {
		// Preserve blank lines as empty bubble rows.
		if len(strings.TrimSpace(para)) == 0 {
			m.addLine("")
			continue
		}

		words := strings.Fields(para)
		currentLine := ""
		currentWidth := 0

		flush := func() {
			if currentLine != "" {
				m.addLine(currentLine)
				currentLine = ""
				currentWidth = 0
			}
		}

		for _, word := range words {
			// Break words longer than the width.
			if displayWidth(word) > width {
				flush()
				for _, chunk := range splitWord(word, width) {
					m.addLine(chunk)
				}
				continue
			}

			w := displayWidth(word)
			sep := 0
			if currentWidth > 0 {
				sep = 1
			}
			if currentWidth+sep+w <= width {
				if currentWidth > 0 {
					currentLine += " "
				}
				currentLine += word
				currentWidth += sep + w
			} else {
				flush()
				currentLine = word
				currentWidth = w
			}
		}
		flush()
	}
}

// bubbleStyle defines the border characters for a speech bubble.
type bubbleStyle struct {
	topLeft     string
	horizontal  string
	topRight    string
	side        string
	bottomLeft  string
	bottomRight string
}

// Pretty unicode box styles. "classic" keeps the old cowsay look.
var bubbleStyles = map[string]bubbleStyle{
	"rounded": {"╭", "─", "╮", "│", "╰", "╯"},
	"sharp":   {"┌", "─", "┐", "│", "└", "┘"},
	"double":  {"╔", "═", "╗", "║", "╚", "╝"},
	"heavy":   {"┏", "━", "┓", "┃", "┗", "┛"},
}

func printBorder(width int, left, middle, right rune) {
	fmt.Printf("%c", left)
	fmt.Print(strings.Repeat(string(middle), width+2))
	fmt.Printf("%c\n", right)
}

func (m *Message) printClassicBubble() {
	width := m.maxWidth

	// Top border.
	printBorder(width, ' ', '_', ' ')

	for i, line := range m.lines {
		var leftChar, rightChar rune

		switch {
		case len(m.lines) == 1:
			leftChar = '<'
			rightChar = '>'
		case i == 0:
			leftChar = '/'
			rightChar = '\\'
		case i == len(m.lines)-1:
			leftChar = '\\'
			rightChar = '/'
		default:
			leftChar = '|'
			rightChar = '|'
		}

		padding := width - displayWidth(line)
		fmt.Printf("%c %s%s %c\n", leftChar, line, strings.Repeat(" ", padding), rightChar)
	}

	// Bottom border.
	printBorder(width, ' ', '-', ' ')
}

func (m *Message) printBubble(styleName string) {
	// Old cowsay look on request.
	if styleName == "classic" || styleName == "ascii" || styleName == "cowsay" {
		m.printClassicBubble()
		return
	}

	style, ok := bubbleStyles[styleName]
	if !ok {
		style = bubbleStyles["rounded"]
	}
	width := m.maxWidth

	// Top border.
	fmt.Printf("%s%s%s\n", style.topLeft, strings.Repeat(style.horizontal, width+2), style.topRight)

	for _, line := range m.lines {
		padding := width - displayWidth(line)
		fmt.Printf("%s %s%s %s\n", style.side, line, strings.Repeat(" ", padding), style.side)
	}

	// Bottom border.
	fmt.Printf("%s%s%s\n", style.bottomLeft, strings.Repeat(style.horizontal, width+2), style.bottomRight)
}

func readStdin() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	// Allow long piped input.
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return strings.Join(lines, "\n"), nil
}

func stdinIsPipe() bool {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice == 0
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [message]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nPrints MESSAGE (or stdin) in a speech bubble next to a femboy.\n")
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\nExamples:\n")
	fmt.Fprintf(os.Stderr, "  %s hello world\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  echo hello | %s\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s -w 60 hello world\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s -style double hello world\n", os.Args[0])
}

func main() {
	width := flag.Int("w", defaultWidth, "maximum `width` of the speech bubble")
	style := flag.String("style", defaultStyle, "bubble `style`: rounded, sharp, double, heavy, classic")
	showHelp := flag.Bool("h", false, "show this help message")
	showVersion := flag.Bool("v", false, "print version and exit")

	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		usage()
		os.Exit(0)
	}

	if *showVersion {
		fmt.Printf("femboysay %s\n", version)
		os.Exit(0)
	}

	if *width <= 0 {
		fmt.Fprintf(os.Stderr, "Error: width must be positive, got %d\n", *width)
		os.Exit(1)
	}

	*style = strings.ToLower(strings.TrimSpace(*style))
	switch *style {
	case "rounded", "sharp", "double", "heavy", "classic", "ascii", "cowsay":
		// valid
	default:
		fmt.Fprintf(os.Stderr, "Error: unknown style %q (choose: rounded, sharp, double, heavy, classic)\n", *style)
		os.Exit(1)
	}

	var messageText string
	if flag.NArg() > 0 {
		messageText = strings.Join(flag.Args(), " ")
	} else if stdinIsPipe() {
		var err error
		messageText, err = readStdin()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	} else {
		usage()
		fmt.Fprintf(os.Stderr, "\nError: no message provided\n")
		os.Exit(1)
	}

	messageText = strings.TrimSpace(messageText)
	if len(messageText) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no message provided\n")
		os.Exit(1)
	}

	msg := newMessage()
	msg.wrapText(messageText, *width)

	msg.printBubble(*style)
	printArt()
}
