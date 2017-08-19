# ASE2017-Project
This is the course project codebase and the required tasks.

Firstly, install [git](https://git-scm.com/downloads)

Then download or git clone the whole project into your local machine and run it locally.
```
git clone https://github.com/tiejianluo/ASE2017-Project.git
```

## Project 1 HTML + CSS + Javascript + jQuery
This is a self-exercise task.
### 1. Setup 
* OS: Linux / macOS
* Editor: [vim](http://www.vim.org/download.php) / [sublime-text](http://www.sublimetext.com/3) / [atom](https://atom.io/) / [VS code](https://code.visualstudio.com)
* Browser: [Chrome](http://www.google.cn/chrome/browser/desktop) / [Firefox](http://www.firefox.com.cn)
## Project 2 ninja-code for Javascript
###
## Project 3 blockly-games

## Project 4 blockly-games on Beego Framework
### 1. install golang
Download the [archive](https://golang.org/dl) and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:
```
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```
Add /usr/local/go/bin to the PATH environment variable. 
You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:
```
export PATH=/usr/local/go/bin:$PATH
```
### 2. install [Beego](https://beego.me/quickstart)
Setup $GOPATH enviroment, it <span style="color: #FFFF00">MUST BE</span> under your workspace directory. 
You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile, for example:
```
export GOPATH=/home/[your-username]/ASE2017-Project/go/workspace
```
Enter your $GOPATH directory:
```
cd $GOPATH
```
using go to get beego
```
go get -u github.com/astaxie/beego
```
using go to get bee
```
go get -u github.com/beego/bee
```
Add $GOPATH/bin to your PATH environment variable. 
You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:
```
export PATH=/usr/local/go/bin:$GOPATH/bin:$PATH
```
Enter $GOPATH/src and start your first beego project!
```
bee run
```
