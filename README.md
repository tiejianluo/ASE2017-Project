# ASE2017-Project
This is the course project codebase and the required tasks.

# Task 1 Local test
Run the following 4 project locally:

Firstly, install [git](https://git-scm.com/downloads)

Then download or git clone the whole project into your local machine and run it locally.
```
git clone https://github.com/tiejianluo/ASE2017-Project.git
```

# Task 2 Online demo
Setup your own github account and open a [GitHub Pages](https://pages.github.com/) for final assessment.

## Project 1 HTML + CSS + Javascript + jQuery
This is a self-exercise task.
### 1. Setup 
* OS: Linux / macOS
* Editor: [vim](http://www.vim.org/download.php) / [sublime-text](http://www.sublimetext.com/3) / [atom](https://atom.io/) / [VS code](https://code.visualstudio.com)
* Browser: [Chrome](http://www.google.cn/chrome/browser/desktop) / [Firefox](http://www.firefox.com.cn)
* Debug within browser: developer tools (F12 / right click -> inspect element) 
## Project 2 ninja-code for Javascript
###
## Project 3 blockly-games

## Project 4 blockly-games on Beego Framework
### 1. Install golang
Download the [archive](https://golang.org/dl) and extract it into /usr/local, creating a Go tree in /usr/local/go. For example:
```
tar -C /usr/local -xzf go$VERSION.$OS-$ARCH.tar.gz
```
Add /usr/local/go/bin to the PATH environment variable. 
You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:
```
export PATH=/usr/local/go/bin:$PATH
```
### 2. Setup [Beego](https://beego.me/quickstart) for blockly-games
Setup $GOPATH enviroment, it **MUST BE** under [your-download-path]/ASE2017-Project/go/workspace directory. 
You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile, for example:
```
export GOPATH=/home/[your-username]/ASE2017-Project/go/workspace
```
Add $GOPATH/bin to your PATH environment variable. 
You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:
```
export PATH=/usr/local/go/bin:$GOPATH/bin:$PATH
```
Enter $GOPATH/src 
```
cd $GOPATH/src
```
Start your first beego project!
```
bee run blocklyGame
```
