import sys
import os
from datetime import datetime

TEMPLATE = """package main
func readFileLines(filepath string, split string) []string {
	content, _ := ioutil.ReadFile(filepath)
	return strings.Split(string(content), split)
}
func main() {
    lines := readFileLines("./test", "\\n")
}
"""

if __name__ == '__main__':
    now = datetime.now()
    if now.month != 12:
        sys.exit(1)

    folder = f"day{str(now.day).zfill(2)}"

    if os.path.exists(folder):
        sys.exit(1)

    os.mkdir(folder)
    with open(folder + "/main.go", "w") as file:
        file.write(TEMPLATE)

    open(folder + "/test", "w").close()
    open(folder + "/input", "w").close()