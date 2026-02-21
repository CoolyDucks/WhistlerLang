# WhistlerLang

**WhistlerLang** is a custom programming language for **learning, prototyping, and creating scripts**.  
—MATE!

---

## Features 

- Interactive **REPL terminal** for rapid experimentation  
- Run `.whlst` scripts with `run <file.whlst>`  
- `say "text"` → prints text in REPL  
- Strong typing for variables declared with `let`  
- Conditional blocks with `if`, `elif`, `else`, and `end`  
- Time commands:  
  - `time.print` → prints current time  
  - `time.set "<FORMAT>" "<PREF>"` → change time format  

---

## Installation 

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

## Example Usage 

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

## License

```licence
----------------------
----------------------


LinkFurry License v1.0


----------------------
----------------------
  
  
  
Copyright (c) 2026, CoolyDucks (All rights reserved)  
Permission is granted, without charge, to any individual, organisation, or entity obtaining a copy of this software (hereinafter referred to as "the Software") and its documentation, to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit others to whom the Software is provided to do the same, provided that all of the following conditions are strictly observed:  
  
  
Preservation of Notices  
All copies, modified versions, derivative works, or adaptations of the Software, whether in source code, binary form, or any other form, must retain this copyright notice, the complete text of this license, all disclaimers, and any associated legal notices.  
Use of Name  
The name "LinkFurry", the copyright holder, or any contributor may not be used to endorse, promote, advertise, or imply association with any derivative works, software products, or services without explicit prior written consent from the copyright holder. No person or entity may claim that the original creator endorses, sponsors, or supports any modified version or derivative project.  
No Warranty  
The Software is provided "as is", without any express or implied warranties of any kind, including but not limited to merchantability, fitness for a particular purpose, non-infringement, accuracy, reliability, security, or suitability. The user assumes all risk regarding performance, use, or outcome from any action related to the Software.  
Liability Disclaimer  
Under no circumstances shall the copyright holder, contributors, distributors, or licensors of the Software be held responsible for any form of damages, claims, losses, or liabilities, whether direct, indirect, incidental, special, exemplary, or consequential, including without limitation:  
loss of profits, revenue, or data  
business interruption, system failure, or downtime  
breach of privacy, security, or confidentiality  
procurement of substitute goods, services, or remedies  
This applies regardless of the cause, including negligence, breach of contract, tort, strict liability, or any other legal theory, even if the possibility of such damages has been communicated or foreseen.  
Absolute Freedom and Protection  
Users are granted full freedom to adapt, modify, enhance, or redistribute the Software for any purpose, commercial or non-commercial, subject only to compliance with this license. The author and contributors are fully protected from claims, disputes, or liabilities arising from any use or distribution of the Software, in whole or in part.  
Derivative Works  
Any derivative works must explicitly reference LinkFurry v1.0 and retain the above copyright notice, disclaimers, and this license. Modifications or enhancements may not imply endorsement, sponsorship, or support by the original author without prior written consent.  
Jurisdiction and Severability  
This license shall be governed by the laws of England and Wales. Should any provision of this license be held invalid, illegal, or unenforceable, the remaining provisions shall continue in full force and effect.  
Complete Agreement  
This license represents the entire agreement between the copyright holder and the user regarding the Software. No other terms, conditions, or representations shall apply unless explicitly included in writing.
```

---

## GitHub 

https://github.com/CoolyDucks/WhistlerLang
