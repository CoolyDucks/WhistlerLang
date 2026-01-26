
# WhistlerLang  🚀

Welcome to WhistlerLang by CoolyDucks 🐤  
The first version of my programming language 🎉  

WhistlerLang is easy medium level language 💻  
It run scripts fast and compile to ELF for Linux amd64 amd32 arm64 arm32 ⚡  

Type `quit` for exit  
Type `run <script.whlst>` for run a script  
Type `help` for learning WhistlerLang 📖  (SOOOO EASY!!)

## Example 💡

say("Hello World")  
sift(1 2 3 4)  
math.add(5 6)  

Prompt look like this 🟢  (the best thing to read)

WhistlerLang 1.0 >>> say("Hi there")
Hi there
WhistlerLang 1.0 >>> run example.whlst
...
WhistlerLang 1.0 >>> quit

How to download and run ⬇️

1️⃣ Clone the repository

git clone https://github.com/CoolyDucks/WhistlerLang.git
cd WhistlerLang

2️⃣ Build WhistlerLang ELF

cd source
GOARCH=amd64 GOOS=linux go build -o ../WhistlerLang
GOARCH=386 GOOS=linux go build -o ../WhistlerLang386
GOARCH=arm64 GOOS=linux go build -o ../WhistlerLangArm64
GOARCH=arm GOOS=linux go build -o ../WhistlerLangArm

3️⃣ Run WhistlerLang

./WhistlerLang

Type your scripts or run .whlst files easily

License 📜

MIT License
Copyright (c) 2026 CoolyDucks 🐤

Permission is hereby granted free of charge to any person obtaining a copy
of this software and associated documentation files the Software to deal
in the Software without restriction including without limitation the rights
to use copy modify merge publish distribute sublicense and or sell copies
of the Software and to permit persons to whom the Software is furnished to do so
subject to the following conditions

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software

THE SOFTWARE IS PROVIDED AS IS WITHOUT WARRANTY OF ANY KIND EXPRESS OR IMPLIED
INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY FITNESS FOR A
PARTICULAR PURPOSE AND NONINFRINGEMENT IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
HOLDERS BE LIABLE FOR ANY CLAIM DAMAGES OR OTHER LIABILITY WHETHER IN AN ACTION OF CONTRACT
TORT OR OTHERWISE ARISING OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE
