# 📝 AutoNote

AutoNote is a lightweight command-line application written in Go that lets you save quick notes instantly with automatic timestamps.  

It is designed for fast idea capture without switching context or opening heavy applications.

---

## ✨ Features
- Save notes directly from the command line  
- Automatically adds date and time  
- Appends notes to a Markdown file (`notes.md`)  
- Minimal, fast, and easy to use  

---

## 🚀 Usage
Add a note using a single command:

```bash
autonote "Meeting with ML team, discuss project proposal"

# 📝 AutoNote

AutoNote is a lightweight command-line application written in Go that lets you save quick notes instantly with automatic timestamps.

It is designed for fast idea capture without switching context or opening heavy applications.

---

## ✨ Features

- Save notes directly from the command line
- Automatically adds date and time
- Appends notes to a Markdown file (`notes.md`)
- Minimal, fast, and easy to use

---

## 🚀 Usage

Add a note using a single command:


autonote "Meeting with ML team, discuss dataset bias"
The note will be saved to notes.md with a timestamp.

## 📂 Project Structure
go

autonote/
├── main.go
├── go.mod
└── notes.md   # created automatically

## 🛠️ How to Run
1. Clone or create the project
git clone <your-repo-url>
cd autonote

2. Initialize Go module (if not already done)
go mod init autonote

3. Run the program
go run main.go "Write proposal introduction"

## ⚙️ Build as a CLI Tool
Build the executable:

go build -o autonote
Run it:

./autonote "Experiment idea for lung X-ray model"
(Optional) Move it to your PATH for global access:

sudo mv autonote /usr/local/bin/

##📄 Output Example (notes.md)
 
- **2025-11-26 23:10**: Meeting with ML team, discuss dataset bias
- **2025-11-26 23:25**: Prepare slides for proposal presentation

## 🎯 Why AutoNote?
Encourages fast knowledge capture
Reduces workflow interruption
Perfect for developers, students, and researchers

## 🔮 Future Improvements
Tags (e.g., --tag ML)
Daily note files
Search and list notes
Custom note directory

## 📜 License
MIT License
