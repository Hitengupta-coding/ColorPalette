# 🎨 Color Palette CLI

Color Palette CLI is a terminal-based program that allows users to explore a vibrant palette of colors, save their favorites, and interact with color groups.

---

## 🌟 Features
- 🎲 Random Color: Fetch a random color shade from a comprehensive list.
- 🧭 Specify a Group: Select a color group (e.g., red, blue, green) for random shade generation.
- 💾 Save Colors: Save custom colors with names and hex codes.
- 📋 View Saved Colors: Display saved colors stored in the program.
- 🔒 User-Friendly Commands: Offers intuitive and easy-to-use commands.
- 🚀 Fast and Lightweight: Designed to be quick and efficient with minimal dependencies.
- 🔧 Easily Customizable: Add more color groups or extend functionality as needed.

---

## 🛠️ Installation
1. Ensure [Go](https://golang.org/doc/install) is installed on your system.
2. Clone the repository:
3. ```bash
    git clone <your-repo-link>
    cd <your-repo-folder>
   ```

4. Build and run the program:
   ```bash
   go build colorpalette.go
   ./colorpalette
   ```

🎯 How to Use
- Run the program:go run main.go
   ```
   go run colorpalette.go
   ```

2. Available commands:
   - **getcolor**: Fetches a random color from all predefined groups.
   - **specify**: Specify a group (red, orange, yellow, green, blue, indigo, violet) to get a random color.
   - **add**: Save a custom color with its name and hex code.
   - **saved**: Displays saved colors.
   - **exit**: Exits the application.

---

## 🤝 Contributing
Contributions are welcome! To contribute:
1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Make your changes and commit them:
   ```bash
   git commit -m "Add your feature or fix"
   ```
4. Push your branch:
   ```bash
   git push origin feature/your-feature-name
   ```
5. Open a Pull Request.

Make sure your contributions follow the coding standards and are well-documented.

---

## 📜 **License Terms**
This project is licensed under the following terms:
- **No Redistribution or Reposting**: Users may not redistribute, repost, or share this project in any form.
- **Modifications**: Users may suggest modifications, but they must obtain explicit approval before implementation.
- **Personal Use Only**: This chatbot is intended for personal, non-commercial use.

---

## 📚 Example
Commands:
- **getcolor**: Get a random color shade.
- **specify**: Specify a color group to get a random shade.
- **add**: Save a color of your choice.
- **saved**: Display saved colors.
- **exit**: Exit the program.

Example interaction:
```plaintext
Enter a command: specify  
Enter the color group (red, orange, yellow, green, blue, indigo, violet): blue  
Random color from blue group: Dodger Blue: #1E90FF  
```

---

## 💡 Notes
- Requires Go to run the script.
- Customize by adding more color groups or extending functionality.

---
