# Quran File Renamer

This is a **Golang CLI tool** that renames **Quran audio files** based on Surah numbers.  
It adds **both English and Arabic Surah names** to file names.

## 🚀 Features

- Accepts **file patterns** (e.g., `*.mp3`).
- Matches **numeric filenames** (`1.mp3`, `2.mp3`, etc.) to **Surah names**.
- Renames files to **"001 - Al-Fathiha - الفاتحة.mp3"** format.
- Works for **multiple files at once**.

---

## 📦 Installation

1. Install **Go** (if not already installed):

   ```sh
   sudo apt install golang  # Ubuntu/Debian
   brew install go          # macOS
