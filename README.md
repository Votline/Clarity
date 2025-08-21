# Clarity - 20-20-20 Eye Care Reminder

Clarity is a minimalist and lightweight application that helps you protect your eyesight by enforcing the 20-20-20 rule. Built with Go and OpenGL(Go-gl), pops up when it's time to take a break.

## Description

Every 20 minutes, Clarity will play a sound signal and display a simple timer, reminding you to look at something 20 feet (6 meters) away for 20 seconds. This helps reduce eye strain caused by prolonged screen time.

### Key Features:
- **20-20-20 Rule Enforcement**: Accurate 20-minute work / 20-second break cycles
- **Procedural UI**: Renders clean, font-less graphics
- **Audible Alerts**: Clear sound signals for the start and end of each break
- **Lightweight**: Minimal resource consumption
- 
## Controls & Interaction
- **Timer Window**: Automatically appears when a break is due
- **Automatic Dismissal**: The break window closes automatically after 20 seconds with a sound

## Technologies
- **Go** - Primary programming language
- **[Go-gl](https://github.com/go-gl/gl)** - Graphics rendering
- **[Gl-glfw](https://github.com/go-gl/glfw)** - Window management

## Installation & Building

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/Votline/Clarity
    cd Clarity
    ```

2.  **Download dependencies**:
    ```bash
    go mod download
    ```

3.  **Build the application**:
    ```bash
    go build -o clarity
    ```

4.  **Run**:
    ```bash
    ./clarity 
    ```

## Licenses
This project is licensed under [GNU AGPL v3](LICENSE).

The full license texts are available in the [licenses directory](licenses/)
