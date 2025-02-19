package main

import (
	"encoding/json"
	"net/http"
	"html/template"
	"log"
)

type KeyData struct {
	Keyname   string `json:"Keyname"`
	Keycode   int    `json:"Keycode"`
	PressCount int   `json:"PressCount"`
}

var keyData = []KeyData{
	{Keyname: "Enter", Keycode: 65293, PressCount: 23},
	{Keyname: "LAlt", Keycode: 65513, PressCount: 14},
	{Keyname: "j", Keycode: 106, PressCount: 12},
	{Keyname: "l", Keycode: 108, PressCount: 10},
	{Keyname: "k", Keycode: 107, PressCount: 6},
	{Keyname: "2", Keycode: 50, PressCount: 5},
	{Keyname: "4", Keycode: 52, PressCount: 4},
	{Keyname: "Space", Keycode: 32, PressCount: 4},
	{Keyname: "1", Keycode: 49, PressCount: 3},
	{Keyname: "LeftShift", Keycode: 65505, PressCount: 3},
	{Keyname: "s", Keycode: 115, PressCount: 3},
	{Keyname: "3", Keycode: 51, PressCount: 2},
	{Keyname: "e", Keycode: 101, PressCount: 2},
	{Keyname: "n", Keycode: 110, PressCount: 2},
	{Keyname: "t", Keycode: 116, PressCount: 2},
	{Keyname: "Backspace", Keycode: 65288, PressCount: 1},
	{Keyname: "I", Keycode: 73, PressCount: 1},
	{Keyname: "LeftCtrl", Keycode: 65507, PressCount: 1},
	{Keyname: "RightArrow", Keycode: 65363, PressCount: 1},
	{Keyname: "Unknown: 65474", Keycode: 65474, PressCount: 1},
	{Keyname: "Unknown: 65481", Keycode: 65481, PressCount: 1},
	{Keyname: "c", Keycode: 99, PressCount: 1},
	{Keyname: "d", Keycode: 100, PressCount: 1},
	{Keyname: "h", Keycode: 104, PressCount: 1},
	{Keyname: "o", Keycode: 111, PressCount: 1},
	{Keyname: "u", Keycode: 117, PressCount: 1},
}

// Home handler to render the page
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("index").Parse(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Key Press Heatmap</title>
    <style>
        .keyboard {
            display: grid;
            grid-template-columns: repeat(10, 50px);
            grid-gap: 5px;
            margin: 20px;
        }
        .key {
            width: 50px;
            height: 50px;
            background-color: lightgray;
            text-align: center;
            line-height: 50px;
            border-radius: 5px;
            cursor: pointer;
            transition: background-color 0.3s;
        }
        .key:hover {
            opacity: 0.8;
        }
    </style>
</head>
<body>

<div id="keyboard" class="keyboard"></div>

<script>
    const keyData = {{.}};

    const keyboardLayout = [
        "1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
        "q", "w", "e", "r", "t", "y", "u", "i", "o", "p",
        "a", "s", "d", "f", "g", "h", "j", "k", "l",
        "z", "x", "c", "v", "b", "n", "m", "Space"
    ];

    function generateKeyboard() {
        const keyboardContainer = document.getElementById('keyboard');

        keyboardLayout.forEach(key => {
            const keyElement = document.createElement('div');
            keyElement.classList.add('key');
            keyElement.textContent = key;
            keyboardContainer.appendChild(keyElement);

            const keyDataItem = keyData.find(item => item.Keyname.toLowerCase() === key);

            if (keyDataItem) {
                const pressCount = keyDataItem.PressCount;
                const intensity = Math.min(pressCount * 10, 255);
                keyElement.style.backgroundColor = `rgb(${intensity}, 0, 0)`;
            }
        });
    }

    generateKeyboard();
</script>

</body>
</html>
`)

    if err != nil {
        log.Fatal(err)
    }

    err = tmpl.Execute(w, keyData)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    http.HandleFunc("/", homeHandler)
    log.Fatal(http.ListenAndServe(":8999", nil))
}
