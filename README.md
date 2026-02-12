# WhistlerLang

**WhistlerLang** is a custom programming language for **learning, prototyping, and creating scripts**.  
—MATE!

---

## Features ✨

- Interactive **REPL terminal** for rapid experimentation  
- Run `.whlst` scripts with `run <file.whlst>`  
- `say "text"` → prints text in REPL  
- Strong typing for variables declared with `let`  
- Conditional blocks with `if`, `elif`, `else`, and `end`  
- Time commands:  
  - `time.print` → prints current time  
  - `time.set "<FORMAT>" "<PREF>"` → change time format  

---

## Installation 🚀

Clone the repository:

```bash
git clone https://github.com/CoolyDucks/WhistlerLang
cd WhistlerLang
./build.sh
```

---

## REPL Commands

- `quit` / `exit` → exit WhistlerLang  
- `run <file.whlst>` → run a WhistlerLang script  
- `say "text"` → print a string  
- `time.print` → show current time  
- `time.set "<FORMAT>" "<PREF>"` → change time format  
- `help` → show help menu  

---

## Example Usage 🎨

```
say "Hello World from WhistlerLang!"
let user = "Alice"
say "Welcome " + user + " to Syntexly That Beautiful"
time.print
time.set "{date} {hou}:{min}:{sec}" "ms"
time.print
let score = 85
if score >= 90
    say "Grade: A"
elif score >= 80
    say "Grade: B"
else
    say "Grade: C or below"
end
```

---

## License 📝

```licence
MIT License

Copyright (c) 2026 CoolyDucks

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

---

## GitHub 🔗

https://github.com/CoolyDucks/WhistlerLang
