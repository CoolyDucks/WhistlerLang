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
Furry Public Licence v1.0

YOURNAME ^ 2026

- When modifying a project, the names of project authors cannot be changed
- License of the original project cannot be altered
- Comments can be edited but cannot mention the original script owner
- Projects must remain safe and not harm devices
- Not crediting the original creator is considered a violation
```

---

## GitHub 🔗

https://github.com/CoolyDucks/WhistlerLang
